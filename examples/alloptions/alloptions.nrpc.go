// This code was autogenerated from alloptions.proto, do not edit.
package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/wyy-go/gogo-nrpc"
	github_com_wyy_go_gogo_nrpc "github.com/wyy-go/gogo-nrpc"
)

// SvcCustomSubjectServer is the interface that providers of the service
// SvcCustomSubject should implement.
type SvcCustomSubjectServer interface {
	MtSimpleReply(ctx context.Context, req *StringArg) (resp *SimpleStringReply, err error)
	MtVoidReply(ctx context.Context, req *StringArg) (err error)
	MtStreamedReply(ctx context.Context, req *StringArg, pushRep func(*SimpleStringReply)) (err error)
	MtVoidReqStreamedReply(ctx context.Context, pushRep func(*SimpleStringReply)) (err error)
}
type UnimplementedSvcCustomSubjectServer struct{}

func (*UnimplementedSvcCustomSubjectServer) MtSimpleReply(ctx context.Context, req *StringArg) (*SimpleStringReply, error) {
	return nil, errors.New("method MtSimpleReply not implemented")
}
func (*UnimplementedSvcCustomSubjectServer) MtVoidReply(ctx context.Context, req *StringArg) error {
	return errors.New("method MtVoidReply not implemented")
}
func (*UnimplementedSvcCustomSubjectServer) MtStreamedReply(ctx context.Context, req *StringArg, pushRep func(*SimpleStringReply)) error {
	return errors.New("method MtStreamedReply not implemented")
}
func (*UnimplementedSvcCustomSubjectServer) MtVoidReqStreamedReply(ctx context.Context, pushRep func(*SimpleStringReply)) error {
	return errors.New("method MtVoidReqStreamedReply not implemented")
}

// SvcCustomSubjectHandler provides a NATS subscription handler that can serve a
// subscription using a given SvcCustomSubjectServer implementation.
type SvcCustomSubjectHandler struct {
	ctx     context.Context
	workers *nrpc.WorkerPool
	nc      nrpc.NatsConn
	server  SvcCustomSubjectServer

	encodings []string
}

func NewSvcCustomSubjectHandler(ctx context.Context, nc nrpc.NatsConn, s SvcCustomSubjectServer) *SvcCustomSubjectHandler {
	return &SvcCustomSubjectHandler{
		ctx:    ctx,
		nc:     nc,
		server: s,

		encodings: []string{"protobuf"},
	}
}

func NewSvcCustomSubjectConcurrentHandler(workers *nrpc.WorkerPool, nc nrpc.NatsConn, s SvcCustomSubjectServer) *SvcCustomSubjectHandler {
	return &SvcCustomSubjectHandler{
		workers: workers,
		nc:      nc,
		server:  s,
	}
}

// SetEncodings sets the output encodings when using a '*Publish' function
func (h *SvcCustomSubjectHandler) SetEncodings(encodings []string) {
	h.encodings = encodings
}

func (h *SvcCustomSubjectHandler) Subject() string {
	return "root.*.custom_subject.>"
}

func (h *SvcCustomSubjectHandler) MtNoRequestPublish(pkginstance string, msg SimpleStringReply) error {
	for _, encoding := range h.encodings {
		rawMsg, err := nrpc.Marshal(encoding, &msg)
		if err != nil {
			log.Printf("SvcCustomSubjectHandler.MtNoRequestPublish: error marshaling the message: %s", err)
			return err
		}
		subject := "root." + pkginstance + "." + "custom_subject." + "mtnorequest"
		if encoding != "protobuf" {
			subject += "." + encoding
		}
		if err := h.nc.Publish(subject, rawMsg); err != nil {
			return err
		}
	}
	return nil
}

func (h *SvcCustomSubjectHandler) Handler(msg *nats.Msg) {
	var ctx context.Context
	if h.workers != nil {
		ctx = h.workers.Context
	} else {
		ctx = h.ctx
	}
	request := nrpc.NewRequest(ctx, h.nc, msg.Subject, msg.Reply)
	// extract method name & encoding from subject
	pkgParams, _, name, tail, err := nrpc.ParseSubject(
		"root", 1, "custom_subject", 0, msg.Subject)
	if err != nil {
		log.Printf("SvcCustomSubjectHanlder: SvcCustomSubject subject parsing failed: %v", err)
		return
	}

	request.MethodName = name
	request.SubjectTail = tail
	request.SetPackageParam("instance", pkgParams[0])

	// call handler and form response
	var immediateError *nrpc.Error
	switch name {
	case "mt_simple_reply":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("MtSimpleReplyHanlder: MtSimpleReply subject parsing failed: %v", err)
			break
		}
		var req StringArg
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("MtSimpleReplyHandler: MtSimpleReply request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type:    nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context) (proto.Message, error) {
				innerResp, err := h.server.MtSimpleReply(ctx, &req)
				if err != nil {
					return nil, err
				}
				return innerResp, err
			}
		}
	case "mtvoidreply":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("MtVoidReplyHanlder: MtVoidReply subject parsing failed: %v", err)
			break
		}
		var req StringArg
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("MtVoidReplyHandler: MtVoidReply request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type:    nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context) (proto.Message, error) {
				var innerResp nrpc.Void
				err := h.server.MtVoidReply(ctx, &req)
				if err != nil {
					return nil, err
				}
				return &innerResp, err
			}
		}
	case "mtnorequest":
		// MtNoRequest is a no-request method. Ignore it.
		return
	case "mtstreamedreply":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("MtStreamedReplyHanlder: MtStreamedReply subject parsing failed: %v", err)
			break
		}
		var req StringArg
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("MtStreamedReplyHandler: MtStreamedReply request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type:    nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.EnableStreamedReply()
			request.Handler = func(ctx context.Context) (proto.Message, error) {
				err := h.server.MtStreamedReply(ctx, &req, func(rep *SimpleStringReply) {
					request.SendStreamReply(rep)
				})
				return nil, err
			}
		}
	case "mtvoidreqstreamedreply":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("MtVoidReqStreamedReplyHanlder: MtVoidReqStreamedReply subject parsing failed: %v", err)
			break
		}
		var req github_com_wyy_go_gogo_nrpc.Void
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("MtVoidReqStreamedReplyHandler: MtVoidReqStreamedReply request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type:    nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.EnableStreamedReply()
			request.Handler = func(ctx context.Context) (proto.Message, error) {
				err := h.server.MtVoidReqStreamedReply(ctx, func(rep *SimpleStringReply) {
					request.SendStreamReply(rep)
				})
				return nil, err
			}
		}
	default:
		log.Printf("SvcCustomSubjectHandler: unknown name %q", name)
		immediateError = &nrpc.Error{
			Type:    nrpc.Error_CLIENT,
			Message: "unknown name: " + name,
		}
	}
	if immediateError == nil {
		if h.workers != nil {
			// Try queuing the request
			if err := h.workers.QueueRequest(request); err != nil {
				log.Printf("nrpc: Error queuing the request: %s", err)
			}
		} else {
			// Run the handler synchronously
			request.RunAndReply()
		}
	}

	if immediateError != nil {
		if err := request.SendReply(nil, immediateError); err != nil {
			log.Printf("SvcCustomSubjectHandler: SvcCustomSubject handler failed to publish the response: %s", err)
		}
	} else {
	}
}

type SvcCustomSubjectClient struct {
	nc               nrpc.NatsConn
	PkgSubject       string
	PkgParaminstance string
	Subject          string
	Encoding         string
	Timeout          time.Duration
}

func NewSvcCustomSubjectClient(nc nrpc.NatsConn, pkgParaminstance string) *SvcCustomSubjectClient {
	return &SvcCustomSubjectClient{
		nc:               nc,
		PkgSubject:       "root",
		PkgParaminstance: pkgParaminstance,
		Subject:          "custom_subject",
		Encoding:         "protobuf",
		Timeout:          5 * time.Second,
	}
}

func (c *SvcCustomSubjectClient) MtSimpleReply(req StringArg) (resp SimpleStringReply, err error) {

	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + "mt_simple_reply"

	// call
	err = nrpc.Call(&req, &resp, c.nc, subject, c.Encoding, c.Timeout)
	if err != nil {
		return // already logged
	}

	return
}

func (c *SvcCustomSubjectClient) MtVoidReply(req StringArg) (err error) {

	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + "mtvoidreply"

	// call
	var resp github_com_wyy_go_gogo_nrpc.Void
	err = nrpc.Call(&req, &resp, c.nc, subject, c.Encoding, c.Timeout)
	if err != nil {
		return // already logged
	}

	return
}

func (c *SvcCustomSubjectClient) MtNoRequestSubject() string {
	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + "mtnorequest"
	if c.Encoding != "protobuf" {
		subject += "." + c.Encoding
	}
	return subject
}

type SvcCustomSubjectMtNoRequestSubscription struct {
	*nats.Subscription

	encoding string
}

func (s *SvcCustomSubjectMtNoRequestSubscription) Next(timeout time.Duration) (next SimpleStringReply, err error) {
	msg, err := s.Subscription.NextMsg(timeout)
	if err != nil {
		return
	}
	err = nrpc.Unmarshal(s.encoding, msg.Data, &next)
	return
}

func (c *SvcCustomSubjectClient) MtNoRequestSubscribeSync() (sub *SvcCustomSubjectMtNoRequestSubscription, err error) {
	subject := c.MtNoRequestSubject()
	natsSub, err := c.nc.SubscribeSync(subject)
	if err != nil {
		return
	}
	sub = &SvcCustomSubjectMtNoRequestSubscription{natsSub, c.Encoding}
	return
}

func (c *SvcCustomSubjectClient) MtNoRequestSubscribe(

	handler func(SimpleStringReply),
) (sub *nats.Subscription, err error) {
	subject := c.MtNoRequestSubject()
	sub, err = c.nc.Subscribe(subject, func(msg *nats.Msg) {
		var pmsg SimpleStringReply
		err := nrpc.Unmarshal(c.Encoding, msg.Data, &pmsg)
		if err != nil {
			log.Printf("SvcCustomSubjectClient.MtNoRequestSubscribe: Error decoding, %s", err)
			return
		}
		handler(pmsg)
	})
	return
}

func (c *SvcCustomSubjectClient) MtNoRequestSubscribeChan() (<-chan SimpleStringReply, *nats.Subscription, error) {
	ch := make(chan SimpleStringReply)
	sub, err := c.MtNoRequestSubscribe(func(msg SimpleStringReply) {
		ch <- msg
	})
	return ch, sub, err
}

func (c *SvcCustomSubjectClient) MtStreamedReply(
	ctx context.Context,
	req StringArg,
	cb func(context.Context, SimpleStringReply),
) error {
	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + "mtstreamedreply"

	sub, err := nrpc.StreamCall(ctx, c.nc, subject, &req, c.Encoding, c.Timeout)
	if err != nil {
		return err
	}

	var res SimpleStringReply
	for {
		err = sub.Next(&res)
		if err != nil {
			break
		}
		cb(ctx, res)
	}
	if err == nrpc.ErrEOS {
		err = nil
	}
	return err
}

func (c *SvcCustomSubjectClient) MtVoidReqStreamedReply(
	ctx context.Context,
	cb func(context.Context, SimpleStringReply),
) error {
	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + "mtvoidreqstreamedreply"

	sub, err := nrpc.StreamCall(ctx, c.nc, subject, &nrpc.Void{}, c.Encoding, c.Timeout)
	if err != nil {
		return err
	}

	var res SimpleStringReply
	for {
		err = sub.Next(&res)
		if err != nil {
			break
		}
		cb(ctx, res)
	}
	if err == nrpc.ErrEOS {
		err = nil
	}
	return err
}

// SvcSubjectParamsServer is the interface that providers of the service
// SvcSubjectParams should implement.
type SvcSubjectParamsServer interface {
	MtWithSubjectParams(ctx context.Context, mp1 string, mp2 string) (resp *SimpleStringReply, err error)
	MtStreamedReplyWithSubjectParams(ctx context.Context, mp1 string, mp2 string, pushRep func(*SimpleStringReply)) (err error)
	MtNoReply(ctx context.Context)
}
type UnimplementedSvcSubjectParamsServer struct{}

func (*UnimplementedSvcSubjectParamsServer) MtWithSubjectParams(ctx context.Context, mp1 string, mp2 string) (*SimpleStringReply, error) {
	return nil, errors.New("method MtWithSubjectParams not implemented")
}
func (*UnimplementedSvcSubjectParamsServer) MtStreamedReplyWithSubjectParams(ctx context.Context, mp1 string, mp2 string, pushRep func(*SimpleStringReply)) error {
	return errors.New("method MtStreamedReplyWithSubjectParams not implemented")
}
func (*UnimplementedSvcSubjectParamsServer) MtNoReply(ctx context.Context) {
}

// SvcSubjectParamsHandler provides a NATS subscription handler that can serve a
// subscription using a given SvcSubjectParamsServer implementation.
type SvcSubjectParamsHandler struct {
	ctx     context.Context
	workers *nrpc.WorkerPool
	nc      nrpc.NatsConn
	server  SvcSubjectParamsServer

	encodings []string
}

func NewSvcSubjectParamsHandler(ctx context.Context, nc nrpc.NatsConn, s SvcSubjectParamsServer) *SvcSubjectParamsHandler {
	return &SvcSubjectParamsHandler{
		ctx:    ctx,
		nc:     nc,
		server: s,

		encodings: []string{"protobuf"},
	}
}

func NewSvcSubjectParamsConcurrentHandler(workers *nrpc.WorkerPool, nc nrpc.NatsConn, s SvcSubjectParamsServer) *SvcSubjectParamsHandler {
	return &SvcSubjectParamsHandler{
		workers: workers,
		nc:      nc,
		server:  s,
	}
}

// SetEncodings sets the output encodings when using a '*Publish' function
func (h *SvcSubjectParamsHandler) SetEncodings(encodings []string) {
	h.encodings = encodings
}

func (h *SvcSubjectParamsHandler) Subject() string {
	return "root.*.svcsubjectparams.*.>"
}

func (h *SvcSubjectParamsHandler) MtNoRequestWParamsPublish(pkginstance string, svcclientid string, mtmp1 string, msg SimpleStringReply) error {
	for _, encoding := range h.encodings {
		rawMsg, err := nrpc.Marshal(encoding, &msg)
		if err != nil {
			log.Printf("SvcSubjectParamsHandler.MtNoRequestWParamsPublish: error marshaling the message: %s", err)
			return err
		}
		subject := "root." + pkginstance + "." + "svcsubjectparams." + svcclientid + "." + "mtnorequestwparams" + "." + mtmp1
		if encoding != "protobuf" {
			subject += "." + encoding
		}
		if err := h.nc.Publish(subject, rawMsg); err != nil {
			return err
		}
	}
	return nil
}

func (h *SvcSubjectParamsHandler) Handler(msg *nats.Msg) {
	var ctx context.Context
	if h.workers != nil {
		ctx = h.workers.Context
	} else {
		ctx = h.ctx
	}
	request := nrpc.NewRequest(ctx, h.nc, msg.Subject, msg.Reply)
	// extract method name & encoding from subject
	pkgParams, svcParams, name, tail, err := nrpc.ParseSubject(
		"root", 1, "svcsubjectparams", 1, msg.Subject)
	if err != nil {
		log.Printf("SvcSubjectParamsHanlder: SvcSubjectParams subject parsing failed: %v", err)
		return
	}

	request.MethodName = name
	request.SubjectTail = tail
	request.SetPackageParam("instance", pkgParams[0])
	request.SetServiceParam("clientid", svcParams[0])

	// call handler and form response
	var immediateError *nrpc.Error
	switch name {
	case "mtwithsubjectparams":
		var mtParams []string
		mtParams, request.Encoding, err = nrpc.ParseSubjectTail(2, request.SubjectTail)
		if err != nil {
			log.Printf("MtWithSubjectParamsHanlder: MtWithSubjectParams subject parsing failed: %v", err)
			break
		}
		var req github_com_wyy_go_gogo_nrpc.Void
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("MtWithSubjectParamsHandler: MtWithSubjectParams request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type:    nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context) (proto.Message, error) {
				innerResp, err := h.server.MtWithSubjectParams(ctx, mtParams[0], mtParams[1])
				if err != nil {
					return nil, err
				}
				return innerResp, err
			}
		}
	case "mtstreamedreplywithsubjectparams":
		var mtParams []string
		mtParams, request.Encoding, err = nrpc.ParseSubjectTail(2, request.SubjectTail)
		if err != nil {
			log.Printf("MtStreamedReplyWithSubjectParamsHanlder: MtStreamedReplyWithSubjectParams subject parsing failed: %v", err)
			break
		}
		var req github_com_wyy_go_gogo_nrpc.Void
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("MtStreamedReplyWithSubjectParamsHandler: MtStreamedReplyWithSubjectParams request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type:    nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.EnableStreamedReply()
			request.Handler = func(ctx context.Context) (proto.Message, error) {
				err := h.server.MtStreamedReplyWithSubjectParams(ctx, mtParams[0], mtParams[1], func(rep *SimpleStringReply) {
					request.SendStreamReply(rep)
				})
				return nil, err
			}
		}
	case "mtnoreply":
		request.NoReply = true
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("MtNoReplyHanlder: MtNoReply subject parsing failed: %v", err)
			break
		}
		var req github_com_wyy_go_gogo_nrpc.Void
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, &req); err != nil {
			log.Printf("MtNoReplyHandler: MtNoReply request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type:    nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context) (proto.Message, error) {
				var innerResp nrpc.NoReply
				h.server.MtNoReply(ctx)
				if err != nil {
					return nil, err
				}
				return &innerResp, err
			}
		}
	case "mtnorequestwparams":
		// MtNoRequestWParams is a no-request method. Ignore it.
		return
	default:
		log.Printf("SvcSubjectParamsHandler: unknown name %q", name)
		immediateError = &nrpc.Error{
			Type:    nrpc.Error_CLIENT,
			Message: "unknown name: " + name,
		}
	}
	if immediateError == nil {
		if h.workers != nil {
			// Try queuing the request
			if err := h.workers.QueueRequest(request); err != nil {
				log.Printf("nrpc: Error queuing the request: %s", err)
			}
		} else {
			// Run the handler synchronously
			request.RunAndReply()
		}
	}

	if immediateError != nil {
		if err := request.SendReply(nil, immediateError); err != nil {
			log.Printf("SvcSubjectParamsHandler: SvcSubjectParams handler failed to publish the response: %s", err)
		}
	} else {
	}
}

type SvcSubjectParamsClient struct {
	nc               nrpc.NatsConn
	PkgSubject       string
	PkgParaminstance string
	Subject          string
	SvcParamclientid string
	Encoding         string
	Timeout          time.Duration
}

func NewSvcSubjectParamsClient(nc nrpc.NatsConn, pkgParaminstance string, svcParamclientid string) *SvcSubjectParamsClient {
	return &SvcSubjectParamsClient{
		nc:               nc,
		PkgSubject:       "root",
		PkgParaminstance: pkgParaminstance,
		Subject:          "svcsubjectparams",
		SvcParamclientid: svcParamclientid,
		Encoding:         "protobuf",
		Timeout:          5 * time.Second,
	}
}

func (c *SvcSubjectParamsClient) MtWithSubjectParams(mp1 string, mp2 string) (resp SimpleStringReply, err error) {

	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + c.SvcParamclientid + "." + "mtwithsubjectparams" + "." + mp1 + "." + mp2

	// call
	var req github_com_wyy_go_gogo_nrpc.Void
	err = nrpc.Call(&req, &resp, c.nc, subject, c.Encoding, c.Timeout)
	if err != nil {
		return // already logged
	}

	return
}

func (c *SvcSubjectParamsClient) MtStreamedReplyWithSubjectParams(
	ctx context.Context, mp1 string, mp2 string,
	cb func(context.Context, SimpleStringReply),
) error {
	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + c.SvcParamclientid + "." + "mtstreamedreplywithsubjectparams" + "." + mp1 + "." + mp2

	sub, err := nrpc.StreamCall(ctx, c.nc, subject, &nrpc.Void{}, c.Encoding, c.Timeout)
	if err != nil {
		return err
	}

	var res SimpleStringReply
	for {
		err = sub.Next(&res)
		if err != nil {
			break
		}
		cb(ctx, res)
	}
	if err == nrpc.ErrEOS {
		err = nil
	}
	return err
}

func (c *SvcSubjectParamsClient) MtNoReply() (err error) {

	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + c.SvcParamclientid + "." + "mtnoreply"

	// call
	var req github_com_wyy_go_gogo_nrpc.Void
	var resp github_com_wyy_go_gogo_nrpc.NoReply
	err = nrpc.Call(&req, &resp, c.nc, subject, c.Encoding, c.Timeout)
	if err != nil {
		return // already logged
	}

	return
}

func (c *SvcSubjectParamsClient) MtNoRequestWParamsSubject(
	mtmp1 string,
) string {
	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + c.SvcParamclientid + "." + "mtnorequestwparams" + "." + mtmp1
	if c.Encoding != "protobuf" {
		subject += "." + c.Encoding
	}
	return subject
}

type SvcSubjectParamsMtNoRequestWParamsSubscription struct {
	*nats.Subscription

	encoding string
}

func (s *SvcSubjectParamsMtNoRequestWParamsSubscription) Next(timeout time.Duration) (next SimpleStringReply, err error) {
	msg, err := s.Subscription.NextMsg(timeout)
	if err != nil {
		return
	}
	err = nrpc.Unmarshal(s.encoding, msg.Data, &next)
	return
}

func (c *SvcSubjectParamsClient) MtNoRequestWParamsSubscribeSync(
	mtmp1 string,
) (sub *SvcSubjectParamsMtNoRequestWParamsSubscription, err error) {
	subject := c.MtNoRequestWParamsSubject(
		mtmp1,
	)
	natsSub, err := c.nc.SubscribeSync(subject)
	if err != nil {
		return
	}
	sub = &SvcSubjectParamsMtNoRequestWParamsSubscription{natsSub, c.Encoding}
	return
}

func (c *SvcSubjectParamsClient) MtNoRequestWParamsSubscribe(
	mtmp1 string,
	handler func(SimpleStringReply),
) (sub *nats.Subscription, err error) {
	subject := c.MtNoRequestWParamsSubject(
		mtmp1,
	)
	sub, err = c.nc.Subscribe(subject, func(msg *nats.Msg) {
		var pmsg SimpleStringReply
		err := nrpc.Unmarshal(c.Encoding, msg.Data, &pmsg)
		if err != nil {
			log.Printf("SvcSubjectParamsClient.MtNoRequestWParamsSubscribe: Error decoding, %s", err)
			return
		}
		handler(pmsg)
	})
	return
}

func (c *SvcSubjectParamsClient) MtNoRequestWParamsSubscribeChan(
	mtmp1 string,
) (<-chan SimpleStringReply, *nats.Subscription, error) {
	ch := make(chan SimpleStringReply)
	sub, err := c.MtNoRequestWParamsSubscribe(mtmp1, func(msg SimpleStringReply) {
		ch <- msg
	})
	return ch, sub, err
}

// NoRequestServiceServer is the interface that providers of the service
// NoRequestService should implement.
type NoRequestServiceServer interface {
}
type UnimplementedNoRequestServiceServer struct{}

// NoRequestServiceHandler provides a NATS subscription handler that can serve a
// subscription using a given NoRequestServiceServer implementation.
type NoRequestServiceHandler struct {
	ctx     context.Context
	workers *nrpc.WorkerPool
	nc      nrpc.NatsConn
	server  NoRequestServiceServer

	encodings []string
}

func NewNoRequestServiceHandler(ctx context.Context, nc nrpc.NatsConn, s NoRequestServiceServer) *NoRequestServiceHandler {
	return &NoRequestServiceHandler{
		ctx:    ctx,
		nc:     nc,
		server: s,

		encodings: []string{"protobuf"},
	}
}

func NewNoRequestServiceConcurrentHandler(workers *nrpc.WorkerPool, nc nrpc.NatsConn, s NoRequestServiceServer) *NoRequestServiceHandler {
	return &NoRequestServiceHandler{
		workers: workers,
		nc:      nc,
		server:  s,
	}
}

// SetEncodings sets the output encodings when using a '*Publish' function
func (h *NoRequestServiceHandler) SetEncodings(encodings []string) {
	h.encodings = encodings
}

func (h *NoRequestServiceHandler) Subject() string {
	return "root.*.norequestservice.>"
}

func (h *NoRequestServiceHandler) MtNoRequestPublish(pkginstance string, msg SimpleStringReply) error {
	for _, encoding := range h.encodings {
		rawMsg, err := nrpc.Marshal(encoding, &msg)
		if err != nil {
			log.Printf("NoRequestServiceHandler.MtNoRequestPublish: error marshaling the message: %s", err)
			return err
		}
		subject := "root." + pkginstance + "." + "norequestservice." + "mtnorequest"
		if encoding != "protobuf" {
			subject += "." + encoding
		}
		if err := h.nc.Publish(subject, rawMsg); err != nil {
			return err
		}
	}
	return nil
}

type NoRequestServiceClient struct {
	nc               nrpc.NatsConn
	PkgSubject       string
	PkgParaminstance string
	Subject          string
	Encoding         string
	Timeout          time.Duration
}

func NewNoRequestServiceClient(nc nrpc.NatsConn, pkgParaminstance string) *NoRequestServiceClient {
	return &NoRequestServiceClient{
		nc:               nc,
		PkgSubject:       "root",
		PkgParaminstance: pkgParaminstance,
		Subject:          "norequestservice",
		Encoding:         "protobuf",
		Timeout:          5 * time.Second,
	}
}

func (c *NoRequestServiceClient) MtNoRequestSubject() string {
	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + "mtnorequest"
	if c.Encoding != "protobuf" {
		subject += "." + c.Encoding
	}
	return subject
}

type NoRequestServiceMtNoRequestSubscription struct {
	*nats.Subscription

	encoding string
}

func (s *NoRequestServiceMtNoRequestSubscription) Next(timeout time.Duration) (next SimpleStringReply, err error) {
	msg, err := s.Subscription.NextMsg(timeout)
	if err != nil {
		return
	}
	err = nrpc.Unmarshal(s.encoding, msg.Data, &next)
	return
}

func (c *NoRequestServiceClient) MtNoRequestSubscribeSync() (sub *NoRequestServiceMtNoRequestSubscription, err error) {
	subject := c.MtNoRequestSubject()
	natsSub, err := c.nc.SubscribeSync(subject)
	if err != nil {
		return
	}
	sub = &NoRequestServiceMtNoRequestSubscription{natsSub, c.Encoding}
	return
}

func (c *NoRequestServiceClient) MtNoRequestSubscribe(

	handler func(SimpleStringReply),
) (sub *nats.Subscription, err error) {
	subject := c.MtNoRequestSubject()
	sub, err = c.nc.Subscribe(subject, func(msg *nats.Msg) {
		var pmsg SimpleStringReply
		err := nrpc.Unmarshal(c.Encoding, msg.Data, &pmsg)
		if err != nil {
			log.Printf("NoRequestServiceClient.MtNoRequestSubscribe: Error decoding, %s", err)
			return
		}
		handler(pmsg)
	})
	return
}

func (c *NoRequestServiceClient) MtNoRequestSubscribeChan() (<-chan SimpleStringReply, *nats.Subscription, error) {
	ch := make(chan SimpleStringReply)
	sub, err := c.MtNoRequestSubscribe(func(msg SimpleStringReply) {
		ch <- msg
	})
	return ch, sub, err
}

type Client struct {
	nc               nrpc.NatsConn
	defaultEncoding  string
	defaultTimeout   time.Duration
	pkgSubject       string
	pkgParaminstance string
	SvcCustomSubject *SvcCustomSubjectClient
	SvcSubjectParams *SvcSubjectParamsClient
	NoRequestService *NoRequestServiceClient
}

func NewClient(nc nrpc.NatsConn, pkgParaminstance string) *Client {
	c := Client{
		nc:               nc,
		defaultEncoding:  "protobuf",
		defaultTimeout:   5 * time.Second,
		pkgSubject:       "root",
		pkgParaminstance: pkgParaminstance,
	}
	c.SvcCustomSubject = NewSvcCustomSubjectClient(nc, c.pkgParaminstance)
	c.NoRequestService = NewNoRequestServiceClient(nc, c.pkgParaminstance)
	return &c
}

func (c *Client) SetEncoding(encoding string) {
	c.defaultEncoding = encoding
	if c.SvcCustomSubject != nil {
		c.SvcCustomSubject.Encoding = encoding
	}
	if c.SvcSubjectParams != nil {
		c.SvcSubjectParams.Encoding = encoding
	}
	if c.NoRequestService != nil {
		c.NoRequestService.Encoding = encoding
	}
}

func (c *Client) SetTimeout(t time.Duration) {
	c.defaultTimeout = t
	if c.SvcCustomSubject != nil {
		c.SvcCustomSubject.Timeout = t
	}
	if c.SvcSubjectParams != nil {
		c.SvcSubjectParams.Timeout = t
	}
	if c.NoRequestService != nil {
		c.NoRequestService.Timeout = t
	}
}

func (c *Client) SetSvcSubjectParamsParams(
	clientid string,
) {
	c.SvcSubjectParams = NewSvcSubjectParamsClient(
		c.nc,
		c.pkgParaminstance,
		clientid,
	)
	c.SvcSubjectParams.Encoding = c.defaultEncoding
	c.SvcSubjectParams.Timeout = c.defaultTimeout
}

func (c *Client) NewSvcSubjectParams(
	clientid string,
) *SvcSubjectParamsClient {
	client := NewSvcSubjectParamsClient(
		c.nc,
		c.pkgParaminstance,
		clientid,
	)
	client.Encoding = c.defaultEncoding
	client.Timeout = c.defaultTimeout
	return client
}
