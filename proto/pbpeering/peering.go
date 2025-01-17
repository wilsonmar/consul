package pbpeering

import (
	"time"

	"github.com/hashicorp/consul/api"
)

// TODO(peering): These are byproducts of not embedding
// types in our protobuf definitions and are temporary;
// Hoping to replace them with 1 or 2 methods per request
// using https://github.com/hashicorp/consul/pull/12507

// RequestDatacenter implements structs.RPCInfo
func (req *GenerateTokenRequest) RequestDatacenter() string {
	return req.Datacenter
}

// IsRead implements structs.RPCInfo
func (req *GenerateTokenRequest) IsRead() bool {
	return false
}

// AllowStaleRead implements structs.RPCInfo
func (req *GenerateTokenRequest) AllowStaleRead() bool {
	return false
}

// TokenSecret implements structs.RPCInfo
func (req *GenerateTokenRequest) TokenSecret() string {
	return req.Token
}

// SetTokenSecret implements structs.RPCInfo
func (req *GenerateTokenRequest) SetTokenSecret(token string) {
	req.Token = token
}

// HasTimedOut implements structs.RPCInfo
func (req *GenerateTokenRequest) HasTimedOut(start time.Time, rpcHoldTimeout, _, _ time.Duration) (bool, error) {
	return time.Since(start) > rpcHoldTimeout, nil
}

// Timeout implements structs.RPCInfo
func (msg *GenerateTokenRequest) Timeout(rpcHoldTimeout time.Duration, maxQueryTime time.Duration, defaultQueryTime time.Duration) time.Duration {
	return rpcHoldTimeout
}

// IsRead implements structs.RPCInfo
func (req *InitiateRequest) IsRead() bool {
	return false
}

// AllowStaleRead implements structs.RPCInfo
func (req *InitiateRequest) AllowStaleRead() bool {
	return false
}

// TokenSecret implements structs.RPCInfo
func (req *InitiateRequest) TokenSecret() string {
	return req.Token
}

// SetTokenSecret implements structs.RPCInfo
func (req *InitiateRequest) SetTokenSecret(token string) {
	req.Token = token
}

// HasTimedOut implements structs.RPCInfo
func (req *InitiateRequest) HasTimedOut(start time.Time, rpcHoldTimeout, _, _ time.Duration) (bool, error) {
	return time.Since(start) > rpcHoldTimeout, nil
}

// Timeout implements structs.RPCInfo
func (msg *InitiateRequest) Timeout(rpcHoldTimeout time.Duration, maxQueryTime time.Duration, defaultQueryTime time.Duration) time.Duration {
	return rpcHoldTimeout
}

// ShouldDial returns true when the peering was stored via the peering initiation endpoint,
// AND the peering is not marked as terminated by our peer.
// If we generated a token for this peer we did not store our server addresses under PeerServerAddresses.
// These server addresses are for dialing, and only the peer initiating the peering will do the dialing.
func (p *Peering) ShouldDial() bool {
	return len(p.PeerServerAddresses) > 0 && p.State != PeeringState_TERMINATED
}

func (x ReplicationMessage_Response_Operation) GoString() string {
	return x.String()
}

// enumcover:PeeringState
func PeeringStateToAPI(s PeeringState) api.PeeringState {
	switch s {
	case PeeringState_INITIAL:
		return api.PeeringStateInitial
	case PeeringState_ACTIVE:
		return api.PeeringStateActive
	case PeeringState_FAILING:
		return api.PeeringStateFailing
	case PeeringState_TERMINATED:
		return api.PeeringStateTerminated
	case PeeringState_UNDEFINED:
		fallthrough
	default:
		return api.PeeringStateUndefined
	}
}

// enumcover:api.PeeringState
func PeeringStateFromAPI(t api.PeeringState) PeeringState {
	switch t {
	case api.PeeringStateInitial:
		return PeeringState_INITIAL
	case api.PeeringStateActive:
		return PeeringState_ACTIVE
	case api.PeeringStateFailing:
		return PeeringState_FAILING
	case api.PeeringStateTerminated:
		return PeeringState_TERMINATED
	case api.PeeringStateUndefined:
		fallthrough
	default:
		return PeeringState_UNDEFINED
	}
}

func (p *Peering) ToAPI() *api.Peering {
	var t api.Peering
	PeeringToAPI(p, &t)
	return &t
}

// TODO consider using mog for this
func (resp *PeeringListResponse) ToAPI() []*api.Peering {
	list := make([]*api.Peering, len(resp.Peerings))
	for i, p := range resp.Peerings {
		list[i] = p.ToAPI()
	}
	return list
}

// TODO consider using mog for this
func (resp *GenerateTokenResponse) ToAPI() *api.PeeringGenerateTokenResponse {
	var t api.PeeringGenerateTokenResponse
	GenerateTokenResponseToAPI(resp, &t)
	return &t
}

// TODO consider using mog for this
func (resp *InitiateResponse) ToAPI() *api.PeeringInitiateResponse {
	var t api.PeeringInitiateResponse
	InitiateResponseToAPI(resp, &t)
	return &t
}

// convenience
func NewGenerateTokenRequestFromAPI(req *api.PeeringGenerateTokenRequest) *GenerateTokenRequest {
	if req == nil {
		return nil
	}
	t := &GenerateTokenRequest{}
	GenerateTokenRequestFromAPI(req, t)
	return t
}

// convenience
func NewInitiateRequestFromAPI(req *api.PeeringInitiateRequest) *InitiateRequest {
	if req == nil {
		return nil
	}
	t := &InitiateRequest{}
	InitiateRequestFromAPI(req, t)
	return t
}
