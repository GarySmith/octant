package api

import (
	"context"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/heptio/developer-dash/internal/cluster"
	"github.com/heptio/developer-dash/internal/log"
	"github.com/heptio/developer-dash/internal/module"
	"k8s.io/apimachinery/pkg/labels"
)

type contentHandler struct {
	modulePaths map[string]module.Module
	modules     []module.Module
	logger      log.Logger
	prefix      string
	nsClient    cluster.NamespaceInterface
}

func (h *contentHandler) RegisterRoutes(ctx context.Context, router *mux.Router) error {
	s := router.PathPrefix("/content").Subrouter() // Expands to "/api/v1/content"
	sm := module.MuxRouter{Router: s}

	// Register content paths for all modules
	for _, m := range h.modules {
		// Registers a content handler for the module - adds up to /api/v1/content/{module}/.....
		h.registerModuleRoute(ctx, sm, m)
	}
	return nil
}

// Register content routes for a specified module
func (h *contentHandler) registerModuleRoute(ctx context.Context, router module.Router, m module.Module) {
	h.logger.Infof("Registering routes for %v", m.Name())
	parent := router.PathPrefix(path.Join("/", m.Name())).Subrouter() // e.g. /overview

	ns := parent.PathPrefix("/namespace/{namespace}").Subrouter()

	for path, handler := range m.Handlers(ctx) {
		// Namespace is optional, so register two alternatives
		ns.Handle(path, handler)
		parent.Handle(path, handler)
	}

	// Namespace is optional, so register two alternatives
	contentHandler := h.handlerForModule(m)
	ns.HandleFunc("/{contentPath:.*?}", contentHandler)
	parent.HandleFunc("/{contentPath:.*?}", contentHandler)
}

// Returns a content http handler for the specified module
func (h *contentHandler) handlerForModule(m module.Module) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		namespace := vars["namespace"] // optional
		if namespace == "" {
			// Fallback to legacy query parameter
			namespace = r.URL.Query().Get("namespace")
		}
		contentPath := path.Join("/", vars["contentPath"]) // the trailing path after optional namespace

		ctx := log.WithLoggerContext(r.Context(), h.logger)
		q := r.URL.Query()
		poll := q.Get("poll")

		filters := q["filter"]

		h.logger.With(
			"module", m.Name(),
			"path", r.URL.Path,
			"namespace", namespace,
			"contentPath", contentPath,
			"poll", poll,
			"filters", fmt.Sprintf("%v", filters),
		).Debugf("content")

		set, err := selectorFromFilters(filters)
		if err != nil {
			h.logger.Errorf("invalid filters: %v", err)
			RespondWithError(w, http.StatusInternalServerError, err.Error(), h.logger)
			return
		}

		if len(set) > 0 {
			h.logger.Debugf("Label Set: %s", set)
		}

		if poll != "" {
			h.handlePoll(ctx, poll, namespace, &set, contentPath, w, r, m)
			return
		}

		resp, err := m.Content(ctx, contentPath, h.prefix, namespace, module.ContentOptions{LabelSet: &set})
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error(), h.logger)
			return
		}

		serveAsJSON(w, resp, h.logger)
	}
}

func (h *contentHandler) handlePoll(ctx context.Context, poll, namespace string, labelSet *labels.Set, contentPath string, w http.ResponseWriter, r *http.Request, m module.Module) {
	eventTimeout := defaultEventTimeout
	timeout, err := strconv.Atoi(poll)
	if err == nil {
		eventTimeout = time.Duration(timeout) * time.Second
	}

	eventGenerators := []eventGenerator{
		&contentEventGenerator{
			generatorFn: m.Content,
			path:        contentPath,
			prefix:      h.prefix,
			namespace:   namespace,
			labelSet:    labelSet,
			runEvery:    eventTimeout,
		},
		&navigationEventGenerator{
			modules:   h.modules,
			namespace: namespace,
		},
		&namespaceEventGenerator{
			nsClient: h.nsClient,
		},
	}

	cs := contentStreamer{
		eventGenerators: eventGenerators,
		w:               w,
		logger:          h.logger,
		streamFn:        stream,
		contentPath:     contentPath,
	}

	if err = cs.content(ctx); err != nil {
		h.logger.Errorf("content error: %v", err)
	}
}

// selectorFromFilters builds a labels.Selector from a list of
// "key:value" formatted strings
func selectorFromFilters(filters []string) (labels.Set, error) {
	set := labels.Set{}

	for _, f := range filters {
		// Filters will be of the format "key:value"
		split := strings.SplitN(f, ":", 2)
		if len(split) < 2 {
			return nil, fmt.Errorf("invalid filter: %s", f)
		}

		set[split[0]] = split[1]
	}

	return set, nil
}
