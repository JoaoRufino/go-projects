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

typedef struct {
	char *buffer;
	int  buffer_size;
	int  buffer_filled;
}xer_buffer_t;

static CAM_t *cam;
static xer_buffer_t *xerbuffer;

void init_xer_buffer(xer_buffer_t* xer_buffer) {
    xer_buffer->buffer = malloc(1024);
    assert(xer_buffer->buffer != NULL);
    xer_buffer->buffer_size = 1024;
    xer_buffer->buffer_filled = 0;
}

void free_xer_buffer(xer_buffer_t* xer_buffer) {
    free(xer_buffer->buffer);
    xer_buffer->buffer_size = 0;
    xer_buffer->buffer_filled = 0;
}

static int xer_print2xerbuf_cb(const void *buffer, size_t size, void *app_key) {
    xer_buffer_t* xb = (xer_buffer_t*) app_key;
    while (xb->buffer_size - xb->buffer_filled <= size+1) {
        xb->buffer_size *= 2;
        xb->buffer_size += 1;
        xb->buffer = realloc(xb->buffer, xb->buffer_size);
        assert(xb->buffer != NULL);
    }
    memcpy(xb->buffer+xb->buffer_filled, buffer, size);
    xb->buffer_filled += size;
    *( xb->buffer + xb->buffer_filled ) = '\0';
}

int xer_encode_to_buffer(xer_buffer_t* xb, asn_TYPE_descriptor_t *td, void *sptr) {
    asn_enc_rval_t er;
    if (!td || !sptr) return -1;
    er = xer_encode(td, sptr, XER_F_BASIC, xer_print2xerbuf_cb, xb);
    if (er.encoded == -1) return -1;
    return 0;
}


int decode(uint8_t *buffer,int file_size){
	asn_dec_rval_t dec;
	int er;
	asn_codec_ctx_t *opt_codec_ctx = 0;
	cam = calloc(1, sizeof(CAM_t));
	if(cam == NULL) {
		syslog_emerg("calloc() failed: %m");
	}
	dec =  uper_decode_complete(opt_codec_ctx,&asn_DEF_CAM, (void **) &cam, buffer, file_size);
	printf("DONE!!!!\n");
	init_xer_buffer(xerbuffer);
	free_xer_buffer(xerbuffer);

	switch(dec.code) {
		case RC_OK:
			xer_fprint(stdout, &asn_DEF_CAM, cam);
			er = xer_encode_to_buffer(xerbuffer,&asn_DEF_CAM,cam);
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
	//var xml string , C.CString(xml)
	server.Get("/cam/json", func(req canopus.Request) canopus.Response {
		msg := canopus.NewMessageOfType(canopus.MessageAcknowledgment, req.GetMessage().GetMessageId(), canopus.NewPlainTextPayload("Acknowledged"))
		res := canopus.NewResponse(msg, nil)
		return res
	})
	server.Get("/denm/json", func(req canopus.Request) canopus.Response {
		msg := canopus.NewMessageOfType(canopus.MessageAcknowledgment, req.GetMessage().GetMessageId(), canopus.NewPlainTextPayload("Acknowledged"))
		res := canopus.NewResponse(msg, nil)
		return res
	})

	server.Post("/cam", func(req canopus.Request) canopus.Response {
		log.Println("cam")

		msg := canopus.ContentMessage(req.GetMessage().GetMessageId(), canopus.MessageAcknowledgment)
		msg.SetStringPayload("DONE! ")
		res := canopus.NewResponse(msg, nil)

		// Save to file
		payload := req.GetMessage().GetPayload().GetBytes()
		log.Println("len", len(payload))
		C.decode((*C.uint8_t)(unsafe.Pointer(&payload[0])), (C.int)(len(payload))) //BEWARE VERY UNSAFE SIFILIS!!! AIDS!!!!
		fmt.Println(C.get_lat(), C.get_lon())
		changeVal := fmt.Sprint((uint32)(C.get_lat())) + "|" + fmt.Sprint((uint32)(C.get_lon()))
		server.NotifyChange("/cam/pos", changeVal, false)
		return res

	})
	/*	server.Post("/denm", func(req canopus.Request) canopus.Response {
		log.Println("denm")
		var bxml *[]byte
		msg := canopus.ContentMessage(req.GetMessage().GetMessageId(), canopus.MessageAcknowledgment)
		msg.SetStringPayload("DONE! ")
		res := canopus.NewResponse(msg, nil)

		// Save to file
		payload := req.GetMessage().GetPayload().GetBytes()
		log.Println("len", len(payload))
		C.DENMdecode((*C.uint8_t)(unsafe.Pointer(&payload[0])), (C.int)(len(payload)), (*C.uint8_t)(unsafe.Pointer(&bxml))) //BEWARE VERY UNSAFE SIFILIS!!! AIDS!!!!
		log.Println(&bxml)

		changeVal := fmt.Sprint((uint32)(C.get_lat())) + "|" + fmt.Sprint((uint32)(C.get_lon()))
		server.NotifyChange("/cam/pos", changeVal, false)
		return res

	})*/

	server.OnMessage(func(msg canopus.Message, inbound bool) {
		//canopus.PrintMessage(msg)
	})

	server.OnObserve(func(resource string, msg canopus.Message) {
		fmt.Println("Observe Requested for " + resource)
	})
	server.OnBlockMessage(func(msg canopus.Message, inbound bool) {
		canopus.PrintMessage(msg)
	})

	server.ListenAndServe(":5683")
	<-make(chan struct{})
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
