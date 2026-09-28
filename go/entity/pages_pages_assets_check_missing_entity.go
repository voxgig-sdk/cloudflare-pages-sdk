package entity

import (
	"github.com/voxgig-sdk/cloudflare-pages-sdk/go/core"

	vs "github.com/voxgig-sdk/cloudflare-pages-sdk/go/utility/struct"
)

type PagesPagesAssetsCheckMissingEntity struct {
	name    string
	client  *core.CloudflarePagesSDK
	utility *core.Utility
	entopts map[string]any
	data    map[string]any
	match   map[string]any
	entctx  *core.Context
	deleted bool
}

func NewPagesPagesAssetsCheckMissingEntity(client *core.CloudflarePagesSDK, entopts map[string]any) *PagesPagesAssetsCheckMissingEntity {
	if entopts == nil {
		entopts = map[string]any{}
	}
	if _, ok := entopts["active"]; !ok {
		entopts["active"] = true
	} else if entopts["active"] == false {
		// keep false
	} else {
		entopts["active"] = true
	}

	e := &PagesPagesAssetsCheckMissingEntity{
		name:    "pages_pages_assets_check_missing",
		client:  client,
		utility: client.GetUtility(),
		entopts: entopts,
		data:    map[string]any{},
		match:   map[string]any{},
	}

	e.entctx = e.utility.MakeContext(map[string]any{
		"entity":  e,
		"entopts": entopts,
	}, client.GetRootCtx())

	e.utility.FeatureHook(e.entctx, "PostConstructEntity")

	return e
}

func (e *PagesPagesAssetsCheckMissingEntity) GetName() string { return e.name }

func (e *PagesPagesAssetsCheckMissingEntity) MarkDeleted() {
	e.deleted = true
}


// Deleted reports whether a successful Remove has resolved on this instance.
func (e *PagesPagesAssetsCheckMissingEntity) Deleted() bool {
	return e.deleted
}


func (e *PagesPagesAssetsCheckMissingEntity) Make() core.Entity {
	opts := map[string]any{}
	for k, v := range e.entopts {
		opts[k] = v
	}
	return NewPagesPagesAssetsCheckMissingEntity(e.client, opts)
}

func (e *PagesPagesAssetsCheckMissingEntity) Data(args ...any) any {
	if len(args) > 0 && args[0] != nil {
		e.data = core.ToMapAny(vs.Clone(args[0]))
		if e.data == nil {
			e.data = map[string]any{}
		}
		e.utility.FeatureHook(e.entctx, "SetData")
	}

	e.utility.FeatureHook(e.entctx, "GetData")
	out := vs.Clone(e.data)
	return out
}

func (e *PagesPagesAssetsCheckMissingEntity) Match(args ...any) any {
	if len(args) > 0 && args[0] != nil {
		e.match = core.ToMapAny(vs.Clone(args[0]))
		if e.match == nil {
			e.match = map[string]any{}
		}
		e.utility.FeatureHook(e.entctx, "SetMatch")
	}

	e.utility.FeatureHook(e.entctx, "GetMatch")
	out := vs.Clone(e.match)
	return out
}

// DataTyped is the statically-typed accessor for this entity's data. With no
// argument it returns the current data as an PagesPagesAssetsCheckMissing; with an argument it
// sets the data and returns the stored value. It delegates to the untyped Data
// (identical runtime) and converts at the typed boundary.
func (e *PagesPagesAssetsCheckMissingEntity) DataTyped(data ...PagesPagesAssetsCheckMissing) PagesPagesAssetsCheckMissing {
	if len(data) > 0 {
		return typedFrom[PagesPagesAssetsCheckMissing](e.Data(asMap(data[0])))
	}
	return typedFrom[PagesPagesAssetsCheckMissing](e.Data())
}

// MatchTyped mirrors DataTyped for the entity's match filter. The match is a
// partial of the entity, so it round-trips through PagesPagesAssetsCheckMissing (all fields
// optional at the wire level).
func (e *PagesPagesAssetsCheckMissingEntity) MatchTyped(match ...PagesPagesAssetsCheckMissing) PagesPagesAssetsCheckMissing {
	if len(match) > 0 {
		return typedFrom[PagesPagesAssetsCheckMissing](e.Match(asMap(match[0])))
	}
	return typedFrom[PagesPagesAssetsCheckMissing](e.Match())
}

func (e *PagesPagesAssetsCheckMissingEntity) Stream(action string, args map[string]any, callopts map[string]any) <-chan any {
	out := make(chan any)

	if callopts == nil {
		callopts = map[string]any{}
	}

	var signal <-chan struct{}
	switch s := callopts["signal"].(type) {
	case <-chan struct{}:
		signal = s
	case chan struct{}:
		signal = s
	}

	ctrl := map[string]any{}
	if c := core.ToMapAny(callopts["ctrl"]); c != nil {
		for k, v := range c {
			ctrl[k] = v
		}
	}

	ctxmap := map[string]any{
		"opname": action,
		"ctrl":   ctrl,
		"match":  e.match,
		"data":   e.data,
	}
	for k, v := range args {
		ctxmap[k] = v
	}

	utility := e.utility
	ctx := utility.MakeContext(ctxmap, e.entctx)
	ctx.Meta["stream"] = callopts

	// Outbound: expose the caller's payload so the request builder / transport
	// can stream it as the request body.
	if body := callopts["body"]; body != nil {
		ctx.Reqdata["body$"] = body
		ctx.Meta["stream_out"] = body
	}

	send := func(item any) bool {
		select {
		case <-signal:
			return false
		case out <- item:
			return true
		}
	}

	go func() {
		defer close(out)

		utility.FeatureHook(ctx, "PrePoint")
		point, err := utility.MakePoint(ctx)
		ctx.Out["point"] = point
		if err != nil {
			return
		}

		utility.FeatureHook(ctx, "PreSpec")
		spec, err := utility.MakeSpec(ctx)
		ctx.Out["spec"] = spec
		if err != nil {
			return
		}

		utility.FeatureHook(ctx, "PreRequest")
		req, err := utility.MakeRequest(ctx)
		ctx.Out["request"] = req
		if err != nil {
			return
		}

		utility.FeatureHook(ctx, "PreResponse")
		resp, err := utility.MakeResponse(ctx)
		ctx.Out["response"] = resp
		if err != nil {
			return
		}

		utility.FeatureHook(ctx, "PreResult")
		result, err := utility.MakeResult(ctx)
		ctx.Out["result"] = result
		if err != nil {
			return
		}

		utility.FeatureHook(ctx, "PreDone")

		// Inbound: prefer the streaming feature's incremental iterator; else
		// fall back to the materialised items so Stream always yields.
		if ctx.Result != nil && ctx.Result.Stream != nil {
			for item := range ctx.Result.Stream() {
				if !send(item) {
					return
				}
			}
			return
		}

		data, derr := utility.Done(ctx)
		if derr != nil {
			return
		}
		switch d := data.(type) {
		case []any:
			for _, item := range d {
				if !send(item) {
					return
				}
			}
		case nil:
			// nothing to yield
		default:
			send(d)
		}
	}()

	return out
}

func (e *PagesPagesAssetsCheckMissingEntity) Load(_ map[string]any, _ map[string]any) (any, error) {
	return core.UnsupportedOp("load", e.name)
}


func (e *PagesPagesAssetsCheckMissingEntity) List(_ map[string]any, _ map[string]any) (any, error) {
	return core.UnsupportedOp("list", e.name)
}



func (e *PagesPagesAssetsCheckMissingEntity) Create(reqdata map[string]any, ctrl map[string]any) (any, error) {
	utility := e.utility
	ctx := utility.MakeContext(map[string]any{
		"opname":  "create",
		"ctrl":    ctrl,
		"match":   e.match,
		"data":    e.data,
		"reqdata": reqdata,
	}, e.entctx)

	return e.runOp(ctx, func() {
		if ctx.Result != nil {
			if ctx.Result.Resdata != nil {
				e.data = core.ToMapAny(vs.Clone(ctx.Result.Resdata))
				if e.data == nil {
					e.data = map[string]any{}
				}
			}
		}
	})
}

// CreateTyped is the statically-typed variant of Create: it takes an
// PagesPagesAssetsCheckMissingCreateData and returns an PagesPagesAssetsCheckMissing. It delegates to the untyped
// Create (identical runtime) and converts at the typed boundary.
func (e *PagesPagesAssetsCheckMissingEntity) CreateTyped(reqdata PagesPagesAssetsCheckMissingCreateData, ctrl map[string]any) (PagesPagesAssetsCheckMissing, error) {
	res, err := e.Create(asMap(reqdata), ctrl)
	if err != nil {
		return PagesPagesAssetsCheckMissing{}, err
	}
	return typedFrom[PagesPagesAssetsCheckMissing](res), nil
}



func (e *PagesPagesAssetsCheckMissingEntity) Update(_ map[string]any, _ map[string]any) (any, error) {
	return core.UnsupportedOp("update", e.name)
}


func (e *PagesPagesAssetsCheckMissingEntity) Remove(_ map[string]any, _ map[string]any) (any, error) {
	return core.UnsupportedOp("remove", e.name)
}


func (e *PagesPagesAssetsCheckMissingEntity) runOp(ctx *core.Context, postDone func()) (any, error) {
	utility := e.utility

	utility.FeatureHook(ctx, "PrePoint")
	point, err := utility.MakePoint(ctx)
	ctx.Out["point"] = point
	if err != nil {
		return utility.MakeError(ctx, err)
	}

	utility.FeatureHook(ctx, "PreSpec")
	spec, err := utility.MakeSpec(ctx)
	ctx.Out["spec"] = spec
	if err != nil {
		return utility.MakeError(ctx, err)
	}

	utility.FeatureHook(ctx, "PreRequest")
	resp, err := utility.MakeRequest(ctx)
	ctx.Out["request"] = resp
	if err != nil {
		return utility.MakeError(ctx, err)
	}

	utility.FeatureHook(ctx, "PreResponse")
	resp2, err := utility.MakeResponse(ctx)
	ctx.Out["response"] = resp2
	if err != nil {
		return utility.MakeError(ctx, err)
	}

	utility.FeatureHook(ctx, "PreResult")
	result, err := utility.MakeResult(ctx)
	ctx.Out["result"] = result
	if err != nil {
		return utility.MakeError(ctx, err)
	}

	utility.FeatureHook(ctx, "PreDone")
	postDone()

	out, doneErr := utility.Done(ctx)
	if doneErr != nil {
		return out, doneErr
	}

	opname := ""
	if ctx.Op != nil {
		opname = ctx.Op.Name
	}

	if ctx.Result != nil && ctx.Result.Ok && opname != "list" {
		if opname == "remove" {
			e.MarkDeleted()
		}
		return e, nil
	}

	return out, nil
}
