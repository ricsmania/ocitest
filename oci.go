package main

/*
#cgo LDFLAGS: -loci -L./oracle
#cgo CFLAGS: -I./oracle
#include <oci.h>
#include <stdio.h>

CreateEnv() {

OCIEnv *env;
env = 0;

  OCIEnvCreate(
    &env,
    0,
    0,
    0,
    0,
    0,
    0,
    0);
  printf("worked\n");
}
*/
import "C"

func main() {
	C.CreateEnv()
}
