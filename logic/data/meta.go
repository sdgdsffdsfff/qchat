// Code generated by tool/convert
// DO NOT EDIT!
package data

type container struct {
	s map[string]interface{}
}

func (c *container) Update(name string, st interface{}) error {
	switch name {

	case "Coordinator":
		tmp := st.(*Coordinator)
		UpdateCoordinator(tmp)

	case "Gateway":
		tmp := st.(*Gateway)
		UpdateGateway(tmp)

	case "Global":
		tmp := st.(*Global)
		UpdateGlobal(tmp)

	case "Saver":
		tmp := st.(*Saver)
		UpdateSaver(tmp)

	case "Session":
		tmp := st.(*Session)
		UpdateSession(tmp)

	case "Agent":
		tmp := st.(*Agent)
		UpdateAgent(tmp)

	case "Dispatcher":
		tmp := st.(*Dispatcher)
		UpdateDispatcher(tmp)

	case "Router":
		tmp := st.(*Router)
		UpdateRouter(tmp)

	case "Center":
		tmp := st.(*Center)
		UpdateCenter(tmp)

	default:
		panic("unsupported struct " + name)
	}
	return nil
}

func (c *container) GetStructs() map[string]interface{} {
	return c.s
}

var ObjectsContainer = &container{
	s: map[string]interface{}{

		"Coordinator": Coordinator{},

		"Gateway": Gateway{},

		"Global": Global{},

		"Saver": Saver{},

		"Session": Session{},

		"Agent": Agent{},

		"Dispatcher": Dispatcher{},

		"Router": Router{},

		"Center": Center{},
	},
}