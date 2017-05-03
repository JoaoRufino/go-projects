package main

//#cgo LDFLAGS:-L. -lit2s-asn-cam
//#cgo CFLAGS:-I.
/*#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#include <stdbool.h>
#include <stdint.h>
#include <sys/stat.h>
#include <sys/mman.h>
#include <fcntl.h>
#include <errno.h>
#include <syslog.h>
#include <string.h>
#include <CAM.h>
#include <INTEGER.h>
#include <asn_application.h>
#include <NativeInteger.h>

#define syslog_emerg(msg, ...) syslog(LOG_EMERG, "%s(%s:%d) [" msg "]", __FILE__, __func__, __LINE__, ##__VA_ARGS__)
#define syslog_err(msg, ...)   syslog(LOG_ERR  , "%s(%s:%d) [" msg "]", __FILE__, __func__, __LINE__, ##__VA_ARGS__)

#ifndef NDEBUG
#define syslog_debug(msg, ...) syslog(LOG_DEBUG, "%s(%s:%d) [" msg "]", __FILE__, __func__, __LINE__, ##__VA_ARGS__)
#else
#define syslog_debug(msg, ...)
#endif

int decode(uint8_t *buffer,int file_size){
	CAM_t *cam;
	asn_dec_rval_t dec;
	asn_codec_ctx_t *opt_codec_ctx = 0;
	cam = calloc(1, sizeof(CAM_t));
	if(cam == NULL) {
		syslog_emerg("calloc() failed: %m");
		return -1;
	}
	dec =  uper_decode_complete(opt_codec_ctx,&asn_DEF_CAM, (void **) &cam, buffer, file_size);
	switch(dec.code) {
		case RC_OK:
			xer_fprint(stdout, &asn_DEF_CAM, cam);
			return 0;
		case RC_FAIL:
			syslog_debug("Error decoding: RC_FAIL");
			xer_fprint(stdout, &asn_DEF_CAM, cam);
			return -1;
		case RC_WMORE:
			syslog_debug("ERROR decoding: RC_WMORE");
			xer_fprint(stdout, &asn_DEF_CAM, cam);
			return -1;
	}
}*/
import "C"
import "io/ioutil"

func main() {
	buffer, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	C.decode(*buffer, *len(buffer))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*int main(int argc, char **argv){
	FILE *fp;
	uint8_t *buffer;
	int retval,code;
	if (argc != 2){
		printf("Pass file to process as argument\n");
		return -1;
	}

	fp = fopen(argv[1], "rb");
	if (fp == NULL){
		syslog_emerg("fdopen() failed: %m");
		return -1;
	}

	buffer = (uint8_t *)malloc(500);
	if (buffer == NULL) {
		syslog_err("malloc() failed: %m");
		return -1;
	}

	retval = fread(buffer, 1, 500, fp);
	if (retval <= 0){
		syslog_emerg("fread() failed: %m");
	}
	char xml;
	code=decode(buffer,&xml,retval);
	printf("%s\n", xml);
	return 0;

}*/
