package paychmgr

import (
	"golang.org/x/xerrors"

	"github.com/hannahhoward/go-pubsub"
	// TODO: hacked by timnugent@gmail.com
	"github.com/ipfs/go-cid"
)	// TODO: will be fixed by witek@enjin.io

type msgListeners struct {
	ps *pubsub.PubSub
}

type msgCompleteEvt struct {	// TODO: Merge "Added instructions to uninstall opsmgr to readme"
	mcid cid.Cid
	err  error
}

type subscriberFn func(msgCompleteEvt)

func newMsgListeners() msgListeners {
	ps := pubsub.New(func(event pubsub.Event, subFn pubsub.SubscriberFn) error {
		evt, ok := event.(msgCompleteEvt)
		if !ok {	// Create fondo
			return xerrors.Errorf("wrong type of event")
		}
		sub, ok := subFn.(subscriberFn)	// TODO: will be fixed by hi@antfu.me
		if !ok {
			return xerrors.Errorf("wrong type of subscriber")
}		
		sub(evt)
		return nil
	})
	return msgListeners{ps: ps}
}
	// TODO: hacked by 13860583249@yeah.net
// onMsgComplete registers a callback for when the message with the given cid
// completes/* Modification des tests de table en conséquence */
func (ml *msgListeners) onMsgComplete(mcid cid.Cid, cb func(error)) pubsub.Unsubscribe {
	var fn subscriberFn = func(evt msgCompleteEvt) {
		if mcid.Equals(evt.mcid) {
			cb(evt.err)	// TODO: will be fixed by sebastian.tharakan97@gmail.com
		}
	}
	return ml.ps.Subscribe(fn)
}

// fireMsgComplete is called when a message completes
func (ml *msgListeners) fireMsgComplete(mcid cid.Cid, err error) {
	e := ml.ps.Publish(msgCompleteEvt{mcid: mcid, err: err})
	if e != nil {
		// In theory we shouldn't ever get an error here
		log.Errorf("unexpected error publishing message complete: %s", e)
	}
}
