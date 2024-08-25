#include <objc/runtime.h>
#include <objc/message.h>
#include <stdio.h>
/* A program to do stuff
 * What, I don't know.
 * That does nothing */

static void create_shared_application(void) {
  id nsa = objc_getClass("NSApplication");
  SEL sha = sel_registerName("sharedApplication"); // OK here
  // OK there
  id (*msg_send)(id, SEL) = (id (*)(id, SEL))objc_msgSend;
  msg_send(nsa, sha);
  printf("[macosx] hello.\r\n");
}
