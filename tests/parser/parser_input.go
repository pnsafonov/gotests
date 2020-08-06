package parser

import "github.com/golang/protobuf/ptypes/wrappers"

func SayHello() string {
	return "Hello!!!!"
}

var (
	id1 = 1
	id2 = 2
	someName = "some_name"
)

const (
	c1 = 1
	c2 = 3
	c3 = "abcdef"
)

var a1 = "a1"

func (r *GetAgentsRequest) Hello(msg1 string, id int) (int, int, string, error) {
	return 0, 0, "", nil
}

// some text
// a1

// before doc

// is this comment?
//
// line of docs
//
// end of docs
//
/*
strange comments
 end of comments */
type GetAgentsRequest struct {

	//id:44

	//AgentIds             []int32              `protobuf:"varint,1,rep,packed,name=agent_ids,json=agentIds,proto3" json:"agent_ids,omitempty"`
	//id:768
	AgentId              int32                `protobuf:"varint,1,rep,packed,name=agent_ids,json=agentIds,proto3" json:"agent_ids,omitempty"`

	//id:533
	//name string

	ParentAgentIds       []int32              `protobuf:"varint,2,rep,packed,name=parent_agent_ids,json=parentAgentIds,proto3" json:"parent_agent_ids,omitempty"`
	UserId               *wrappers.Int32Value `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

/*
Just a multi-line
comment
 */

type structA struct {
	a1 int
	a2 int
	v1 string
}
