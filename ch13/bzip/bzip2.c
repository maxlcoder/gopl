/* 文件是 bzip2.c */
/* 对 libbzip2 的简单包装适合 cgo */
#include <bzlib.h>

int bz2compress(bz_stream *s, int action,
                char *in, unsigned *inlen, char *out, unsigned *outlen) {
  s->next_in = in;
  s->avail_in = *inlen;
  s->next_out = out;
  s->avail_out = *outlen;
  int r = BZ2_bzCompress(s, action);
  *inlen -= s->avail_in;
  *outlen -= s->avail_out;
  s->next_in = s->next_out = NULL;
  return r;
}
