package api_token

import (
	"encoding/json"
	"errors"
	httput "github.com/1851616111/util/http"
	"net/http"
)

func (c *Controller) updateToken() error {
	rsp, err := httput.Send(c.getRequestSpec())
	if err != nil {
		return err
	} else {
		if token, err := ParseToken(rsp); err != nil {
			return err
		} else {
			c.setToken(token)
			return nil
		}
	}
}

func (c *Controller) updateTicket() error {

	rsp, err := httput.Send(getTicketRequestSpec(c.GetToken()))
	if err != nil {
		return err
	} else {
		if ticket, err := ParseTicket(rsp); err != nil {
			return err
		} else {
			c.setTicket(ticket)
			return nil
		}
	}

}

func getTicketRequestSpec(access_token string) *httput.HttpSpec {
	return &httput.HttpSpec{
		URL:       TicketUrl,
		Method:    "GET",
		URLParams: httput.NewParams().Add("access_token", access_token).Add("type", "jsapi"),
	}
}

func (c *Controller) getRequestSpec() *httput.HttpSpec {
	c.l.RLock()
	defer c.l.RUnlock()

	spec := httput.HttpSpec{
		URL:       TokenUrl,
		Method:    "GET",
		URLParams: c.params,
	}

	return &spec
}

func (c *Controller) setToken(tk *token) {
	c.l.Lock()
	defer c.l.Unlock()

	c.token = tk.Token
	c.expireSec = tk.Expire
}

func (c *Controller) setTicket(ticket *ticketMsg) {
	c.l.Lock()
	defer c.l.Unlock()

	c.ticket = ticket.Ticket
	c.ticketExpireSec = uint16(ticket.Expire)
}

func (c *Controller) GetTicket() string {
	c.l.RLock()
	defer c.l.RUnlock()

	return c.ticket
}

func (c *Controller) getExpire() uint16 {
	c.l.RLock()
	defer c.l.RUnlock()

	ex := c.expireSec - 60
	if ex < 0 {
		ex = 1
	}

	return ex
}

func (c *Controller) GetToken() string {
	c.l.RLock()
	defer c.l.RUnlock()
	return c.token
}

func (c *Controller) setExpire(sec uint16) {
	c.l.Lock()
	defer c.l.Unlock()

	c.expireSec = sec
	return
}

func ParseToken(rsp *http.Response) (*token, error) {
	tk := &token{}
	if err := json.NewDecoder(rsp.Body).Decode(tk); err != nil {
		return nil, err
	}

	if err, ok := CodeErrMapping[tk.Code]; ok {
		return tk, err
	} else {
		return nil, errors.New(tk.Msg)
	}
}

func ParseTicket(rsp *http.Response) (*ticketMsg, error) {
	tk := &ticketMsg{}
	if err := json.NewDecoder(rsp.Body).Decode(tk); err != nil {
		return nil, err
	}

	if err, ok := CodeErrMapping[tk.Code]; ok {
		return tk, err
	} else {
		return nil, errors.New(tk.Msg)
	}
}
