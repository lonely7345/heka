/***** BEGIN LICENSE BLOCK *****
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this file,
# You can obtain one at http://mozilla.org/MPL/2.0/.
#
# The Initial Developer of the Original Code is the Mozilla Foundation.
# Portions created by the Initial Developer are Copyright (C) 2012
# the Initial Developer. All Rights Reserved.
#
# Contributor(s):
#   Rob Miller (rmiller@mozilla.com)
#
# ***** END LICENSE BLOCK *****/
package hekagrater

import (
	"github.com/bitly/go-simplejson"
	"log"
	"time"
)

type Decoder interface {
	Decode(msgBytes *[]byte) *Message
}

type JsonDecoder struct {
}

func (self *JsonDecoder) Decode(msgBytes *[]byte) *Message {
	var msg Message
	msgJson, err := simplejson.NewJson(*msgBytes)
	if err != nil {
		log.Printf("Error decoding message: %s\n", err.Error())
		return nil
	}

	msg.Type = msgJson.Get("type").MustString()
	timeStr := msgJson.Get("timestamp").MustString()
	msg.Timestamp, err = time.Parse(timeFormat, timeStr)
	if err != nil {
		msg.Timestamp, err = time.Parse(timeFormatFullSecond, timeStr)
		if err != nil {
			log.Printf("Timestamp parsing error: %s\n", err.Error())
		}
	}
	msg.Logger = msgJson.Get("logger").MustString()
	msg.Severity = msgJson.Get("severity").MustInt()
	msg.Payload, _ = msgJson.Get("payload").String()
	msg.Fields = msgJson.Get("fields")
	msg.Env_version = msgJson.Get("env_version").MustString()
	msg.Pid, _ = msgJson.Get("metlog_pid").Int()
	msg.Hostname, _ = msgJson.Get("metlog_hostname").String()

	return &msg
}
