package main

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/JoaoRufino/canopus"
)

//#cgo LDFLAGS:-L/home/jplr/gitrepos/Golang-projects/asn/Final/libs -lit2s-asn-cam
//#cgo CFLAGS:-I/home/jplr/gitrepos/Golang-projects/asn/Final/libs
/*#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#include <stdint.h>
#include <errno.h>
#include <syslog.h>
#include <CAM.h>
#include <INTEGER.h>
#include <asn_application.h>

#define syslog_emerg(msg, ...) syslog(LOG_EMERG, "%s(%s:%d) [" msg "]", __FILE__, __func__, __LINE__, ##__VA_ARGS__)
#define syslog_err(msg, ...)   syslog(LOG_ERR  , "%s(%s:%d) [" msg "]", __FILE__, __func__, __LINE__, ##__VA_ARGS__)

#ifndef NDEBUG
#define syslog_debug(msg, ...) syslog(LOG_DEBUG, "%s(%s:%d) [" msg "]", __FILE__, __func__, __LINE__, ##__VA_ARGS__)
#else
#define syslog_debug(msg, ...)
#endif
static CAM_t *cam;

int decode(uint8_t *buffer,int file_size){
	asn_dec_rval_t dec;
	asn_codec_ctx_t *opt_codec_ctx = 0;
	cam = calloc(1, sizeof(CAM_t));
	if(cam == NULL) {
		syslog_emerg("calloc() failed: %m");
	}
	dec =  uper_decode_complete(opt_codec_ctx,&asn_DEF_CAM, (void **) &cam, buffer, file_size);
	switch(dec.code) {
		case RC_OK:
			xer_fprint(stdout, &asn_DEF_CAM, cam);
			return 0;
		case RC_FAIL:
			syslog_debug("Error decoding: RC_FAIL");
			xer_fprint(stdout, &asn_DEF_CAM, cam);
			return 0;
		case RC_WMORE:
			syslog_debug("ERROR decoding: RC_WMORE");
			xer_fprint(stdout, &asn_DEF_CAM, cam);
			return 0;
	}
}
uint32_t get_lat(void) {
	return (cam->cam.camParameters.basicContainer.referencePosition.latitude);
}

uint32_t get_lon(void) {
	return (cam->cam.camParameters.basicContainer.referencePosition.longitude);
}

*/
import "C"

func main() {
	type CAM struct {
		cam *C.CAM_t
	}

	/*	var err error
		var buffer []byte
		buffer, err = ioutil.ReadFile("./it2s-CAM_Tx.code")*/

	server := canopus.NewServer()

	server.Post("/cam", func(req canopus.Request) canopus.Response {
		msg := canopus.ContentMessage(req.GetMessage().GetMessageId(), canopus.MessageAcknowledgment)
		msg.SetStringPayload("DONE! ")
		res := canopus.NewResponse(msg, nil)

		// Save to file
		payload := req.GetMessage().GetPayload().GetBytes()
		log.Println("len", len(payload))
		C.decode((*C.uint8_t)(unsafe.Pointer(&payload[0])), (C.int)(len(payload))) //BEWARE VERY UNSAFE SIFILIS!!! AIDS!!!!
		log.Println(C.get_lat(), C.get_lon())

		changeVal := fmt.Sprint((uint32)(C.get_lat())) + "|" + fmt.Sprint((uint32)(C.get_lon()))
		server.NotifyChange("/cam/pos", changeVal, false)
		return res

	})
	server.Get("/cam/pos", func(req canopus.Request) canopus.Response {
		msg := canopus.NewMessageOfType(canopus.MessageAcknowledgment, req.GetMessage().GetMessageId(), canopus.NewPlainTextPayload("Acknowledged"))
		res := canopus.NewResponse(msg, nil)

		return res
	})

	server.OnBlockMessage(func(msg canopus.Message, inbound bool) {
		// log.Println("Incoming Block Message:")
		// canopus.PrintMessage(msg)
	})

	server.OnObserve(func(resource string, msg canopus.Message) {
		fmt.Println("[SERVER << ] Observe Requested for " + resource)
	})

	server.ListenAndServe(":5683")
	<-make(chan struct{})
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
