!(function (t) {
  'function' == typeof define && define.amd ? define(t) : t();
})(function () {
  'use strict';
  var t =
      'undefined' != typeof globalThis
        ? globalThis
        : 'undefined' != typeof window
        ? window
        : 'undefined' != typeof global
        ? global
        : 'undefined' != typeof self
        ? self
        : {},
    e = { exports: {} };
  (function (t, e) {
    function r(t, e) {
      if (!t) throw new Error(e || 'Assertion failed');
    }
    function n(t, e) {
      t.super_ = e;
      var r = function () {};
      (r.prototype = e.prototype),
        (t.prototype = new r()),
        (t.prototype.constructor = t);
    }
    function i(t, e, r) {
      if (i.isBN(t)) return t;
      (this.negative = 0),
        (this.words = null),
        (this.length = 0),
        (this.red = null),
        null !== t &&
          (('le' !== e && 'be' !== e) || ((r = e), (e = 10)),
          this._init(t || 0, e || 10, r || 'be'));
    }
    var o;
    'object' == typeof t ? (t.exports = i) : (e.BN = i),
      (i.BN = i),
      (i.wordSize = 26);
    try {
      o =
        'undefined' != typeof window && void 0 !== window.Buffer
          ? window.Buffer
          : require('buffer').Buffer;
    } catch (t) {}
    function s(t, e) {
      var n = t.charCodeAt(e);
      return n >= 48 && n <= 57
        ? n - 48
        : n >= 65 && n <= 70
        ? n - 55
        : n >= 97 && n <= 102
        ? n - 87
        : void r(!1, 'Invalid character in ' + t);
    }
    function a(t, e, r) {
      var n = s(t, r);
      return r - 1 >= e && (n |= s(t, r - 1) << 4), n;
    }
    function u(t, e, n, i) {
      for (var o = 0, s = 0, a = Math.min(t.length, n), u = e; u < a; u++) {
        var h = t.charCodeAt(u) - 48;
        (o *= i),
          (s = h >= 49 ? h - 49 + 10 : h >= 17 ? h - 17 + 10 : h),
          r(h >= 0 && s < i, 'Invalid character'),
          (o += s);
      }
      return o;
    }
    function h(t, e) {
      (t.words = e.words),
        (t.length = e.length),
        (t.negative = e.negative),
        (t.red = e.red);
    }
    if (
      ((i.isBN = function (t) {
        return (
          t instanceof i ||
          (null !== t &&
            'object' == typeof t &&
            t.constructor.wordSize === i.wordSize &&
            Array.isArray(t.words))
        );
      }),
      (i.max = function (t, e) {
        return t.cmp(e) > 0 ? t : e;
      }),
      (i.min = function (t, e) {
        return t.cmp(e) < 0 ? t : e;
      }),
      (i.prototype._init = function (t, e, n) {
        if ('number' == typeof t) return this._initNumber(t, e, n);
        if ('object' == typeof t) return this._initArray(t, e, n);
        'hex' === e && (e = 16), r(e === (0 | e) && e >= 2 && e <= 36);
        var i = 0;
        '-' === (t = t.toString().replace(/\s+/g, ''))[0] &&
          (i++, (this.negative = 1)),
          i < t.length &&
            (16 === e
              ? this._parseHex(t, i, n)
              : (this._parseBase(t, e, i),
                'le' === n && this._initArray(this.toArray(), e, n)));
      }),
      (i.prototype._initNumber = function (t, e, n) {
        t < 0 && ((this.negative = 1), (t = -t)),
          t < 67108864
            ? ((this.words = [67108863 & t]), (this.length = 1))
            : t < 4503599627370496
            ? ((this.words = [67108863 & t, (t / 67108864) & 67108863]),
              (this.length = 2))
            : (r(t < 9007199254740992),
              (this.words = [67108863 & t, (t / 67108864) & 67108863, 1]),
              (this.length = 3)),
          'le' === n && this._initArray(this.toArray(), e, n);
      }),
      (i.prototype._initArray = function (t, e, n) {
        if ((r('number' == typeof t.length), t.length <= 0))
          return (this.words = [0]), (this.length = 1), this;
        (this.length = Math.ceil(t.length / 3)),
          (this.words = new Array(this.length));
        for (var i = 0; i < this.length; i++) this.words[i] = 0;
        var o,
          s,
          a = 0;
        if ('be' === n)
          for (i = t.length - 1, o = 0; i >= 0; i -= 3)
            (s = t[i] | (t[i - 1] << 8) | (t[i - 2] << 16)),
              (this.words[o] |= (s << a) & 67108863),
              (this.words[o + 1] = (s >>> (26 - a)) & 67108863),
              (a += 24) >= 26 && ((a -= 26), o++);
        else if ('le' === n)
          for (i = 0, o = 0; i < t.length; i += 3)
            (s = t[i] | (t[i + 1] << 8) | (t[i + 2] << 16)),
              (this.words[o] |= (s << a) & 67108863),
              (this.words[o + 1] = (s >>> (26 - a)) & 67108863),
              (a += 24) >= 26 && ((a -= 26), o++);
        return this._strip();
      }),
      (i.prototype._parseHex = function (t, e, r) {
        (this.length = Math.ceil((t.length - e) / 6)),
          (this.words = new Array(this.length));
        for (var n = 0; n < this.length; n++) this.words[n] = 0;
        var i,
          o = 0,
          s = 0;
        if ('be' === r)
          for (n = t.length - 1; n >= e; n -= 2)
            (i = a(t, e, n) << o),
              (this.words[s] |= 67108863 & i),
              o >= 18
                ? ((o -= 18), (s += 1), (this.words[s] |= i >>> 26))
                : (o += 8);
        else
          for (n = (t.length - e) % 2 == 0 ? e + 1 : e; n < t.length; n += 2)
            (i = a(t, e, n) << o),
              (this.words[s] |= 67108863 & i),
              o >= 18
                ? ((o -= 18), (s += 1), (this.words[s] |= i >>> 26))
                : (o += 8);
        this._strip();
      }),
      (i.prototype._parseBase = function (t, e, r) {
        (this.words = [0]), (this.length = 1);
        for (var n = 0, i = 1; i <= 67108863; i *= e) n++;
        n--, (i = (i / e) | 0);
        for (
          var o = t.length - r,
            s = o % n,
            a = Math.min(o, o - s) + r,
            h = 0,
            l = r;
          l < a;
          l += n
        )
          (h = u(t, l, l + n, e)),
            this.imuln(i),
            this.words[0] + h < 67108864
              ? (this.words[0] += h)
              : this._iaddn(h);
        if (0 !== s) {
          var c = 1;
          for (h = u(t, l, t.length, e), l = 0; l < s; l++) c *= e;
          this.imuln(c),
            this.words[0] + h < 67108864
              ? (this.words[0] += h)
              : this._iaddn(h);
        }
        this._strip();
      }),
      (i.prototype.copy = function (t) {
        t.words = new Array(this.length);
        for (var e = 0; e < this.length; e++) t.words[e] = this.words[e];
        (t.length = this.length),
          (t.negative = this.negative),
          (t.red = this.red);
      }),
      (i.prototype._move = function (t) {
        h(t, this);
      }),
      (i.prototype.clone = function () {
        var t = new i(null);
        return this.copy(t), t;
      }),
      (i.prototype._expand = function (t) {
        for (; this.length < t; ) this.words[this.length++] = 0;
        return this;
      }),
      (i.prototype._strip = function () {
        for (; this.length > 1 && 0 === this.words[this.length - 1]; )
          this.length--;
        return this._normSign();
      }),
      (i.prototype._normSign = function () {
        return (
          1 === this.length && 0 === this.words[0] && (this.negative = 0), this
        );
      }),
      'undefined' != typeof Symbol && 'function' == typeof Symbol.for)
    )
      try {
        i.prototype[Symbol.for('nodejs.util.inspect.custom')] = l;
      } catch (t) {
        i.prototype.inspect = l;
      }
    else i.prototype.inspect = l;
    function l() {
      return (this.red ? '<BN-R: ' : '<BN: ') + this.toString(16) + '>';
    }
    var c = [
        '',
        '0',
        '00',
        '000',
        '0000',
        '00000',
        '000000',
        '0000000',
        '00000000',
        '000000000',
        '0000000000',
        '00000000000',
        '000000000000',
        '0000000000000',
        '00000000000000',
        '000000000000000',
        '0000000000000000',
        '00000000000000000',
        '000000000000000000',
        '0000000000000000000',
        '00000000000000000000',
        '000000000000000000000',
        '0000000000000000000000',
        '00000000000000000000000',
        '000000000000000000000000',
        '0000000000000000000000000',
      ],
      f = [
        0, 0, 25, 16, 12, 11, 10, 9, 8, 8, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5,
        5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
      ],
      d = [
        0, 0, 33554432, 43046721, 16777216, 48828125, 60466176, 40353607,
        16777216, 43046721, 1e7, 19487171, 35831808, 62748517, 7529536,
        11390625, 16777216, 24137569, 34012224, 47045881, 64e6, 4084101,
        5153632, 6436343, 7962624, 9765625, 11881376, 14348907, 17210368,
        20511149, 243e5, 28629151, 33554432, 39135393, 45435424, 52521875,
        60466176,
      ];
    function m(t, e, r) {
      r.negative = e.negative ^ t.negative;
      var n = (t.length + e.length) | 0;
      (r.length = n), (n = (n - 1) | 0);
      var i = 0 | t.words[0],
        o = 0 | e.words[0],
        s = i * o,
        a = 67108863 & s,
        u = (s / 67108864) | 0;
      r.words[0] = a;
      for (var h = 1; h < n; h++) {
        for (
          var l = u >>> 26,
            c = 67108863 & u,
            f = Math.min(h, e.length - 1),
            d = Math.max(0, h - t.length + 1);
          d <= f;
          d++
        ) {
          var m = (h - d) | 0;
          (l +=
            ((s = (i = 0 | t.words[m]) * (o = 0 | e.words[d]) + c) / 67108864) |
            0),
            (c = 67108863 & s);
        }
        (r.words[h] = 0 | c), (u = 0 | l);
      }
      return 0 !== u ? (r.words[h] = 0 | u) : r.length--, r._strip();
    }
    (i.prototype.toString = function (t, e) {
      var n;
      if (((e = 0 | e || 1), 16 === (t = t || 10) || 'hex' === t)) {
        n = '';
        for (var i = 0, o = 0, s = 0; s < this.length; s++) {
          var a = this.words[s],
            u = (16777215 & ((a << i) | o)).toString(16);
          (o = (a >>> (24 - i)) & 16777215),
            (i += 2) >= 26 && ((i -= 26), s--),
            (n =
              0 !== o || s !== this.length - 1
                ? c[6 - u.length] + u + n
                : u + n);
        }
        for (0 !== o && (n = o.toString(16) + n); n.length % e != 0; )
          n = '0' + n;
        return 0 !== this.negative && (n = '-' + n), n;
      }
      if (t === (0 | t) && t >= 2 && t <= 36) {
        var h = f[t],
          l = d[t];
        n = '';
        var m = this.clone();
        for (m.negative = 0; !m.isZero(); ) {
          var p = m.modrn(l).toString(t);
          n = (m = m.idivn(l)).isZero() ? p + n : c[h - p.length] + p + n;
        }
        for (this.isZero() && (n = '0' + n); n.length % e != 0; ) n = '0' + n;
        return 0 !== this.negative && (n = '-' + n), n;
      }
      r(!1, 'Base should be between 2 and 36');
    }),
      (i.prototype.toNumber = function () {
        var t = this.words[0];
        return (
          2 === this.length
            ? (t += 67108864 * this.words[1])
            : 3 === this.length && 1 === this.words[2]
            ? (t += 4503599627370496 + 67108864 * this.words[1])
            : this.length > 2 &&
              r(!1, 'Number can only safely store up to 53 bits'),
          0 !== this.negative ? -t : t
        );
      }),
      (i.prototype.toJSON = function () {
        return this.toString(16, 2);
      }),
      o &&
        (i.prototype.toBuffer = function (t, e) {
          return this.toArrayLike(o, t, e);
        }),
      (i.prototype.toArray = function (t, e) {
        return this.toArrayLike(Array, t, e);
      }),
      (i.prototype.toArrayLike = function (t, e, n) {
        this._strip();
        var i = this.byteLength(),
          o = n || Math.max(1, i);
        r(i <= o, 'byte array longer than desired length'),
          r(o > 0, 'Requested array length <= 0');
        var s = (function (t, e) {
          return t.allocUnsafe ? t.allocUnsafe(e) : new t(e);
        })(t, o);
        return this['_toArrayLike' + ('le' === e ? 'LE' : 'BE')](s, i), s;
      }),
      (i.prototype._toArrayLikeLE = function (t, e) {
        for (var r = 0, n = 0, i = 0, o = 0; i < this.length; i++) {
          var s = (this.words[i] << o) | n;
          (t[r++] = 255 & s),
            r < t.length && (t[r++] = (s >> 8) & 255),
            r < t.length && (t[r++] = (s >> 16) & 255),
            6 === o
              ? (r < t.length && (t[r++] = (s >> 24) & 255), (n = 0), (o = 0))
              : ((n = s >>> 24), (o += 2));
        }
        if (r < t.length) for (t[r++] = n; r < t.length; ) t[r++] = 0;
      }),
      (i.prototype._toArrayLikeBE = function (t, e) {
        for (var r = t.length - 1, n = 0, i = 0, o = 0; i < this.length; i++) {
          var s = (this.words[i] << o) | n;
          (t[r--] = 255 & s),
            r >= 0 && (t[r--] = (s >> 8) & 255),
            r >= 0 && (t[r--] = (s >> 16) & 255),
            6 === o
              ? (r >= 0 && (t[r--] = (s >> 24) & 255), (n = 0), (o = 0))
              : ((n = s >>> 24), (o += 2));
        }
        if (r >= 0) for (t[r--] = n; r >= 0; ) t[r--] = 0;
      }),
      Math.clz32
        ? (i.prototype._countBits = function (t) {
            return 32 - Math.clz32(t);
          })
        : (i.prototype._countBits = function (t) {
            var e = t,
              r = 0;
            return (
              e >= 4096 && ((r += 13), (e >>>= 13)),
              e >= 64 && ((r += 7), (e >>>= 7)),
              e >= 8 && ((r += 4), (e >>>= 4)),
              e >= 2 && ((r += 2), (e >>>= 2)),
              r + e
            );
          }),
      (i.prototype._zeroBits = function (t) {
        if (0 === t) return 26;
        var e = t,
          r = 0;
        return (
          8191 & e || ((r += 13), (e >>>= 13)),
          127 & e || ((r += 7), (e >>>= 7)),
          15 & e || ((r += 4), (e >>>= 4)),
          3 & e || ((r += 2), (e >>>= 2)),
          1 & e || r++,
          r
        );
      }),
      (i.prototype.bitLength = function () {
        var t = this.words[this.length - 1],
          e = this._countBits(t);
        return 26 * (this.length - 1) + e;
      }),
      (i.prototype.zeroBits = function () {
        if (this.isZero()) return 0;
        for (var t = 0, e = 0; e < this.length; e++) {
          var r = this._zeroBits(this.words[e]);
          if (((t += r), 26 !== r)) break;
        }
        return t;
      }),
      (i.prototype.byteLength = function () {
        return Math.ceil(this.bitLength() / 8);
      }),
      (i.prototype.toTwos = function (t) {
        return 0 !== this.negative
          ? this.abs().inotn(t).iaddn(1)
          : this.clone();
      }),
      (i.prototype.fromTwos = function (t) {
        return this.testn(t - 1) ? this.notn(t).iaddn(1).ineg() : this.clone();
      }),
      (i.prototype.isNeg = function () {
        return 0 !== this.negative;
      }),
      (i.prototype.neg = function () {
        return this.clone().ineg();
      }),
      (i.prototype.ineg = function () {
        return this.isZero() || (this.negative ^= 1), this;
      }),
      (i.prototype.iuor = function (t) {
        for (; this.length < t.length; ) this.words[this.length++] = 0;
        for (var e = 0; e < t.length; e++)
          this.words[e] = this.words[e] | t.words[e];
        return this._strip();
      }),
      (i.prototype.ior = function (t) {
        return r(!(this.negative | t.negative)), this.iuor(t);
      }),
      (i.prototype.or = function (t) {
        return this.length > t.length
          ? this.clone().ior(t)
          : t.clone().ior(this);
      }),
      (i.prototype.uor = function (t) {
        return this.length > t.length
          ? this.clone().iuor(t)
          : t.clone().iuor(this);
      }),
      (i.prototype.iuand = function (t) {
        var e;
        e = this.length > t.length ? t : this;
        for (var r = 0; r < e.length; r++)
          this.words[r] = this.words[r] & t.words[r];
        return (this.length = e.length), this._strip();
      }),
      (i.prototype.iand = function (t) {
        return r(!(this.negative | t.negative)), this.iuand(t);
      }),
      (i.prototype.and = function (t) {
        return this.length > t.length
          ? this.clone().iand(t)
          : t.clone().iand(this);
      }),
      (i.prototype.uand = function (t) {
        return this.length > t.length
          ? this.clone().iuand(t)
          : t.clone().iuand(this);
      }),
      (i.prototype.iuxor = function (t) {
        var e, r;
        this.length > t.length ? ((e = this), (r = t)) : ((e = t), (r = this));
        for (var n = 0; n < r.length; n++)
          this.words[n] = e.words[n] ^ r.words[n];
        if (this !== e) for (; n < e.length; n++) this.words[n] = e.words[n];
        return (this.length = e.length), this._strip();
      }),
      (i.prototype.ixor = function (t) {
        return r(!(this.negative | t.negative)), this.iuxor(t);
      }),
      (i.prototype.xor = function (t) {
        return this.length > t.length
          ? this.clone().ixor(t)
          : t.clone().ixor(this);
      }),
      (i.prototype.uxor = function (t) {
        return this.length > t.length
          ? this.clone().iuxor(t)
          : t.clone().iuxor(this);
      }),
      (i.prototype.inotn = function (t) {
        r('number' == typeof t && t >= 0);
        var e = 0 | Math.ceil(t / 26),
          n = t % 26;
        this._expand(e), n > 0 && e--;
        for (var i = 0; i < e; i++) this.words[i] = 67108863 & ~this.words[i];
        return (
          n > 0 && (this.words[i] = ~this.words[i] & (67108863 >> (26 - n))),
          this._strip()
        );
      }),
      (i.prototype.notn = function (t) {
        return this.clone().inotn(t);
      }),
      (i.prototype.setn = function (t, e) {
        r('number' == typeof t && t >= 0);
        var n = (t / 26) | 0,
          i = t % 26;
        return (
          this._expand(n + 1),
          (this.words[n] = e
            ? this.words[n] | (1 << i)
            : this.words[n] & ~(1 << i)),
          this._strip()
        );
      }),
      (i.prototype.iadd = function (t) {
        var e, r, n;
        if (0 !== this.negative && 0 === t.negative)
          return (
            (this.negative = 0),
            (e = this.isub(t)),
            (this.negative ^= 1),
            this._normSign()
          );
        if (0 === this.negative && 0 !== t.negative)
          return (
            (t.negative = 0),
            (e = this.isub(t)),
            (t.negative = 1),
            e._normSign()
          );
        this.length > t.length ? ((r = this), (n = t)) : ((r = t), (n = this));
        for (var i = 0, o = 0; o < n.length; o++)
          (e = (0 | r.words[o]) + (0 | n.words[o]) + i),
            (this.words[o] = 67108863 & e),
            (i = e >>> 26);
        for (; 0 !== i && o < r.length; o++)
          (e = (0 | r.words[o]) + i),
            (this.words[o] = 67108863 & e),
            (i = e >>> 26);
        if (((this.length = r.length), 0 !== i))
          (this.words[this.length] = i), this.length++;
        else if (r !== this)
          for (; o < r.length; o++) this.words[o] = r.words[o];
        return this;
      }),
      (i.prototype.add = function (t) {
        var e;
        return 0 !== t.negative && 0 === this.negative
          ? ((t.negative = 0), (e = this.sub(t)), (t.negative ^= 1), e)
          : 0 === t.negative && 0 !== this.negative
          ? ((this.negative = 0), (e = t.sub(this)), (this.negative = 1), e)
          : this.length > t.length
          ? this.clone().iadd(t)
          : t.clone().iadd(this);
      }),
      (i.prototype.isub = function (t) {
        if (0 !== t.negative) {
          t.negative = 0;
          var e = this.iadd(t);
          return (t.negative = 1), e._normSign();
        }
        if (0 !== this.negative)
          return (
            (this.negative = 0),
            this.iadd(t),
            (this.negative = 1),
            this._normSign()
          );
        var r,
          n,
          i = this.cmp(t);
        if (0 === i)
          return (
            (this.negative = 0), (this.length = 1), (this.words[0] = 0), this
          );
        i > 0 ? ((r = this), (n = t)) : ((r = t), (n = this));
        for (var o = 0, s = 0; s < n.length; s++)
          (o = (e = (0 | r.words[s]) - (0 | n.words[s]) + o) >> 26),
            (this.words[s] = 67108863 & e);
        for (; 0 !== o && s < r.length; s++)
          (o = (e = (0 | r.words[s]) + o) >> 26),
            (this.words[s] = 67108863 & e);
        if (0 === o && s < r.length && r !== this)
          for (; s < r.length; s++) this.words[s] = r.words[s];
        return (
          (this.length = Math.max(this.length, s)),
          r !== this && (this.negative = 1),
          this._strip()
        );
      }),
      (i.prototype.sub = function (t) {
        return this.clone().isub(t);
      });
    var p = function (t, e, r) {
      var n,
        i,
        o,
        s = t.words,
        a = e.words,
        u = r.words,
        h = 0,
        l = 0 | s[0],
        c = 8191 & l,
        f = l >>> 13,
        d = 0 | s[1],
        m = 8191 & d,
        p = d >>> 13,
        g = 0 | s[2],
        y = 8191 & g,
        v = g >>> 13,
        w = 0 | s[3],
        b = 8191 & w,
        A = w >>> 13,
        E = 0 | s[4],
        M = 8191 & E,
        _ = E >>> 13,
        k = 0 | s[5],
        S = 8191 & k,
        N = k >>> 13,
        P = 0 | s[6],
        x = 8191 & P,
        T = P >>> 13,
        C = 0 | s[7],
        R = 8191 & C,
        I = C >>> 13,
        B = 0 | s[8],
        O = 8191 & B,
        F = B >>> 13,
        D = 0 | s[9],
        L = 8191 & D,
        U = D >>> 13,
        j = 0 | a[0],
        z = 8191 & j,
        G = j >>> 13,
        q = 0 | a[1],
        K = 8191 & q,
        H = q >>> 13,
        V = 0 | a[2],
        J = 8191 & V,
        W = V >>> 13,
        Q = 0 | a[3],
        Y = 8191 & Q,
        Z = Q >>> 13,
        X = 0 | a[4],
        $ = 8191 & X,
        tt = X >>> 13,
        et = 0 | a[5],
        rt = 8191 & et,
        nt = et >>> 13,
        it = 0 | a[6],
        ot = 8191 & it,
        st = it >>> 13,
        at = 0 | a[7],
        ut = 8191 & at,
        ht = at >>> 13,
        lt = 0 | a[8],
        ct = 8191 & lt,
        ft = lt >>> 13,
        dt = 0 | a[9],
        mt = 8191 & dt,
        pt = dt >>> 13;
      (r.negative = t.negative ^ e.negative), (r.length = 19);
      var gt =
        (((h + (n = Math.imul(c, z))) | 0) +
          ((8191 & (i = ((i = Math.imul(c, G)) + Math.imul(f, z)) | 0)) <<
            13)) |
        0;
      (h = ((((o = Math.imul(f, G)) + (i >>> 13)) | 0) + (gt >>> 26)) | 0),
        (gt &= 67108863),
        (n = Math.imul(m, z)),
        (i = ((i = Math.imul(m, G)) + Math.imul(p, z)) | 0),
        (o = Math.imul(p, G));
      var yt =
        (((h + (n = (n + Math.imul(c, K)) | 0)) | 0) +
          ((8191 &
            (i = ((i = (i + Math.imul(c, H)) | 0) + Math.imul(f, K)) | 0)) <<
            13)) |
        0;
      (h =
        ((((o = (o + Math.imul(f, H)) | 0) + (i >>> 13)) | 0) + (yt >>> 26)) |
        0),
        (yt &= 67108863),
        (n = Math.imul(y, z)),
        (i = ((i = Math.imul(y, G)) + Math.imul(v, z)) | 0),
        (o = Math.imul(v, G)),
        (n = (n + Math.imul(m, K)) | 0),
        (i = ((i = (i + Math.imul(m, H)) | 0) + Math.imul(p, K)) | 0),
        (o = (o + Math.imul(p, H)) | 0);
      var vt =
        (((h + (n = (n + Math.imul(c, J)) | 0)) | 0) +
          ((8191 &
            (i = ((i = (i + Math.imul(c, W)) | 0) + Math.imul(f, J)) | 0)) <<
            13)) |
        0;
      (h =
        ((((o = (o + Math.imul(f, W)) | 0) + (i >>> 13)) | 0) + (vt >>> 26)) |
        0),
        (vt &= 67108863),
        (n = Math.imul(b, z)),
        (i = ((i = Math.imul(b, G)) + Math.imul(A, z)) | 0),
        (o = Math.imul(A, G)),
        (n = (n + Math.imul(y, K)) | 0),
        (i = ((i = (i + Math.imul(y, H)) | 0) + Math.imul(v, K)) | 0),
        (o = (o + Math.imul(v, H)) | 0),
        (n = (n + Math.imul(m, J)) | 0),
        (i = ((i = (i + Math.imul(m, W)) | 0) + Math.imul(p, J)) | 0),
        (o = (o + Math.imul(p, W)) | 0);
      var wt =
        (((h + (n = (n + Math.imul(c, Y)) | 0)) | 0) +
          ((8191 &
            (i = ((i = (i + Math.imul(c, Z)) | 0) + Math.imul(f, Y)) | 0)) <<
            13)) |
        0;
      (h =
        ((((o = (o + Math.imul(f, Z)) | 0) + (i >>> 13)) | 0) + (wt >>> 26)) |
        0),
        (wt &= 67108863),
        (n = Math.imul(M, z)),
        (i = ((i = Math.imul(M, G)) + Math.imul(_, z)) | 0),
        (o = Math.imul(_, G)),
        (n = (n + Math.imul(b, K)) | 0),
        (i = ((i = (i + Math.imul(b, H)) | 0) + Math.imul(A, K)) | 0),
        (o = (o + Math.imul(A, H)) | 0),
        (n = (n + Math.imul(y, J)) | 0),
        (i = ((i = (i + Math.imul(y, W)) | 0) + Math.imul(v, J)) | 0),
        (o = (o + Math.imul(v, W)) | 0),
        (n = (n + Math.imul(m, Y)) | 0),
        (i = ((i = (i + Math.imul(m, Z)) | 0) + Math.imul(p, Y)) | 0),
        (o = (o + Math.imul(p, Z)) | 0);
      var bt =
        (((h + (n = (n + Math.imul(c, $)) | 0)) | 0) +
          ((8191 &
            (i = ((i = (i + Math.imul(c, tt)) | 0) + Math.imul(f, $)) | 0)) <<
            13)) |
        0;
      (h =
        ((((o = (o + Math.imul(f, tt)) | 0) + (i >>> 13)) | 0) + (bt >>> 26)) |
        0),
        (bt &= 67108863),
        (n = Math.imul(S, z)),
        (i = ((i = Math.imul(S, G)) + Math.imul(N, z)) | 0),
        (o = Math.imul(N, G)),
        (n = (n + Math.imul(M, K)) | 0),
        (i = ((i = (i + Math.imul(M, H)) | 0) + Math.imul(_, K)) | 0),
        (o = (o + Math.imul(_, H)) | 0),
        (n = (n + Math.imul(b, J)) | 0),
        (i = ((i = (i + Math.imul(b, W)) | 0) + Math.imul(A, J)) | 0),
        (o = (o + Math.imul(A, W)) | 0),
        (n = (n + Math.imul(y, Y)) | 0),
        (i = ((i = (i + Math.imul(y, Z)) | 0) + Math.imul(v, Y)) | 0),
        (o = (o + Math.imul(v, Z)) | 0),
        (n = (n + Math.imul(m, $)) | 0),
        (i = ((i = (i + Math.imul(m, tt)) | 0) + Math.imul(p, $)) | 0),
        (o = (o + Math.imul(p, tt)) | 0);
      var At =
        (((h + (n = (n + Math.imul(c, rt)) | 0)) | 0) +
          ((8191 &
            (i = ((i = (i + Math.imul(c, nt)) | 0) + Math.imul(f, rt)) | 0)) <<
            13)) |
        0;
      (h =
        ((((o = (o + Math.imul(f, nt)) | 0) + (i >>> 13)) | 0) + (At >>> 26)) |
        0),
        (At &= 67108863),
        (n = Math.imul(x, z)),
        (i = ((i = Math.imul(x, G)) + Math.imul(T, z)) | 0),
        (o = Math.imul(T, G)),
        (n = (n + Math.imul(S, K)) | 0),
        (i = ((i = (i + Math.imul(S, H)) | 0) + Math.imul(N, K)) | 0),
        (o = (o + Math.imul(N, H)) | 0),
        (n = (n + Math.imul(M, J)) | 0),
        (i = ((i = (i + Math.imul(M, W)) | 0) + Math.imul(_, J)) | 0),
        (o = (o + Math.imul(_, W)) | 0),
        (n = (n + Math.imul(b, Y)) | 0),
        (i = ((i = (i + Math.imul(b, Z)) | 0) + Math.imul(A, Y)) | 0),
        (o = (o + Math.imul(A, Z)) | 0),
        (n = (n + Math.imul(y, $)) | 0),
        (i = ((i = (i + Math.imul(y, tt)) | 0) + Math.imul(v, $)) | 0),
        (o = (o + Math.imul(v, tt)) | 0),
        (n = (n + Math.imul(m, rt)) | 0),
        (i = ((i = (i + Math.imul(m, nt)) | 0) + Math.imul(p, rt)) | 0),
        (o = (o + Math.imul(p, nt)) | 0);
      var Et =
        (((h + (n = (n + Math.imul(c, ot)) | 0)) | 0) +
          ((8191 &
            (i = ((i = (i + Math.imul(c, st)) | 0) + Math.imul(f, ot)) | 0)) <<
            13)) |
        0;
      (h =
        ((((o = (o + Math.imul(f, st)) | 0) + (i >>> 13)) | 0) + (Et >>> 26)) |
        0),
        (Et &= 67108863),
        (n = Math.imul(R, z)),
        (i = ((i = Math.imul(R, G)) + Math.imul(I, z)) | 0),
        (o = Math.imul(I, G)),
        (n = (n + Math.imul(x, K)) | 0),
        (i = ((i = (i + Math.imul(x, H)) | 0) + Math.imul(T, K)) | 0),
        (o = (o + Math.imul(T, H)) | 0),
        (n = (n + Math.imul(S, J)) | 0),
        (i = ((i = (i + Math.imul(S, W)) | 0) + Math.imul(N, J)) | 0),
        (o = (o + Math.imul(N, W)) | 0),
        (n = (n + Math.imul(M, Y)) | 0),
        (i = ((i = (i + Math.imul(M, Z)) | 0) + Math.imul(_, Y)) | 0),
        (o = (o + Math.imul(_, Z)) | 0),
        (n = (n + Math.imul(b, $)) | 0),
        (i = ((i = (i + Math.imul(b, tt)) | 0) + Math.imul(A, $)) | 0),
        (o = (o + Math.imul(A, tt)) | 0),
        (n = (n + Math.imul(y, rt)) | 0),
        (i = ((i = (i + Math.imul(y, nt)) | 0) + Math.imul(v, rt)) | 0),
        (o = (o + Math.imul(v, nt)) | 0),
        (n = (n + Math.imul(m, ot)) | 0),
        (i = ((i = (i + Math.imul(m, st)) | 0) + Math.imul(p, ot)) | 0),
        (o = (o + Math.imul(p, st)) | 0);
      var Mt =
        (((h + (n = (n + Math.imul(c, ut)) | 0)) | 0) +
          ((8191 &
            (i = ((i = (i + Math.imul(c, ht)) | 0) + Math.imul(f, ut)) | 0)) <<
            13)) |
        0;
      (h =
        ((((o = (o + Math.imul(f, ht)) | 0) + (i >>> 13)) | 0) + (Mt >>> 26)) |
        0),
        (Mt &= 67108863),
        (n = Math.imul(O, z)),
        (i = ((i = Math.imul(O, G)) + Math.imul(F, z)) | 0),
        (o = Math.imul(F, G)),
        (n = (n + Math.imul(R, K)) | 0),
        (i = ((i = (i + Math.imul(R, H)) | 0) + Math.imul(I, K)) | 0),
        (o = (o + Math.imul(I, H)) | 0),
        (n = (n + Math.imul(x, J)) | 0),
        (i = ((i = (i + Math.imul(x, W)) | 0) + Math.imul(T, J)) | 0),
        (o = (o + Math.imul(T, W)) | 0),
        (n = (n + Math.imul(S, Y)) | 0),
        (i = ((i = (i + Math.imul(S, Z)) | 0) + Math.imul(N, Y)) | 0),
        (o = (o + Math.imul(N, Z)) | 0),
        (n = (n + Math.imul(M, $)) | 0),
        (i = ((i = (i + Math.imul(M, tt)) | 0) + Math.imul(_, $)) | 0),
        (o = (o + Math.imul(_, tt)) | 0),
        (n = (n + Math.imul(b, rt)) | 0),
        (i = ((i = (i + Math.imul(b, nt)) | 0) + Math.imul(A, rt)) | 0),
        (o = (o + Math.imul(A, nt)) | 0),
        (n = (n + Math.imul(y, ot)) | 0),
        (i = ((i = (i + Math.imul(y, st)) | 0) + Math.imul(v, ot)) | 0),
        (o = (o + Math.imul(v, st)) | 0),
        (n = (n + Math.imul(m, ut)) | 0),
        (i = ((i = (i + Math.imul(m, ht)) | 0) + Math.imul(p, ut)) | 0),
        (o = (o + Math.imul(p, ht)) | 0);
      var _t =
        (((h + (n = (n + Math.imul(c, ct)) | 0)) | 0) +
          ((8191 &
            (i = ((i = (i + Math.imul(c, ft)) | 0) + Math.imul(f, ct)) | 0)) <<
            13)) |
        0;
      (h =
        ((((o = (o + Math.imul(f, ft)) | 0) + (i >>> 13)) | 0) + (_t >>> 26)) |
        0),
        (_t &= 67108863),
        (n = Math.imul(L, z)),
        (i = ((i = Math.imul(L, G)) + Math.imul(U, z)) | 0),
        (o = Math.imul(U, G)),
        (n = (n + Math.imul(O, K)) | 0),
        (i = ((i = (i + Math.imul(O, H)) | 0) + Math.imul(F, K)) | 0),
        (o = (o + Math.imul(F, H)) | 0),
        (n = (n + Math.imul(R, J)) | 0),
        (i = ((i = (i + Math.imul(R, W)) | 0) + Math.imul(I, J)) | 0),
        (o = (o + Math.imul(I, W)) | 0),
        (n = (n + Math.imul(x, Y)) | 0),
        (i = ((i = (i + Math.imul(x, Z)) | 0) + Math.imul(T, Y)) | 0),
        (o = (o + Math.imul(T, Z)) | 0),
        (n = (n + Math.imul(S, $)) | 0),
        (i = ((i = (i + Math.imul(S, tt)) | 0) + Math.imul(N, $)) | 0),
        (o = (o + Math.imul(N, tt)) | 0),
        (n = (n + Math.imul(M, rt)) | 0),
        (i = ((i = (i + Math.imul(M, nt)) | 0) + Math.imul(_, rt)) | 0),
        (o = (o + Math.imul(_, nt)) | 0),
        (n = (n + Math.imul(b, ot)) | 0),
        (i = ((i = (i + Math.imul(b, st)) | 0) + Math.imul(A, ot)) | 0),
        (o = (o + Math.imul(A, st)) | 0),
        (n = (n + Math.imul(y, ut)) | 0),
        (i = ((i = (i + Math.imul(y, ht)) | 0) + Math.imul(v, ut)) | 0),
        (o = (o + Math.imul(v, ht)) | 0),
        (n = (n + Math.imul(m, ct)) | 0),
        (i = ((i = (i + Math.imul(m, ft)) | 0) + Math.imul(p, ct)) | 0),
        (o = (o + Math.imul(p, ft)) | 0);
      var kt =
        (((h + (n = (n + Math.imul(c, mt)) | 0)) | 0) +
          ((8191 &
            (i = ((i = (i + Math.imul(c, pt)) | 0) + Math.imul(f, mt)) | 0)) <<
            13)) |
        0;
      (h =
        ((((o = (o + Math.imul(f, pt)) | 0) + (i >>> 13)) | 0) + (kt >>> 26)) |
        0),
        (kt &= 67108863),
        (n = Math.imul(L, K)),
        (i = ((i = Math.imul(L, H)) + Math.imul(U, K)) | 0),
        (o = Math.imul(U, H)),
        (n = (n + Math.imul(O, J)) | 0),
        (i = ((i = (i + Math.imul(O, W)) | 0) + Math.imul(F, J)) | 0),
        (o = (o + Math.imul(F, W)) | 0),
        (n = (n + Math.imul(R, Y)) | 0),
        (i = ((i = (i + Math.imul(R, Z)) | 0) + Math.imul(I, Y)) | 0),
        (o = (o + Math.imul(I, Z)) | 0),
        (n = (n + Math.imul(x, $)) | 0),
        (i = ((i = (i + Math.imul(x, tt)) | 0) + Math.imul(T, $)) | 0),
        (o = (o + Math.imul(T, tt)) | 0),
        (n = (n + Math.imul(S, rt)) | 0),
        (i = ((i = (i + Math.imul(S, nt)) | 0) + Math.imul(N, rt)) | 0),
        (o = (o + Math.imul(N, nt)) | 0),
        (n = (n + Math.imul(M, ot)) | 0),
        (i = ((i = (i + Math.imul(M, st)) | 0) + Math.imul(_, ot)) | 0),
        (o = (o + Math.imul(_, st)) | 0),
        (n = (n + Math.imul(b, ut)) | 0),
        (i = ((i = (i + Math.imul(b, ht)) | 0) + Math.imul(A, ut)) | 0),
        (o = (o + Math.imul(A, ht)) | 0),
        (n = (n + Math.imul(y, ct)) | 0),
        (i = ((i = (i + Math.imul(y, ft)) | 0) + Math.imul(v, ct)) | 0),
        (o = (o + Math.imul(v, ft)) | 0);
      var St =
        (((h + (n = (n + Math.imul(m, mt)) | 0)) | 0) +
          ((8191 &
            (i = ((i = (i + Math.imul(m, pt)) | 0) + Math.imul(p, mt)) | 0)) <<
            13)) |
        0;
      (h =
        ((((o = (o + Math.imul(p, pt)) | 0) + (i >>> 13)) | 0) + (St >>> 26)) |
        0),
        (St &= 67108863),
        (n = Math.imul(L, J)),
        (i = ((i = Math.imul(L, W)) + Math.imul(U, J)) | 0),
        (o = Math.imul(U, W)),
        (n = (n + Math.imul(O, Y)) | 0),
        (i = ((i = (i + Math.imul(O, Z)) | 0) + Math.imul(F, Y)) | 0),
        (o = (o + Math.imul(F, Z)) | 0),
        (n = (n + Math.imul(R, $)) | 0),
        (i = ((i = (i + Math.imul(R, tt)) | 0) + Math.imul(I, $)) | 0),
        (o = (o + Math.imul(I, tt)) | 0),
        (n = (n + Math.imul(x, rt)) | 0),
        (i = ((i = (i + Math.imul(x, nt)) | 0) + Math.imul(T, rt)) | 0),
        (o = (o + Math.imul(T, nt)) | 0),
        (n = (n + Math.imul(S, ot)) | 0),
        (i = ((i = (i + Math.imul(S, st)) | 0) + Math.imul(N, ot)) | 0),
        (o = (o + Math.imul(N, st)) | 0),
        (n = (n + Math.imul(M, ut)) | 0),
        (i = ((i = (i + Math.imul(M, ht)) | 0) + Math.imul(_, ut)) | 0),
        (o = (o + Math.imul(_, ht)) | 0),
        (n = (n + Math.imul(b, ct)) | 0),
        (i = ((i = (i + Math.imul(b, ft)) | 0) + Math.imul(A, ct)) | 0),
        (o = (o + Math.imul(A, ft)) | 0);
      var Nt =
        (((h + (n = (n + Math.imul(y, mt)) | 0)) | 0) +
          ((8191 &
            (i = ((i = (i + Math.imul(y, pt)) | 0) + Math.imul(v, mt)) | 0)) <<
            13)) |
        0;
      (h =
        ((((o = (o + Math.imul(v, pt)) | 0) + (i >>> 13)) | 0) + (Nt >>> 26)) |
        0),
        (Nt &= 67108863),
        (n = Math.imul(L, Y)),
        (i = ((i = Math.imul(L, Z)) + Math.imul(U, Y)) | 0),
        (o = Math.imul(U, Z)),
        (n = (n + Math.imul(O, $)) | 0),
        (i = ((i = (i + Math.imul(O, tt)) | 0) + Math.imul(F, $)) | 0),
        (o = (o + Math.imul(F, tt)) | 0),
        (n = (n + Math.imul(R, rt)) | 0),
        (i = ((i = (i + Math.imul(R, nt)) | 0) + Math.imul(I, rt)) | 0),
        (o = (o + Math.imul(I, nt)) | 0),
        (n = (n + Math.imul(x, ot)) | 0),
        (i = ((i = (i + Math.imul(x, st)) | 0) + Math.imul(T, ot)) | 0),
        (o = (o + Math.imul(T, st)) | 0),
        (n = (n + Math.imul(S, ut)) | 0),
        (i = ((i = (i + Math.imul(S, ht)) | 0) + Math.imul(N, ut)) | 0),
        (o = (o + Math.imul(N, ht)) | 0),
        (n = (n + Math.imul(M, ct)) | 0),
        (i = ((i = (i + Math.imul(M, ft)) | 0) + Math.imul(_, ct)) | 0),
        (o = (o + Math.imul(_, ft)) | 0);
      var Pt =
        (((h + (n = (n + Math.imul(b, mt)) | 0)) | 0) +
          ((8191 &
            (i = ((i = (i + Math.imul(b, pt)) | 0) + Math.imul(A, mt)) | 0)) <<
            13)) |
        0;
      (h =
        ((((o = (o + Math.imul(A, pt)) | 0) + (i >>> 13)) | 0) + (Pt >>> 26)) |
        0),
        (Pt &= 67108863),
        (n = Math.imul(L, $)),
        (i = ((i = Math.imul(L, tt)) + Math.imul(U, $)) | 0),
        (o = Math.imul(U, tt)),
        (n = (n + Math.imul(O, rt)) | 0),
        (i = ((i = (i + Math.imul(O, nt)) | 0) + Math.imul(F, rt)) | 0),
        (o = (o + Math.imul(F, nt)) | 0),
        (n = (n + Math.imul(R, ot)) | 0),
        (i = ((i = (i + Math.imul(R, st)) | 0) + Math.imul(I, ot)) | 0),
        (o = (o + Math.imul(I, st)) | 0),
        (n = (n + Math.imul(x, ut)) | 0),
        (i = ((i = (i + Math.imul(x, ht)) | 0) + Math.imul(T, ut)) | 0),
        (o = (o + Math.imul(T, ht)) | 0),
        (n = (n + Math.imul(S, ct)) | 0),
        (i = ((i = (i + Math.imul(S, ft)) | 0) + Math.imul(N, ct)) | 0),
        (o = (o + Math.imul(N, ft)) | 0);
      var xt =
        (((h + (n = (n + Math.imul(M, mt)) | 0)) | 0) +
          ((8191 &
            (i = ((i = (i + Math.imul(M, pt)) | 0) + Math.imul(_, mt)) | 0)) <<
            13)) |
        0;
      (h =
        ((((o = (o + Math.imul(_, pt)) | 0) + (i >>> 13)) | 0) + (xt >>> 26)) |
        0),
        (xt &= 67108863),
        (n = Math.imul(L, rt)),
        (i = ((i = Math.imul(L, nt)) + Math.imul(U, rt)) | 0),
        (o = Math.imul(U, nt)),
        (n = (n + Math.imul(O, ot)) | 0),
        (i = ((i = (i + Math.imul(O, st)) | 0) + Math.imul(F, ot)) | 0),
        (o = (o + Math.imul(F, st)) | 0),
        (n = (n + Math.imul(R, ut)) | 0),
        (i = ((i = (i + Math.imul(R, ht)) | 0) + Math.imul(I, ut)) | 0),
        (o = (o + Math.imul(I, ht)) | 0),
        (n = (n + Math.imul(x, ct)) | 0),
        (i = ((i = (i + Math.imul(x, ft)) | 0) + Math.imul(T, ct)) | 0),
        (o = (o + Math.imul(T, ft)) | 0);
      var Tt =
        (((h + (n = (n + Math.imul(S, mt)) | 0)) | 0) +
          ((8191 &
            (i = ((i = (i + Math.imul(S, pt)) | 0) + Math.imul(N, mt)) | 0)) <<
            13)) |
        0;
      (h =
        ((((o = (o + Math.imul(N, pt)) | 0) + (i >>> 13)) | 0) + (Tt >>> 26)) |
        0),
        (Tt &= 67108863),
        (n = Math.imul(L, ot)),
        (i = ((i = Math.imul(L, st)) + Math.imul(U, ot)) | 0),
        (o = Math.imul(U, st)),
        (n = (n + Math.imul(O, ut)) | 0),
        (i = ((i = (i + Math.imul(O, ht)) | 0) + Math.imul(F, ut)) | 0),
        (o = (o + Math.imul(F, ht)) | 0),
        (n = (n + Math.imul(R, ct)) | 0),
        (i = ((i = (i + Math.imul(R, ft)) | 0) + Math.imul(I, ct)) | 0),
        (o = (o + Math.imul(I, ft)) | 0);
      var Ct =
        (((h + (n = (n + Math.imul(x, mt)) | 0)) | 0) +
          ((8191 &
            (i = ((i = (i + Math.imul(x, pt)) | 0) + Math.imul(T, mt)) | 0)) <<
            13)) |
        0;
      (h =
        ((((o = (o + Math.imul(T, pt)) | 0) + (i >>> 13)) | 0) + (Ct >>> 26)) |
        0),
        (Ct &= 67108863),
        (n = Math.imul(L, ut)),
        (i = ((i = Math.imul(L, ht)) + Math.imul(U, ut)) | 0),
        (o = Math.imul(U, ht)),
        (n = (n + Math.imul(O, ct)) | 0),
        (i = ((i = (i + Math.imul(O, ft)) | 0) + Math.imul(F, ct)) | 0),
        (o = (o + Math.imul(F, ft)) | 0);
      var Rt =
        (((h + (n = (n + Math.imul(R, mt)) | 0)) | 0) +
          ((8191 &
            (i = ((i = (i + Math.imul(R, pt)) | 0) + Math.imul(I, mt)) | 0)) <<
            13)) |
        0;
      (h =
        ((((o = (o + Math.imul(I, pt)) | 0) + (i >>> 13)) | 0) + (Rt >>> 26)) |
        0),
        (Rt &= 67108863),
        (n = Math.imul(L, ct)),
        (i = ((i = Math.imul(L, ft)) + Math.imul(U, ct)) | 0),
        (o = Math.imul(U, ft));
      var It =
        (((h + (n = (n + Math.imul(O, mt)) | 0)) | 0) +
          ((8191 &
            (i = ((i = (i + Math.imul(O, pt)) | 0) + Math.imul(F, mt)) | 0)) <<
            13)) |
        0;
      (h =
        ((((o = (o + Math.imul(F, pt)) | 0) + (i >>> 13)) | 0) + (It >>> 26)) |
        0),
        (It &= 67108863);
      var Bt =
        (((h + (n = Math.imul(L, mt))) | 0) +
          ((8191 & (i = ((i = Math.imul(L, pt)) + Math.imul(U, mt)) | 0)) <<
            13)) |
        0;
      return (
        (h = ((((o = Math.imul(U, pt)) + (i >>> 13)) | 0) + (Bt >>> 26)) | 0),
        (Bt &= 67108863),
        (u[0] = gt),
        (u[1] = yt),
        (u[2] = vt),
        (u[3] = wt),
        (u[4] = bt),
        (u[5] = At),
        (u[6] = Et),
        (u[7] = Mt),
        (u[8] = _t),
        (u[9] = kt),
        (u[10] = St),
        (u[11] = Nt),
        (u[12] = Pt),
        (u[13] = xt),
        (u[14] = Tt),
        (u[15] = Ct),
        (u[16] = Rt),
        (u[17] = It),
        (u[18] = Bt),
        0 !== h && ((u[19] = h), r.length++),
        r
      );
    };
    function g(t, e, r) {
      (r.negative = e.negative ^ t.negative), (r.length = t.length + e.length);
      for (var n = 0, i = 0, o = 0; o < r.length - 1; o++) {
        var s = i;
        i = 0;
        for (
          var a = 67108863 & n,
            u = Math.min(o, e.length - 1),
            h = Math.max(0, o - t.length + 1);
          h <= u;
          h++
        ) {
          var l = o - h,
            c = (0 | t.words[l]) * (0 | e.words[h]),
            f = 67108863 & c;
          (a = 67108863 & (f = (f + a) | 0)),
            (i +=
              (s = ((s = (s + ((c / 67108864) | 0)) | 0) + (f >>> 26)) | 0) >>>
              26),
            (s &= 67108863);
        }
        (r.words[o] = a), (n = s), (s = i);
      }
      return 0 !== n ? (r.words[o] = n) : r.length--, r._strip();
    }
    function y(t, e, r) {
      return g(t, e, r);
    }
    Math.imul || (p = m),
      (i.prototype.mulTo = function (t, e) {
        var r = this.length + t.length;
        return 10 === this.length && 10 === t.length
          ? p(this, t, e)
          : r < 63
          ? m(this, t, e)
          : r < 1024
          ? g(this, t, e)
          : y(this, t, e);
      }),
      (i.prototype.mul = function (t) {
        var e = new i(null);
        return (e.words = new Array(this.length + t.length)), this.mulTo(t, e);
      }),
      (i.prototype.mulf = function (t) {
        var e = new i(null);
        return (e.words = new Array(this.length + t.length)), y(this, t, e);
      }),
      (i.prototype.imul = function (t) {
        return this.clone().mulTo(t, this);
      }),
      (i.prototype.imuln = function (t) {
        var e = t < 0;
        e && (t = -t), r('number' == typeof t), r(t < 67108864);
        for (var n = 0, i = 0; i < this.length; i++) {
          var o = (0 | this.words[i]) * t,
            s = (67108863 & o) + (67108863 & n);
          (n >>= 26),
            (n += (o / 67108864) | 0),
            (n += s >>> 26),
            (this.words[i] = 67108863 & s);
        }
        return (
          0 !== n && ((this.words[i] = n), this.length++),
          e ? this.ineg() : this
        );
      }),
      (i.prototype.muln = function (t) {
        return this.clone().imuln(t);
      }),
      (i.prototype.sqr = function () {
        return this.mul(this);
      }),
      (i.prototype.isqr = function () {
        return this.imul(this.clone());
      }),
      (i.prototype.pow = function (t) {
        var e = (function (t) {
          for (var e = new Array(t.bitLength()), r = 0; r < e.length; r++) {
            var n = (r / 26) | 0,
              i = r % 26;
            e[r] = (t.words[n] >>> i) & 1;
          }
          return e;
        })(t);
        if (0 === e.length) return new i(1);
        for (var r = this, n = 0; n < e.length && 0 === e[n]; n++, r = r.sqr());
        if (++n < e.length)
          for (var o = r.sqr(); n < e.length; n++, o = o.sqr())
            0 !== e[n] && (r = r.mul(o));
        return r;
      }),
      (i.prototype.iushln = function (t) {
        r('number' == typeof t && t >= 0);
        var e,
          n = t % 26,
          i = (t - n) / 26,
          o = (67108863 >>> (26 - n)) << (26 - n);
        if (0 !== n) {
          var s = 0;
          for (e = 0; e < this.length; e++) {
            var a = this.words[e] & o,
              u = ((0 | this.words[e]) - a) << n;
            (this.words[e] = u | s), (s = a >>> (26 - n));
          }
          s && ((this.words[e] = s), this.length++);
        }
        if (0 !== i) {
          for (e = this.length - 1; e >= 0; e--)
            this.words[e + i] = this.words[e];
          for (e = 0; e < i; e++) this.words[e] = 0;
          this.length += i;
        }
        return this._strip();
      }),
      (i.prototype.ishln = function (t) {
        return r(0 === this.negative), this.iushln(t);
      }),
      (i.prototype.iushrn = function (t, e, n) {
        var i;
        r('number' == typeof t && t >= 0), (i = e ? (e - (e % 26)) / 26 : 0);
        var o = t % 26,
          s = Math.min((t - o) / 26, this.length),
          a = 67108863 ^ ((67108863 >>> o) << o),
          u = n;
        if (((i -= s), (i = Math.max(0, i)), u)) {
          for (var h = 0; h < s; h++) u.words[h] = this.words[h];
          u.length = s;
        }
        if (0 === s);
        else if (this.length > s)
          for (this.length -= s, h = 0; h < this.length; h++)
            this.words[h] = this.words[h + s];
        else (this.words[0] = 0), (this.length = 1);
        var l = 0;
        for (h = this.length - 1; h >= 0 && (0 !== l || h >= i); h--) {
          var c = 0 | this.words[h];
          (this.words[h] = (l << (26 - o)) | (c >>> o)), (l = c & a);
        }
        return (
          u && 0 !== l && (u.words[u.length++] = l),
          0 === this.length && ((this.words[0] = 0), (this.length = 1)),
          this._strip()
        );
      }),
      (i.prototype.ishrn = function (t, e, n) {
        return r(0 === this.negative), this.iushrn(t, e, n);
      }),
      (i.prototype.shln = function (t) {
        return this.clone().ishln(t);
      }),
      (i.prototype.ushln = function (t) {
        return this.clone().iushln(t);
      }),
      (i.prototype.shrn = function (t) {
        return this.clone().ishrn(t);
      }),
      (i.prototype.ushrn = function (t) {
        return this.clone().iushrn(t);
      }),
      (i.prototype.testn = function (t) {
        r('number' == typeof t && t >= 0);
        var e = t % 26,
          n = (t - e) / 26,
          i = 1 << e;
        return !(this.length <= n || !(this.words[n] & i));
      }),
      (i.prototype.imaskn = function (t) {
        r('number' == typeof t && t >= 0);
        var e = t % 26,
          n = (t - e) / 26;
        if (
          (r(0 === this.negative, 'imaskn works only with positive numbers'),
          this.length <= n)
        )
          return this;
        if (
          (0 !== e && n++, (this.length = Math.min(n, this.length)), 0 !== e)
        ) {
          var i = 67108863 ^ ((67108863 >>> e) << e);
          this.words[this.length - 1] &= i;
        }
        return this._strip();
      }),
      (i.prototype.maskn = function (t) {
        return this.clone().imaskn(t);
      }),
      (i.prototype.iaddn = function (t) {
        return (
          r('number' == typeof t),
          r(t < 67108864),
          t < 0
            ? this.isubn(-t)
            : 0 !== this.negative
            ? 1 === this.length && (0 | this.words[0]) <= t
              ? ((this.words[0] = t - (0 | this.words[0])),
                (this.negative = 0),
                this)
              : ((this.negative = 0), this.isubn(t), (this.negative = 1), this)
            : this._iaddn(t)
        );
      }),
      (i.prototype._iaddn = function (t) {
        this.words[0] += t;
        for (var e = 0; e < this.length && this.words[e] >= 67108864; e++)
          (this.words[e] -= 67108864),
            e === this.length - 1
              ? (this.words[e + 1] = 1)
              : this.words[e + 1]++;
        return (this.length = Math.max(this.length, e + 1)), this;
      }),
      (i.prototype.isubn = function (t) {
        if ((r('number' == typeof t), r(t < 67108864), t < 0))
          return this.iaddn(-t);
        if (0 !== this.negative)
          return (this.negative = 0), this.iaddn(t), (this.negative = 1), this;
        if (((this.words[0] -= t), 1 === this.length && this.words[0] < 0))
          (this.words[0] = -this.words[0]), (this.negative = 1);
        else
          for (var e = 0; e < this.length && this.words[e] < 0; e++)
            (this.words[e] += 67108864), (this.words[e + 1] -= 1);
        return this._strip();
      }),
      (i.prototype.addn = function (t) {
        return this.clone().iaddn(t);
      }),
      (i.prototype.subn = function (t) {
        return this.clone().isubn(t);
      }),
      (i.prototype.iabs = function () {
        return (this.negative = 0), this;
      }),
      (i.prototype.abs = function () {
        return this.clone().iabs();
      }),
      (i.prototype._ishlnsubmul = function (t, e, n) {
        var i,
          o,
          s = t.length + n;
        this._expand(s);
        var a = 0;
        for (i = 0; i < t.length; i++) {
          o = (0 | this.words[i + n]) + a;
          var u = (0 | t.words[i]) * e;
          (a = ((o -= 67108863 & u) >> 26) - ((u / 67108864) | 0)),
            (this.words[i + n] = 67108863 & o);
        }
        for (; i < this.length - n; i++)
          (a = (o = (0 | this.words[i + n]) + a) >> 26),
            (this.words[i + n] = 67108863 & o);
        if (0 === a) return this._strip();
        for (r(-1 === a), a = 0, i = 0; i < this.length; i++)
          (a = (o = -(0 | this.words[i]) + a) >> 26),
            (this.words[i] = 67108863 & o);
        return (this.negative = 1), this._strip();
      }),
      (i.prototype._wordDiv = function (t, e) {
        var r = (this.length, t.length),
          n = this.clone(),
          o = t,
          s = 0 | o.words[o.length - 1];
        0 != (r = 26 - this._countBits(s)) &&
          ((o = o.ushln(r)), n.iushln(r), (s = 0 | o.words[o.length - 1]));
        var a,
          u = n.length - o.length;
        if ('mod' !== e) {
          ((a = new i(null)).length = u + 1), (a.words = new Array(a.length));
          for (var h = 0; h < a.length; h++) a.words[h] = 0;
        }
        var l = n.clone()._ishlnsubmul(o, 1, u);
        0 === l.negative && ((n = l), a && (a.words[u] = 1));
        for (var c = u - 1; c >= 0; c--) {
          var f =
            67108864 * (0 | n.words[o.length + c]) +
            (0 | n.words[o.length + c - 1]);
          for (
            f = Math.min((f / s) | 0, 67108863), n._ishlnsubmul(o, f, c);
            0 !== n.negative;

          )
            f--,
              (n.negative = 0),
              n._ishlnsubmul(o, 1, c),
              n.isZero() || (n.negative ^= 1);
          a && (a.words[c] = f);
        }
        return (
          a && a._strip(),
          n._strip(),
          'div' !== e && 0 !== r && n.iushrn(r),
          { div: a || null, mod: n }
        );
      }),
      (i.prototype.divmod = function (t, e, n) {
        return (
          r(!t.isZero()),
          this.isZero()
            ? { div: new i(0), mod: new i(0) }
            : 0 !== this.negative && 0 === t.negative
            ? ((a = this.neg().divmod(t, e)),
              'mod' !== e && (o = a.div.neg()),
              'div' !== e &&
                ((s = a.mod.neg()), n && 0 !== s.negative && s.iadd(t)),
              { div: o, mod: s })
            : 0 === this.negative && 0 !== t.negative
            ? ((a = this.divmod(t.neg(), e)),
              'mod' !== e && (o = a.div.neg()),
              { div: o, mod: a.mod })
            : this.negative & t.negative
            ? ((a = this.neg().divmod(t.neg(), e)),
              'div' !== e &&
                ((s = a.mod.neg()), n && 0 !== s.negative && s.isub(t)),
              { div: a.div, mod: s })
            : t.length > this.length || this.cmp(t) < 0
            ? { div: new i(0), mod: this }
            : 1 === t.length
            ? 'div' === e
              ? { div: this.divn(t.words[0]), mod: null }
              : 'mod' === e
              ? { div: null, mod: new i(this.modrn(t.words[0])) }
              : {
                  div: this.divn(t.words[0]),
                  mod: new i(this.modrn(t.words[0])),
                }
            : this._wordDiv(t, e)
        );
        var o, s, a;
      }),
      (i.prototype.div = function (t) {
        return this.divmod(t, 'div', !1).div;
      }),
      (i.prototype.mod = function (t) {
        return this.divmod(t, 'mod', !1).mod;
      }),
      (i.prototype.umod = function (t) {
        return this.divmod(t, 'mod', !0).mod;
      }),
      (i.prototype.divRound = function (t) {
        var e = this.divmod(t);
        if (e.mod.isZero()) return e.div;
        var r = 0 !== e.div.negative ? e.mod.isub(t) : e.mod,
          n = t.ushrn(1),
          i = t.andln(1),
          o = r.cmp(n);
        return o < 0 || (1 === i && 0 === o)
          ? e.div
          : 0 !== e.div.negative
          ? e.div.isubn(1)
          : e.div.iaddn(1);
      }),
      (i.prototype.modrn = function (t) {
        var e = t < 0;
        e && (t = -t), r(t <= 67108863);
        for (var n = (1 << 26) % t, i = 0, o = this.length - 1; o >= 0; o--)
          i = (n * i + (0 | this.words[o])) % t;
        return e ? -i : i;
      }),
      (i.prototype.modn = function (t) {
        return this.modrn(t);
      }),
      (i.prototype.idivn = function (t) {
        var e = t < 0;
        e && (t = -t), r(t <= 67108863);
        for (var n = 0, i = this.length - 1; i >= 0; i--) {
          var o = (0 | this.words[i]) + 67108864 * n;
          (this.words[i] = (o / t) | 0), (n = o % t);
        }
        return this._strip(), e ? this.ineg() : this;
      }),
      (i.prototype.divn = function (t) {
        return this.clone().idivn(t);
      }),
      (i.prototype.egcd = function (t) {
        r(0 === t.negative), r(!t.isZero());
        var e = this,
          n = t.clone();
        e = 0 !== e.negative ? e.umod(t) : e.clone();
        for (
          var o = new i(1), s = new i(0), a = new i(0), u = new i(1), h = 0;
          e.isEven() && n.isEven();

        )
          e.iushrn(1), n.iushrn(1), ++h;
        for (var l = n.clone(), c = e.clone(); !e.isZero(); ) {
          for (var f = 0, d = 1; !(e.words[0] & d) && f < 26; ++f, d <<= 1);
          if (f > 0)
            for (e.iushrn(f); f-- > 0; )
              (o.isOdd() || s.isOdd()) && (o.iadd(l), s.isub(c)),
                o.iushrn(1),
                s.iushrn(1);
          for (var m = 0, p = 1; !(n.words[0] & p) && m < 26; ++m, p <<= 1);
          if (m > 0)
            for (n.iushrn(m); m-- > 0; )
              (a.isOdd() || u.isOdd()) && (a.iadd(l), u.isub(c)),
                a.iushrn(1),
                u.iushrn(1);
          e.cmp(n) >= 0
            ? (e.isub(n), o.isub(a), s.isub(u))
            : (n.isub(e), a.isub(o), u.isub(s));
        }
        return { a: a, b: u, gcd: n.iushln(h) };
      }),
      (i.prototype._invmp = function (t) {
        r(0 === t.negative), r(!t.isZero());
        var e = this,
          n = t.clone();
        e = 0 !== e.negative ? e.umod(t) : e.clone();
        for (
          var o, s = new i(1), a = new i(0), u = n.clone();
          e.cmpn(1) > 0 && n.cmpn(1) > 0;

        ) {
          for (var h = 0, l = 1; !(e.words[0] & l) && h < 26; ++h, l <<= 1);
          if (h > 0)
            for (e.iushrn(h); h-- > 0; ) s.isOdd() && s.iadd(u), s.iushrn(1);
          for (var c = 0, f = 1; !(n.words[0] & f) && c < 26; ++c, f <<= 1);
          if (c > 0)
            for (n.iushrn(c); c-- > 0; ) a.isOdd() && a.iadd(u), a.iushrn(1);
          e.cmp(n) >= 0 ? (e.isub(n), s.isub(a)) : (n.isub(e), a.isub(s));
        }
        return (o = 0 === e.cmpn(1) ? s : a).cmpn(0) < 0 && o.iadd(t), o;
      }),
      (i.prototype.gcd = function (t) {
        if (this.isZero()) return t.abs();
        if (t.isZero()) return this.abs();
        var e = this.clone(),
          r = t.clone();
        (e.negative = 0), (r.negative = 0);
        for (var n = 0; e.isEven() && r.isEven(); n++) e.iushrn(1), r.iushrn(1);
        for (;;) {
          for (; e.isEven(); ) e.iushrn(1);
          for (; r.isEven(); ) r.iushrn(1);
          var i = e.cmp(r);
          if (i < 0) {
            var o = e;
            (e = r), (r = o);
          } else if (0 === i || 0 === r.cmpn(1)) break;
          e.isub(r);
        }
        return r.iushln(n);
      }),
      (i.prototype.invm = function (t) {
        return this.egcd(t).a.umod(t);
      }),
      (i.prototype.isEven = function () {
        return !(1 & this.words[0]);
      }),
      (i.prototype.isOdd = function () {
        return !(1 & ~this.words[0]);
      }),
      (i.prototype.andln = function (t) {
        return this.words[0] & t;
      }),
      (i.prototype.bincn = function (t) {
        r('number' == typeof t);
        var e = t % 26,
          n = (t - e) / 26,
          i = 1 << e;
        if (this.length <= n)
          return this._expand(n + 1), (this.words[n] |= i), this;
        for (var o = i, s = n; 0 !== o && s < this.length; s++) {
          var a = 0 | this.words[s];
          (o = (a += o) >>> 26), (a &= 67108863), (this.words[s] = a);
        }
        return 0 !== o && ((this.words[s] = o), this.length++), this;
      }),
      (i.prototype.isZero = function () {
        return 1 === this.length && 0 === this.words[0];
      }),
      (i.prototype.cmpn = function (t) {
        var e,
          n = t < 0;
        if (0 !== this.negative && !n) return -1;
        if (0 === this.negative && n) return 1;
        if ((this._strip(), this.length > 1)) e = 1;
        else {
          n && (t = -t), r(t <= 67108863, 'Number is too big');
          var i = 0 | this.words[0];
          e = i === t ? 0 : i < t ? -1 : 1;
        }
        return 0 !== this.negative ? 0 | -e : e;
      }),
      (i.prototype.cmp = function (t) {
        if (0 !== this.negative && 0 === t.negative) return -1;
        if (0 === this.negative && 0 !== t.negative) return 1;
        var e = this.ucmp(t);
        return 0 !== this.negative ? 0 | -e : e;
      }),
      (i.prototype.ucmp = function (t) {
        if (this.length > t.length) return 1;
        if (this.length < t.length) return -1;
        for (var e = 0, r = this.length - 1; r >= 0; r--) {
          var n = 0 | this.words[r],
            i = 0 | t.words[r];
          if (n !== i) {
            n < i ? (e = -1) : n > i && (e = 1);
            break;
          }
        }
        return e;
      }),
      (i.prototype.gtn = function (t) {
        return 1 === this.cmpn(t);
      }),
      (i.prototype.gt = function (t) {
        return 1 === this.cmp(t);
      }),
      (i.prototype.gten = function (t) {
        return this.cmpn(t) >= 0;
      }),
      (i.prototype.gte = function (t) {
        return this.cmp(t) >= 0;
      }),
      (i.prototype.ltn = function (t) {
        return -1 === this.cmpn(t);
      }),
      (i.prototype.lt = function (t) {
        return -1 === this.cmp(t);
      }),
      (i.prototype.lten = function (t) {
        return this.cmpn(t) <= 0;
      }),
      (i.prototype.lte = function (t) {
        return this.cmp(t) <= 0;
      }),
      (i.prototype.eqn = function (t) {
        return 0 === this.cmpn(t);
      }),
      (i.prototype.eq = function (t) {
        return 0 === this.cmp(t);
      }),
      (i.red = function (t) {
        return new _(t);
      }),
      (i.prototype.toRed = function (t) {
        return (
          r(!this.red, 'Already a number in reduction context'),
          r(0 === this.negative, 'red works only with positives'),
          t.convertTo(this)._forceRed(t)
        );
      }),
      (i.prototype.fromRed = function () {
        return (
          r(this.red, 'fromRed works only with numbers in reduction context'),
          this.red.convertFrom(this)
        );
      }),
      (i.prototype._forceRed = function (t) {
        return (this.red = t), this;
      }),
      (i.prototype.forceRed = function (t) {
        return (
          r(!this.red, 'Already a number in reduction context'),
          this._forceRed(t)
        );
      }),
      (i.prototype.redAdd = function (t) {
        return (
          r(this.red, 'redAdd works only with red numbers'),
          this.red.add(this, t)
        );
      }),
      (i.prototype.redIAdd = function (t) {
        return (
          r(this.red, 'redIAdd works only with red numbers'),
          this.red.iadd(this, t)
        );
      }),
      (i.prototype.redSub = function (t) {
        return (
          r(this.red, 'redSub works only with red numbers'),
          this.red.sub(this, t)
        );
      }),
      (i.prototype.redISub = function (t) {
        return (
          r(this.red, 'redISub works only with red numbers'),
          this.red.isub(this, t)
        );
      }),
      (i.prototype.redShl = function (t) {
        return (
          r(this.red, 'redShl works only with red numbers'),
          this.red.shl(this, t)
        );
      }),
      (i.prototype.redMul = function (t) {
        return (
          r(this.red, 'redMul works only with red numbers'),
          this.red._verify2(this, t),
          this.red.mul(this, t)
        );
      }),
      (i.prototype.redIMul = function (t) {
        return (
          r(this.red, 'redMul works only with red numbers'),
          this.red._verify2(this, t),
          this.red.imul(this, t)
        );
      }),
      (i.prototype.redSqr = function () {
        return (
          r(this.red, 'redSqr works only with red numbers'),
          this.red._verify1(this),
          this.red.sqr(this)
        );
      }),
      (i.prototype.redISqr = function () {
        return (
          r(this.red, 'redISqr works only with red numbers'),
          this.red._verify1(this),
          this.red.isqr(this)
        );
      }),
      (i.prototype.redSqrt = function () {
        return (
          r(this.red, 'redSqrt works only with red numbers'),
          this.red._verify1(this),
          this.red.sqrt(this)
        );
      }),
      (i.prototype.redInvm = function () {
        return (
          r(this.red, 'redInvm works only with red numbers'),
          this.red._verify1(this),
          this.red.invm(this)
        );
      }),
      (i.prototype.redNeg = function () {
        return (
          r(this.red, 'redNeg works only with red numbers'),
          this.red._verify1(this),
          this.red.neg(this)
        );
      }),
      (i.prototype.redPow = function (t) {
        return (
          r(this.red && !t.red, 'redPow(normalNum)'),
          this.red._verify1(this),
          this.red.pow(this, t)
        );
      });
    var v = { k256: null, p224: null, p192: null, p25519: null };
    function w(t, e) {
      (this.name = t),
        (this.p = new i(e, 16)),
        (this.n = this.p.bitLength()),
        (this.k = new i(1).iushln(this.n).isub(this.p)),
        (this.tmp = this._tmp());
    }
    function b() {
      w.call(
        this,
        'k256',
        'ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff fffffffe fffffc2f'
      );
    }
    function A() {
      w.call(
        this,
        'p224',
        'ffffffff ffffffff ffffffff ffffffff 00000000 00000000 00000001'
      );
    }
    function E() {
      w.call(
        this,
        'p192',
        'ffffffff ffffffff ffffffff fffffffe ffffffff ffffffff'
      );
    }
    function M() {
      w.call(
        this,
        '25519',
        '7fffffffffffffff ffffffffffffffff ffffffffffffffff ffffffffffffffed'
      );
    }
    function _(t) {
      if ('string' == typeof t) {
        var e = i._prime(t);
        (this.m = e.p), (this.prime = e);
      } else r(t.gtn(1), 'modulus must be greater than 1'), (this.m = t), (this.prime = null);
    }
    function k(t) {
      _.call(this, t),
        (this.shift = this.m.bitLength()),
        this.shift % 26 != 0 && (this.shift += 26 - (this.shift % 26)),
        (this.r = new i(1).iushln(this.shift)),
        (this.r2 = this.imod(this.r.sqr())),
        (this.rinv = this.r._invmp(this.m)),
        (this.minv = this.rinv.mul(this.r).isubn(1).div(this.m)),
        (this.minv = this.minv.umod(this.r)),
        (this.minv = this.r.sub(this.minv));
    }
    (w.prototype._tmp = function () {
      var t = new i(null);
      return (t.words = new Array(Math.ceil(this.n / 13))), t;
    }),
      (w.prototype.ireduce = function (t) {
        var e,
          r = t;
        do {
          this.split(r, this.tmp),
            (e = (r = (r = this.imulK(r)).iadd(this.tmp)).bitLength());
        } while (e > this.n);
        var n = e < this.n ? -1 : r.ucmp(this.p);
        return (
          0 === n
            ? ((r.words[0] = 0), (r.length = 1))
            : n > 0
            ? r.isub(this.p)
            : void 0 !== r.strip
            ? r.strip()
            : r._strip(),
          r
        );
      }),
      (w.prototype.split = function (t, e) {
        t.iushrn(this.n, 0, e);
      }),
      (w.prototype.imulK = function (t) {
        return t.imul(this.k);
      }),
      n(b, w),
      (b.prototype.split = function (t, e) {
        for (var r = 4194303, n = Math.min(t.length, 9), i = 0; i < n; i++)
          e.words[i] = t.words[i];
        if (((e.length = n), t.length <= 9))
          return (t.words[0] = 0), void (t.length = 1);
        var o = t.words[9];
        for (e.words[e.length++] = o & r, i = 10; i < t.length; i++) {
          var s = 0 | t.words[i];
          (t.words[i - 10] = ((s & r) << 4) | (o >>> 22)), (o = s);
        }
        (o >>>= 22),
          (t.words[i - 10] = o),
          0 === o && t.length > 10 ? (t.length -= 10) : (t.length -= 9);
      }),
      (b.prototype.imulK = function (t) {
        (t.words[t.length] = 0), (t.words[t.length + 1] = 0), (t.length += 2);
        for (var e = 0, r = 0; r < t.length; r++) {
          var n = 0 | t.words[r];
          (e += 977 * n),
            (t.words[r] = 67108863 & e),
            (e = 64 * n + ((e / 67108864) | 0));
        }
        return (
          0 === t.words[t.length - 1] &&
            (t.length--, 0 === t.words[t.length - 1] && t.length--),
          t
        );
      }),
      n(A, w),
      n(E, w),
      n(M, w),
      (M.prototype.imulK = function (t) {
        for (var e = 0, r = 0; r < t.length; r++) {
          var n = 19 * (0 | t.words[r]) + e,
            i = 67108863 & n;
          (n >>>= 26), (t.words[r] = i), (e = n);
        }
        return 0 !== e && (t.words[t.length++] = e), t;
      }),
      (i._prime = function (t) {
        if (v[t]) return v[t];
        var e;
        if ('k256' === t) e = new b();
        else if ('p224' === t) e = new A();
        else if ('p192' === t) e = new E();
        else {
          if ('p25519' !== t) throw new Error('Unknown prime ' + t);
          e = new M();
        }
        return (v[t] = e), e;
      }),
      (_.prototype._verify1 = function (t) {
        r(0 === t.negative, 'red works only with positives'),
          r(t.red, 'red works only with red numbers');
      }),
      (_.prototype._verify2 = function (t, e) {
        r(!(t.negative | e.negative), 'red works only with positives'),
          r(t.red && t.red === e.red, 'red works only with red numbers');
      }),
      (_.prototype.imod = function (t) {
        return this.prime
          ? this.prime.ireduce(t)._forceRed(this)
          : (h(t, t.umod(this.m)._forceRed(this)), t);
      }),
      (_.prototype.neg = function (t) {
        return t.isZero() ? t.clone() : this.m.sub(t)._forceRed(this);
      }),
      (_.prototype.add = function (t, e) {
        this._verify2(t, e);
        var r = t.add(e);
        return r.cmp(this.m) >= 0 && r.isub(this.m), r._forceRed(this);
      }),
      (_.prototype.iadd = function (t, e) {
        this._verify2(t, e);
        var r = t.iadd(e);
        return r.cmp(this.m) >= 0 && r.isub(this.m), r;
      }),
      (_.prototype.sub = function (t, e) {
        this._verify2(t, e);
        var r = t.sub(e);
        return r.cmpn(0) < 0 && r.iadd(this.m), r._forceRed(this);
      }),
      (_.prototype.isub = function (t, e) {
        this._verify2(t, e);
        var r = t.isub(e);
        return r.cmpn(0) < 0 && r.iadd(this.m), r;
      }),
      (_.prototype.shl = function (t, e) {
        return this._verify1(t), this.imod(t.ushln(e));
      }),
      (_.prototype.imul = function (t, e) {
        return this._verify2(t, e), this.imod(t.imul(e));
      }),
      (_.prototype.mul = function (t, e) {
        return this._verify2(t, e), this.imod(t.mul(e));
      }),
      (_.prototype.isqr = function (t) {
        return this.imul(t, t.clone());
      }),
      (_.prototype.sqr = function (t) {
        return this.mul(t, t);
      }),
      (_.prototype.sqrt = function (t) {
        if (t.isZero()) return t.clone();
        var e = this.m.andln(3);
        if ((r(e % 2 == 1), 3 === e)) {
          var n = this.m.add(new i(1)).iushrn(2);
          return this.pow(t, n);
        }
        for (var o = this.m.subn(1), s = 0; !o.isZero() && 0 === o.andln(1); )
          s++, o.iushrn(1);
        r(!o.isZero());
        var a = new i(1).toRed(this),
          u = a.redNeg(),
          h = this.m.subn(1).iushrn(1),
          l = this.m.bitLength();
        for (l = new i(2 * l * l).toRed(this); 0 !== this.pow(l, h).cmp(u); )
          l.redIAdd(u);
        for (
          var c = this.pow(l, o),
            f = this.pow(t, o.addn(1).iushrn(1)),
            d = this.pow(t, o),
            m = s;
          0 !== d.cmp(a);

        ) {
          for (var p = d, g = 0; 0 !== p.cmp(a); g++) p = p.redSqr();
          r(g < m);
          var y = this.pow(c, new i(1).iushln(m - g - 1));
          (f = f.redMul(y)), (c = y.redSqr()), (d = d.redMul(c)), (m = g);
        }
        return f;
      }),
      (_.prototype.invm = function (t) {
        var e = t._invmp(this.m);
        return 0 !== e.negative
          ? ((e.negative = 0), this.imod(e).redNeg())
          : this.imod(e);
      }),
      (_.prototype.pow = function (t, e) {
        if (e.isZero()) return new i(1).toRed(this);
        if (0 === e.cmpn(1)) return t.clone();
        var r = new Array(16);
        (r[0] = new i(1).toRed(this)), (r[1] = t);
        for (var n = 2; n < r.length; n++) r[n] = this.mul(r[n - 1], t);
        var o = r[0],
          s = 0,
          a = 0,
          u = e.bitLength() % 26;
        for (0 === u && (u = 26), n = e.length - 1; n >= 0; n--) {
          for (var h = e.words[n], l = u - 1; l >= 0; l--) {
            var c = (h >> l) & 1;
            o !== r[0] && (o = this.sqr(o)),
              0 !== c || 0 !== s
                ? ((s <<= 1),
                  (s |= c),
                  (4 == ++a || (0 === n && 0 === l)) &&
                    ((o = this.mul(o, r[s])), (a = 0), (s = 0)))
                : (a = 0);
          }
          u = 26;
        }
        return o;
      }),
      (_.prototype.convertTo = function (t) {
        var e = t.umod(this.m);
        return e === t ? e.clone() : e;
      }),
      (_.prototype.convertFrom = function (t) {
        var e = t.clone();
        return (e.red = null), e;
      }),
      (i.mont = function (t) {
        return new k(t);
      }),
      n(k, _),
      (k.prototype.convertTo = function (t) {
        return this.imod(t.ushln(this.shift));
      }),
      (k.prototype.convertFrom = function (t) {
        var e = this.imod(t.mul(this.rinv));
        return (e.red = null), e;
      }),
      (k.prototype.imul = function (t, e) {
        if (t.isZero() || e.isZero())
          return (t.words[0] = 0), (t.length = 1), t;
        var r = t.imul(e),
          n = r.maskn(this.shift).mul(this.minv).imaskn(this.shift).mul(this.m),
          i = r.isub(n).iushrn(this.shift),
          o = i;
        return (
          i.cmp(this.m) >= 0
            ? (o = i.isub(this.m))
            : i.cmpn(0) < 0 && (o = i.iadd(this.m)),
          o._forceRed(this)
        );
      }),
      (k.prototype.mul = function (t, e) {
        if (t.isZero() || e.isZero()) return new i(0)._forceRed(this);
        var r = t.mul(e),
          n = r.maskn(this.shift).mul(this.minv).imaskn(this.shift).mul(this.m),
          o = r.isub(n).iushrn(this.shift),
          s = o;
        return (
          o.cmp(this.m) >= 0
            ? (s = o.isub(this.m))
            : o.cmpn(0) < 0 && (s = o.iadd(this.m)),
          s._forceRed(this)
        );
      }),
      (k.prototype.invm = function (t) {
        return this.imod(t._invmp(this.m).mul(this.r2))._forceRed(this);
      });
  })(e, t);
  var r = e.exports;
  let n = !1,
    i = !1;
  const o = { debug: 1, default: 2, info: 2, warning: 3, error: 4, off: 5 };
  let s = o.default,
    a = null;
  const u = (function () {
    try {
      const t = [];
      if (
        (['NFD', 'NFC', 'NFKD', 'NFKC'].forEach((e) => {
          try {
            if ('test' !== 'test'.normalize(e))
              throw new Error('bad normalize');
          } catch (r) {
            t.push(e);
          }
        }),
        t.length)
      )
        throw new Error('missing ' + t.join(', '));
      if (
        String.fromCharCode(233).normalize('NFD') !==
        String.fromCharCode(101, 769)
      )
        throw new Error('broken implementation');
    } catch (t) {
      return t.message;
    }
    return null;
  })();
  var h, l;
  !(function (t) {
    (t.DEBUG = 'DEBUG'),
      (t.INFO = 'INFO'),
      (t.WARNING = 'WARNING'),
      (t.ERROR = 'ERROR'),
      (t.OFF = 'OFF');
  })(h || (h = {})),
    (function (t) {
      (t.UNKNOWN_ERROR = 'UNKNOWN_ERROR'),
        (t.NOT_IMPLEMENTED = 'NOT_IMPLEMENTED'),
        (t.UNSUPPORTED_OPERATION = 'UNSUPPORTED_OPERATION'),
        (t.NETWORK_ERROR = 'NETWORK_ERROR'),
        (t.SERVER_ERROR = 'SERVER_ERROR'),
        (t.TIMEOUT = 'TIMEOUT'),
        (t.BUFFER_OVERRUN = 'BUFFER_OVERRUN'),
        (t.NUMERIC_FAULT = 'NUMERIC_FAULT'),
        (t.MISSING_NEW = 'MISSING_NEW'),
        (t.INVALID_ARGUMENT = 'INVALID_ARGUMENT'),
        (t.MISSING_ARGUMENT = 'MISSING_ARGUMENT'),
        (t.UNEXPECTED_ARGUMENT = 'UNEXPECTED_ARGUMENT'),
        (t.CALL_EXCEPTION = 'CALL_EXCEPTION'),
        (t.INSUFFICIENT_FUNDS = 'INSUFFICIENT_FUNDS'),
        (t.NONCE_EXPIRED = 'NONCE_EXPIRED'),
        (t.REPLACEMENT_UNDERPRICED = 'REPLACEMENT_UNDERPRICED'),
        (t.UNPREDICTABLE_GAS_LIMIT = 'UNPREDICTABLE_GAS_LIMIT'),
        (t.TRANSACTION_REPLACED = 'TRANSACTION_REPLACED'),
        (t.ACTION_REJECTED = 'ACTION_REJECTED');
    })(l || (l = {}));
  const c = '0123456789abcdef';
  class f {
    constructor(t) {
      Object.defineProperty(this, 'version', {
        enumerable: !0,
        value: t,
        writable: !1,
      });
    }
    _log(t, e) {
      const r = t.toLowerCase();
      null == o[r] &&
        this.throwArgumentError('invalid log level name', 'logLevel', t),
        s > o[r] || console.log.apply(console, e);
    }
    debug(...t) {
      this._log(f.levels.DEBUG, t);
    }
    info(...t) {
      this._log(f.levels.INFO, t);
    }
    warn(...t) {
      this._log(f.levels.WARNING, t);
    }
    makeError(t, e, r) {
      if (i) return this.makeError('censored error', e, {});
      e || (e = f.errors.UNKNOWN_ERROR), r || (r = {});
      const n = [];
      Object.keys(r).forEach((t) => {
        const e = r[t];
        try {
          if (e instanceof Uint8Array) {
            let r = '';
            for (let t = 0; t < e.length; t++)
              (r += c[e[t] >> 4]), (r += c[15 & e[t]]);
            n.push(t + '=Uint8Array(0x' + r + ')');
          } else n.push(t + '=' + JSON.stringify(e));
        } catch (e) {
          n.push(t + '=' + JSON.stringify(r[t].toString()));
        }
      }),
        n.push(`code=${e}`),
        n.push(`version=${this.version}`);
      const o = t;
      let s = '';
      switch (e) {
        case l.NUMERIC_FAULT: {
          s = 'NUMERIC_FAULT';
          const e = t;
          switch (e) {
            case 'overflow':
            case 'underflow':
            case 'division-by-zero':
              s += '-' + e;
              break;
            case 'negative-power':
            case 'negative-width':
              s += '-unsupported';
              break;
            case 'unbound-bitwise-result':
              s += '-unbound-result';
          }
          break;
        }
        case l.CALL_EXCEPTION:
        case l.INSUFFICIENT_FUNDS:
        case l.MISSING_NEW:
        case l.NONCE_EXPIRED:
        case l.REPLACEMENT_UNDERPRICED:
        case l.TRANSACTION_REPLACED:
        case l.UNPREDICTABLE_GAS_LIMIT:
          s = e;
      }
      s && (t += ' [ See: https://links.ethers.org/v5-errors-' + s + ' ]'),
        n.length && (t += ' (' + n.join(', ') + ')');
      const a = new Error(t);
      return (
        (a.reason = o),
        (a.code = e),
        Object.keys(r).forEach(function (t) {
          a[t] = r[t];
        }),
        a
      );
    }
    throwError(t, e, r) {
      throw this.makeError(t, e, r);
    }
    throwArgumentError(t, e, r) {
      return this.throwError(t, f.errors.INVALID_ARGUMENT, {
        argument: e,
        value: r,
      });
    }
    assert(t, e, r, n) {
      t || this.throwError(e, r, n);
    }
    assertArgument(t, e, r, n) {
      t || this.throwArgumentError(e, r, n);
    }
    checkNormalize(t) {
      u &&
        this.throwError(
          'platform missing String.prototype.normalize',
          f.errors.UNSUPPORTED_OPERATION,
          { operation: 'String.prototype.normalize', form: u }
        );
    }
    checkSafeUint53(t, e) {
      'number' == typeof t &&
        (null == e && (e = 'value not safe'),
        (t < 0 || t >= 9007199254740991) &&
          this.throwError(e, f.errors.NUMERIC_FAULT, {
            operation: 'checkSafeInteger',
            fault: 'out-of-safe-range',
            value: t,
          }),
        t % 1 &&
          this.throwError(e, f.errors.NUMERIC_FAULT, {
            operation: 'checkSafeInteger',
            fault: 'non-integer',
            value: t,
          }));
    }
    checkArgumentCount(t, e, r) {
      (r = r ? ': ' + r : ''),
        t < e &&
          this.throwError('missing argument' + r, f.errors.MISSING_ARGUMENT, {
            count: t,
            expectedCount: e,
          }),
        t > e &&
          this.throwError(
            'too many arguments' + r,
            f.errors.UNEXPECTED_ARGUMENT,
            { count: t, expectedCount: e }
          );
    }
    checkNew(t, e) {
      (t !== Object && null != t) ||
        this.throwError('missing new', f.errors.MISSING_NEW, { name: e.name });
    }
    checkAbstract(t, e) {
      t === e
        ? this.throwError(
            'cannot instantiate abstract class ' +
              JSON.stringify(e.name) +
              ' directly; use a sub-class',
            f.errors.UNSUPPORTED_OPERATION,
            { name: t.name, operation: 'new' }
          )
        : (t !== Object && null != t) ||
          this.throwError('missing new', f.errors.MISSING_NEW, {
            name: e.name,
          });
    }
    static globalLogger() {
      return a || (a = new f('logger/5.7.0')), a;
    }
    static setCensorship(t, e) {
      if (
        (!t &&
          e &&
          this.globalLogger().throwError(
            'cannot permanently disable censorship',
            f.errors.UNSUPPORTED_OPERATION,
            { operation: 'setCensorship' }
          ),
        n)
      ) {
        if (!t) return;
        this.globalLogger().throwError(
          'error censorship permanent',
          f.errors.UNSUPPORTED_OPERATION,
          { operation: 'setCensorship' }
        );
      }
      (i = !!t), (n = !!e);
    }
    static setLogLevel(t) {
      const e = o[t.toLowerCase()];
      null != e ? (s = e) : f.globalLogger().warn('invalid log level - ' + t);
    }
    static from(t) {
      return new f(t);
    }
  }
  (f.errors = l), (f.levels = h);
  const d = new f('bytes/5.7.0');
  function m(t) {
    return !!t.toHexString;
  }
  function p(t) {
    return (
      t.slice ||
        (t.slice = function () {
          const e = Array.prototype.slice.call(arguments);
          return p(new Uint8Array(Array.prototype.slice.apply(t, e)));
        }),
      t
    );
  }
  function g(t) {
    return (M(t) && !(t.length % 2)) || v(t);
  }
  function y(t) {
    return 'number' == typeof t && t == t && t % 1 == 0;
  }
  function v(t) {
    if (null == t) return !1;
    if (t.constructor === Uint8Array) return !0;
    if ('string' == typeof t) return !1;
    if (!y(t.length) || t.length < 0) return !1;
    for (let e = 0; e < t.length; e++) {
      const r = t[e];
      if (!y(r) || r < 0 || r >= 256) return !1;
    }
    return !0;
  }
  function w(t, e) {
    if ((e || (e = {}), 'number' == typeof t)) {
      d.checkSafeUint53(t, 'invalid arrayify value');
      const e = [];
      for (; t; ) e.unshift(255 & t), (t = parseInt(String(t / 256)));
      return 0 === e.length && e.push(0), p(new Uint8Array(e));
    }
    if (
      (e.allowMissingPrefix &&
        'string' == typeof t &&
        '0x' !== t.substring(0, 2) &&
        (t = '0x' + t),
      m(t) && (t = t.toHexString()),
      M(t))
    ) {
      let r = t.substring(2);
      r.length % 2 &&
        ('left' === e.hexPad
          ? (r = '0' + r)
          : 'right' === e.hexPad
          ? (r += '0')
          : d.throwArgumentError('hex data is odd-length', 'value', t));
      const n = [];
      for (let t = 0; t < r.length; t += 2)
        n.push(parseInt(r.substring(t, t + 2), 16));
      return p(new Uint8Array(n));
    }
    return v(t)
      ? p(new Uint8Array(t))
      : d.throwArgumentError('invalid arrayify value', 'value', t);
  }
  function b(t) {
    const e = t.map((t) => w(t)),
      r = e.reduce((t, e) => t + e.length, 0),
      n = new Uint8Array(r);
    return e.reduce((t, e) => (n.set(e, t), t + e.length), 0), p(n);
  }
  function A(t) {
    let e = w(t);
    if (0 === e.length) return e;
    let r = 0;
    for (; r < e.length && 0 === e[r]; ) r++;
    return r && (e = e.slice(r)), e;
  }
  function E(t, e) {
    (t = w(t)).length > e &&
      d.throwArgumentError('value out of range', 'value', arguments[0]);
    const r = new Uint8Array(e);
    return r.set(t, e - t.length), p(r);
  }
  function M(t, e) {
    return (
      !('string' != typeof t || !t.match(/^0x[0-9A-Fa-f]*$/)) &&
      (!e || t.length === 2 + 2 * e)
    );
  }
  const _ = '0123456789abcdef';
  function k(t, e) {
    if ((e || (e = {}), 'number' == typeof t)) {
      d.checkSafeUint53(t, 'invalid hexlify value');
      let e = '';
      for (; t; ) (e = _[15 & t] + e), (t = Math.floor(t / 16));
      return e.length ? (e.length % 2 && (e = '0' + e), '0x' + e) : '0x00';
    }
    if ('bigint' == typeof t)
      return (t = t.toString(16)).length % 2 ? '0x0' + t : '0x' + t;
    if (
      (e.allowMissingPrefix &&
        'string' == typeof t &&
        '0x' !== t.substring(0, 2) &&
        (t = '0x' + t),
      m(t))
    )
      return t.toHexString();
    if (M(t))
      return (
        t.length % 2 &&
          ('left' === e.hexPad
            ? (t = '0x0' + t.substring(2))
            : 'right' === e.hexPad
            ? (t += '0')
            : d.throwArgumentError('hex data is odd-length', 'value', t)),
        t.toLowerCase()
      );
    if (v(t)) {
      let e = '0x';
      for (let r = 0; r < t.length; r++) {
        let n = t[r];
        e += _[(240 & n) >> 4] + _[15 & n];
      }
      return e;
    }
    return d.throwArgumentError('invalid hexlify value', 'value', t);
  }
  function S(t) {
    if ('string' != typeof t) t = k(t);
    else if (!M(t) || t.length % 2) return null;
    return (t.length - 2) / 2;
  }
  function N(t, e, r) {
    return (
      'string' != typeof t
        ? (t = k(t))
        : (!M(t) || t.length % 2) &&
          d.throwArgumentError('invalid hexData', 'value', t),
      (e = 2 + 2 * e),
      null != r ? '0x' + t.substring(e, 2 + 2 * r) : '0x' + t.substring(e)
    );
  }
  function P(t) {
    let e = '0x';
    return (
      t.forEach((t) => {
        e += k(t).substring(2);
      }),
      e
    );
  }
  function x(t) {
    const e = T(k(t, { hexPad: 'left' }));
    return '0x' === e ? '0x0' : e;
  }
  function T(t) {
    'string' != typeof t && (t = k(t)),
      M(t) || d.throwArgumentError('invalid hex string', 'value', t),
      (t = t.substring(2));
    let e = 0;
    for (; e < t.length && '0' === t[e]; ) e++;
    return '0x' + t.substring(e);
  }
  function C(t, e) {
    for (
      'string' != typeof t
        ? (t = k(t))
        : M(t) || d.throwArgumentError('invalid hex string', 'value', t),
        t.length > 2 * e + 2 &&
          d.throwArgumentError('value out of range', 'value', arguments[1]);
      t.length < 2 * e + 2;

    )
      t = '0x0' + t.substring(2);
    return t;
  }
  function R(t) {
    const e = {
      r: '0x',
      s: '0x',
      _vs: '0x',
      recoveryParam: 0,
      v: 0,
      yParityAndS: '0x',
      compact: '0x',
    };
    if (g(t)) {
      let r = w(t);
      64 === r.length
        ? ((e.v = 27 + (r[32] >> 7)),
          (r[32] &= 127),
          (e.r = k(r.slice(0, 32))),
          (e.s = k(r.slice(32, 64))))
        : 65 === r.length
        ? ((e.r = k(r.slice(0, 32))), (e.s = k(r.slice(32, 64))), (e.v = r[64]))
        : d.throwArgumentError('invalid signature string', 'signature', t),
        e.v < 27 &&
          (0 === e.v || 1 === e.v
            ? (e.v += 27)
            : d.throwArgumentError('signature invalid v byte', 'signature', t)),
        (e.recoveryParam = 1 - (e.v % 2)),
        e.recoveryParam && (r[32] |= 128),
        (e._vs = k(r.slice(32, 64)));
    } else {
      if (
        ((e.r = t.r),
        (e.s = t.s),
        (e.v = t.v),
        (e.recoveryParam = t.recoveryParam),
        (e._vs = t._vs),
        null != e._vs)
      ) {
        const r = E(w(e._vs), 32);
        e._vs = k(r);
        const n = r[0] >= 128 ? 1 : 0;
        null == e.recoveryParam
          ? (e.recoveryParam = n)
          : e.recoveryParam !== n &&
            d.throwArgumentError(
              'signature recoveryParam mismatch _vs',
              'signature',
              t
            ),
          (r[0] &= 127);
        const i = k(r);
        null == e.s
          ? (e.s = i)
          : e.s !== i &&
            d.throwArgumentError('signature v mismatch _vs', 'signature', t);
      }
      if (null == e.recoveryParam)
        null == e.v
          ? d.throwArgumentError(
              'signature missing v and recoveryParam',
              'signature',
              t
            )
          : 0 === e.v || 1 === e.v
          ? (e.recoveryParam = e.v)
          : (e.recoveryParam = 1 - (e.v % 2));
      else if (null == e.v) e.v = 27 + e.recoveryParam;
      else {
        const r = 0 === e.v || 1 === e.v ? e.v : 1 - (e.v % 2);
        e.recoveryParam !== r &&
          d.throwArgumentError(
            'signature recoveryParam mismatch v',
            'signature',
            t
          );
      }
      null != e.r && M(e.r)
        ? (e.r = C(e.r, 32))
        : d.throwArgumentError(
            'signature missing or invalid r',
            'signature',
            t
          ),
        null != e.s && M(e.s)
          ? (e.s = C(e.s, 32))
          : d.throwArgumentError(
              'signature missing or invalid s',
              'signature',
              t
            );
      const r = w(e.s);
      r[0] >= 128 &&
        d.throwArgumentError('signature s out of range', 'signature', t),
        e.recoveryParam && (r[0] |= 128);
      const n = k(r);
      e._vs &&
        (M(e._vs) ||
          d.throwArgumentError('signature invalid _vs', 'signature', t),
        (e._vs = C(e._vs, 32))),
        null == e._vs
          ? (e._vs = n)
          : e._vs !== n &&
            d.throwArgumentError(
              'signature _vs mismatch v and s',
              'signature',
              t
            );
    }
    return (
      (e.yParityAndS = e._vs), (e.compact = e.r + e.yParityAndS.substring(2)), e
    );
  }
  function I(t) {
    return k(b([(t = R(t)).r, t.s, t.recoveryParam ? '0x1c' : '0x1b']));
  }
  const B = 'bignumber/5.7.0';
  var O = r.BN;
  const F = new f(B),
    D = {},
    L = 9007199254740991;
  let U = !1;
  class j {
    constructor(t, e) {
      t !== D &&
        F.throwError(
          'cannot call constructor directly; use BigNumber.from',
          f.errors.UNSUPPORTED_OPERATION,
          { operation: 'new (BigNumber)' }
        ),
        (this._hex = e),
        (this._isBigNumber = !0),
        Object.freeze(this);
    }
    fromTwos(t) {
      return G(q(this).fromTwos(t));
    }
    toTwos(t) {
      return G(q(this).toTwos(t));
    }
    abs() {
      return '-' === this._hex[0] ? j.from(this._hex.substring(1)) : this;
    }
    add(t) {
      return G(q(this).add(q(t)));
    }
    sub(t) {
      return G(q(this).sub(q(t)));
    }
    div(t) {
      return (
        j.from(t).isZero() && K('division-by-zero', 'div'), G(q(this).div(q(t)))
      );
    }
    mul(t) {
      return G(q(this).mul(q(t)));
    }
    mod(t) {
      const e = q(t);
      return e.isNeg() && K('division-by-zero', 'mod'), G(q(this).umod(e));
    }
    pow(t) {
      const e = q(t);
      return e.isNeg() && K('negative-power', 'pow'), G(q(this).pow(e));
    }
    and(t) {
      const e = q(t);
      return (
        (this.isNegative() || e.isNeg()) && K('unbound-bitwise-result', 'and'),
        G(q(this).and(e))
      );
    }
    or(t) {
      const e = q(t);
      return (
        (this.isNegative() || e.isNeg()) && K('unbound-bitwise-result', 'or'),
        G(q(this).or(e))
      );
    }
    xor(t) {
      const e = q(t);
      return (
        (this.isNegative() || e.isNeg()) && K('unbound-bitwise-result', 'xor'),
        G(q(this).xor(e))
      );
    }
    mask(t) {
      return (
        (this.isNegative() || t < 0) && K('negative-width', 'mask'),
        G(q(this).maskn(t))
      );
    }
    shl(t) {
      return (
        (this.isNegative() || t < 0) && K('negative-width', 'shl'),
        G(q(this).shln(t))
      );
    }
    shr(t) {
      return (
        (this.isNegative() || t < 0) && K('negative-width', 'shr'),
        G(q(this).shrn(t))
      );
    }
    eq(t) {
      return q(this).eq(q(t));
    }
    lt(t) {
      return q(this).lt(q(t));
    }
    lte(t) {
      return q(this).lte(q(t));
    }
    gt(t) {
      return q(this).gt(q(t));
    }
    gte(t) {
      return q(this).gte(q(t));
    }
    isNegative() {
      return '-' === this._hex[0];
    }
    isZero() {
      return q(this).isZero();
    }
    toNumber() {
      try {
        return q(this).toNumber();
      } catch (t) {
        K('overflow', 'toNumber', this.toString());
      }
      return null;
    }
    toBigInt() {
      try {
        return BigInt(this.toString());
      } catch (t) {}
      return F.throwError(
        'this platform does not support BigInt',
        f.errors.UNSUPPORTED_OPERATION,
        { value: this.toString() }
      );
    }
    toString() {
      return (
        arguments.length > 0 &&
          (10 === arguments[0]
            ? U ||
              ((U = !0),
              F.warn(
                'BigNumber.toString does not accept any parameters; base-10 is assumed'
              ))
            : 16 === arguments[0]
            ? F.throwError(
                'BigNumber.toString does not accept any parameters; use bigNumber.toHexString()',
                f.errors.UNEXPECTED_ARGUMENT,
                {}
              )
            : F.throwError(
                'BigNumber.toString does not accept parameters',
                f.errors.UNEXPECTED_ARGUMENT,
                {}
              )),
        q(this).toString(10)
      );
    }
    toHexString() {
      return this._hex;
    }
    toJSON(t) {
      return { type: 'BigNumber', hex: this.toHexString() };
    }
    static from(t) {
      if (t instanceof j) return t;
      if ('string' == typeof t)
        return t.match(/^-?0x[0-9a-f]+$/i)
          ? new j(D, z(t))
          : t.match(/^-?[0-9]+$/)
          ? new j(D, z(new O(t)))
          : F.throwArgumentError('invalid BigNumber string', 'value', t);
      if ('number' == typeof t)
        return (
          t % 1 && K('underflow', 'BigNumber.from', t),
          (t >= L || t <= -L) && K('overflow', 'BigNumber.from', t),
          j.from(String(t))
        );
      const e = t;
      if ('bigint' == typeof e) return j.from(e.toString());
      if (v(e)) return j.from(k(e));
      if (e)
        if (e.toHexString) {
          const t = e.toHexString();
          if ('string' == typeof t) return j.from(t);
        } else {
          let t = e._hex;
          if (
            (null == t && 'BigNumber' === e.type && (t = e.hex),
            'string' == typeof t &&
              (M(t) || ('-' === t[0] && M(t.substring(1)))))
          )
            return j.from(t);
        }
      return F.throwArgumentError('invalid BigNumber value', 'value', t);
    }
    static isBigNumber(t) {
      return !(!t || !t._isBigNumber);
    }
  }
  function z(t) {
    if ('string' != typeof t) return z(t.toString(16));
    if ('-' === t[0])
      return (
        '-' === (t = t.substring(1))[0] &&
          F.throwArgumentError('invalid hex', 'value', t),
        '0x00' === (t = z(t)) ? t : '-' + t
      );
    if (('0x' !== t.substring(0, 2) && (t = '0x' + t), '0x' === t))
      return '0x00';
    for (
      t.length % 2 && (t = '0x0' + t.substring(2));
      t.length > 4 && '0x00' === t.substring(0, 4);

    )
      t = '0x' + t.substring(4);
    return t;
  }
  function G(t) {
    return j.from(z(t));
  }
  function q(t) {
    const e = j.from(t).toHexString();
    return '-' === e[0]
      ? new O('-' + e.substring(3), 16)
      : new O(e.substring(2), 16);
  }
  function K(t, e, r) {
    const n = { fault: t, operation: e };
    return (
      null != r && (n.value = r), F.throwError(t, f.errors.NUMERIC_FAULT, n)
    );
  }
  const H = new f(B),
    V = {},
    J = j.from(0),
    W = j.from(-1);
  function Q(t, e, r, n) {
    const i = { fault: e, operation: r };
    return (
      void 0 !== n && (i.value = n), H.throwError(t, f.errors.NUMERIC_FAULT, i)
    );
  }
  let Y = '0';
  for (; Y.length < 256; ) Y += Y;
  function Z(t) {
    if ('number' != typeof t)
      try {
        t = j.from(t).toNumber();
      } catch (t) {}
    return 'number' == typeof t && t >= 0 && t <= 256 && !(t % 1)
      ? '1' + Y.substring(0, t)
      : H.throwArgumentError('invalid decimal size', 'decimals', t);
  }
  function X(t, e) {
    null == e && (e = 0);
    const r = Z(e),
      n = (t = j.from(t)).lt(J);
    n && (t = t.mul(W));
    let i = t.mod(r).toString();
    for (; i.length < r.length - 1; ) i = '0' + i;
    i = i.match(/^([0-9]*[1-9]|0)(0*)/)[1];
    const o = t.div(r).toString();
    return (t = 1 === r.length ? o : o + '.' + i), n && (t = '-' + t), t;
  }
  function $(t, e) {
    null == e && (e = 0);
    const r = Z(e);
    ('string' == typeof t && t.match(/^-?[0-9.]+$/)) ||
      H.throwArgumentError('invalid decimal value', 'value', t);
    const n = '-' === t.substring(0, 1);
    n && (t = t.substring(1)),
      '.' === t && H.throwArgumentError('missing value', 'value', t);
    const i = t.split('.');
    i.length > 2 && H.throwArgumentError('too many decimal points', 'value', t);
    let o = i[0],
      s = i[1];
    for (o || (o = '0'), s || (s = '0'); '0' === s[s.length - 1]; )
      s = s.substring(0, s.length - 1);
    for (
      s.length > r.length - 1 &&
        Q('fractional component exceeds decimals', 'underflow', 'parseFixed'),
        '' === s && (s = '0');
      s.length < r.length - 1;

    )
      s += '0';
    const a = j.from(o),
      u = j.from(s);
    let h = a.mul(r).add(u);
    return n && (h = h.mul(W)), h;
  }
  class tt {
    constructor(t, e, r, n) {
      t !== V &&
        H.throwError(
          'cannot use FixedFormat constructor; use FixedFormat.from',
          f.errors.UNSUPPORTED_OPERATION,
          { operation: 'new FixedFormat' }
        ),
        (this.signed = e),
        (this.width = r),
        (this.decimals = n),
        (this.name = (e ? '' : 'u') + 'fixed' + String(r) + 'x' + String(n)),
        (this._multiplier = Z(n)),
        Object.freeze(this);
    }
    static from(t) {
      if (t instanceof tt) return t;
      'number' == typeof t && (t = `fixed128x${t}`);
      let e = !0,
        r = 128,
        n = 18;
      if ('string' == typeof t)
        if ('fixed' === t);
        else if ('ufixed' === t) e = !1;
        else {
          const i = t.match(/^(u?)fixed([0-9]+)x([0-9]+)$/);
          i || H.throwArgumentError('invalid fixed format', 'format', t),
            (e = 'u' !== i[1]),
            (r = parseInt(i[2])),
            (n = parseInt(i[3]));
        }
      else if (t) {
        const i = (e, r, n) =>
          null == t[e]
            ? n
            : (typeof t[e] !== r &&
                H.throwArgumentError(
                  'invalid fixed format (' + e + ' not ' + r + ')',
                  'format.' + e,
                  t[e]
                ),
              t[e]);
        (e = i('signed', 'boolean', e)),
          (r = i('width', 'number', r)),
          (n = i('decimals', 'number', n));
      }
      return (
        r % 8 &&
          H.throwArgumentError(
            'invalid fixed format width (not byte aligned)',
            'format.width',
            r
          ),
        n > 80 &&
          H.throwArgumentError(
            'invalid fixed format (decimals too large)',
            'format.decimals',
            n
          ),
        new tt(V, e, r, n)
      );
    }
  }
  class et {
    constructor(t, e, r, n) {
      t !== V &&
        H.throwError(
          'cannot use FixedNumber constructor; use FixedNumber.from',
          f.errors.UNSUPPORTED_OPERATION,
          { operation: 'new FixedFormat' }
        ),
        (this.format = n),
        (this._hex = e),
        (this._value = r),
        (this._isFixedNumber = !0),
        Object.freeze(this);
    }
    _checkFormat(t) {
      this.format.name !== t.format.name &&
        H.throwArgumentError(
          'incompatible format; use fixedNumber.toFormat',
          'other',
          t
        );
    }
    addUnsafe(t) {
      this._checkFormat(t);
      const e = $(this._value, this.format.decimals),
        r = $(t._value, t.format.decimals);
      return et.fromValue(e.add(r), this.format.decimals, this.format);
    }
    subUnsafe(t) {
      this._checkFormat(t);
      const e = $(this._value, this.format.decimals),
        r = $(t._value, t.format.decimals);
      return et.fromValue(e.sub(r), this.format.decimals, this.format);
    }
    mulUnsafe(t) {
      this._checkFormat(t);
      const e = $(this._value, this.format.decimals),
        r = $(t._value, t.format.decimals);
      return et.fromValue(
        e.mul(r).div(this.format._multiplier),
        this.format.decimals,
        this.format
      );
    }
    divUnsafe(t) {
      this._checkFormat(t);
      const e = $(this._value, this.format.decimals),
        r = $(t._value, t.format.decimals);
      return et.fromValue(
        e.mul(this.format._multiplier).div(r),
        this.format.decimals,
        this.format
      );
    }
    floor() {
      const t = this.toString().split('.');
      1 === t.length && t.push('0');
      let e = et.from(t[0], this.format);
      const r = !t[1].match(/^(0*)$/);
      return (
        this.isNegative() && r && (e = e.subUnsafe(rt.toFormat(e.format))), e
      );
    }
    ceiling() {
      const t = this.toString().split('.');
      1 === t.length && t.push('0');
      let e = et.from(t[0], this.format);
      const r = !t[1].match(/^(0*)$/);
      return (
        !this.isNegative() && r && (e = e.addUnsafe(rt.toFormat(e.format))), e
      );
    }
    round(t) {
      null == t && (t = 0);
      const e = this.toString().split('.');
      if (
        (1 === e.length && e.push('0'),
        (t < 0 || t > 80 || t % 1) &&
          H.throwArgumentError('invalid decimal count', 'decimals', t),
        e[1].length <= t)
      )
        return this;
      const r = et.from('1' + Y.substring(0, t), this.format),
        n = nt.toFormat(this.format);
      return this.mulUnsafe(r).addUnsafe(n).floor().divUnsafe(r);
    }
    isZero() {
      return '0.0' === this._value || '0' === this._value;
    }
    isNegative() {
      return '-' === this._value[0];
    }
    toString() {
      return this._value;
    }
    toHexString(t) {
      if (null == t) return this._hex;
      t % 8 && H.throwArgumentError('invalid byte width', 'width', t);
      return C(
        j.from(this._hex).fromTwos(this.format.width).toTwos(t).toHexString(),
        t / 8
      );
    }
    toUnsafeFloat() {
      return parseFloat(this.toString());
    }
    toFormat(t) {
      return et.fromString(this._value, t);
    }
    static fromValue(t, e, r) {
      return (
        null != r ||
          null == e ||
          (function (t) {
            return (
              null != t &&
              (j.isBigNumber(t) ||
                ('number' == typeof t && t % 1 == 0) ||
                ('string' == typeof t && !!t.match(/^-?[0-9]+$/)) ||
                M(t) ||
                'bigint' == typeof t ||
                v(t))
            );
          })(e) ||
          ((r = e), (e = null)),
        null == e && (e = 0),
        null == r && (r = 'fixed'),
        et.fromString(X(t, e), tt.from(r))
      );
    }
    static fromString(t, e) {
      null == e && (e = 'fixed');
      const r = tt.from(e),
        n = $(t, r.decimals);
      !r.signed &&
        n.lt(J) &&
        Q('unsigned value cannot be negative', 'overflow', 'value', t);
      let i = null;
      r.signed
        ? (i = n.toTwos(r.width).toHexString())
        : ((i = n.toHexString()), (i = C(i, r.width / 8)));
      const o = X(n, r.decimals);
      return new et(V, i, o, r);
    }
    static fromBytes(t, e) {
      null == e && (e = 'fixed');
      const r = tt.from(e);
      if (w(t).length > r.width / 8) throw new Error('overflow');
      let n = j.from(t);
      r.signed && (n = n.fromTwos(r.width));
      const i = n.toTwos((r.signed ? 0 : 1) + r.width).toHexString(),
        o = X(n, r.decimals);
      return new et(V, i, o, r);
    }
    static from(t, e) {
      if ('string' == typeof t) return et.fromString(t, e);
      if (v(t)) return et.fromBytes(t, e);
      try {
        return et.fromValue(t, 0, e);
      } catch (t) {
        if (t.code !== f.errors.INVALID_ARGUMENT) throw t;
      }
      return H.throwArgumentError('invalid FixedNumber value', 'value', t);
    }
    static isFixedNumber(t) {
      return !(!t || !t._isFixedNumber);
    }
  }
  const rt = et.from(1),
    nt = et.from('0.5');
  var it = function (t, e, r, n) {
    return new (r || (r = Promise))(function (i, o) {
      function s(t) {
        try {
          u(n.next(t));
        } catch (t) {
          o(t);
        }
      }
      function a(t) {
        try {
          u(n.throw(t));
        } catch (t) {
          o(t);
        }
      }
      function u(t) {
        var e;
        t.done
          ? i(t.value)
          : ((e = t.value),
            e instanceof r
              ? e
              : new r(function (t) {
                  t(e);
                })).then(s, a);
      }
      u((n = n.apply(t, e || [])).next());
    });
  };
  const ot = new f('properties/5.7.0');
  function st(t, e, r) {
    Object.defineProperty(t, e, { enumerable: !0, value: r, writable: !1 });
  }
  function at(t, e) {
    for (let r = 0; r < 32; r++) {
      if (t[e]) return t[e];
      if (!t.prototype || 'object' != typeof t.prototype) break;
      t = Object.getPrototypeOf(t.prototype).constructor;
    }
    return null;
  }
  function ut(t) {
    return it(this, void 0, void 0, function* () {
      const e = Object.keys(t).map((e) => {
        const r = t[e];
        return Promise.resolve(r).then((t) => ({ key: e, value: t }));
      });
      return (yield Promise.all(e)).reduce(
        (t, e) => ((t[e.key] = e.value), t),
        {}
      );
    });
  }
  function ht(t, e) {
    (t && 'object' == typeof t) ||
      ot.throwArgumentError('invalid object', 'object', t),
      Object.keys(t).forEach((r) => {
        e[r] ||
          ot.throwArgumentError(
            'invalid object key - ' + r,
            'transaction:' + r,
            t
          );
      });
  }
  function lt(t) {
    const e = {};
    for (const r in t) e[r] = t[r];
    return e;
  }
  const ct = { bigint: !0, boolean: !0, function: !0, number: !0, string: !0 };
  function ft(t) {
    if (null == t || ct[typeof t]) return !0;
    if (Array.isArray(t) || 'object' == typeof t) {
      if (!Object.isFrozen(t)) return !1;
      const e = Object.keys(t);
      for (let r = 0; r < e.length; r++) {
        let n = null;
        try {
          n = t[e[r]];
        } catch (t) {
          continue;
        }
        if (!ft(n)) return !1;
      }
      return !0;
    }
    return ot.throwArgumentError('Cannot deepCopy ' + typeof t, 'object', t);
  }
  function dt(t) {
    if (ft(t)) return t;
    if (Array.isArray(t)) return Object.freeze(t.map((t) => mt(t)));
    if ('object' == typeof t) {
      const e = {};
      for (const r in t) {
        const n = t[r];
        void 0 !== n && st(e, r, mt(n));
      }
      return e;
    }
    return ot.throwArgumentError('Cannot deepCopy ' + typeof t, 'object', t);
  }
  function mt(t) {
    return dt(t);
  }
  class pt {
    constructor(t) {
      for (const e in t) this[e] = mt(t[e]);
    }
  }
  const gt = 'abi/5.7.0',
    yt = new f(gt),
    vt = {};
  let wt = { calldata: !0, memory: !0, storage: !0 },
    bt = { calldata: !0, memory: !0 };
  function At(t, e) {
    if ('bytes' === t || 'string' === t) {
      if (wt[e]) return !0;
    } else if ('address' === t) {
      if ('payable' === e) return !0;
    } else if ((t.indexOf('[') >= 0 || 'tuple' === t) && bt[e]) return !0;
    return (
      (wt[e] || 'payable' === e) &&
        yt.throwArgumentError('invalid modifier', 'name', e),
      !1
    );
  }
  function Et(t, e) {
    for (let r in e) st(t, r, e[r]);
  }
  const Mt = Object.freeze({
      sighash: 'sighash',
      minimal: 'minimal',
      full: 'full',
      json: 'json',
    }),
    _t = new RegExp(/^(.*)\[([0-9]*)\]$/);
  class kt {
    constructor(t, e) {
      t !== vt &&
        yt.throwError('use fromString', f.errors.UNSUPPORTED_OPERATION, {
          operation: 'new ParamType()',
        }),
        Et(this, e);
      let r = this.type.match(_t);
      Et(
        this,
        r
          ? {
              arrayLength: parseInt(r[2] || '-1'),
              arrayChildren: kt.fromObject({
                type: r[1],
                components: this.components,
              }),
              baseType: 'array',
            }
          : {
              arrayLength: null,
              arrayChildren: null,
              baseType: null != this.components ? 'tuple' : this.type,
            }
      ),
        (this._isParamType = !0),
        Object.freeze(this);
    }
    format(t) {
      if (
        (t || (t = Mt.sighash),
        Mt[t] || yt.throwArgumentError('invalid format type', 'format', t),
        t === Mt.json)
      ) {
        let e = {
          type: 'tuple' === this.baseType ? 'tuple' : this.type,
          name: this.name || void 0,
        };
        return (
          'boolean' == typeof this.indexed && (e.indexed = this.indexed),
          this.components &&
            (e.components = this.components.map((e) =>
              JSON.parse(e.format(t))
            )),
          JSON.stringify(e)
        );
      }
      let e = '';
      return (
        'array' === this.baseType
          ? ((e += this.arrayChildren.format(t)),
            (e +=
              '[' +
              (this.arrayLength < 0 ? '' : String(this.arrayLength)) +
              ']'))
          : 'tuple' === this.baseType
          ? (t !== Mt.sighash && (e += this.type),
            (e +=
              '(' +
              this.components
                .map((e) => e.format(t))
                .join(t === Mt.full ? ', ' : ',') +
              ')'))
          : (e += this.type),
        t !== Mt.sighash &&
          (!0 === this.indexed && (e += ' indexed'),
          t === Mt.full && this.name && (e += ' ' + this.name)),
        e
      );
    }
    static from(t, e) {
      return 'string' == typeof t ? kt.fromString(t, e) : kt.fromObject(t);
    }
    static fromObject(t) {
      return kt.isParamType(t)
        ? t
        : new kt(vt, {
            name: t.name || null,
            type: Ft(t.type),
            indexed: null == t.indexed ? null : !!t.indexed,
            components: t.components ? t.components.map(kt.fromObject) : null,
          });
    }
    static fromString(t, e) {
      return (
        (r = (function (t, e) {
          let r = t;
          function n(e) {
            yt.throwArgumentError(
              `unexpected character at position ${e}`,
              'param',
              t
            );
          }
          function i(t) {
            let r = { type: '', name: '', parent: t, state: { allowType: !0 } };
            return e && (r.indexed = !1), r;
          }
          t = t.replace(/\s/g, ' ');
          let o = { type: '', name: '', state: { allowType: !0 } },
            s = o;
          for (let r = 0; r < t.length; r++) {
            let o = t[r];
            switch (o) {
              case '(':
                s.state.allowType && '' === s.type
                  ? (s.type = 'tuple')
                  : s.state.allowParams || n(r),
                  (s.state.allowType = !1),
                  (s.type = Ft(s.type)),
                  (s.components = [i(s)]),
                  (s = s.components[0]);
                break;
              case ')':
                delete s.state,
                  'indexed' === s.name &&
                    (e || n(r), (s.indexed = !0), (s.name = '')),
                  At(s.type, s.name) && (s.name = ''),
                  (s.type = Ft(s.type));
                let t = s;
                (s = s.parent),
                  s || n(r),
                  delete t.parent,
                  (s.state.allowParams = !1),
                  (s.state.allowName = !0),
                  (s.state.allowArray = !0);
                break;
              case ',':
                delete s.state,
                  'indexed' === s.name &&
                    (e || n(r), (s.indexed = !0), (s.name = '')),
                  At(s.type, s.name) && (s.name = ''),
                  (s.type = Ft(s.type));
                let a = i(s.parent);
                s.parent.components.push(a), delete s.parent, (s = a);
                break;
              case ' ':
                s.state.allowType &&
                  '' !== s.type &&
                  ((s.type = Ft(s.type)),
                  delete s.state.allowType,
                  (s.state.allowName = !0),
                  (s.state.allowParams = !0)),
                  s.state.allowName &&
                    '' !== s.name &&
                    ('indexed' === s.name
                      ? (e || n(r),
                        s.indexed && n(r),
                        (s.indexed = !0),
                        (s.name = ''))
                      : At(s.type, s.name)
                      ? (s.name = '')
                      : (s.state.allowName = !1));
                break;
              case '[':
                s.state.allowArray || n(r),
                  (s.type += o),
                  (s.state.allowArray = !1),
                  (s.state.allowName = !1),
                  (s.state.readArray = !0);
                break;
              case ']':
                s.state.readArray || n(r),
                  (s.type += o),
                  (s.state.readArray = !1),
                  (s.state.allowArray = !0),
                  (s.state.allowName = !0);
                break;
              default:
                s.state.allowType
                  ? ((s.type += o),
                    (s.state.allowParams = !0),
                    (s.state.allowArray = !0))
                  : s.state.allowName
                  ? ((s.name += o), delete s.state.allowArray)
                  : s.state.readArray
                  ? (s.type += o)
                  : n(r);
            }
          }
          return (
            s.parent && yt.throwArgumentError('unexpected eof', 'param', t),
            delete o.state,
            'indexed' === s.name
              ? (e || n(r.length - 7),
                s.indexed && n(r.length - 7),
                (s.indexed = !0),
                (s.name = ''))
              : At(s.type, s.name) && (s.name = ''),
            (o.type = Ft(o.type)),
            o
          );
        })(t, !!e)),
        kt.fromObject({
          name: r.name,
          type: r.type,
          indexed: r.indexed,
          components: r.components,
        })
      );
      var r;
    }
    static isParamType(t) {
      return !(null == t || !t._isParamType);
    }
  }
  function St(t, e) {
    return (function (t) {
      t = t.trim();
      let e = [],
        r = '',
        n = 0;
      for (let i = 0; i < t.length; i++) {
        let o = t[i];
        ',' === o && 0 === n
          ? (e.push(r), (r = ''))
          : ((r += o),
            '(' === o
              ? n++
              : ')' === o &&
                (n--,
                -1 === n &&
                  yt.throwArgumentError('unbalanced parenthesis', 'value', t)));
      }
      r && e.push(r);
      return e;
    })(t).map((t) => kt.fromString(t, e));
  }
  class Nt {
    constructor(t, e) {
      t !== vt &&
        yt.throwError(
          'use a static from method',
          f.errors.UNSUPPORTED_OPERATION,
          { operation: 'new Fragment()' }
        ),
        Et(this, e),
        (this._isFragment = !0),
        Object.freeze(this);
    }
    static from(t) {
      return Nt.isFragment(t)
        ? t
        : 'string' == typeof t
        ? Nt.fromString(t)
        : Nt.fromObject(t);
    }
    static fromObject(t) {
      if (Nt.isFragment(t)) return t;
      switch (t.type) {
        case 'function':
          return It.fromObject(t);
        case 'event':
          return Pt.fromObject(t);
        case 'constructor':
          return Rt.fromObject(t);
        case 'error':
          return Ot.fromObject(t);
        case 'fallback':
        case 'receive':
          return null;
      }
      return yt.throwArgumentError('invalid fragment object', 'value', t);
    }
    static fromString(t) {
      return 'event' ===
        (t = (t = (t = t.replace(/\s/g, ' '))
          .replace(/\(/g, ' (')
          .replace(/\)/g, ') ')
          .replace(/\s+/g, ' ')).trim()).split(' ')[0]
        ? Pt.fromString(t.substring(5).trim())
        : 'function' === t.split(' ')[0]
        ? It.fromString(t.substring(8).trim())
        : 'constructor' === t.split('(')[0].trim()
        ? Rt.fromString(t.trim())
        : 'error' === t.split(' ')[0]
        ? Ot.fromString(t.substring(5).trim())
        : yt.throwArgumentError('unsupported fragment', 'value', t);
    }
    static isFragment(t) {
      return !(!t || !t._isFragment);
    }
  }
  class Pt extends Nt {
    format(t) {
      if (
        (t || (t = Mt.sighash),
        Mt[t] || yt.throwArgumentError('invalid format type', 'format', t),
        t === Mt.json)
      )
        return JSON.stringify({
          type: 'event',
          anonymous: this.anonymous,
          name: this.name,
          inputs: this.inputs.map((e) => JSON.parse(e.format(t))),
        });
      let e = '';
      return (
        t !== Mt.sighash && (e += 'event '),
        (e +=
          this.name +
          '(' +
          this.inputs.map((e) => e.format(t)).join(t === Mt.full ? ', ' : ',') +
          ') '),
        t !== Mt.sighash && this.anonymous && (e += 'anonymous '),
        e.trim()
      );
    }
    static from(t) {
      return 'string' == typeof t ? Pt.fromString(t) : Pt.fromObject(t);
    }
    static fromObject(t) {
      if (Pt.isEventFragment(t)) return t;
      'event' !== t.type &&
        yt.throwArgumentError('invalid event object', 'value', t);
      const e = {
        name: Lt(t.name),
        anonymous: t.anonymous,
        inputs: t.inputs ? t.inputs.map(kt.fromObject) : [],
        type: 'event',
      };
      return new Pt(vt, e);
    }
    static fromString(t) {
      let e = t.match(Ut);
      e || yt.throwArgumentError('invalid event string', 'value', t);
      let r = !1;
      return (
        e[3].split(' ').forEach((t) => {
          switch (t.trim()) {
            case 'anonymous':
              r = !0;
              break;
            case '':
              break;
            default:
              yt.warn('unknown modifier: ' + t);
          }
        }),
        Pt.fromObject({
          name: e[1].trim(),
          anonymous: r,
          inputs: St(e[2], !0),
          type: 'event',
        })
      );
    }
    static isEventFragment(t) {
      return t && t._isFragment && 'event' === t.type;
    }
  }
  function xt(t, e) {
    e.gas = null;
    let r = t.split('@');
    return 1 !== r.length
      ? (r.length > 2 &&
          yt.throwArgumentError(
            'invalid human-readable ABI signature',
            'value',
            t
          ),
        r[1].match(/^[0-9]+$/) ||
          yt.throwArgumentError(
            'invalid human-readable ABI signature gas',
            'value',
            t
          ),
        (e.gas = j.from(r[1])),
        r[0])
      : t;
  }
  function Tt(t, e) {
    (e.constant = !1),
      (e.payable = !1),
      (e.stateMutability = 'nonpayable'),
      t.split(' ').forEach((t) => {
        switch (t.trim()) {
          case 'constant':
            e.constant = !0;
            break;
          case 'payable':
            (e.payable = !0), (e.stateMutability = 'payable');
            break;
          case 'nonpayable':
            (e.payable = !1), (e.stateMutability = 'nonpayable');
            break;
          case 'pure':
            (e.constant = !0), (e.stateMutability = 'pure');
            break;
          case 'view':
            (e.constant = !0), (e.stateMutability = 'view');
            break;
          case 'external':
          case 'public':
          case '':
            break;
          default:
            console.log('unknown modifier: ' + t);
        }
      });
  }
  function Ct(t) {
    let e = { constant: !1, payable: !0, stateMutability: 'payable' };
    return (
      null != t.stateMutability
        ? ((e.stateMutability = t.stateMutability),
          (e.constant =
            'view' === e.stateMutability || 'pure' === e.stateMutability),
          null != t.constant &&
            !!t.constant !== e.constant &&
            yt.throwArgumentError(
              'cannot have constant function with mutability ' +
                e.stateMutability,
              'value',
              t
            ),
          (e.payable = 'payable' === e.stateMutability),
          null != t.payable &&
            !!t.payable !== e.payable &&
            yt.throwArgumentError(
              'cannot have payable function with mutability ' +
                e.stateMutability,
              'value',
              t
            ))
        : null != t.payable
        ? ((e.payable = !!t.payable),
          null != t.constant ||
            e.payable ||
            'constructor' === t.type ||
            yt.throwArgumentError(
              'unable to determine stateMutability',
              'value',
              t
            ),
          (e.constant = !!t.constant),
          e.constant
            ? (e.stateMutability = 'view')
            : (e.stateMutability = e.payable ? 'payable' : 'nonpayable'),
          e.payable &&
            e.constant &&
            yt.throwArgumentError(
              'cannot have constant payable function',
              'value',
              t
            ))
        : null != t.constant
        ? ((e.constant = !!t.constant),
          (e.payable = !e.constant),
          (e.stateMutability = e.constant ? 'view' : 'payable'))
        : 'constructor' !== t.type &&
          yt.throwArgumentError(
            'unable to determine stateMutability',
            'value',
            t
          ),
      e
    );
  }
  class Rt extends Nt {
    format(t) {
      if (
        (t || (t = Mt.sighash),
        Mt[t] || yt.throwArgumentError('invalid format type', 'format', t),
        t === Mt.json)
      )
        return JSON.stringify({
          type: 'constructor',
          stateMutability:
            'nonpayable' !== this.stateMutability
              ? this.stateMutability
              : void 0,
          payable: this.payable,
          gas: this.gas ? this.gas.toNumber() : void 0,
          inputs: this.inputs.map((e) => JSON.parse(e.format(t))),
        });
      t === Mt.sighash &&
        yt.throwError(
          'cannot format a constructor for sighash',
          f.errors.UNSUPPORTED_OPERATION,
          { operation: 'format(sighash)' }
        );
      let e =
        'constructor(' +
        this.inputs.map((e) => e.format(t)).join(t === Mt.full ? ', ' : ',') +
        ') ';
      return (
        this.stateMutability &&
          'nonpayable' !== this.stateMutability &&
          (e += this.stateMutability + ' '),
        e.trim()
      );
    }
    static from(t) {
      return 'string' == typeof t ? Rt.fromString(t) : Rt.fromObject(t);
    }
    static fromObject(t) {
      if (Rt.isConstructorFragment(t)) return t;
      'constructor' !== t.type &&
        yt.throwArgumentError('invalid constructor object', 'value', t);
      let e = Ct(t);
      e.constant &&
        yt.throwArgumentError('constructor cannot be constant', 'value', t);
      const r = {
        name: null,
        type: t.type,
        inputs: t.inputs ? t.inputs.map(kt.fromObject) : [],
        payable: e.payable,
        stateMutability: e.stateMutability,
        gas: t.gas ? j.from(t.gas) : null,
      };
      return new Rt(vt, r);
    }
    static fromString(t) {
      let e = { type: 'constructor' },
        r = (t = xt(t, e)).match(Ut);
      return (
        (r && 'constructor' === r[1].trim()) ||
          yt.throwArgumentError('invalid constructor string', 'value', t),
        (e.inputs = St(r[2].trim(), !1)),
        Tt(r[3].trim(), e),
        Rt.fromObject(e)
      );
    }
    static isConstructorFragment(t) {
      return t && t._isFragment && 'constructor' === t.type;
    }
  }
  class It extends Rt {
    format(t) {
      if (
        (t || (t = Mt.sighash),
        Mt[t] || yt.throwArgumentError('invalid format type', 'format', t),
        t === Mt.json)
      )
        return JSON.stringify({
          type: 'function',
          name: this.name,
          constant: this.constant,
          stateMutability:
            'nonpayable' !== this.stateMutability
              ? this.stateMutability
              : void 0,
          payable: this.payable,
          gas: this.gas ? this.gas.toNumber() : void 0,
          inputs: this.inputs.map((e) => JSON.parse(e.format(t))),
          outputs: this.outputs.map((e) => JSON.parse(e.format(t))),
        });
      let e = '';
      return (
        t !== Mt.sighash && (e += 'function '),
        (e +=
          this.name +
          '(' +
          this.inputs.map((e) => e.format(t)).join(t === Mt.full ? ', ' : ',') +
          ') '),
        t !== Mt.sighash &&
          (this.stateMutability
            ? 'nonpayable' !== this.stateMutability &&
              (e += this.stateMutability + ' ')
            : this.constant && (e += 'view '),
          this.outputs &&
            this.outputs.length &&
            (e +=
              'returns (' +
              this.outputs.map((e) => e.format(t)).join(', ') +
              ') '),
          null != this.gas && (e += '@' + this.gas.toString() + ' ')),
        e.trim()
      );
    }
    static from(t) {
      return 'string' == typeof t ? It.fromString(t) : It.fromObject(t);
    }
    static fromObject(t) {
      if (It.isFunctionFragment(t)) return t;
      'function' !== t.type &&
        yt.throwArgumentError('invalid function object', 'value', t);
      let e = Ct(t);
      const r = {
        type: t.type,
        name: Lt(t.name),
        constant: e.constant,
        inputs: t.inputs ? t.inputs.map(kt.fromObject) : [],
        outputs: t.outputs ? t.outputs.map(kt.fromObject) : [],
        payable: e.payable,
        stateMutability: e.stateMutability,
        gas: t.gas ? j.from(t.gas) : null,
      };
      return new It(vt, r);
    }
    static fromString(t) {
      let e = { type: 'function' },
        r = (t = xt(t, e)).split(' returns ');
      r.length > 2 &&
        yt.throwArgumentError('invalid function string', 'value', t);
      let n = r[0].match(Ut);
      if (
        (n || yt.throwArgumentError('invalid function signature', 'value', t),
        (e.name = n[1].trim()),
        e.name && Lt(e.name),
        (e.inputs = St(n[2], !1)),
        Tt(n[3].trim(), e),
        r.length > 1)
      ) {
        let n = r[1].match(Ut);
        ('' == n[1].trim() && '' == n[3].trim()) ||
          yt.throwArgumentError('unexpected tokens', 'value', t),
          (e.outputs = St(n[2], !1));
      } else e.outputs = [];
      return It.fromObject(e);
    }
    static isFunctionFragment(t) {
      return t && t._isFragment && 'function' === t.type;
    }
  }
  function Bt(t) {
    const e = t.format();
    return (
      ('Error(string)' !== e && 'Panic(uint256)' !== e) ||
        yt.throwArgumentError(
          `cannot specify user defined ${e} error`,
          'fragment',
          t
        ),
      t
    );
  }
  class Ot extends Nt {
    format(t) {
      if (
        (t || (t = Mt.sighash),
        Mt[t] || yt.throwArgumentError('invalid format type', 'format', t),
        t === Mt.json)
      )
        return JSON.stringify({
          type: 'error',
          name: this.name,
          inputs: this.inputs.map((e) => JSON.parse(e.format(t))),
        });
      let e = '';
      return (
        t !== Mt.sighash && (e += 'error '),
        (e +=
          this.name +
          '(' +
          this.inputs.map((e) => e.format(t)).join(t === Mt.full ? ', ' : ',') +
          ') '),
        e.trim()
      );
    }
    static from(t) {
      return 'string' == typeof t ? Ot.fromString(t) : Ot.fromObject(t);
    }
    static fromObject(t) {
      if (Ot.isErrorFragment(t)) return t;
      'error' !== t.type &&
        yt.throwArgumentError('invalid error object', 'value', t);
      const e = {
        type: t.type,
        name: Lt(t.name),
        inputs: t.inputs ? t.inputs.map(kt.fromObject) : [],
      };
      return Bt(new Ot(vt, e));
    }
    static fromString(t) {
      let e = { type: 'error' },
        r = t.match(Ut);
      return (
        r || yt.throwArgumentError('invalid error signature', 'value', t),
        (e.name = r[1].trim()),
        e.name && Lt(e.name),
        (e.inputs = St(r[2], !1)),
        Bt(Ot.fromObject(e))
      );
    }
    static isErrorFragment(t) {
      return t && t._isFragment && 'error' === t.type;
    }
  }
  function Ft(t) {
    return (
      t.match(/^uint($|[^1-9])/)
        ? (t = 'uint256' + t.substring(4))
        : t.match(/^int($|[^1-9])/) && (t = 'int256' + t.substring(3)),
      t
    );
  }
  const Dt = new RegExp('^[a-zA-Z$_][a-zA-Z0-9$_]*$');
  function Lt(t) {
    return (
      (t && t.match(Dt)) ||
        yt.throwArgumentError(`invalid identifier "${t}"`, 'value', t),
      t
    );
  }
  const Ut = new RegExp('^([^)(]*)\\((.*)\\)([^)(]*)$');
  const jt = new f(gt);
  class zt {
    constructor(t, e, r, n) {
      (this.name = t),
        (this.type = e),
        (this.localName = r),
        (this.dynamic = n);
    }
    _throwError(t, e) {
      jt.throwArgumentError(t, this.localName, e);
    }
  }
  class Gt {
    constructor(t) {
      st(this, 'wordSize', t || 32),
        (this._data = []),
        (this._dataLength = 0),
        (this._padding = new Uint8Array(t));
    }
    get data() {
      return P(this._data);
    }
    get length() {
      return this._dataLength;
    }
    _writeData(t) {
      return this._data.push(t), (this._dataLength += t.length), t.length;
    }
    appendWriter(t) {
      return this._writeData(b(t._data));
    }
    writeBytes(t) {
      let e = w(t);
      const r = e.length % this.wordSize;
      return r && (e = b([e, this._padding.slice(r)])), this._writeData(e);
    }
    _getValue(t) {
      let e = w(j.from(t));
      return (
        e.length > this.wordSize &&
          jt.throwError('value out-of-bounds', f.errors.BUFFER_OVERRUN, {
            length: this.wordSize,
            offset: e.length,
          }),
        e.length % this.wordSize &&
          (e = b([this._padding.slice(e.length % this.wordSize), e])),
        e
      );
    }
    writeValue(t) {
      return this._writeData(this._getValue(t));
    }
    writeUpdatableValue() {
      const t = this._data.length;
      return (
        this._data.push(this._padding),
        (this._dataLength += this.wordSize),
        (e) => {
          this._data[t] = this._getValue(e);
        }
      );
    }
  }
  class qt {
    constructor(t, e, r, n) {
      st(this, '_data', w(t)),
        st(this, 'wordSize', e || 32),
        st(this, '_coerceFunc', r),
        st(this, 'allowLoose', n),
        (this._offset = 0);
    }
    get data() {
      return k(this._data);
    }
    get consumed() {
      return this._offset;
    }
    static coerce(t, e) {
      let r = t.match('^u?int([0-9]+)$');
      return r && parseInt(r[1]) <= 48 && (e = e.toNumber()), e;
    }
    coerce(t, e) {
      return this._coerceFunc ? this._coerceFunc(t, e) : qt.coerce(t, e);
    }
    _peekBytes(t, e, r) {
      let n = Math.ceil(e / this.wordSize) * this.wordSize;
      return (
        this._offset + n > this._data.length &&
          (this.allowLoose && r && this._offset + e <= this._data.length
            ? (n = e)
            : jt.throwError('data out-of-bounds', f.errors.BUFFER_OVERRUN, {
                length: this._data.length,
                offset: this._offset + n,
              })),
        this._data.slice(this._offset, this._offset + n)
      );
    }
    subReader(t) {
      return new qt(
        this._data.slice(this._offset + t),
        this.wordSize,
        this._coerceFunc,
        this.allowLoose
      );
    }
    readBytes(t, e) {
      let r = this._peekBytes(0, t, !!e);
      return (this._offset += r.length), r.slice(0, t);
    }
    readValue() {
      return j.from(this.readBytes(this.wordSize));
    }
  }
  var Kt =
    'undefined' != typeof global
      ? global
      : 'undefined' != typeof self
      ? self
      : 'undefined' != typeof window
      ? window
      : {};
  function Ht() {
    throw new Error('setTimeout has not been defined');
  }
  function Vt() {
    throw new Error('clearTimeout has not been defined');
  }
  var Jt = Ht,
    Wt = Vt;
  function Qt(t) {
    if (Jt === setTimeout) return setTimeout(t, 0);
    if ((Jt === Ht || !Jt) && setTimeout)
      return (Jt = setTimeout), setTimeout(t, 0);
    try {
      return Jt(t, 0);
    } catch (e) {
      try {
        return Jt.call(null, t, 0);
      } catch (e) {
        return Jt.call(this, t, 0);
      }
    }
  }
  'function' == typeof Kt.setTimeout && (Jt = setTimeout),
    'function' == typeof Kt.clearTimeout && (Wt = clearTimeout);
  var Yt,
    Zt = [],
    Xt = !1,
    $t = -1;
  function te() {
    Xt &&
      Yt &&
      ((Xt = !1),
      Yt.length ? (Zt = Yt.concat(Zt)) : ($t = -1),
      Zt.length && ee());
  }
  function ee() {
    if (!Xt) {
      var t = Qt(te);
      Xt = !0;
      for (var e = Zt.length; e; ) {
        for (Yt = Zt, Zt = []; ++$t < e; ) Yt && Yt[$t].run();
        ($t = -1), (e = Zt.length);
      }
      (Yt = null),
        (Xt = !1),
        (function (t) {
          if (Wt === clearTimeout) return clearTimeout(t);
          if ((Wt === Vt || !Wt) && clearTimeout)
            return (Wt = clearTimeout), clearTimeout(t);
          try {
            return Wt(t);
          } catch (e) {
            try {
              return Wt.call(null, t);
            } catch (e) {
              return Wt.call(this, t);
            }
          }
        })(t);
    }
  }
  function re(t, e) {
    (this.fun = t), (this.array = e);
  }
  re.prototype.run = function () {
    this.fun.apply(null, this.array);
  };
  function ne() {}
  var ie = ne,
    oe = ne,
    se = ne,
    ae = ne,
    ue = ne,
    he = ne,
    le = ne;
  var ce = Kt.performance || {},
    fe =
      ce.now ||
      ce.mozNow ||
      ce.msNow ||
      ce.oNow ||
      ce.webkitNow ||
      function () {
        return new Date().getTime();
      };
  var de = new Date();
  var me = {
      nextTick: function (t) {
        var e = new Array(arguments.length - 1);
        if (arguments.length > 1)
          for (var r = 1; r < arguments.length; r++) e[r - 1] = arguments[r];
        Zt.push(new re(t, e)), 1 !== Zt.length || Xt || Qt(ee);
      },
      title: 'browser',
      browser: !0,
      env: {},
      argv: [],
      version: '',
      versions: {},
      on: ie,
      addListener: oe,
      once: se,
      off: ae,
      removeListener: ue,
      removeAllListeners: he,
      emit: le,
      binding: function (t) {
        throw new Error('process.binding is not supported');
      },
      cwd: function () {
        return '/';
      },
      chdir: function (t) {
        throw new Error('process.chdir is not supported');
      },
      umask: function () {
        return 0;
      },
      hrtime: function (t) {
        var e = 0.001 * fe.call(ce),
          r = Math.floor(e),
          n = Math.floor((e % 1) * 1e9);
        return t && ((r -= t[0]), (n -= t[1]) < 0 && (r--, (n += 1e9))), [r, n];
      },
      platform: 'browser',
      release: {},
      config: {},
      uptime: function () {
        return (new Date() - de) / 1e3;
      },
    },
    pe = { exports: {} };
  !(function (e) {
    !(function () {
      var r = 'input is invalid type',
        n = 'object' == typeof window,
        i = n ? window : {};
      i.JS_SHA3_NO_WINDOW && (n = !1);
      var o = !n && 'object' == typeof self;
      !i.JS_SHA3_NO_NODE_JS && me.versions && me.versions.node
        ? (i = t)
        : o && (i = self);
      var s = !i.JS_SHA3_NO_COMMON_JS && e.exports,
        a = !i.JS_SHA3_NO_ARRAY_BUFFER && 'undefined' != typeof ArrayBuffer,
        u = '0123456789abcdef'.split(''),
        h = [4, 1024, 262144, 67108864],
        l = [0, 8, 16, 24],
        c = [
          1, 0, 32898, 0, 32906, 2147483648, 2147516416, 2147483648, 32907, 0,
          2147483649, 0, 2147516545, 2147483648, 32777, 2147483648, 138, 0, 136,
          0, 2147516425, 0, 2147483658, 0, 2147516555, 0, 139, 2147483648,
          32905, 2147483648, 32771, 2147483648, 32770, 2147483648, 128,
          2147483648, 32778, 0, 2147483658, 2147483648, 2147516545, 2147483648,
          32896, 2147483648, 2147483649, 0, 2147516424, 2147483648,
        ],
        f = [224, 256, 384, 512],
        d = [128, 256],
        m = ['hex', 'buffer', 'arrayBuffer', 'array', 'digest'],
        p = { 128: 168, 256: 136 };
      (!i.JS_SHA3_NO_NODE_JS && Array.isArray) ||
        (Array.isArray = function (t) {
          return '[object Array]' === Object.prototype.toString.call(t);
        }),
        !a ||
          (!i.JS_SHA3_NO_ARRAY_BUFFER_IS_VIEW && ArrayBuffer.isView) ||
          (ArrayBuffer.isView = function (t) {
            return (
              'object' == typeof t &&
              t.buffer &&
              t.buffer.constructor === ArrayBuffer
            );
          });
      for (
        var g = function (t, e, r) {
            return function (n) {
              return new C(t, e, t).update(n)[r]();
            };
          },
          y = function (t, e, r) {
            return function (n, i) {
              return new C(t, e, i).update(n)[r]();
            };
          },
          v = function (t, e, r) {
            return function (e, n, i, o) {
              return M['cshake' + t].update(e, n, i, o)[r]();
            };
          },
          w = function (t, e, r) {
            return function (e, n, i, o) {
              return M['kmac' + t].update(e, n, i, o)[r]();
            };
          },
          b = function (t, e, r, n) {
            for (var i = 0; i < m.length; ++i) {
              var o = m[i];
              t[o] = e(r, n, o);
            }
            return t;
          },
          A = function (t, e) {
            var r = g(t, e, 'hex');
            return (
              (r.create = function () {
                return new C(t, e, t);
              }),
              (r.update = function (t) {
                return r.create().update(t);
              }),
              b(r, g, t, e)
            );
          },
          E = [
            {
              name: 'keccak',
              padding: [1, 256, 65536, 16777216],
              bits: f,
              createMethod: A,
            },
            {
              name: 'sha3',
              padding: [6, 1536, 393216, 100663296],
              bits: f,
              createMethod: A,
            },
            {
              name: 'shake',
              padding: [31, 7936, 2031616, 520093696],
              bits: d,
              createMethod: function (t, e) {
                var r = y(t, e, 'hex');
                return (
                  (r.create = function (r) {
                    return new C(t, e, r);
                  }),
                  (r.update = function (t, e) {
                    return r.create(e).update(t);
                  }),
                  b(r, y, t, e)
                );
              },
            },
            {
              name: 'cshake',
              padding: h,
              bits: d,
              createMethod: function (t, e) {
                var r = p[t],
                  n = v(t, 0, 'hex');
                return (
                  (n.create = function (n, i, o) {
                    return i || o
                      ? new C(t, e, n).bytepad([i, o], r)
                      : M['shake' + t].create(n);
                  }),
                  (n.update = function (t, e, r, i) {
                    return n.create(e, r, i).update(t);
                  }),
                  b(n, v, t, e)
                );
              },
            },
            {
              name: 'kmac',
              padding: h,
              bits: d,
              createMethod: function (t, e) {
                var r = p[t],
                  n = w(t, 0, 'hex');
                return (
                  (n.create = function (n, i, o) {
                    return new R(t, e, i)
                      .bytepad(['KMAC', o], r)
                      .bytepad([n], r);
                  }),
                  (n.update = function (t, e, r, i) {
                    return n.create(t, r, i).update(e);
                  }),
                  b(n, w, t, e)
                );
              },
            },
          ],
          M = {},
          _ = [],
          k = 0;
        k < E.length;
        ++k
      )
        for (var S = E[k], N = S.bits, P = 0; P < N.length; ++P) {
          var x = S.name + '_' + N[P];
          if (
            (_.push(x),
            (M[x] = S.createMethod(N[P], S.padding)),
            'sha3' !== S.name)
          ) {
            var T = S.name + N[P];
            _.push(T), (M[T] = M[x]);
          }
        }
      function C(t, e, r) {
        (this.blocks = []),
          (this.s = []),
          (this.padding = e),
          (this.outputBits = r),
          (this.reset = !0),
          (this.finalized = !1),
          (this.block = 0),
          (this.start = 0),
          (this.blockCount = (1600 - (t << 1)) >> 5),
          (this.byteCount = this.blockCount << 2),
          (this.outputBlocks = r >> 5),
          (this.extraBytes = (31 & r) >> 3);
        for (var n = 0; n < 50; ++n) this.s[n] = 0;
      }
      function R(t, e, r) {
        C.call(this, t, e, r);
      }
      (C.prototype.update = function (t) {
        if (this.finalized) throw new Error('finalize already called');
        var e,
          n = typeof t;
        if ('string' !== n) {
          if ('object' !== n) throw new Error(r);
          if (null === t) throw new Error(r);
          if (a && t.constructor === ArrayBuffer) t = new Uint8Array(t);
          else if (!(Array.isArray(t) || (a && ArrayBuffer.isView(t))))
            throw new Error(r);
          e = !0;
        }
        for (
          var i,
            o,
            s = this.blocks,
            u = this.byteCount,
            h = t.length,
            c = this.blockCount,
            f = 0,
            d = this.s;
          f < h;

        ) {
          if (this.reset)
            for (this.reset = !1, s[0] = this.block, i = 1; i < c + 1; ++i)
              s[i] = 0;
          if (e)
            for (i = this.start; f < h && i < u; ++f)
              s[i >> 2] |= t[f] << l[3 & i++];
          else
            for (i = this.start; f < h && i < u; ++f)
              (o = t.charCodeAt(f)) < 128
                ? (s[i >> 2] |= o << l[3 & i++])
                : o < 2048
                ? ((s[i >> 2] |= (192 | (o >> 6)) << l[3 & i++]),
                  (s[i >> 2] |= (128 | (63 & o)) << l[3 & i++]))
                : o < 55296 || o >= 57344
                ? ((s[i >> 2] |= (224 | (o >> 12)) << l[3 & i++]),
                  (s[i >> 2] |= (128 | ((o >> 6) & 63)) << l[3 & i++]),
                  (s[i >> 2] |= (128 | (63 & o)) << l[3 & i++]))
                : ((o =
                    65536 + (((1023 & o) << 10) | (1023 & t.charCodeAt(++f)))),
                  (s[i >> 2] |= (240 | (o >> 18)) << l[3 & i++]),
                  (s[i >> 2] |= (128 | ((o >> 12) & 63)) << l[3 & i++]),
                  (s[i >> 2] |= (128 | ((o >> 6) & 63)) << l[3 & i++]),
                  (s[i >> 2] |= (128 | (63 & o)) << l[3 & i++]));
          if (((this.lastByteIndex = i), i >= u)) {
            for (this.start = i - u, this.block = s[c], i = 0; i < c; ++i)
              d[i] ^= s[i];
            I(d), (this.reset = !0);
          } else this.start = i;
        }
        return this;
      }),
        (C.prototype.encode = function (t, e) {
          var r = 255 & t,
            n = 1,
            i = [r];
          for (r = 255 & (t >>= 8); r > 0; )
            i.unshift(r), (r = 255 & (t >>= 8)), ++n;
          return e ? i.push(n) : i.unshift(n), this.update(i), i.length;
        }),
        (C.prototype.encodeString = function (t) {
          var e,
            n = typeof t;
          if ('string' !== n) {
            if ('object' !== n) throw new Error(r);
            if (null === t) throw new Error(r);
            if (a && t.constructor === ArrayBuffer) t = new Uint8Array(t);
            else if (!(Array.isArray(t) || (a && ArrayBuffer.isView(t))))
              throw new Error(r);
            e = !0;
          }
          var i = 0,
            o = t.length;
          if (e) i = o;
          else
            for (var s = 0; s < t.length; ++s) {
              var u = t.charCodeAt(s);
              u < 128
                ? (i += 1)
                : u < 2048
                ? (i += 2)
                : u < 55296 || u >= 57344
                ? (i += 3)
                : ((u =
                    65536 + (((1023 & u) << 10) | (1023 & t.charCodeAt(++s)))),
                  (i += 4));
            }
          return (i += this.encode(8 * i)), this.update(t), i;
        }),
        (C.prototype.bytepad = function (t, e) {
          for (var r = this.encode(e), n = 0; n < t.length; ++n)
            r += this.encodeString(t[n]);
          var i = e - (r % e),
            o = [];
          return (o.length = i), this.update(o), this;
        }),
        (C.prototype.finalize = function () {
          if (!this.finalized) {
            this.finalized = !0;
            var t = this.blocks,
              e = this.lastByteIndex,
              r = this.blockCount,
              n = this.s;
            if (
              ((t[e >> 2] |= this.padding[3 & e]),
              this.lastByteIndex === this.byteCount)
            )
              for (t[0] = t[r], e = 1; e < r + 1; ++e) t[e] = 0;
            for (t[r - 1] |= 2147483648, e = 0; e < r; ++e) n[e] ^= t[e];
            I(n);
          }
        }),
        (C.prototype.toString = C.prototype.hex =
          function () {
            this.finalize();
            for (
              var t,
                e = this.blockCount,
                r = this.s,
                n = this.outputBlocks,
                i = this.extraBytes,
                o = 0,
                s = 0,
                a = '';
              s < n;

            ) {
              for (o = 0; o < e && s < n; ++o, ++s)
                (t = r[o]),
                  (a +=
                    u[(t >> 4) & 15] +
                    u[15 & t] +
                    u[(t >> 12) & 15] +
                    u[(t >> 8) & 15] +
                    u[(t >> 20) & 15] +
                    u[(t >> 16) & 15] +
                    u[(t >> 28) & 15] +
                    u[(t >> 24) & 15]);
              s % e == 0 && (I(r), (o = 0));
            }
            return (
              i &&
                ((t = r[o]),
                (a += u[(t >> 4) & 15] + u[15 & t]),
                i > 1 && (a += u[(t >> 12) & 15] + u[(t >> 8) & 15]),
                i > 2 && (a += u[(t >> 20) & 15] + u[(t >> 16) & 15])),
              a
            );
          }),
        (C.prototype.arrayBuffer = function () {
          this.finalize();
          var t,
            e = this.blockCount,
            r = this.s,
            n = this.outputBlocks,
            i = this.extraBytes,
            o = 0,
            s = 0,
            a = this.outputBits >> 3;
          t = i ? new ArrayBuffer((n + 1) << 2) : new ArrayBuffer(a);
          for (var u = new Uint32Array(t); s < n; ) {
            for (o = 0; o < e && s < n; ++o, ++s) u[s] = r[o];
            s % e == 0 && I(r);
          }
          return i && ((u[o] = r[o]), (t = t.slice(0, a))), t;
        }),
        (C.prototype.buffer = C.prototype.arrayBuffer),
        (C.prototype.digest = C.prototype.array =
          function () {
            this.finalize();
            for (
              var t,
                e,
                r = this.blockCount,
                n = this.s,
                i = this.outputBlocks,
                o = this.extraBytes,
                s = 0,
                a = 0,
                u = [];
              a < i;

            ) {
              for (s = 0; s < r && a < i; ++s, ++a)
                (t = a << 2),
                  (e = n[s]),
                  (u[t] = 255 & e),
                  (u[t + 1] = (e >> 8) & 255),
                  (u[t + 2] = (e >> 16) & 255),
                  (u[t + 3] = (e >> 24) & 255);
              a % r == 0 && I(n);
            }
            return (
              o &&
                ((t = a << 2),
                (e = n[s]),
                (u[t] = 255 & e),
                o > 1 && (u[t + 1] = (e >> 8) & 255),
                o > 2 && (u[t + 2] = (e >> 16) & 255)),
              u
            );
          }),
        (R.prototype = new C()),
        (R.prototype.finalize = function () {
          return (
            this.encode(this.outputBits, !0), C.prototype.finalize.call(this)
          );
        });
      var I = function (t) {
        var e,
          r,
          n,
          i,
          o,
          s,
          a,
          u,
          h,
          l,
          f,
          d,
          m,
          p,
          g,
          y,
          v,
          w,
          b,
          A,
          E,
          M,
          _,
          k,
          S,
          N,
          P,
          x,
          T,
          C,
          R,
          I,
          B,
          O,
          F,
          D,
          L,
          U,
          j,
          z,
          G,
          q,
          K,
          H,
          V,
          J,
          W,
          Q,
          Y,
          Z,
          X,
          $,
          tt,
          et,
          rt,
          nt,
          it,
          ot,
          st,
          at,
          ut,
          ht,
          lt;
        for (n = 0; n < 48; n += 2)
          (i = t[0] ^ t[10] ^ t[20] ^ t[30] ^ t[40]),
            (o = t[1] ^ t[11] ^ t[21] ^ t[31] ^ t[41]),
            (s = t[2] ^ t[12] ^ t[22] ^ t[32] ^ t[42]),
            (a = t[3] ^ t[13] ^ t[23] ^ t[33] ^ t[43]),
            (u = t[4] ^ t[14] ^ t[24] ^ t[34] ^ t[44]),
            (h = t[5] ^ t[15] ^ t[25] ^ t[35] ^ t[45]),
            (l = t[6] ^ t[16] ^ t[26] ^ t[36] ^ t[46]),
            (f = t[7] ^ t[17] ^ t[27] ^ t[37] ^ t[47]),
            (e =
              (d = t[8] ^ t[18] ^ t[28] ^ t[38] ^ t[48]) ^
              ((s << 1) | (a >>> 31))),
            (r =
              (m = t[9] ^ t[19] ^ t[29] ^ t[39] ^ t[49]) ^
              ((a << 1) | (s >>> 31))),
            (t[0] ^= e),
            (t[1] ^= r),
            (t[10] ^= e),
            (t[11] ^= r),
            (t[20] ^= e),
            (t[21] ^= r),
            (t[30] ^= e),
            (t[31] ^= r),
            (t[40] ^= e),
            (t[41] ^= r),
            (e = i ^ ((u << 1) | (h >>> 31))),
            (r = o ^ ((h << 1) | (u >>> 31))),
            (t[2] ^= e),
            (t[3] ^= r),
            (t[12] ^= e),
            (t[13] ^= r),
            (t[22] ^= e),
            (t[23] ^= r),
            (t[32] ^= e),
            (t[33] ^= r),
            (t[42] ^= e),
            (t[43] ^= r),
            (e = s ^ ((l << 1) | (f >>> 31))),
            (r = a ^ ((f << 1) | (l >>> 31))),
            (t[4] ^= e),
            (t[5] ^= r),
            (t[14] ^= e),
            (t[15] ^= r),
            (t[24] ^= e),
            (t[25] ^= r),
            (t[34] ^= e),
            (t[35] ^= r),
            (t[44] ^= e),
            (t[45] ^= r),
            (e = u ^ ((d << 1) | (m >>> 31))),
            (r = h ^ ((m << 1) | (d >>> 31))),
            (t[6] ^= e),
            (t[7] ^= r),
            (t[16] ^= e),
            (t[17] ^= r),
            (t[26] ^= e),
            (t[27] ^= r),
            (t[36] ^= e),
            (t[37] ^= r),
            (t[46] ^= e),
            (t[47] ^= r),
            (e = l ^ ((i << 1) | (o >>> 31))),
            (r = f ^ ((o << 1) | (i >>> 31))),
            (t[8] ^= e),
            (t[9] ^= r),
            (t[18] ^= e),
            (t[19] ^= r),
            (t[28] ^= e),
            (t[29] ^= r),
            (t[38] ^= e),
            (t[39] ^= r),
            (t[48] ^= e),
            (t[49] ^= r),
            (p = t[0]),
            (g = t[1]),
            (J = (t[11] << 4) | (t[10] >>> 28)),
            (W = (t[10] << 4) | (t[11] >>> 28)),
            (x = (t[20] << 3) | (t[21] >>> 29)),
            (T = (t[21] << 3) | (t[20] >>> 29)),
            (at = (t[31] << 9) | (t[30] >>> 23)),
            (ut = (t[30] << 9) | (t[31] >>> 23)),
            (q = (t[40] << 18) | (t[41] >>> 14)),
            (K = (t[41] << 18) | (t[40] >>> 14)),
            (O = (t[2] << 1) | (t[3] >>> 31)),
            (F = (t[3] << 1) | (t[2] >>> 31)),
            (y = (t[13] << 12) | (t[12] >>> 20)),
            (v = (t[12] << 12) | (t[13] >>> 20)),
            (Q = (t[22] << 10) | (t[23] >>> 22)),
            (Y = (t[23] << 10) | (t[22] >>> 22)),
            (C = (t[33] << 13) | (t[32] >>> 19)),
            (R = (t[32] << 13) | (t[33] >>> 19)),
            (ht = (t[42] << 2) | (t[43] >>> 30)),
            (lt = (t[43] << 2) | (t[42] >>> 30)),
            (et = (t[5] << 30) | (t[4] >>> 2)),
            (rt = (t[4] << 30) | (t[5] >>> 2)),
            (D = (t[14] << 6) | (t[15] >>> 26)),
            (L = (t[15] << 6) | (t[14] >>> 26)),
            (w = (t[25] << 11) | (t[24] >>> 21)),
            (b = (t[24] << 11) | (t[25] >>> 21)),
            (Z = (t[34] << 15) | (t[35] >>> 17)),
            (X = (t[35] << 15) | (t[34] >>> 17)),
            (I = (t[45] << 29) | (t[44] >>> 3)),
            (B = (t[44] << 29) | (t[45] >>> 3)),
            (k = (t[6] << 28) | (t[7] >>> 4)),
            (S = (t[7] << 28) | (t[6] >>> 4)),
            (nt = (t[17] << 23) | (t[16] >>> 9)),
            (it = (t[16] << 23) | (t[17] >>> 9)),
            (U = (t[26] << 25) | (t[27] >>> 7)),
            (j = (t[27] << 25) | (t[26] >>> 7)),
            (A = (t[36] << 21) | (t[37] >>> 11)),
            (E = (t[37] << 21) | (t[36] >>> 11)),
            ($ = (t[47] << 24) | (t[46] >>> 8)),
            (tt = (t[46] << 24) | (t[47] >>> 8)),
            (H = (t[8] << 27) | (t[9] >>> 5)),
            (V = (t[9] << 27) | (t[8] >>> 5)),
            (N = (t[18] << 20) | (t[19] >>> 12)),
            (P = (t[19] << 20) | (t[18] >>> 12)),
            (ot = (t[29] << 7) | (t[28] >>> 25)),
            (st = (t[28] << 7) | (t[29] >>> 25)),
            (z = (t[38] << 8) | (t[39] >>> 24)),
            (G = (t[39] << 8) | (t[38] >>> 24)),
            (M = (t[48] << 14) | (t[49] >>> 18)),
            (_ = (t[49] << 14) | (t[48] >>> 18)),
            (t[0] = p ^ (~y & w)),
            (t[1] = g ^ (~v & b)),
            (t[10] = k ^ (~N & x)),
            (t[11] = S ^ (~P & T)),
            (t[20] = O ^ (~D & U)),
            (t[21] = F ^ (~L & j)),
            (t[30] = H ^ (~J & Q)),
            (t[31] = V ^ (~W & Y)),
            (t[40] = et ^ (~nt & ot)),
            (t[41] = rt ^ (~it & st)),
            (t[2] = y ^ (~w & A)),
            (t[3] = v ^ (~b & E)),
            (t[12] = N ^ (~x & C)),
            (t[13] = P ^ (~T & R)),
            (t[22] = D ^ (~U & z)),
            (t[23] = L ^ (~j & G)),
            (t[32] = J ^ (~Q & Z)),
            (t[33] = W ^ (~Y & X)),
            (t[42] = nt ^ (~ot & at)),
            (t[43] = it ^ (~st & ut)),
            (t[4] = w ^ (~A & M)),
            (t[5] = b ^ (~E & _)),
            (t[14] = x ^ (~C & I)),
            (t[15] = T ^ (~R & B)),
            (t[24] = U ^ (~z & q)),
            (t[25] = j ^ (~G & K)),
            (t[34] = Q ^ (~Z & $)),
            (t[35] = Y ^ (~X & tt)),
            (t[44] = ot ^ (~at & ht)),
            (t[45] = st ^ (~ut & lt)),
            (t[6] = A ^ (~M & p)),
            (t[7] = E ^ (~_ & g)),
            (t[16] = C ^ (~I & k)),
            (t[17] = R ^ (~B & S)),
            (t[26] = z ^ (~q & O)),
            (t[27] = G ^ (~K & F)),
            (t[36] = Z ^ (~$ & H)),
            (t[37] = X ^ (~tt & V)),
            (t[46] = at ^ (~ht & et)),
            (t[47] = ut ^ (~lt & rt)),
            (t[8] = M ^ (~p & y)),
            (t[9] = _ ^ (~g & v)),
            (t[18] = I ^ (~k & N)),
            (t[19] = B ^ (~S & P)),
            (t[28] = q ^ (~O & D)),
            (t[29] = K ^ (~F & L)),
            (t[38] = $ ^ (~H & J)),
            (t[39] = tt ^ (~V & W)),
            (t[48] = ht ^ (~et & nt)),
            (t[49] = lt ^ (~rt & it)),
            (t[0] ^= c[n]),
            (t[1] ^= c[n + 1]);
      };
      if (s) e.exports = M;
      else for (k = 0; k < _.length; ++k) i[_[k]] = M[_[k]];
    })();
  })(pe);
  var ge = pe.exports;
  function ye(t) {
    return '0x' + ge.keccak_256(w(t));
  }
  const ve = new f('rlp/5.7.0');
  function we(t) {
    const e = [];
    for (; t; ) e.unshift(255 & t), (t >>= 8);
    return e;
  }
  function be(t, e, r) {
    let n = 0;
    for (let i = 0; i < r; i++) n = 256 * n + t[e + i];
    return n;
  }
  function Ae(t) {
    if (Array.isArray(t)) {
      let e = [];
      if (
        (t.forEach(function (t) {
          e = e.concat(Ae(t));
        }),
        e.length <= 55)
      )
        return e.unshift(192 + e.length), e;
      const r = we(e.length);
      return r.unshift(247 + r.length), r.concat(e);
    }
    g(t) || ve.throwArgumentError('RLP object must be BytesLike', 'object', t);
    const e = Array.prototype.slice.call(w(t));
    if (1 === e.length && e[0] <= 127) return e;
    if (e.length <= 55) return e.unshift(128 + e.length), e;
    const r = we(e.length);
    return r.unshift(183 + r.length), r.concat(e);
  }
  function Ee(t) {
    return k(Ae(t));
  }
  function Me(t, e, r, n) {
    const i = [];
    for (; r < e + 1 + n; ) {
      const o = _e(t, r);
      i.push(o.result),
        (r += o.consumed) > e + 1 + n &&
          ve.throwError('child data too short', f.errors.BUFFER_OVERRUN, {});
    }
    return { consumed: 1 + n, result: i };
  }
  function _e(t, e) {
    if (
      (0 === t.length &&
        ve.throwError('data too short', f.errors.BUFFER_OVERRUN, {}),
      t[e] >= 248)
    ) {
      const r = t[e] - 247;
      e + 1 + r > t.length &&
        ve.throwError(
          'data short segment too short',
          f.errors.BUFFER_OVERRUN,
          {}
        );
      const n = be(t, e + 1, r);
      return (
        e + 1 + r + n > t.length &&
          ve.throwError(
            'data long segment too short',
            f.errors.BUFFER_OVERRUN,
            {}
          ),
        Me(t, e, e + 1 + r, r + n)
      );
    }
    if (t[e] >= 192) {
      const r = t[e] - 192;
      return (
        e + 1 + r > t.length &&
          ve.throwError('data array too short', f.errors.BUFFER_OVERRUN, {}),
        Me(t, e, e + 1, r)
      );
    }
    if (t[e] >= 184) {
      const r = t[e] - 183;
      e + 1 + r > t.length &&
        ve.throwError('data array too short', f.errors.BUFFER_OVERRUN, {});
      const n = be(t, e + 1, r);
      e + 1 + r + n > t.length &&
        ve.throwError('data array too short', f.errors.BUFFER_OVERRUN, {});
      return {
        consumed: 1 + r + n,
        result: k(t.slice(e + 1 + r, e + 1 + r + n)),
      };
    }
    if (t[e] >= 128) {
      const r = t[e] - 128;
      e + 1 + r > t.length &&
        ve.throwError('data too short', f.errors.BUFFER_OVERRUN, {});
      return { consumed: 1 + r, result: k(t.slice(e + 1, e + 1 + r)) };
    }
    return { consumed: 1, result: k(t[e]) };
  }
  function ke(t) {
    const e = w(t),
      r = _e(e, 0);
    return (
      r.consumed !== e.length &&
        ve.throwArgumentError('invalid rlp data', 'data', t),
      r.result
    );
  }
  var Se = Object.freeze({ __proto__: null, encode: Ee, decode: ke });
  const Ne = new f('address/5.7.0');
  function Pe(t) {
    M(t, 20) || Ne.throwArgumentError('invalid address', 'address', t);
    const e = (t = t.toLowerCase()).substring(2).split(''),
      r = new Uint8Array(40);
    for (let t = 0; t < 40; t++) r[t] = e[t].charCodeAt(0);
    const n = w(ye(r));
    for (let t = 0; t < 40; t += 2)
      n[t >> 1] >> 4 >= 8 && (e[t] = e[t].toUpperCase()),
        (15 & n[t >> 1]) >= 8 && (e[t + 1] = e[t + 1].toUpperCase());
    return '0x' + e.join('');
  }
  const xe = {};
  for (let t = 0; t < 10; t++) xe[String(t)] = String(t);
  for (let t = 0; t < 26; t++) xe[String.fromCharCode(65 + t)] = String(10 + t);
  const Te = Math.floor(
    (function (t) {
      return Math.log10 ? Math.log10(t) : Math.log(t) / Math.LN10;
    })(9007199254740991)
  );
  function Ce(t) {
    let e = (t = (t = t.toUpperCase()).substring(4) + t.substring(0, 2) + '00')
      .split('')
      .map((t) => xe[t])
      .join('');
    for (; e.length >= Te; ) {
      let t = e.substring(0, Te);
      e = (parseInt(t, 10) % 97) + e.substring(t.length);
    }
    let r = String(98 - (parseInt(e, 10) % 97));
    for (; r.length < 2; ) r = '0' + r;
    return r;
  }
  function Re(t) {
    let e = null;
    if (
      ('string' != typeof t &&
        Ne.throwArgumentError('invalid address', 'address', t),
      t.match(/^(0x)?[0-9a-fA-F]{40}$/))
    )
      '0x' !== t.substring(0, 2) && (t = '0x' + t),
        (e = Pe(t)),
        t.match(/([A-F].*[a-f])|([a-f].*[A-F])/) &&
          e !== t &&
          Ne.throwArgumentError('bad address checksum', 'address', t);
    else if (t.match(/^XE[0-9]{2}[0-9A-Za-z]{30,31}$/)) {
      for (
        t.substring(2, 4) !== Ce(t) &&
          Ne.throwArgumentError('bad icap checksum', 'address', t),
          r = t.substring(4),
          e = new O(r, 36).toString(16);
        e.length < 40;

      )
        e = '0' + e;
      e = Pe('0x' + e);
    } else Ne.throwArgumentError('invalid address', 'address', t);
    var r;
    return e;
  }
  function Ie(t) {
    let e = null;
    try {
      e = Re(t.from);
    } catch (e) {
      Ne.throwArgumentError('missing from address', 'transaction', t);
    }
    return Re(N(ye(Ee([e, A(w(j.from(t.nonce).toHexString()))])), 12));
  }
  class Be extends zt {
    constructor(t) {
      super('address', 'address', t, !1);
    }
    defaultValue() {
      return '0x0000000000000000000000000000000000000000';
    }
    encode(t, e) {
      try {
        e = Re(e);
      } catch (t) {
        this._throwError(t.message, e);
      }
      return t.writeValue(e);
    }
    decode(t) {
      return Re(C(t.readValue().toHexString(), 20));
    }
  }
  class Oe extends zt {
    constructor(t) {
      super(t.name, t.type, void 0, t.dynamic), (this.coder = t);
    }
    defaultValue() {
      return this.coder.defaultValue();
    }
    encode(t, e) {
      return this.coder.encode(t, e);
    }
    decode(t) {
      return this.coder.decode(t);
    }
  }
  const Fe = new f(gt);
  function De(t, e, r) {
    let n = null;
    if (Array.isArray(r)) n = r;
    else if (r && 'object' == typeof r) {
      let t = {};
      n = e.map((e) => {
        const n = e.localName;
        return (
          n ||
            Fe.throwError(
              'cannot encode object for signature with missing names',
              f.errors.INVALID_ARGUMENT,
              { argument: 'values', coder: e, value: r }
            ),
          t[n] &&
            Fe.throwError(
              'cannot encode object for signature with duplicate names',
              f.errors.INVALID_ARGUMENT,
              { argument: 'values', coder: e, value: r }
            ),
          (t[n] = !0),
          r[n]
        );
      });
    } else Fe.throwArgumentError('invalid tuple value', 'tuple', r);
    e.length !== n.length &&
      Fe.throwArgumentError('types/value length mismatch', 'tuple', r);
    let i = new Gt(t.wordSize),
      o = new Gt(t.wordSize),
      s = [];
    e.forEach((t, e) => {
      let r = n[e];
      if (t.dynamic) {
        let e = o.length;
        t.encode(o, r);
        let n = i.writeUpdatableValue();
        s.push((t) => {
          n(t + e);
        });
      } else t.encode(i, r);
    }),
      s.forEach((t) => {
        t(i.length);
      });
    let a = t.appendWriter(i);
    return (a += t.appendWriter(o)), a;
  }
  function Le(t, e) {
    let r = [],
      n = t.subReader(0);
    e.forEach((e) => {
      let i = null;
      if (e.dynamic) {
        let r = t.readValue(),
          o = n.subReader(r.toNumber());
        try {
          i = e.decode(o);
        } catch (t) {
          if (t.code === f.errors.BUFFER_OVERRUN) throw t;
          (i = t),
            (i.baseType = e.name),
            (i.name = e.localName),
            (i.type = e.type);
        }
      } else
        try {
          i = e.decode(t);
        } catch (t) {
          if (t.code === f.errors.BUFFER_OVERRUN) throw t;
          (i = t),
            (i.baseType = e.name),
            (i.name = e.localName),
            (i.type = e.type);
        }
      null != i && r.push(i);
    });
    const i = e.reduce((t, e) => {
      const r = e.localName;
      return r && (t[r] || (t[r] = 0), t[r]++), t;
    }, {});
    e.forEach((t, e) => {
      let n = t.localName;
      if (!n || 1 !== i[n]) return;
      if (('length' === n && (n = '_length'), null != r[n])) return;
      const o = r[e];
      o instanceof Error
        ? Object.defineProperty(r, n, {
            enumerable: !0,
            get: () => {
              throw o;
            },
          })
        : (r[n] = o);
    });
    for (let t = 0; t < r.length; t++) {
      const e = r[t];
      e instanceof Error &&
        Object.defineProperty(r, t, {
          enumerable: !0,
          get: () => {
            throw e;
          },
        });
    }
    return Object.freeze(r);
  }
  class Ue extends zt {
    constructor(t, e, r) {
      super(
        'array',
        t.type + '[' + (e >= 0 ? e : '') + ']',
        r,
        -1 === e || t.dynamic
      ),
        (this.coder = t),
        (this.length = e);
    }
    defaultValue() {
      const t = this.coder.defaultValue(),
        e = [];
      for (let r = 0; r < this.length; r++) e.push(t);
      return e;
    }
    encode(t, e) {
      Array.isArray(e) || this._throwError('expected array value', e);
      let r = this.length;
      -1 === r && ((r = e.length), t.writeValue(e.length)),
        Fe.checkArgumentCount(
          e.length,
          r,
          'coder array' + (this.localName ? ' ' + this.localName : '')
        );
      let n = [];
      for (let t = 0; t < e.length; t++) n.push(this.coder);
      return De(t, n, e);
    }
    decode(t) {
      let e = this.length;
      -1 === e &&
        ((e = t.readValue().toNumber()),
        32 * e > t._data.length &&
          Fe.throwError('insufficient data length', f.errors.BUFFER_OVERRUN, {
            length: t._data.length,
            count: e,
          }));
      let r = [];
      for (let t = 0; t < e; t++) r.push(new Oe(this.coder));
      return t.coerce(this.name, Le(t, r));
    }
  }
  class je extends zt {
    constructor(t) {
      super('bool', 'bool', t, !1);
    }
    defaultValue() {
      return !1;
    }
    encode(t, e) {
      return t.writeValue(e ? 1 : 0);
    }
    decode(t) {
      return t.coerce(this.type, !t.readValue().isZero());
    }
  }
  class ze extends zt {
    constructor(t, e) {
      super(t, t, e, !0);
    }
    defaultValue() {
      return '0x';
    }
    encode(t, e) {
      e = w(e);
      let r = t.writeValue(e.length);
      return (r += t.writeBytes(e)), r;
    }
    decode(t) {
      return t.readBytes(t.readValue().toNumber(), !0);
    }
  }
  class Ge extends ze {
    constructor(t) {
      super('bytes', t);
    }
    decode(t) {
      return t.coerce(this.name, k(super.decode(t)));
    }
  }
  class qe extends zt {
    constructor(t, e) {
      let r = 'bytes' + String(t);
      super(r, r, e, !1), (this.size = t);
    }
    defaultValue() {
      return '0x0000000000000000000000000000000000000000000000000000000000000000'.substring(
        0,
        2 + 2 * this.size
      );
    }
    encode(t, e) {
      let r = w(e);
      return (
        r.length !== this.size && this._throwError('incorrect data length', e),
        t.writeBytes(r)
      );
    }
    decode(t) {
      return t.coerce(this.name, k(t.readBytes(this.size)));
    }
  }
  class Ke extends zt {
    constructor(t) {
      super('null', '', t, !1);
    }
    defaultValue() {
      return null;
    }
    encode(t, e) {
      return null != e && this._throwError('not null', e), t.writeBytes([]);
    }
    decode(t) {
      return t.readBytes(0), t.coerce(this.name, null);
    }
  }
  const He = '0x0000000000000000000000000000000000000000',
    Ve = j.from(-1),
    Je = j.from(0),
    We = j.from(1),
    Qe = j.from(2),
    Ye = j.from('1000000000000000000'),
    Ze = j.from(
      '0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff'
    ),
    Xe = j.from(
      '-0x8000000000000000000000000000000000000000000000000000000000000000'
    ),
    $e = j.from(
      '0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff'
    ),
    tr = '0x0000000000000000000000000000000000000000000000000000000000000000';
  var er = Object.freeze({
    __proto__: null,
    AddressZero: He,
    NegativeOne: Ve,
    Zero: Je,
    One: We,
    Two: Qe,
    WeiPerEther: Ye,
    MaxUint256: Ze,
    MinInt256: Xe,
    MaxInt256: $e,
    HashZero: tr,
    EtherSymbol: 'Ξ',
  });
  class rr extends zt {
    constructor(t, e, r) {
      const n = (e ? 'int' : 'uint') + 8 * t;
      super(n, n, r, !1), (this.size = t), (this.signed = e);
    }
    defaultValue() {
      return 0;
    }
    encode(t, e) {
      let r = j.from(e),
        n = Ze.mask(8 * t.wordSize);
      if (this.signed) {
        let t = n.mask(8 * this.size - 1);
        (r.gt(t) || r.lt(t.add(We).mul(Ve))) &&
          this._throwError('value out-of-bounds', e);
      } else (r.lt(Je) || r.gt(n.mask(8 * this.size))) && this._throwError('value out-of-bounds', e);
      return (
        (r = r.toTwos(8 * this.size).mask(8 * this.size)),
        this.signed && (r = r.fromTwos(8 * this.size).toTwos(8 * t.wordSize)),
        t.writeValue(r)
      );
    }
    decode(t) {
      let e = t.readValue().mask(8 * this.size);
      return (
        this.signed && (e = e.fromTwos(8 * this.size)), t.coerce(this.name, e)
      );
    }
  }
  const nr = new f('strings/5.7.0');
  var ir, or;
  function sr(t, e, r, n, i) {
    if (t === or.BAD_PREFIX || t === or.UNEXPECTED_CONTINUE) {
      let t = 0;
      for (let n = e + 1; n < r.length && r[n] >> 6 == 2; n++) t++;
      return t;
    }
    return t === or.OVERRUN ? r.length - e - 1 : 0;
  }
  !(function (t) {
    (t.current = ''),
      (t.NFC = 'NFC'),
      (t.NFD = 'NFD'),
      (t.NFKC = 'NFKC'),
      (t.NFKD = 'NFKD');
  })(ir || (ir = {})),
    (function (t) {
      (t.UNEXPECTED_CONTINUE = 'unexpected continuation byte'),
        (t.BAD_PREFIX = 'bad codepoint prefix'),
        (t.OVERRUN = 'string overrun'),
        (t.MISSING_CONTINUE = 'missing continuation byte'),
        (t.OUT_OF_RANGE = 'out of UTF-8 range'),
        (t.UTF16_SURROGATE = 'UTF-16 surrogate'),
        (t.OVERLONG = 'overlong representation');
    })(or || (or = {}));
  const ar = Object.freeze({
    error: function (t, e, r, n, i) {
      return nr.throwArgumentError(
        `invalid codepoint at offset ${e}; ${t}`,
        'bytes',
        r
      );
    },
    ignore: sr,
    replace: function (t, e, r, n, i) {
      return t === or.OVERLONG ? (n.push(i), 0) : (n.push(65533), sr(t, e, r));
    },
  });
  function ur(t, e) {
    null == e && (e = ar.error), (t = w(t));
    const r = [];
    let n = 0;
    for (; n < t.length; ) {
      const i = t[n++];
      if (!(i >> 7)) {
        r.push(i);
        continue;
      }
      let o = null,
        s = null;
      if (192 == (224 & i)) (o = 1), (s = 127);
      else if (224 == (240 & i)) (o = 2), (s = 2047);
      else {
        if (240 != (248 & i)) {
          n += e(
            128 == (192 & i) ? or.UNEXPECTED_CONTINUE : or.BAD_PREFIX,
            n - 1,
            t,
            r
          );
          continue;
        }
        (o = 3), (s = 65535);
      }
      if (n - 1 + o >= t.length) {
        n += e(or.OVERRUN, n - 1, t, r);
        continue;
      }
      let a = i & ((1 << (8 - o - 1)) - 1);
      for (let i = 0; i < o; i++) {
        let i = t[n];
        if (128 != (192 & i)) {
          (n += e(or.MISSING_CONTINUE, n, t, r)), (a = null);
          break;
        }
        (a = (a << 6) | (63 & i)), n++;
      }
      null !== a &&
        (a > 1114111
          ? (n += e(or.OUT_OF_RANGE, n - 1 - o, t, r, a))
          : a >= 55296 && a <= 57343
          ? (n += e(or.UTF16_SURROGATE, n - 1 - o, t, r, a))
          : a <= s
          ? (n += e(or.OVERLONG, n - 1 - o, t, r, a))
          : r.push(a));
    }
    return r;
  }
  function hr(t, e = ir.current) {
    e != ir.current && (nr.checkNormalize(), (t = t.normalize(e)));
    let r = [];
    for (let e = 0; e < t.length; e++) {
      const n = t.charCodeAt(e);
      if (n < 128) r.push(n);
      else if (n < 2048) r.push((n >> 6) | 192), r.push((63 & n) | 128);
      else if (55296 == (64512 & n)) {
        e++;
        const i = t.charCodeAt(e);
        if (e >= t.length || 56320 != (64512 & i))
          throw new Error('invalid utf-8 string');
        const o = 65536 + ((1023 & n) << 10) + (1023 & i);
        r.push((o >> 18) | 240),
          r.push(((o >> 12) & 63) | 128),
          r.push(((o >> 6) & 63) | 128),
          r.push((63 & o) | 128);
      } else
        r.push((n >> 12) | 224),
          r.push(((n >> 6) & 63) | 128),
          r.push((63 & n) | 128);
    }
    return w(r);
  }
  function lr(t) {
    const e = '0000' + t.toString(16);
    return '\\u' + e.substring(e.length - 4);
  }
  function cr(t) {
    return t
      .map((t) =>
        t <= 65535
          ? String.fromCharCode(t)
          : ((t -= 65536),
            String.fromCharCode(55296 + ((t >> 10) & 1023), 56320 + (1023 & t)))
      )
      .join('');
  }
  function fr(t, e) {
    return cr(ur(t, e));
  }
  function dr(t, e = ir.current) {
    return ur(hr(t, e));
  }
  function mr(t, e) {
    e ||
      (e = function (t) {
        return [parseInt(t, 16)];
      });
    let r = 0,
      n = {};
    return (
      t.split(',').forEach((t) => {
        let i = t.split(':');
        (r += parseInt(i[0], 16)), (n[r] = e(i[1]));
      }),
      n
    );
  }
  function pr(t) {
    let e = 0;
    return t.split(',').map((t) => {
      let r = t.split('-');
      1 === r.length ? (r[1] = '0') : '' === r[1] && (r[1] = '1');
      let n = e + parseInt(r[0], 16);
      return (e = parseInt(r[1], 16)), { l: n, h: e };
    });
  }
  function gr(t, e) {
    let r = 0;
    for (let n = 0; n < e.length; n++) {
      let i = e[n];
      if (((r += i.l), t >= r && t <= r + i.h && (t - r) % (i.d || 1) == 0)) {
        if (i.e && -1 !== i.e.indexOf(t - r)) continue;
        return i;
      }
    }
    return null;
  }
  const yr = pr(
      '221,13-1b,5f-,40-10,51-f,11-3,3-3,2-2,2-4,8,2,15,2d,28-8,88,48,27-,3-5,11-20,27-,8,28,3-5,12,18,b-a,1c-4,6-16,2-d,2-2,2,1b-4,17-9,8f-,10,f,1f-2,1c-34,33-14e,4,36-,13-,6-2,1a-f,4,9-,3-,17,8,2-2,5-,2,8-,3-,4-8,2-3,3,6-,16-6,2-,7-3,3-,17,8,3,3,3-,2,6-3,3-,4-a,5,2-6,10-b,4,8,2,4,17,8,3,6-,b,4,4-,2-e,2-4,b-10,4,9-,3-,17,8,3-,5-,9-2,3-,4-7,3-3,3,4-3,c-10,3,7-2,4,5-2,3,2,3-2,3-2,4-2,9,4-3,6-2,4,5-8,2-e,d-d,4,9,4,18,b,6-3,8,4,5-6,3-8,3-3,b-11,3,9,4,18,b,6-3,8,4,5-6,3-6,2,3-3,b-11,3,9,4,18,11-3,7-,4,5-8,2-7,3-3,b-11,3,13-2,19,a,2-,8-2,2-3,7,2,9-11,4-b,3b-3,1e-24,3,2-,3,2-,2-5,5,8,4,2,2-,3,e,4-,6,2,7-,b-,3-21,49,23-5,1c-3,9,25,10-,2-2f,23,6,3,8-2,5-5,1b-45,27-9,2a-,2-3,5b-4,45-4,53-5,8,40,2,5-,8,2,5-,28,2,5-,20,2,5-,8,2,5-,8,8,18,20,2,5-,8,28,14-5,1d-22,56-b,277-8,1e-2,52-e,e,8-a,18-8,15-b,e,4,3-b,5e-2,b-15,10,b-5,59-7,2b-555,9d-3,5b-5,17-,7-,27-,7-,9,2,2,2,20-,36,10,f-,7,14-,4,a,54-3,2-6,6-5,9-,1c-10,13-1d,1c-14,3c-,10-6,32-b,240-30,28-18,c-14,a0,115-,3,66-,b-76,5,5-,1d,24,2,5-2,2,8-,35-2,19,f-10,1d-3,311-37f,1b,5a-b,d7-19,d-3,41,57-,68-4,29-3,5f,29-37,2e-2,25-c,2c-2,4e-3,30,78-3,64-,20,19b7-49,51a7-59,48e-2,38-738,2ba5-5b,222f-,3c-94,8-b,6-4,1b,6,2,3,3,6d-20,16e-f,41-,37-7,2e-2,11-f,5-b,18-,b,14,5-3,6,88-,2,bf-2,7-,7-,7-,4-2,8,8-9,8-2ff,20,5-b,1c-b4,27-,27-cbb1,f7-9,28-2,b5-221,56,48,3-,2-,3-,5,d,2,5,3,42,5-,9,8,1d,5,6,2-2,8,153-3,123-3,33-27fd,a6da-5128,21f-5df,3-fffd,3-fffd,3-fffd,3-fffd,3-fffd,3-fffd,3-fffd,3-fffd,3-fffd,3-fffd,3-fffd,3,2-1d,61-ff7d'
    ),
    vr = 'ad,34f,1806,180b,180c,180d,200b,200c,200d,2060,feff'
      .split(',')
      .map((t) => parseInt(t, 16)),
    wr = [
      { h: 25, s: 32, l: 65 },
      { h: 30, s: 32, e: [23], l: 127 },
      { h: 54, s: 1, e: [48], l: 64, d: 2 },
      { h: 14, s: 1, l: 57, d: 2 },
      { h: 44, s: 1, l: 17, d: 2 },
      { h: 10, s: 1, e: [2, 6, 8], l: 61, d: 2 },
      { h: 16, s: 1, l: 68, d: 2 },
      { h: 84, s: 1, e: [18, 24, 66], l: 19, d: 2 },
      { h: 26, s: 32, e: [17], l: 435 },
      { h: 22, s: 1, l: 71, d: 2 },
      { h: 15, s: 80, l: 40 },
      { h: 31, s: 32, l: 16 },
      { h: 32, s: 1, l: 80, d: 2 },
      { h: 52, s: 1, l: 42, d: 2 },
      { h: 12, s: 1, l: 55, d: 2 },
      { h: 40, s: 1, e: [38], l: 15, d: 2 },
      { h: 14, s: 1, l: 48, d: 2 },
      { h: 37, s: 48, l: 49 },
      { h: 148, s: 1, l: 6351, d: 2 },
      { h: 88, s: 1, l: 160, d: 2 },
      { h: 15, s: 16, l: 704 },
      { h: 25, s: 26, l: 854 },
      { h: 25, s: 32, l: 55915 },
      { h: 37, s: 40, l: 1247 },
      { h: 25, s: -119711, l: 53248 },
      { h: 25, s: -119763, l: 52 },
      { h: 25, s: -119815, l: 52 },
      { h: 25, s: -119867, e: [1, 4, 5, 7, 8, 11, 12, 17], l: 52 },
      { h: 25, s: -119919, l: 52 },
      { h: 24, s: -119971, e: [2, 7, 8, 17], l: 52 },
      { h: 24, s: -120023, e: [2, 7, 13, 15, 16, 17], l: 52 },
      { h: 25, s: -120075, l: 52 },
      { h: 25, s: -120127, l: 52 },
      { h: 25, s: -120179, l: 52 },
      { h: 25, s: -120231, l: 52 },
      { h: 25, s: -120283, l: 52 },
      { h: 25, s: -120335, l: 52 },
      { h: 24, s: -119543, e: [17], l: 56 },
      { h: 24, s: -119601, e: [17], l: 58 },
      { h: 24, s: -119659, e: [17], l: 58 },
      { h: 24, s: -119717, e: [17], l: 58 },
      { h: 24, s: -119775, e: [17], l: 58 },
    ],
    br = mr(
      'b5:3bc,c3:ff,7:73,2:253,5:254,3:256,1:257,5:259,1:25b,3:260,1:263,2:269,1:268,5:26f,1:272,2:275,7:280,3:283,5:288,3:28a,1:28b,5:292,3f:195,1:1bf,29:19e,125:3b9,8b:3b2,1:3b8,1:3c5,3:3c6,1:3c0,1a:3ba,1:3c1,1:3c3,2:3b8,1:3b5,1bc9:3b9,1c:1f76,1:1f77,f:1f7a,1:1f7b,d:1f78,1:1f79,1:1f7c,1:1f7d,107:63,5:25b,4:68,1:68,1:68,3:69,1:69,1:6c,3:6e,4:70,1:71,1:72,1:72,1:72,7:7a,2:3c9,2:7a,2:6b,1:e5,1:62,1:63,3:65,1:66,2:6d,b:3b3,1:3c0,6:64,1b574:3b8,1a:3c3,20:3b8,1a:3c3,20:3b8,1a:3c3,20:3b8,1a:3c3,20:3b8,1a:3c3'
    ),
    Ar = mr(
      '179:1,2:1,2:1,5:1,2:1,a:4f,a:1,8:1,2:1,2:1,3:1,5:1,3:1,4:1,2:1,3:1,4:1,8:2,1:1,2:2,1:1,2:2,27:2,195:26,2:25,1:25,1:25,2:40,2:3f,1:3f,33:1,11:-6,1:-9,1ac7:-3a,6d:-8,1:-8,1:-8,1:-8,1:-8,1:-8,1:-8,1:-8,9:-8,1:-8,1:-8,1:-8,1:-8,1:-8,b:-8,1:-8,1:-8,1:-8,1:-8,1:-8,1:-8,1:-8,9:-8,1:-8,1:-8,1:-8,1:-8,1:-8,1:-8,1:-8,9:-8,1:-8,1:-8,1:-8,1:-8,1:-8,c:-8,2:-8,2:-8,2:-8,9:-8,1:-8,1:-8,1:-8,1:-8,1:-8,1:-8,1:-8,49:-8,1:-8,1:-4a,1:-4a,d:-56,1:-56,1:-56,1:-56,d:-8,1:-8,f:-8,1:-8,3:-7'
    ),
    Er = mr(
      'df:00730073,51:00690307,19:02BC006E,a7:006A030C,18a:002003B9,16:03B903080301,20:03C503080301,1d7:05650582,190f:00680331,1:00740308,1:0077030A,1:0079030A,1:006102BE,b6:03C50313,2:03C503130300,2:03C503130301,2:03C503130342,2a:1F0003B9,1:1F0103B9,1:1F0203B9,1:1F0303B9,1:1F0403B9,1:1F0503B9,1:1F0603B9,1:1F0703B9,1:1F0003B9,1:1F0103B9,1:1F0203B9,1:1F0303B9,1:1F0403B9,1:1F0503B9,1:1F0603B9,1:1F0703B9,1:1F2003B9,1:1F2103B9,1:1F2203B9,1:1F2303B9,1:1F2403B9,1:1F2503B9,1:1F2603B9,1:1F2703B9,1:1F2003B9,1:1F2103B9,1:1F2203B9,1:1F2303B9,1:1F2403B9,1:1F2503B9,1:1F2603B9,1:1F2703B9,1:1F6003B9,1:1F6103B9,1:1F6203B9,1:1F6303B9,1:1F6403B9,1:1F6503B9,1:1F6603B9,1:1F6703B9,1:1F6003B9,1:1F6103B9,1:1F6203B9,1:1F6303B9,1:1F6403B9,1:1F6503B9,1:1F6603B9,1:1F6703B9,3:1F7003B9,1:03B103B9,1:03AC03B9,2:03B10342,1:03B1034203B9,5:03B103B9,6:1F7403B9,1:03B703B9,1:03AE03B9,2:03B70342,1:03B7034203B9,5:03B703B9,6:03B903080300,1:03B903080301,3:03B90342,1:03B903080342,b:03C503080300,1:03C503080301,1:03C10313,2:03C50342,1:03C503080342,b:1F7C03B9,1:03C903B9,1:03CE03B9,2:03C90342,1:03C9034203B9,5:03C903B9,ac:00720073,5b:00B00063,6:00B00066,d:006E006F,a:0073006D,1:00740065006C,1:0074006D,124f:006800700061,2:00610075,2:006F0076,b:00700061,1:006E0061,1:03BC0061,1:006D0061,1:006B0061,1:006B0062,1:006D0062,1:00670062,3:00700066,1:006E0066,1:03BC0066,4:0068007A,1:006B0068007A,1:006D0068007A,1:00670068007A,1:00740068007A,15:00700061,1:006B00700061,1:006D00700061,1:006700700061,8:00700076,1:006E0076,1:03BC0076,1:006D0076,1:006B0076,1:006D0076,1:00700077,1:006E0077,1:03BC0077,1:006D0077,1:006B0077,1:006D0077,1:006B03C9,1:006D03C9,2:00620071,3:00632215006B0067,1:0063006F002E,1:00640062,1:00670079,2:00680070,2:006B006B,1:006B006D,9:00700068,2:00700070006D,1:00700072,2:00730076,1:00770062,c723:00660066,1:00660069,1:0066006C,1:006600660069,1:00660066006C,1:00730074,1:00730074,d:05740576,1:05740565,1:0574056B,1:057E0576,1:0574056D',
      function (t) {
        if (t.length % 4 != 0) throw new Error('bad data');
        let e = [];
        for (let r = 0; r < t.length; r += 4)
          e.push(parseInt(t.substring(r, r + 4), 16));
        return e;
      }
    ),
    Mr = pr(
      '80-20,2a0-,39c,32,f71,18e,7f2-f,19-7,30-4,7-5,f81-b,5,a800-20ff,4d1-1f,110,fa-6,d174-7,2e84-,ffff-,ffff-,ffff-,ffff-,ffff-,ffff-,ffff-,ffff-,ffff-,ffff-,ffff-,ffff-,2,1f-5f,ff7f-20001'
    );
  class _r extends ze {
    constructor(t) {
      super('string', t);
    }
    defaultValue() {
      return '';
    }
    encode(t, e) {
      return super.encode(t, hr(e));
    }
    decode(t) {
      return fr(super.decode(t));
    }
  }
  class kr extends zt {
    constructor(t, e) {
      let r = !1;
      const n = [];
      t.forEach((t) => {
        t.dynamic && (r = !0), n.push(t.type);
      });
      super('tuple', 'tuple(' + n.join(',') + ')', e, r), (this.coders = t);
    }
    defaultValue() {
      const t = [];
      this.coders.forEach((e) => {
        t.push(e.defaultValue());
      });
      const e = this.coders.reduce((t, e) => {
        const r = e.localName;
        return r && (t[r] || (t[r] = 0), t[r]++), t;
      }, {});
      return (
        this.coders.forEach((r, n) => {
          let i = r.localName;
          i &&
            1 === e[i] &&
            ('length' === i && (i = '_length'), null == t[i] && (t[i] = t[n]));
        }),
        Object.freeze(t)
      );
    }
    encode(t, e) {
      return De(t, this.coders, e);
    }
    decode(t) {
      return t.coerce(this.name, Le(t, this.coders));
    }
  }
  const Sr = new f(gt),
    Nr = new RegExp(/^bytes([0-9]*)$/),
    Pr = new RegExp(/^(u?int)([0-9]*)$/);
  const xr = new (class {
    constructor(t) {
      st(this, 'coerceFunc', t || null);
    }
    _getCoder(t) {
      switch (t.baseType) {
        case 'address':
          return new Be(t.name);
        case 'bool':
          return new je(t.name);
        case 'string':
          return new _r(t.name);
        case 'bytes':
          return new Ge(t.name);
        case 'array':
          return new Ue(this._getCoder(t.arrayChildren), t.arrayLength, t.name);
        case 'tuple':
          return new kr(
            (t.components || []).map((t) => this._getCoder(t)),
            t.name
          );
        case '':
          return new Ke(t.name);
      }
      let e = t.type.match(Pr);
      if (e) {
        let r = parseInt(e[2] || '256');
        return (
          (0 === r || r > 256 || r % 8 != 0) &&
            Sr.throwArgumentError(
              'invalid ' + e[1] + ' bit length',
              'param',
              t
            ),
          new rr(r / 8, 'int' === e[1], t.name)
        );
      }
      if (((e = t.type.match(Nr)), e)) {
        let r = parseInt(e[1]);
        return (
          (0 === r || r > 32) &&
            Sr.throwArgumentError('invalid bytes length', 'param', t),
          new qe(r, t.name)
        );
      }
      return Sr.throwArgumentError('invalid type', 'type', t.type);
    }
    _getWordSize() {
      return 32;
    }
    _getReader(t, e) {
      return new qt(t, this._getWordSize(), this.coerceFunc, e);
    }
    _getWriter() {
      return new Gt(this._getWordSize());
    }
    getDefaultValue(t) {
      const e = t.map((t) => this._getCoder(kt.from(t)));
      return new kr(e, '_').defaultValue();
    }
    encode(t, e) {
      t.length !== e.length &&
        Sr.throwError(
          'types/values length mismatch',
          f.errors.INVALID_ARGUMENT,
          {
            count: { types: t.length, values: e.length },
            value: { types: t, values: e },
          }
        );
      const r = t.map((t) => this._getCoder(kt.from(t))),
        n = new kr(r, '_'),
        i = this._getWriter();
      return n.encode(i, e), i.data;
    }
    decode(t, e, r) {
      const n = t.map((t) => this._getCoder(kt.from(t)));
      return new kr(n, '_').decode(this._getReader(w(e), r));
    }
  })();
  function Tr(t) {
    return ye(hr(t));
  }
  const Cr = 'hash/5.7.0';
  function Rr(t) {
    t = atob(t);
    const e = [];
    for (let r = 0; r < t.length; r++) e.push(t.charCodeAt(r));
    return w(e);
  }
  function Ir(t) {
    t = w(t);
    let e = '';
    for (let r = 0; r < t.length; r++) e += String.fromCharCode(t[r]);
    return btoa(e);
  }
  var Br = Object.freeze({ __proto__: null, decode: Rr, encode: Ir });
  function Or(t, e) {
    null == e && (e = 1);
    const r = [],
      n = r.forEach,
      i = function (t, e) {
        n.call(t, function (t) {
          e > 0 && Array.isArray(t) ? i(t, e - 1) : r.push(t);
        });
      };
    return i(t, e), r;
  }
  function Fr(t) {
    return (function (t) {
      let e = 0;
      return () => t[e++];
    })(
      (function (t) {
        let e = 0;
        function r() {
          return (t[e++] << 8) | t[e++];
        }
        let n = r(),
          i = 1,
          o = [0, 1];
        for (let t = 1; t < n; t++) o.push((i += r()));
        let s = r(),
          a = e;
        e += s;
        let u = 0,
          h = 0;
        function l() {
          return 0 == u && ((h = (h << 8) | t[e++]), (u = 8)), (h >> --u) & 1;
        }
        const c = Math.pow(2, 31),
          f = c >>> 1,
          d = f >> 1,
          m = c - 1;
        let p = 0;
        for (let t = 0; t < 31; t++) p = (p << 1) | l();
        let g = [],
          y = 0,
          v = c;
        for (;;) {
          let t = Math.floor(((p - y + 1) * i - 1) / v),
            e = 0,
            r = n;
          for (; r - e > 1; ) {
            let n = (e + r) >>> 1;
            t < o[n] ? (r = n) : (e = n);
          }
          if (0 == e) break;
          g.push(e);
          let s = y + Math.floor((v * o[e]) / i),
            a = y + Math.floor((v * o[e + 1]) / i) - 1;
          for (; !((s ^ a) & f); )
            (p = ((p << 1) & m) | l()),
              (s = (s << 1) & m),
              (a = ((a << 1) & m) | 1);
          for (; s & ~a & d; )
            (p = (p & f) | ((p << 1) & (m >>> 1)) | l()),
              (s = (s << 1) ^ f),
              (a = ((a ^ f) << 1) | f | 1);
          (y = s), (v = 1 + a - s);
        }
        let w = n - 4;
        return g.map((e) => {
          switch (e - w) {
            case 3:
              return w + 65792 + ((t[a++] << 16) | (t[a++] << 8) | t[a++]);
            case 2:
              return w + 256 + ((t[a++] << 8) | t[a++]);
            case 1:
              return w + t[a++];
            default:
              return e - 1;
          }
        });
      })(t)
    );
  }
  function Dr(t) {
    return 1 & t ? ~t >> 1 : t >> 1;
  }
  function Lr(t, e) {
    let r = Array(t);
    for (let n = 0, i = -1; n < t; n++) r[n] = i += 1 + e();
    return r;
  }
  function Ur(t, e) {
    let r = Array(t);
    for (let n = 0, i = 0; n < t; n++) r[n] = i += Dr(e());
    return r;
  }
  function jr(t, e) {
    let r = Lr(t(), t),
      n = t(),
      i = Lr(n, t),
      o = (function (t, e) {
        let r = Array(t);
        for (let n = 0; n < t; n++) r[n] = 1 + e();
        return r;
      })(n, t);
    for (let t = 0; t < n; t++) for (let e = 0; e < o[t]; e++) r.push(i[t] + e);
    return e ? r.map((t) => e[t]) : r;
  }
  function zr(t, e, r) {
    let n = Array(t)
      .fill(void 0)
      .map(() => []);
    for (let i = 0; i < e; i++) Ur(t, r).forEach((t, e) => n[e].push(t));
    return n;
  }
  function Gr(t, e) {
    let r = 1 + e(),
      n = e(),
      i = (function (t) {
        let e = [];
        for (;;) {
          let r = t();
          if (0 == r) break;
          e.push(r);
        }
        return e;
      })(e);
    return Or(
      zr(i.length, 1 + t, e).map((t, e) => {
        const o = t[0],
          s = t.slice(1);
        return Array(i[e])
          .fill(void 0)
          .map((t, e) => {
            let i = e * n;
            return [o + e * r, s.map((t) => t + i)];
          });
      })
    );
  }
  function qr(t, e) {
    return zr(1 + e(), 1 + t, e).map((t) => [t[0], t.slice(1)]);
  }
  const Kr = Fr(
      Rr(
        'AEQF2AO2DEsA2wIrAGsBRABxAN8AZwCcAEwAqgA0AGwAUgByADcATAAVAFYAIQAyACEAKAAYAFgAGwAjABQAMAAmADIAFAAfABQAKwATACoADgAbAA8AHQAYABoAGQAxADgALAAoADwAEwA9ABMAGgARAA4ADwAWABMAFgAIAA8AHgQXBYMA5BHJAS8JtAYoAe4AExozi0UAH21tAaMnBT8CrnIyhrMDhRgDygIBUAEHcoFHUPe8AXBjAewCjgDQR8IICIcEcQLwATXCDgzvHwBmBoHNAqsBdBcUAykgDhAMShskMgo8AY8jqAQfAUAfHw8BDw87MioGlCIPBwZCa4ELatMAAMspJVgsDl8AIhckSg8XAHdvTwBcIQEiDT4OPhUqbyECAEoAS34Aej8Ybx83JgT/Xw8gHxZ/7w8RICxPHA9vBw+Pfw8PHwAPFv+fAsAvCc8vEr8ivwD/EQ8Bol8OEBa/A78hrwAPCU8vESNvvwWfHwNfAVoDHr+ZAAED34YaAdJPAK7PLwSEgDLHAGo1Pz8Pvx9fUwMrpb8O/58VTzAPIBoXIyQJNF8hpwIVAT8YGAUADDNBaX3RAMomJCg9EhUeA29MABsZBTMNJipjOhc19gcIDR8bBwQHEggCWi6DIgLuAQYA+BAFCha3A5XiAEsqM7UFFgFLhAMjFTMYE1Klnw74nRVBG/ASCm0BYRN/BrsU3VoWy+S0vV8LQx+vN8gF2AC2AK5EAWwApgYDKmAAroQ0NDQ0AT+OCg7wAAIHRAbpNgVcBV0APTA5BfbPFgMLzcYL/QqqA82eBALKCjQCjqYCht0/k2+OAsXQAoP3ASTKDgDw6ACKAUYCMpIKJpRaAE4A5womABzZvs0REEKiACIQAd5QdAECAj4Ywg/wGqY2AVgAYADYvAoCGAEubA0gvAY2ALAAbpbvqpyEAGAEpgQAJgAG7gAgAEACmghUFwCqAMpAINQIwC4DthRAAPcycKgApoIdABwBfCisABoATwBqASIAvhnSBP8aH/ECeAKXAq40NjgDBTwFYQU6AXs3oABgAD4XNgmcCY1eCl5tIFZeUqGgyoNHABgAEQAaABNwWQAmABMATPMa3T34ADldyprmM1M2XociUQgLzvwAXT3xABgAEQAaABNwIGFAnADD8AAgAD4BBJWzaCcIAIEBFMAWwKoAAdq9BWAF5wLQpALEtQAKUSGkahR4GnJM+gsAwCgeFAiUAECQ0BQuL8AAIAAAADKeIheclvFqQAAETr4iAMxIARMgAMIoHhQIAn0E0pDQFC4HhznoAAAAIAI2C0/4lvFqQAAETgBJJwYCAy4ABgYAFAA8MBKYEH4eRhTkAjYeFcgACAYAeABsOqyQ5gRwDayqugEgaIIAtgoACgDmEABmBAWGme5OBJJA2m4cDeoAmITWAXwrMgOgAGwBCh6CBXYF1Tzg1wKAAFdiuABRAFwAXQBsAG8AdgBrAHYAbwCEAHEwfxQBVE5TEQADVFhTBwBDANILAqcCzgLTApQCrQL6vAAMAL8APLhNBKkE6glGKTAU4Dr4N2EYEwBCkABKk8rHAbYBmwIoAiU4Ajf/Aq4CowCAANIChzgaNBsCsTgeODcFXrgClQKdAqQBiQGYAqsCsjTsNHsfNPA0ixsAWTWiOAMFPDQSNCk2BDZHNow2TTZUNhk28Jk9VzI3QkEoAoICoQKwAqcAQAAxBV4FXbS9BW47YkIXP1ciUqs05DS/FwABUwJW11e6nHuYZmSh/RAYA8oMKvZ8KASoUAJYWAJ6ILAsAZSoqjpgA0ocBIhmDgDWAAawRDQoAAcuAj5iAHABZiR2AIgiHgCaAU68ACxuHAG0ygM8MiZIAlgBdF4GagJqAPZOHAMuBgoATkYAsABiAHgAMLoGDPj0HpKEBAAOJgAuALggTAHWAeAMEDbd20Uege0ADwAWADkAQgA9OHd+2MUQZBBhBgNNDkxxPxUQArEPqwvqERoM1irQ090ANK4H8ANYB/ADWANYB/AH8ANYB/ADWANYA1gDWBwP8B/YxRBkD00EcgWTBZAE2wiIJk4RhgctCNdUEnQjHEwDSgEBIypJITuYMxAlR0wRTQgIATZHbKx9PQNMMbBU+pCnA9AyVDlxBgMedhKlAC8PeCE1uk6DekxxpQpQT7NX9wBFBgASqwAS5gBJDSgAUCwGPQBI4zTYABNGAE2bAE3KAExdGABKaAbgAFBXAFCOAFBJABI2SWdObALDOq0//QomCZhvwHdTBkIQHCemEPgMNAG2ATwN7kvZBPIGPATKH34ZGg/OlZ0Ipi3eDO4m5C6igFsj9iqEBe5L9TzeC05RaQ9aC2YJ5DpkgU8DIgEOIowK3g06CG4Q9ArKbA3mEUYHOgPWSZsApgcCCxIdNhW2JhFirQsKOXgG/Br3C5AmsBMqev0F1BoiBk4BKhsAANAu6IWxWjJcHU9gBgQLJiPIFKlQIQ0mQLh4SRocBxYlqgKSQ3FKiFE3HpQh9zw+DWcuFFF9B/Y8BhlQC4I8n0asRQ8R0z6OPUkiSkwtBDaALDAnjAnQD4YMunxzAVoJIgmyDHITMhEYN8YIOgcaLpclJxYIIkaWYJsE+KAD9BPSAwwFQAlCBxQDthwuEy8VKgUOgSXYAvQ21i60ApBWgQEYBcwPJh/gEFFH4Q7qCJwCZgOEJewALhUiABginAhEZABgj9lTBi7MCMhqbSN1A2gU6GIRdAeSDlgHqBw0FcAc4nDJXgyGCSiksAlcAXYJmgFgBOQICjVcjKEgQmdUi1kYnCBiQUBd/QIyDGYVoES+h3kCjA9sEhwBNgF0BzoNAgJ4Ee4RbBCWCOyGBTW2M/k6JgRQIYQgEgooA1BszwsoJvoM+WoBpBJjAw00PnfvZ6xgtyUX/gcaMsZBYSHyC5NPzgydGsIYQ1QvGeUHwAP0GvQn60FYBgADpAQUOk4z7wS+C2oIjAlAAEoOpBgH2BhrCnKM0QEyjAG4mgNYkoQCcJAGOAcMAGgMiAV65gAeAqgIpAAGANADWAA6Aq4HngAaAIZCAT4DKDABIuYCkAOUCDLMAZYwAfQqBBzEDBYA+DhuSwLDsgKAa2ajBd5ZAo8CSjYBTiYEBk9IUgOwcuIA3ABMBhTgSAEWrEvMG+REAeBwLADIAPwABjYHBkIBzgH0bgC4AWALMgmjtLYBTuoqAIQAFmwB2AKKAN4ANgCA8gFUAE4FWvoF1AJQSgESMhksWGIBvAMgATQBDgB6BsyOpsoIIARuB9QCEBwV4gLvLwe2AgMi4BPOQsYCvd9WADIXUu5eZwqoCqdeaAC0YTQHMnM9UQAPH6k+yAdy/BZIiQImSwBQ5gBQQzSaNTFWSTYBpwGqKQK38AFtqwBI/wK37gK3rQK3sAK6280C0gK33AK3zxAAUEIAUD9SklKDArekArw5AEQAzAHCO147WTteO1k7XjtZO147WTteO1kDmChYI03AVU0oJqkKbV9GYewMpw3VRMk6ShPcYFJgMxPJLbgUwhXPJVcZPhq9JwYl5VUKDwUt1GYxCC00dhe9AEApaYNCY4ceMQpMHOhTklT5LRwAskujM7ANrRsWREEFSHXuYisWDwojAmSCAmJDXE6wXDchAqH4AmiZAmYKAp+FOBwMAmY8AmYnBG8EgAN/FAN+kzkHOXgYOYM6JCQCbB4CMjc4CwJtyAJtr/CLADRoRiwBaADfAOIASwYHmQyOAP8MwwAOtgJ3MAJ2o0ACeUxEAni7Hl3cRa9G9AJ8QAJ6yQJ9CgJ88UgBSH5kJQAsFklZSlwWGErNAtECAtDNSygDiFADh+dExpEzAvKiXQQDA69Lz0wuJgTQTU1NsAKLQAKK2cIcCB5EaAa4Ao44Ao5dQZiCAo7aAo5deVG1UzYLUtVUhgKT/AKTDQDqAB1VH1WwVdEHLBwplocy4nhnRTw6ApegAu+zWCKpAFomApaQApZ9nQCqWa1aCoJOADwClrYClk9cRVzSApnMApllXMtdCBoCnJw5wzqeApwXAp+cAp65iwAeEDIrEAKd8gKekwC2PmE1YfACntQCoG8BqgKeoCACnk+mY8lkKCYsAiewAiZ/AqD8AqBN2AKmMAKlzwKoAAB+AqfzaH1osgAESmodatICrOQCrK8CrWgCrQMCVx4CVd0CseLYAx9PbJgCsr4OArLpGGzhbWRtSWADJc4Ctl08QG6RAylGArhfArlIFgK5K3hwN3DiAr0aAy2zAzISAr6JcgMDM3ICvhtzI3NQAsPMAsMFc4N0TDZGdOEDPKgDPJsDPcACxX0CxkgCxhGKAshqUgLIRQLJUALJLwJkngLd03h6YniveSZL0QMYpGcDAmH1GfSVJXsMXpNevBICz2wCz20wTFTT9BSgAMeuAs90ASrrA04TfkwGAtwoAtuLAtJQA1JdA1NgAQIDVY2AikABzBfuYUZ2AILPg44C2sgC2d+EEYRKpz0DhqYAMANkD4ZyWvoAVgLfZgLeuXR4AuIw7RUB8zEoAfScAfLTiALr9ALpcXoAAur6AurlAPpIAboC7ooC652Wq5cEAu5AA4XhmHpw4XGiAvMEAGoDjheZlAL3FAORbwOSiAL3mQL52gL4Z5odmqy8OJsfA52EAv77ARwAOp8dn7QDBY4DpmsDptoA0sYDBmuhiaIGCgMMSgFgASACtgNGAJwEgLpoBgC8BGzAEowcggCEDC6kdjoAJAM0C5IKRoABZCgiAIzw3AYBLACkfng9ogigkgNmWAN6AEQCvrkEVqTGAwCsBRbAA+4iQkMCHR072jI2PTbUNsk2RjY5NvA23TZKNiU3EDcZN5I+RTxDRTBCJkK5VBYKFhZfwQCWygU3AJBRHpu+OytgNxa61A40GMsYjsn7BVwFXQVcBV0FaAVdBVwFXQVcBV0FXAVdBVwFXUsaCNyKAK4AAQUHBwKU7oICoW1e7jAEzgPxA+YDwgCkBFDAwADABKzAAOxFLhitA1UFTDeyPkM+bj51QkRCuwTQWWQ8X+0AWBYzsACNA8xwzAGm7EZ/QisoCTAbLDs6fnLfb8H2GccsbgFw13M1HAVkBW/Jxsm9CNRO8E8FDD0FBQw9FkcClOYCoMFegpDfADgcMiA2AJQACB8AsigKAIzIEAJKeBIApY5yPZQIAKQiHb4fvj5BKSRPQrZCOz0oXyxgOywfKAnGbgMClQaCAkILXgdeCD9IIGUgQj5fPoY+dT52Ao5CM0dAX9BTVG9SDzFwWTQAbxBzJF/lOEIQQglCCkKJIAls5AcClQICoKPMODEFxhi6KSAbiyfIRrMjtCgdWCAkPlFBIitCsEJRzAbMAV/OEyQzDg0OAQQEJ36i328/Mk9AybDJsQlq3tDRApUKAkFzXf1d/j9uALYP6hCoFgCTGD8kPsFKQiobrm0+zj0KSD8kPnVCRBwMDyJRTHFgMTJa5rwXQiQ2YfI/JD7BMEJEHGINTw4TOFlIRzwJO0icMQpyPyQ+wzJCRBv6DVgnKB01NgUKj2bwYzMqCoBkznBgEF+zYDIocwRIX+NgHj4HICNfh2C4CwdwFWpTG/lgUhYGAwRfv2Ts8mAaXzVgml/XYIJfuWC4HI1gUF9pYJZgMR6ilQHMAOwLAlDRefC0in4AXAEJA6PjCwc0IamOANMMCAECRQDFNRTZBgd+CwQlRA+r6+gLBDEFBnwUBXgKATIArwAGRAAHA3cDdAN2A3kDdwN9A3oDdQN7A30DfAN4A3oDfQAYEAAlAtYASwMAUAFsAHcKAHcAmgB3AHUAdQB2AHVu8UgAygDAAHcAdQB1AHYAdQALCgB3AAsAmgB3AAsCOwB3AAtu8UgAygDAAHgKAJoAdwB3AHUAdQB2AHUAeAB1AHUAdgB1bvFIAMoAwAALCgCaAHcACwB3AAsCOwB3AAtu8UgAygDAAH4ACwGgALcBpwC6AahdAu0COwLtbvFIAMoAwAALCgCaAu0ACwLtAAsCOwLtAAtu8UgAygDAA24ACwNvAAu0VsQAAzsAABCkjUIpAAsAUIusOggWcgMeBxVsGwL67U/2HlzmWOEeOgALASvuAAseAfpKUpnpGgYJDCIZM6YyARUE9ThqAD5iXQgnAJYJPnOzw0ZAEZxEKsIAkA4DhAHnTAIDxxUDK0lxCQlPYgIvIQVYJQBVqE1GakUAKGYiDToSBA1EtAYAXQJYAIF8GgMHRyAAIAjOe9YncekRAA0KACUrjwE7Ayc6AAYWAqaiKG4McEcqANoN3+Mg9TwCBhIkuCny+JwUQ29L008JluRxu3K+oAdqiHOqFH0AG5SUIfUJ5SxCGfxdipRzqTmT4V5Zb+r1Uo4Vm+NqSSEl2mNvR2JhIa8SpYO6ntdwFXHCWTCK8f2+Hxo7uiG3drDycAuKIMP5bhi06ACnqArH1rz4Rqg//lm6SgJGEVbF9xJHISaR6HxqxSnkw6shDnelHKNEfGUXSJRJ1GcsmtJw25xrZMDK9gXSm1/YMkdX4/6NKYOdtk/NQ3/NnDASjTc3fPjIjW/5sVfVObX2oTDWkr1dF9f3kxBsD3/3aQO8hPfRz+e0uEiJqt1161griu7gz8hDDwtpy+F+BWtefnKHZPAxcZoWbnznhJpy0e842j36bcNzGnIEusgGX0a8ZxsnjcSsPDZ09yZ36fCQbriHeQ72JRMILNl6ePPf2HWoVwgWAm1fb3V2sAY0+B6rAXqSwPBgseVmoqsBTSrm91+XasMYYySI8eeRxH3ZvHkMz3BQ5aJ3iUVbYPNM3/7emRtjlsMgv/9VyTsyt/mK+8fgWeT6SoFaclXqn42dAIsvAarF5vNNWHzKSkKQ/8Hfk5ZWK7r9yliOsooyBjRhfkHP4Q2DkWXQi6FG/9r/IwbmkV5T7JSopHKn1pJwm9tb5Ot0oyN1Z2mPpKXHTxx2nlK08fKk1hEYA8WgVVWL5lgx0iTv+KdojJeU23ZDjmiubXOxVXJKKi2Wjuh2HLZOFLiSC7Tls5SMh4f+Pj6xUSrNjFqLGehRNB8lC0QSLNmkJJx/wSG3MnjE9T1CkPwJI0wH2lfzwETIiVqUxg0dfu5q39Gt+hwdcxkhhNvQ4TyrBceof3Mhs/IxFci1HmHr4FMZgXEEczPiGCx0HRwzAqDq2j9AVm1kwN0mRVLWLylgtoPNapF5cY4Y1wJh/e0BBwZj44YgZrDNqvD/9Hv7GFYdUQeDJuQ3EWI4HaKqavU1XjC/n41kT4L79kqGq0kLhdTZvgP3TA3fS0ozVz+5piZsoOtIvBUFoMKbNcmBL6YxxaUAusHB38XrS8dQMnQwJfUUkpRoGr5AUeWicvBTzyK9g77+yCkf5PAysL7r/JjcZgrbvRpMW9iyaxZvKO6ceZN2EwIxKwVFPuvFuiEPGCoagbMo+SpydLrXqBzNCDGFCrO/rkcwa2xhokQZ5CdZ0AsU3JfSqJ6n5I14YA+P/uAgfhPU84Tlw7cEFfp7AEE8ey4sP12PTt4Cods1GRgDOB5xvyiR5m+Bx8O5nBCNctU8BevfV5A08x6RHd5jcwPTMDSZJOedIZ1cGQ704lxbAzqZOP05ZxaOghzSdvFBHYqomATARyAADK4elP8Ly3IrUZKfWh23Xy20uBUmLS4Pfagu9+oyVa2iPgqRP3F2CTUsvJ7+RYnN8fFZbU/HVvxvcFFDKkiTqV5UBZ3Gz54JAKByi9hkKMZJvuGgcSYXFmw08UyoQyVdfTD1/dMkCHXcTGAKeROgArsvmRrQTLUOXioOHGK2QkjHuoYFgXciZoTJd6Fs5q1QX1G+p/e26hYsEf7QZD1nnIyl/SFkNtYYmmBhpBrxl9WbY0YpHWRuw2Ll/tj9mD8P4snVzJl4F9J+1arVeTb9E5r2ILH04qStjxQNwn3m4YNqxmaNbLAqW2TN6LidwuJRqS+NXbtqxoeDXpxeGWmxzSkWxjkyCkX4NQRme6q5SAcC+M7+9ETfA/EwrzQajKakCwYyeunP6ZFlxU2oMEn1Pz31zeStW74G406ZJFCl1wAXIoUKkWotYEpOuXB1uVNxJ63dpJEqfxBeptwIHNrPz8BllZoIcBoXwgfJ+8VAUnVPvRvexnw0Ma/WiGYuJO5y8QTvEYBigFmhUxY5RqzE8OcywN/8m4UYrlaniJO75XQ6KSo9+tWHlu+hMi0UVdiKQp7NelnoZUzNaIyBPVeOwK6GNp+FfHuPOoyhaWuNvTYFkvxscMQWDh+zeFCFkgwbXftiV23ywJ4+uwRqmg9k3KzwIQpzppt8DBBOMbrqwQM5Gb05sEwdKzMiAqOloaA/lr0KA+1pr0/+HiWoiIjHA/wir2nIuS3PeU/ji3O6ZwoxcR1SZ9FhtLC5S0FIzFhbBWcGVP/KpxOPSiUoAdWUpqKH++6Scz507iCcxYI6rdMBICPJZea7OcmeFw5mObJSiqpjg2UoWNIs+cFhyDSt6geV5qgi3FunmwwDoGSMgerFOZGX1m0dMCYo5XOruxO063dwENK9DbnVM9wYFREzh4vyU1WYYJ/LRRp6oxgjqP/X5a8/4Af6p6NWkQferzBmXme0zY/4nwMJm/wd1tIqSwGz+E3xPEAOoZlJit3XddD7/BT1pllzOx+8bmQtANQ/S6fZexc6qi3W+Q2xcmXTUhuS5mpHQRvcxZUN0S5+PL9lXWUAaRZhEH8hTdAcuNMMCuVNKTEGtSUKNi3O6KhSaTzck8csZ2vWRZ+d7mW8c4IKwXIYd25S/zIftPkwPzufjEvOHWVD1m+FjpDVUTV0DGDuHj6QnaEwLu/dEgdLQOg9E1Sro9XHJ8ykLAwtPu+pxqKDuFexqON1sKQm7rwbE1E68UCfA/erovrTCG+DBSNg0l4goDQvZN6uNlbyLpcZAwj2UclycvLpIZMgv4yRlpb3YuMftozorbcGVHt/VeDV3+Fdf1TP0iuaCsPi2G4XeGhsyF1ubVDxkoJhmniQ0/jSg/eYML9KLfnCFgISWkp91eauR3IQvED0nAPXK+6hPCYs+n3+hCZbiskmVMG2da+0EsZPonUeIY8EbfusQXjsK/eFDaosbPjEfQS0RKG7yj5GG69M7MeO1HmiUYocgygJHL6M1qzUDDwUSmr99V7Sdr2F3JjQAJY+F0yH33Iv3+C9M38eML7gTgmNu/r2bUMiPvpYbZ6v1/IaESirBHNa7mPKn4dEmYg7v/+HQgPN1G79jBQ1+soydfDC2r+h2Bl/KIc5KjMK7OH6nb1jLsNf0EHVe2KBiE51ox636uyG6Lho0t3J34L5QY/ilE3mikaF4HKXG1mG1rCevT1Vv6GavltxoQe/bMrpZvRggnBxSEPEeEzkEdOxTnPXHVjUYdw8JYvjB/o7Eegc3Ma+NUxLLnsK0kJlinPmUHzHGtrk5+CAbVzFOBqpyy3QVUnzTDfC/0XD94/okH+OB+i7g9lolhWIjSnfIb+Eq43ZXOWmwvjyV/qqD+t0e+7mTEM74qP/Ozt8nmC7mRpyu63OB4KnUzFc074SqoyPUAgM+/TJGFo6T44EHnQU4X4z6qannVqgw/U7zCpwcmXV1AubIrvOmkKHazJAR55ePjp5tLBsN8vAqs3NAHdcEHOR2xQ0lsNAFzSUuxFQCFYvXLZJdOj9p4fNq6p0HBGUik2YzaI4xySy91KzhQ0+q1hjxvImRwPRf76tChlRkhRCi74NXZ9qUNeIwP+s5p+3m5nwPdNOHgSLD79n7O9m1n1uDHiMntq4nkYwV5OZ1ENbXxFd4PgrlvavZsyUO4MqYlqqn1O8W/I1dEZq5dXhrbETLaZIbC2Kj/Aa/QM+fqUOHdf0tXAQ1huZ3cmWECWSXy/43j35+Mvq9xws7JKseriZ1pEWKc8qlzNrGPUGcVgOa9cPJYIJsGnJTAUsEcDOEVULO5x0rXBijc1lgXEzQQKhROf8zIV82w8eswc78YX11KYLWQRcgHNJElBxfXr72lS2RBSl07qTKorO2uUDZr3sFhYsvnhLZn0A94KRzJ/7DEGIAhW5ZWFpL8gEwu1aLA9MuWZzNwl8Oze9Y+bX+v9gywRVnoB5I/8kXTXU3141yRLYrIOOz6SOnyHNy4SieqzkBXharjfjqq1q6tklaEbA8Qfm2DaIPs7OTq/nvJBjKfO2H9bH2cCMh1+5gspfycu8f/cuuRmtDjyqZ7uCIMyjdV3a+p3fqmXsRx4C8lujezIFHnQiVTXLXuI1XrwN3+siYYj2HHTvESUx8DlOTXpak9qFRK+L3mgJ1WsD7F4cu1aJoFoYQnu+wGDMOjJM3kiBQWHCcvhJ/HRdxodOQp45YZaOTA22Nb4XKCVxqkbwMYFhzYQYIAnCW8FW14uf98jhUG2zrKhQQ0q0CEq0t5nXyvUyvR8DvD69LU+g3i+HFWQMQ8PqZuHD+sNKAV0+M6EJC0szq7rEr7B5bQ8BcNHzvDMc9eqB5ZCQdTf80Obn4uzjwpYU7SISdtV0QGa9D3Wrh2BDQtpBKxaNFV+/Cy2P/Sv+8s7Ud0Fd74X4+o/TNztWgETUapy+majNQ68Lq3ee0ZO48VEbTZYiH1Co4OlfWef82RWeyUXo7woM03PyapGfikTnQinoNq5z5veLpeMV3HCAMTaZmA1oGLAn7XS3XYsz+XK7VMQsc4XKrmDXOLU/pSXVNUq8dIqTba///3x6LiLS6xs1xuCAYSfcQ3+rQgmu7uvf3THKt5Ooo97TqcbRqxx7EASizaQCBQllG/rYxVapMLgtLbZS64w1MDBMXX+PQpBKNwqUKOf2DDRDUXQf9EhOS0Qj4nTmlA8dzSLz/G1d+Ud8MTy/6ghhdiLpeerGY/UlDOfiuqFsMUU5/UYlP+BAmgRLuNpvrUaLlVkrqDievNVEAwF+4CoM1MZTmjxjJMsKJq+u8Zd7tNCUFy6LiyYXRJQ4VyvEQFFaCGKsxIwQkk7EzZ6LTJq2hUuPhvAW+gQnSG6J+MszC+7QCRHcnqDdyNRJ6T9xyS87A6MDutbzKGvGktpbXqtzWtXb9HsfK2cBMomjN9a4y+TaJLnXxAeX/HWzmf4cR4vALt/P4w4qgKY04ml4ZdLOinFYS6cup3G/1ie4+t1eOnpBNlqGqs75ilzkT4+DsZQxNvaSKJ//6zIbbk/M7LOhFmRc/1R+kBtz7JFGdZm/COotIdvQoXpTqP/1uqEUmCb/QWoGLMwO5ANcHzxdY48IGP5+J+zKOTBFZ4Pid+GTM+Wq12MV/H86xEJptBa6T+p3kgpwLedManBHC2GgNrFpoN2xnrMz9WFWX/8/ygSBkavq2Uv7FdCsLEYLu9LLIvAU0bNRDtzYl+/vXmjpIvuJFYjmI0im6QEYqnIeMsNjXG4vIutIGHijeAG/9EDBozKV5cldkHbLxHh25vT+ZEzbhXlqvpzKJwcEgfNwLAKFeo0/pvEE10XDB+EXRTXtSzJozQKFFAJhMxYkVaCW+E9AL7tMeU8acxidHqzb6lX4691UsDpy/LLRmT+epgW56+5Cw8tB4kMUv6s9lh3eRKbyGs+H/4mQMaYzPTf2OOdokEn+zzgvoD3FqNKk8QqGAXVsqcGdXrT62fSPkR2vROFi68A6se86UxRUk4cajfPyCC4G5wDhD+zNq4jodQ4u4n/m37Lr36n4LIAAsVr02dFi9AiwA81MYs2rm4eDlDNmdMRvEKRHfBwW5DdMNp0jPFZMeARqF/wL4XBfd+EMLBfMzpH5GH6NaW+1vrvMdg+VxDzatk3MXgO3ro3P/DpcC6+Mo4MySJhKJhSR01SGGGp5hPWmrrUgrv3lDnP+HhcI3nt3YqBoVAVTBAQT5iuhTg8nvPtd8ZeYj6w1x6RqGUBrSku7+N1+BaasZvjTk64RoIDlL8brpEcJx3OmY7jLoZsswdtmhfC/G21llXhITOwmvRDDeTTPbyASOa16cF5/A1fZAidJpqju3wYAy9avPR1ya6eNp9K8XYrrtuxlqi+bDKwlfrYdR0RRiKRVTLOH85+ZY7XSmzRpfZBJjaTa81VDcJHpZnZnSQLASGYW9l51ZV/h7eVzTi3Hv6hUsgc/51AqJRTkpbFVLXXszoBL8nBX0u/0jBLT8nH+fJePbrwURT58OY+UieRjd1vs04w0VG5VN2U6MoGZkQzKN/ptz0Q366dxoTGmj7i1NQGHi9GgnquXFYdrCfZBmeb7s0T6yrdlZH5cZuwHFyIJ/kAtGsTg0xH5taAAq44BAk1CPk9KVVbqQzrCUiFdF/6gtlPQ8bHHc1G1W92MXGZ5HEHftyLYs8mbD/9xYRUWkHmlM0zC2ilJlnNgV4bfALpQghxOUoZL7VTqtCHIaQSXm+YUMnpkXybnV+A6xlm2CVy8fn0Xlm2XRa0+zzOa21JWWmixfiPMSCZ7qA4rS93VN3pkpF1s5TonQjisHf7iU9ZGvUPOAKZcR1pbeVf/Ul7OhepGCaId9wOtqo7pJ7yLcBZ0pFkOF28y4zEI/kcUNmutBHaQpBdNM8vjCS6HZRokkeo88TBAjGyG7SR+6vUgTcyK9Imalj0kuxz0wmK+byQU11AiJFk/ya5dNduRClcnU64yGu/ieWSeOos1t3ep+RPIWQ2pyTYVbZltTbsb7NiwSi3AV+8KLWk7LxCnfZUetEM8ThnsSoGH38/nyAwFguJp8FjvlHtcWZuU4hPva0rHfr0UhOOJ/F6vS62FW7KzkmRll2HEc7oUq4fyi5T70Vl7YVIfsPHUCdHesf9Lk7WNVWO75JDkYbMI8TOW8JKVtLY9d6UJRITO8oKo0xS+o99Yy04iniGHAaGj88kEWgwv0OrHdY/nr76DOGNS59hXCGXzTKUvDl9iKpLSWYN1lxIeyywdNpTkhay74w2jFT6NS8qkjo5CxA1yfSYwp6AJIZNKIeEK5PJAW7ORgWgwp0VgzYpqovMrWxbu+DGZ6Lhie1RAqpzm8VUzKJOH3mCzWuTOLsN3VT/dv2eeYe9UjbR8YTBsLz7q60VN1sU51k+um1f8JxD5pPhbhSC8rRaB454tmh6YUWrJI3+GWY0qeWioj/tbkYITOkJaeuGt4JrJvHA+l0Gu7kY7XOaa05alMnRWVCXqFgLIwSY4uF59Ue5SU4QKuc/HamDxbr0x6csCetXGoP7Qn1Bk/J9DsynO/UD6iZ1Hyrz+jit0hDCwi/E9OjgKTbB3ZQKQ/0ZOvevfNHG0NK4Aj3Cp7NpRk07RT1i/S0EL93Ag8GRgKI9CfpajKyK6+Jj/PI1KO5/85VAwz2AwzP8FTBb075IxCXv6T9RVvWT2tUaqxDS92zrGUbWzUYk9mSs82pECH+fkqsDt93VW++4YsR/dHCYcQSYTO/KaBMDj9LSD/J/+z20Kq8XvZUAIHtm9hRPP3ItbuAu2Hm5lkPs92pd7kCxgRs0xOVBnZ13ccdA0aunrwv9SdqElJRC3g+oCu+nXyCgmXUs9yMjTMAIHfxZV+aPKcZeUBWt057Xo85Ks1Ir5gzEHCWqZEhrLZMuF11ziGtFQUds/EESajhagzcKsxamcSZxGth4UII+adPhQkUnx2WyN+4YWR+r3f8MnkyGFuR4zjzxJS8WsQYR5PTyRaD9ixa6Mh741nBHbzfjXHskGDq179xaRNrCIB1z1xRfWfjqw2pHc1zk9xlPpL8sQWAIuETZZhbnmL54rceXVNRvUiKrrqIkeogsl0XXb17ylNb0f4GA9Wd44vffEG8FSZGHEL2fbaTGRcSiCeA8PmA/f6Hz8HCS76fXUHwgwkzSwlI71ekZ7Fapmlk/KC+Hs8hUcw3N2LN5LhkVYyizYFl/uPeVP5lsoJHhhfWvvSWruCUW1ZcJOeuTbrDgywJ/qG07gZJplnTvLcYdNaH0KMYOYMGX+rB4NGPFmQsNaIwlWrfCezxre8zXBrsMT+edVLbLqN1BqB76JH4BvZTqUIMfGwPGEn+EnmTV86fPBaYbFL3DFEhjB45CewkXEAtJxk4/Ms2pPXnaRqdky0HOYdcUcE2zcXq4vaIvW2/v0nHFJH2XXe22ueDmq/18XGtELSq85j9X8q0tcNSSKJIX8FTuJF/Pf8j5PhqG2u+osvsLxYrvvfeVJL+4tkcXcr9JV7v0ERmj/X6fM3NC4j6dS1+9Umr2oPavqiAydTZPLMNRGY23LO9zAVDly7jD+70G5TPPLdhRIl4WxcYjLnM+SNcJ26FOrkrISUtPObIz5Zb3AG612krnpy15RMW+1cQjlnWFI6538qky9axd2oJmHIHP08KyP0ubGO+TQNOYuv2uh17yCIvR8VcStw7o1g0NM60sk+8Tq7YfIBJrtp53GkvzXH7OA0p8/n/u1satf/VJhtR1l8Wa6Gmaug7haSpaCaYQax6ta0mkutlb+eAOSG1aobM81D9A4iS1RRlzBBoVX6tU1S6WE2N9ORY6DfeLRC4l9Rvr5h95XDWB2mR1d4WFudpsgVYwiTwT31ljskD8ZyDOlm5DkGh9N/UB/0AI5Xvb8ZBmai2hQ4BWMqFwYnzxwB26YHSOv9WgY3JXnvoN+2R4rqGVh/LLDMtpFP+SpMGJNWvbIl5SOodbCczW2RKleksPoUeGEzrjtKHVdtZA+kfqO+rVx/iclCqwoopepvJpSTDjT+b9GWylGRF8EDbGlw6eUzmJM95Ovoz+kwLX3c2fTjFeYEsE7vUZm3mqdGJuKh2w9/QGSaqRHs99aScGOdDqkFcACoqdbBoQqqjamhH6Q9ng39JCg3lrGJwd50Qk9ovnqBTr8MME7Ps2wiVfygUmPoUBJJfJWX5Nda0nuncbFkA=='
      )
    ),
    Hr = new Set(jr(Kr)),
    Vr = new Set(jr(Kr)),
    Jr = (function (t) {
      let e = [];
      for (;;) {
        let r = t();
        if (0 == r) break;
        e.push(Gr(r, t));
      }
      for (;;) {
        let r = t() - 1;
        if (r < 0) break;
        e.push(qr(r, t));
      }
      return (function (t) {
        const e = {};
        for (let r = 0; r < t.length; r++) {
          const n = t[r];
          e[n[0]] = n[1];
        }
        return e;
      })(Or(e));
    })(Kr),
    Wr = (function (t) {
      let e = jr(t).sort((t, e) => t - e);
      return (function r() {
        let n = [];
        for (;;) {
          let i = jr(t, e);
          if (0 == i.length) break;
          n.push({ set: new Set(i), node: r() });
        }
        n.sort((t, e) => e.set.size - t.set.size);
        let i = t(),
          o = i % 3;
        i = (i / 3) | 0;
        let s = !!(1 & i);
        return (
          (i >>= 1),
          { branches: n, valid: o, fe0f: s, save: 1 == i, check: 2 == i }
        );
      })();
    })(Kr);
  function Qr(t) {
    return dr(t);
  }
  function Yr(t) {
    return t.filter((t) => 65039 != t);
  }
  function Zr(t) {
    for (let e of t.split('.')) {
      let t = Qr(e);
      try {
        for (let e = t.lastIndexOf(95) - 1; e >= 0; e--)
          if (95 !== t[e]) throw new Error('underscore only allowed at start');
        if (
          t.length >= 4 &&
          t.every((t) => t < 128) &&
          45 === t[2] &&
          45 === t[3]
        )
          throw new Error('invalid label extension');
      } catch (t) {
        throw new Error(`Invalid label "${e}": ${t.message}`);
      }
    }
    return t;
  }
  function Xr(t) {
    return Zr(
      (function (t, e) {
        let r = Qr(t).reverse(),
          n = [];
        for (; r.length; ) {
          let t = $r(r);
          if (t) {
            n.push(...e(t));
            continue;
          }
          let i = r.pop();
          if (Hr.has(i)) {
            n.push(i);
            continue;
          }
          if (Vr.has(i)) continue;
          let o = Jr[i];
          if (!o)
            throw new Error(
              `Disallowed codepoint: 0x${i.toString(16).toUpperCase()}`
            );
          n.push(...o);
        }
        return Zr(
          (function (t) {
            return t.normalize('NFC');
          })(String.fromCodePoint(...n))
        );
      })(t, Yr)
    );
  }
  function $r(t, e) {
    var r;
    let n,
      i,
      o = Wr,
      s = [],
      a = t.length;
    for (e && (e.length = 0); a; ) {
      let u = t[--a];
      if (
        ((o =
          null === (r = o.branches.find((t) => t.set.has(u))) || void 0 === r
            ? void 0
            : r.node),
        !o)
      )
        break;
      if (o.save) i = u;
      else if (o.check && u === i) break;
      s.push(u),
        o.fe0f && (s.push(65039), a > 0 && 65039 == t[a - 1] && a--),
        o.valid &&
          ((n = s.slice()),
          2 == o.valid && n.splice(1, 1),
          e && e.push(...t.slice(a).reverse()),
          (t.length = a));
    }
    return n;
  }
  const tn = new f(Cr),
    en = new Uint8Array(32);
  function rn(t) {
    if (0 === t.length) throw new Error('invalid ENS name; empty component');
    return t;
  }
  function nn(t) {
    const e = hr(Xr(t)),
      r = [];
    if (0 === t.length) return r;
    let n = 0;
    for (let t = 0; t < e.length; t++) {
      46 === e[t] && (r.push(rn(e.slice(n, t))), (n = t + 1));
    }
    if (n >= e.length) throw new Error('invalid ENS name; empty component');
    return r.push(rn(e.slice(n))), r;
  }
  function on(t) {
    'string' != typeof t &&
      tn.throwArgumentError('invalid ENS name; not a string', 'name', t);
    let e = en;
    const r = nn(t);
    for (; r.length; ) e = ye(b([e, ye(r.pop())]));
    return k(e);
  }
  function sn(t) {
    return (
      k(
        b(
          nn(t).map((t) => {
            if (t.length > 63)
              throw new Error(
                'invalid DNS encoded entry; length exceeds 63 bytes'
              );
            const e = new Uint8Array(t.length + 1);
            return e.set(t, 1), (e[0] = e.length - 1), e;
          })
        )
      ) + '00'
    );
  }
  en.fill(0);
  function an(t) {
    return (
      'string' == typeof t && (t = hr(t)),
      ye(b([hr('Ethereum Signed Message:\n'), hr(String(t.length)), t]))
    );
  }
  var un = function (t, e, r, n) {
    return new (r || (r = Promise))(function (i, o) {
      function s(t) {
        try {
          u(n.next(t));
        } catch (t) {
          o(t);
        }
      }
      function a(t) {
        try {
          u(n.throw(t));
        } catch (t) {
          o(t);
        }
      }
      function u(t) {
        var e;
        t.done
          ? i(t.value)
          : ((e = t.value),
            e instanceof r
              ? e
              : new r(function (t) {
                  t(e);
                })).then(s, a);
      }
      u((n = n.apply(t, e || [])).next());
    });
  };
  const hn = new f(Cr),
    ln = new Uint8Array(32);
  ln.fill(0);
  const cn = j.from(-1),
    fn = j.from(0),
    dn = j.from(1),
    mn = j.from(
      '0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff'
    );
  const pn = C(dn.toHexString(), 32),
    gn = C(fn.toHexString(), 32),
    yn = {
      name: 'string',
      version: 'string',
      chainId: 'uint256',
      verifyingContract: 'address',
      salt: 'bytes32',
    },
    vn = ['name', 'version', 'chainId', 'verifyingContract', 'salt'];
  function wn(t) {
    return function (e) {
      return (
        'string' != typeof e &&
          hn.throwArgumentError(
            `invalid domain value for ${JSON.stringify(t)}`,
            `domain.${t}`,
            e
          ),
        e
      );
    };
  }
  const bn = {
    name: wn('name'),
    version: wn('version'),
    chainId: function (t) {
      try {
        return j.from(t).toString();
      } catch (t) {}
      return hn.throwArgumentError(
        'invalid domain value for "chainId"',
        'domain.chainId',
        t
      );
    },
    verifyingContract: function (t) {
      try {
        return Re(t).toLowerCase();
      } catch (t) {}
      return hn.throwArgumentError(
        'invalid domain value "verifyingContract"',
        'domain.verifyingContract',
        t
      );
    },
    salt: function (t) {
      try {
        const e = w(t);
        if (32 !== e.length) throw new Error('bad length');
        return k(e);
      } catch (t) {}
      return hn.throwArgumentError(
        'invalid domain value "salt"',
        'domain.salt',
        t
      );
    },
  };
  function An(t) {
    {
      const e = t.match(/^(u?)int(\d*)$/);
      if (e) {
        const r = '' === e[1],
          n = parseInt(e[2] || '256');
        (n % 8 != 0 || n > 256 || (e[2] && e[2] !== String(n))) &&
          hn.throwArgumentError('invalid numeric width', 'type', t);
        const i = mn.mask(r ? n - 1 : n),
          o = r ? i.add(dn).mul(cn) : fn;
        return function (e) {
          const r = j.from(e);
          return (
            (r.lt(o) || r.gt(i)) &&
              hn.throwArgumentError(`value out-of-bounds for ${t}`, 'value', e),
            C(r.toTwos(256).toHexString(), 32)
          );
        };
      }
    }
    {
      const e = t.match(/^bytes(\d+)$/);
      if (e) {
        const r = parseInt(e[1]);
        return (
          (0 === r || r > 32 || e[1] !== String(r)) &&
            hn.throwArgumentError('invalid bytes width', 'type', t),
          function (e) {
            return (
              w(e).length !== r &&
                hn.throwArgumentError(`invalid length for ${t}`, 'value', e),
              (function (t) {
                const e = w(t),
                  r = e.length % 32;
                return r ? P([e, ln.slice(r)]) : k(e);
              })(e)
            );
          }
        );
      }
    }
    switch (t) {
      case 'address':
        return function (t) {
          return C(Re(t), 32);
        };
      case 'bool':
        return function (t) {
          return t ? pn : gn;
        };
      case 'bytes':
        return function (t) {
          return ye(t);
        };
      case 'string':
        return function (t) {
          return Tr(t);
        };
    }
    return null;
  }
  function En(t, e) {
    return `${t}(${e.map(({ name: t, type: e }) => e + ' ' + t).join(',')})`;
  }
  class Mn {
    constructor(t) {
      st(this, 'types', Object.freeze(mt(t))),
        st(this, '_encoderCache', {}),
        st(this, '_types', {});
      const e = {},
        r = {},
        n = {};
      Object.keys(t).forEach((t) => {
        (e[t] = {}), (r[t] = []), (n[t] = {});
      });
      for (const n in t) {
        const i = {};
        t[n].forEach((o) => {
          i[o.name] &&
            hn.throwArgumentError(
              `duplicate variable name ${JSON.stringify(
                o.name
              )} in ${JSON.stringify(n)}`,
              'types',
              t
            ),
            (i[o.name] = !0);
          const s = o.type.match(/^([^\x5b]*)(\x5b|$)/)[1];
          s === n &&
            hn.throwArgumentError(
              `circular type reference to ${JSON.stringify(s)}`,
              'types',
              t
            );
          An(s) ||
            (r[s] ||
              hn.throwArgumentError(
                `unknown type ${JSON.stringify(s)}`,
                'types',
                t
              ),
            r[s].push(n),
            (e[n][s] = !0));
        });
      }
      const i = Object.keys(r).filter((t) => 0 === r[t].length);
      0 === i.length
        ? hn.throwArgumentError('missing primary type', 'types', t)
        : i.length > 1 &&
          hn.throwArgumentError(
            `ambiguous primary types or unused types: ${i
              .map((t) => JSON.stringify(t))
              .join(', ')}`,
            'types',
            t
          ),
        st(this, 'primaryType', i[0]),
        (function i(o, s) {
          s[o] &&
            hn.throwArgumentError(
              `circular type reference to ${JSON.stringify(o)}`,
              'types',
              t
            ),
            (s[o] = !0),
            Object.keys(e[o]).forEach((t) => {
              r[t] &&
                (i(t, s),
                Object.keys(s).forEach((e) => {
                  n[e][t] = !0;
                }));
            }),
            delete s[o];
        })(this.primaryType, {});
      for (const e in n) {
        const r = Object.keys(n[e]);
        r.sort(),
          (this._types[e] = En(e, t[e]) + r.map((e) => En(e, t[e])).join(''));
      }
    }
    getEncoder(t) {
      let e = this._encoderCache[t];
      return e || (e = this._encoderCache[t] = this._getEncoder(t)), e;
    }
    _getEncoder(t) {
      {
        const e = An(t);
        if (e) return e;
      }
      const e = t.match(/^(.*)(\x5b(\d*)\x5d)$/);
      if (e) {
        const t = e[1],
          r = this.getEncoder(t),
          n = parseInt(e[3]);
        return (e) => {
          n >= 0 &&
            e.length !== n &&
            hn.throwArgumentError(
              'array length mismatch; expected length ${ arrayLength }',
              'value',
              e
            );
          let i = e.map(r);
          return this._types[t] && (i = i.map(ye)), ye(P(i));
        };
      }
      const r = this.types[t];
      if (r) {
        const e = Tr(this._types[t]);
        return (t) => {
          const n = r.map(({ name: e, type: r }) => {
            const n = this.getEncoder(r)(t[e]);
            return this._types[r] ? ye(n) : n;
          });
          return n.unshift(e), P(n);
        };
      }
      return hn.throwArgumentError(`unknown type: ${t}`, 'type', t);
    }
    encodeType(t) {
      const e = this._types[t];
      return (
        e ||
          hn.throwArgumentError(
            `unknown type: ${JSON.stringify(t)}`,
            'name',
            t
          ),
        e
      );
    }
    encodeData(t, e) {
      return this.getEncoder(t)(e);
    }
    hashStruct(t, e) {
      return ye(this.encodeData(t, e));
    }
    encode(t) {
      return this.encodeData(this.primaryType, t);
    }
    hash(t) {
      return this.hashStruct(this.primaryType, t);
    }
    _visit(t, e, r) {
      if (An(t)) return r(t, e);
      const n = t.match(/^(.*)(\x5b(\d*)\x5d)$/);
      if (n) {
        const t = n[1],
          i = parseInt(n[3]);
        return (
          i >= 0 &&
            e.length !== i &&
            hn.throwArgumentError(
              'array length mismatch; expected length ${ arrayLength }',
              'value',
              e
            ),
          e.map((e) => this._visit(t, e, r))
        );
      }
      const i = this.types[t];
      return i
        ? i.reduce(
            (t, { name: n, type: i }) => ((t[n] = this._visit(i, e[n], r)), t),
            {}
          )
        : hn.throwArgumentError(`unknown type: ${t}`, 'type', t);
    }
    visit(t, e) {
      return this._visit(this.primaryType, t, e);
    }
    static from(t) {
      return new Mn(t);
    }
    static getPrimaryType(t) {
      return Mn.from(t).primaryType;
    }
    static hashStruct(t, e, r) {
      return Mn.from(e).hashStruct(t, r);
    }
    static hashDomain(t) {
      const e = [];
      for (const r in t) {
        const n = yn[r];
        n ||
          hn.throwArgumentError(
            `invalid typed-data domain key: ${JSON.stringify(r)}`,
            'domain',
            t
          ),
          e.push({ name: r, type: n });
      }
      return (
        e.sort((t, e) => vn.indexOf(t.name) - vn.indexOf(e.name)),
        Mn.hashStruct('EIP712Domain', { EIP712Domain: e }, t)
      );
    }
    static encode(t, e, r) {
      return P(['0x1901', Mn.hashDomain(t), Mn.from(e).hash(r)]);
    }
    static hash(t, e, r) {
      return ye(Mn.encode(t, e, r));
    }
    static resolveNames(t, e, r, n) {
      return un(this, void 0, void 0, function* () {
        t = lt(t);
        const i = {};
        t.verifyingContract &&
          !M(t.verifyingContract, 20) &&
          (i[t.verifyingContract] = '0x');
        const o = Mn.from(e);
        o.visit(r, (t, e) => ('address' !== t || M(e, 20) || (i[e] = '0x'), e));
        for (const t in i) i[t] = yield n(t);
        return (
          t.verifyingContract &&
            i[t.verifyingContract] &&
            (t.verifyingContract = i[t.verifyingContract]),
          (r = o.visit(r, (t, e) => ('address' === t && i[e] ? i[e] : e))),
          { domain: t, value: r }
        );
      });
    }
    static getPayload(t, e, r) {
      Mn.hashDomain(t);
      const n = {},
        i = [];
      vn.forEach((e) => {
        const r = t[e];
        null != r && ((n[e] = bn[e](r)), i.push({ name: e, type: yn[e] }));
      });
      const o = Mn.from(e),
        s = lt(e);
      return (
        s.EIP712Domain
          ? hn.throwArgumentError(
              'types must not contain EIP712Domain type',
              'types.EIP712Domain',
              e
            )
          : (s.EIP712Domain = i),
        o.encode(r),
        {
          types: s,
          domain: n,
          primaryType: o.primaryType,
          message: o.visit(r, (t, e) => {
            if (t.match(/^bytes(\d*)/)) return k(w(e));
            if (t.match(/^u?int/)) return j.from(e).toString();
            switch (t) {
              case 'address':
                return e.toLowerCase();
              case 'bool':
                return !!e;
              case 'string':
                return (
                  'string' != typeof e &&
                    hn.throwArgumentError('invalid string', 'value', e),
                  e
                );
            }
            return hn.throwArgumentError('unsupported type', 'type', t);
          }),
        }
      );
    }
  }
  const _n = new f(gt);
  class kn extends pt {}
  class Sn extends pt {}
  class Nn extends pt {}
  class Pn extends pt {
    static isIndexed(t) {
      return !(!t || !t._isIndexed);
    }
  }
  const xn = {
    '0x08c379a0': {
      signature: 'Error(string)',
      name: 'Error',
      inputs: ['string'],
      reason: !0,
    },
    '0x4e487b71': {
      signature: 'Panic(uint256)',
      name: 'Panic',
      inputs: ['uint256'],
    },
  };
  function Tn(t, e) {
    const r = new Error(
      `deferred error during ABI decoding triggered accessing ${t}`
    );
    return (r.error = e), r;
  }
  class Cn {
    constructor(t) {
      let e = [];
      (e = 'string' == typeof t ? JSON.parse(t) : t),
        st(
          this,
          'fragments',
          e.map((t) => Nt.from(t)).filter((t) => null != t)
        ),
        st(this, '_abiCoder', at(new.target, 'getAbiCoder')()),
        st(this, 'functions', {}),
        st(this, 'errors', {}),
        st(this, 'events', {}),
        st(this, 'structs', {}),
        this.fragments.forEach((t) => {
          let e = null;
          switch (t.type) {
            case 'constructor':
              return this.deploy
                ? void _n.warn('duplicate definition - constructor')
                : void st(this, 'deploy', t);
            case 'function':
              e = this.functions;
              break;
            case 'event':
              e = this.events;
              break;
            case 'error':
              e = this.errors;
              break;
            default:
              return;
          }
          let r = t.format();
          e[r] ? _n.warn('duplicate definition - ' + r) : (e[r] = t);
        }),
        this.deploy ||
          st(this, 'deploy', Rt.from({ payable: !1, type: 'constructor' })),
        st(this, '_isInterface', !0);
    }
    format(t) {
      t || (t = Mt.full),
        t === Mt.sighash &&
          _n.throwArgumentError(
            'interface does not support formatting sighash',
            'format',
            t
          );
      const e = this.fragments.map((e) => e.format(t));
      return t === Mt.json ? JSON.stringify(e.map((t) => JSON.parse(t))) : e;
    }
    static getAbiCoder() {
      return xr;
    }
    static getAddress(t) {
      return Re(t);
    }
    static getSighash(t) {
      return N(Tr(t.format()), 0, 4);
    }
    static getEventTopic(t) {
      return Tr(t.format());
    }
    getFunction(t) {
      if (M(t)) {
        for (const e in this.functions)
          if (t === this.getSighash(e)) return this.functions[e];
        _n.throwArgumentError('no matching function', 'sighash', t);
      }
      if (-1 === t.indexOf('(')) {
        const e = t.trim(),
          r = Object.keys(this.functions).filter((t) => t.split('(')[0] === e);
        return (
          0 === r.length
            ? _n.throwArgumentError('no matching function', 'name', e)
            : r.length > 1 &&
              _n.throwArgumentError('multiple matching functions', 'name', e),
          this.functions[r[0]]
        );
      }
      const e = this.functions[It.fromString(t).format()];
      return (
        e || _n.throwArgumentError('no matching function', 'signature', t), e
      );
    }
    getEvent(t) {
      if (M(t)) {
        const e = t.toLowerCase();
        for (const t in this.events)
          if (e === this.getEventTopic(t)) return this.events[t];
        _n.throwArgumentError('no matching event', 'topichash', e);
      }
      if (-1 === t.indexOf('(')) {
        const e = t.trim(),
          r = Object.keys(this.events).filter((t) => t.split('(')[0] === e);
        return (
          0 === r.length
            ? _n.throwArgumentError('no matching event', 'name', e)
            : r.length > 1 &&
              _n.throwArgumentError('multiple matching events', 'name', e),
          this.events[r[0]]
        );
      }
      const e = this.events[Pt.fromString(t).format()];
      return e || _n.throwArgumentError('no matching event', 'signature', t), e;
    }
    getError(t) {
      if (M(t)) {
        const e = at(this.constructor, 'getSighash');
        for (const r in this.errors) {
          if (t === e(this.errors[r])) return this.errors[r];
        }
        _n.throwArgumentError('no matching error', 'sighash', t);
      }
      if (-1 === t.indexOf('(')) {
        const e = t.trim(),
          r = Object.keys(this.errors).filter((t) => t.split('(')[0] === e);
        return (
          0 === r.length
            ? _n.throwArgumentError('no matching error', 'name', e)
            : r.length > 1 &&
              _n.throwArgumentError('multiple matching errors', 'name', e),
          this.errors[r[0]]
        );
      }
      const e = this.errors[It.fromString(t).format()];
      return e || _n.throwArgumentError('no matching error', 'signature', t), e;
    }
    getSighash(t) {
      if ('string' == typeof t)
        try {
          t = this.getFunction(t);
        } catch (e) {
          try {
            t = this.getError(t);
          } catch (t) {
            throw e;
          }
        }
      return at(this.constructor, 'getSighash')(t);
    }
    getEventTopic(t) {
      return (
        'string' == typeof t && (t = this.getEvent(t)),
        at(this.constructor, 'getEventTopic')(t)
      );
    }
    _decodeParams(t, e) {
      return this._abiCoder.decode(t, e);
    }
    _encodeParams(t, e) {
      return this._abiCoder.encode(t, e);
    }
    encodeDeploy(t) {
      return this._encodeParams(this.deploy.inputs, t || []);
    }
    decodeErrorResult(t, e) {
      'string' == typeof t && (t = this.getError(t));
      const r = w(e);
      return (
        k(r.slice(0, 4)) !== this.getSighash(t) &&
          _n.throwArgumentError(
            `data signature does not match error ${t.name}.`,
            'data',
            k(r)
          ),
        this._decodeParams(t.inputs, r.slice(4))
      );
    }
    encodeErrorResult(t, e) {
      return (
        'string' == typeof t && (t = this.getError(t)),
        k(b([this.getSighash(t), this._encodeParams(t.inputs, e || [])]))
      );
    }
    decodeFunctionData(t, e) {
      'string' == typeof t && (t = this.getFunction(t));
      const r = w(e);
      return (
        k(r.slice(0, 4)) !== this.getSighash(t) &&
          _n.throwArgumentError(
            `data signature does not match function ${t.name}.`,
            'data',
            k(r)
          ),
        this._decodeParams(t.inputs, r.slice(4))
      );
    }
    encodeFunctionData(t, e) {
      return (
        'string' == typeof t && (t = this.getFunction(t)),
        k(b([this.getSighash(t), this._encodeParams(t.inputs, e || [])]))
      );
    }
    decodeFunctionResult(t, e) {
      'string' == typeof t && (t = this.getFunction(t));
      let r = w(e),
        n = null,
        i = '',
        o = null,
        s = null,
        a = null;
      switch (r.length % this._abiCoder._getWordSize()) {
        case 0:
          try {
            return this._abiCoder.decode(t.outputs, r);
          } catch (t) {}
          break;
        case 4: {
          const t = k(r.slice(0, 4)),
            e = xn[t];
          if (e)
            (o = this._abiCoder.decode(e.inputs, r.slice(4))),
              (s = e.name),
              (a = e.signature),
              e.reason && (n = o[0]),
              'Error' === s
                ? (i = `; VM Exception while processing transaction: reverted with reason string ${JSON.stringify(
                    o[0]
                  )}`)
                : 'Panic' === s &&
                  (i = `; VM Exception while processing transaction: reverted with panic code ${o[0]}`);
          else
            try {
              const e = this.getError(t);
              (o = this._abiCoder.decode(e.inputs, r.slice(4))),
                (s = e.name),
                (a = e.format());
            } catch (t) {}
          break;
        }
      }
      return _n.throwError(
        'call revert exception' + i,
        f.errors.CALL_EXCEPTION,
        {
          method: t.format(),
          data: k(e),
          errorArgs: o,
          errorName: s,
          errorSignature: a,
          reason: n,
        }
      );
    }
    encodeFunctionResult(t, e) {
      return (
        'string' == typeof t && (t = this.getFunction(t)),
        k(this._abiCoder.encode(t.outputs, e || []))
      );
    }
    encodeFilterTopics(t, e) {
      'string' == typeof t && (t = this.getEvent(t)),
        e.length > t.inputs.length &&
          _n.throwError(
            'too many arguments for ' + t.format(),
            f.errors.UNEXPECTED_ARGUMENT,
            { argument: 'values', value: e }
          );
      let r = [];
      t.anonymous || r.push(this.getEventTopic(t));
      const n = (t, e) =>
        'string' === t.type
          ? Tr(e)
          : 'bytes' === t.type
          ? ye(k(e))
          : ('bool' === t.type &&
              'boolean' == typeof e &&
              (e = e ? '0x01' : '0x00'),
            t.type.match(/^u?int/) && (e = j.from(e).toHexString()),
            'address' === t.type && this._abiCoder.encode(['address'], [e]),
            C(k(e), 32));
      for (
        e.forEach((e, i) => {
          let o = t.inputs[i];
          o.indexed
            ? null == e
              ? r.push(null)
              : 'array' === o.baseType || 'tuple' === o.baseType
              ? _n.throwArgumentError(
                  'filtering with tuples or arrays not supported',
                  'contract.' + o.name,
                  e
                )
              : Array.isArray(e)
              ? r.push(e.map((t) => n(o, t)))
              : r.push(n(o, e))
            : null != e &&
              _n.throwArgumentError(
                'cannot filter non-indexed parameters; must be null',
                'contract.' + o.name,
                e
              );
        });
        r.length && null === r[r.length - 1];

      )
        r.pop();
      return r;
    }
    encodeEventLog(t, e) {
      'string' == typeof t && (t = this.getEvent(t));
      const r = [],
        n = [],
        i = [];
      return (
        t.anonymous || r.push(this.getEventTopic(t)),
        e.length !== t.inputs.length &&
          _n.throwArgumentError('event arguments/values mismatch', 'values', e),
        t.inputs.forEach((t, o) => {
          const s = e[o];
          if (t.indexed)
            if ('string' === t.type) r.push(Tr(s));
            else if ('bytes' === t.type) r.push(ye(s));
            else {
              if ('tuple' === t.baseType || 'array' === t.baseType)
                throw new Error('not implemented');
              r.push(this._abiCoder.encode([t.type], [s]));
            }
          else n.push(t), i.push(s);
        }),
        { data: this._abiCoder.encode(n, i), topics: r }
      );
    }
    decodeEventLog(t, e, r) {
      if (
        ('string' == typeof t && (t = this.getEvent(t)),
        null != r && !t.anonymous)
      ) {
        let e = this.getEventTopic(t);
        (M(r[0], 32) && r[0].toLowerCase() === e) ||
          _n.throwError('fragment/topic mismatch', f.errors.INVALID_ARGUMENT, {
            argument: 'topics[0]',
            expected: e,
            value: r[0],
          }),
          (r = r.slice(1));
      }
      let n = [],
        i = [],
        o = [];
      t.inputs.forEach((t, e) => {
        t.indexed
          ? 'string' === t.type ||
            'bytes' === t.type ||
            'tuple' === t.baseType ||
            'array' === t.baseType
            ? (n.push(kt.fromObject({ type: 'bytes32', name: t.name })),
              o.push(!0))
            : (n.push(t), o.push(!1))
          : (i.push(t), o.push(!1));
      });
      let s = null != r ? this._abiCoder.decode(n, b(r)) : null,
        a = this._abiCoder.decode(i, e, !0),
        u = [],
        h = 0,
        l = 0;
      t.inputs.forEach((t, e) => {
        if (t.indexed)
          if (null == s) u[e] = new Pn({ _isIndexed: !0, hash: null });
          else if (o[e]) u[e] = new Pn({ _isIndexed: !0, hash: s[l++] });
          else
            try {
              u[e] = s[l++];
            } catch (t) {
              u[e] = t;
            }
        else
          try {
            u[e] = a[h++];
          } catch (t) {
            u[e] = t;
          }
        if (t.name && null == u[t.name]) {
          const r = u[e];
          r instanceof Error
            ? Object.defineProperty(u, t.name, {
                enumerable: !0,
                get: () => {
                  throw Tn(`property ${JSON.stringify(t.name)}`, r);
                },
              })
            : (u[t.name] = r);
        }
      });
      for (let t = 0; t < u.length; t++) {
        const e = u[t];
        e instanceof Error &&
          Object.defineProperty(u, t, {
            enumerable: !0,
            get: () => {
              throw Tn(`index ${t}`, e);
            },
          });
      }
      return Object.freeze(u);
    }
    parseTransaction(t) {
      let e = this.getFunction(t.data.substring(0, 10).toLowerCase());
      return e
        ? new Sn({
            args: this._abiCoder.decode(e.inputs, '0x' + t.data.substring(10)),
            functionFragment: e,
            name: e.name,
            signature: e.format(),
            sighash: this.getSighash(e),
            value: j.from(t.value || '0'),
          })
        : null;
    }
    parseLog(t) {
      let e = this.getEvent(t.topics[0]);
      return !e || e.anonymous
        ? null
        : new kn({
            eventFragment: e,
            name: e.name,
            signature: e.format(),
            topic: this.getEventTopic(e),
            args: this.decodeEventLog(e, t.data, t.topics),
          });
    }
    parseError(t) {
      const e = k(t);
      let r = this.getError(e.substring(0, 10).toLowerCase());
      return r
        ? new Nn({
            args: this._abiCoder.decode(r.inputs, '0x' + e.substring(10)),
            errorFragment: r,
            name: r.name,
            signature: r.format(),
            sighash: this.getSighash(r),
          })
        : null;
    }
    static isInterface(t) {
      return !(!t || !t._isInterface);
    }
  }
  var Rn = function (t, e, r, n) {
    return new (r || (r = Promise))(function (i, o) {
      function s(t) {
        try {
          u(n.next(t));
        } catch (t) {
          o(t);
        }
      }
      function a(t) {
        try {
          u(n.throw(t));
        } catch (t) {
          o(t);
        }
      }
      function u(t) {
        var e;
        t.done
          ? i(t.value)
          : ((e = t.value),
            e instanceof r
              ? e
              : new r(function (t) {
                  t(e);
                })).then(s, a);
      }
      u((n = n.apply(t, e || [])).next());
    });
  };
  const In = new f('abstract-provider/5.7.0');
  class Bn extends pt {
    static isForkEvent(t) {
      return !(!t || !t._isForkEvent);
    }
  }
  class On {
    constructor() {
      In.checkAbstract(new.target, On), st(this, '_isProvider', !0);
    }
    getFeeData() {
      return Rn(this, void 0, void 0, function* () {
        const { block: t, gasPrice: e } = yield ut({
          block: this.getBlock('latest'),
          gasPrice: this.getGasPrice().catch((t) => null),
        });
        let r = null,
          n = null,
          i = null;
        return (
          t &&
            t.baseFeePerGas &&
            ((r = t.baseFeePerGas),
            (i = j.from('1500000000')),
            (n = t.baseFeePerGas.mul(2).add(i))),
          {
            lastBaseFeePerGas: r,
            maxFeePerGas: n,
            maxPriorityFeePerGas: i,
            gasPrice: e,
          }
        );
      });
    }
    addListener(t, e) {
      return this.on(t, e);
    }
    removeListener(t, e) {
      return this.off(t, e);
    }
    static isProvider(t) {
      return !(!t || !t._isProvider);
    }
  }
  var Fn = function (t, e, r, n) {
    return new (r || (r = Promise))(function (i, o) {
      function s(t) {
        try {
          u(n.next(t));
        } catch (t) {
          o(t);
        }
      }
      function a(t) {
        try {
          u(n.throw(t));
        } catch (t) {
          o(t);
        }
      }
      function u(t) {
        var e;
        t.done
          ? i(t.value)
          : ((e = t.value),
            e instanceof r
              ? e
              : new r(function (t) {
                  t(e);
                })).then(s, a);
      }
      u((n = n.apply(t, e || [])).next());
    });
  };
  const Dn = new f('abstract-signer/5.7.0'),
    Ln = [
      'accessList',
      'ccipReadEnabled',
      'chainId',
      'customData',
      'data',
      'from',
      'gasLimit',
      'gasPrice',
      'maxFeePerGas',
      'maxPriorityFeePerGas',
      'nonce',
      'to',
      'type',
      'value',
    ],
    Un = [
      f.errors.INSUFFICIENT_FUNDS,
      f.errors.NONCE_EXPIRED,
      f.errors.REPLACEMENT_UNDERPRICED,
    ];
  class jn {
    constructor() {
      Dn.checkAbstract(new.target, jn), st(this, '_isSigner', !0);
    }
    getBalance(t) {
      return Fn(this, void 0, void 0, function* () {
        return (
          this._checkProvider('getBalance'),
          yield this.provider.getBalance(this.getAddress(), t)
        );
      });
    }
    getTransactionCount(t) {
      return Fn(this, void 0, void 0, function* () {
        return (
          this._checkProvider('getTransactionCount'),
          yield this.provider.getTransactionCount(this.getAddress(), t)
        );
      });
    }
    estimateGas(t) {
      return Fn(this, void 0, void 0, function* () {
        this._checkProvider('estimateGas');
        const e = yield ut(this.checkTransaction(t));
        return yield this.provider.estimateGas(e);
      });
    }
    call(t, e) {
      return Fn(this, void 0, void 0, function* () {
        this._checkProvider('call');
        const r = yield ut(this.checkTransaction(t));
        return yield this.provider.call(r, e);
      });
    }
    sendTransaction(t) {
      return Fn(this, void 0, void 0, function* () {
        this._checkProvider('sendTransaction');
        const e = yield this.populateTransaction(t),
          r = yield this.signTransaction(e);
        return yield this.provider.sendTransaction(r);
      });
    }
    getChainId() {
      return Fn(this, void 0, void 0, function* () {
        this._checkProvider('getChainId');
        return (yield this.provider.getNetwork()).chainId;
      });
    }
    getGasPrice() {
      return Fn(this, void 0, void 0, function* () {
        return (
          this._checkProvider('getGasPrice'), yield this.provider.getGasPrice()
        );
      });
    }
    getFeeData() {
      return Fn(this, void 0, void 0, function* () {
        return (
          this._checkProvider('getFeeData'), yield this.provider.getFeeData()
        );
      });
    }
    resolveName(t) {
      return Fn(this, void 0, void 0, function* () {
        return (
          this._checkProvider('resolveName'), yield this.provider.resolveName(t)
        );
      });
    }
    checkTransaction(t) {
      for (const e in t)
        -1 === Ln.indexOf(e) &&
          Dn.throwArgumentError(
            'invalid transaction key: ' + e,
            'transaction',
            t
          );
      const e = lt(t);
      return (
        null == e.from
          ? (e.from = this.getAddress())
          : (e.from = Promise.all([
              Promise.resolve(e.from),
              this.getAddress(),
            ]).then(
              (e) => (
                e[0].toLowerCase() !== e[1].toLowerCase() &&
                  Dn.throwArgumentError(
                    'from address mismatch',
                    'transaction',
                    t
                  ),
                e[0]
              )
            )),
        e
      );
    }
    populateTransaction(t) {
      return Fn(this, void 0, void 0, function* () {
        const e = yield ut(this.checkTransaction(t));
        null != e.to &&
          ((e.to = Promise.resolve(e.to).then((t) =>
            Fn(this, void 0, void 0, function* () {
              if (null == t) return null;
              const e = yield this.resolveName(t);
              return (
                null == e &&
                  Dn.throwArgumentError(
                    'provided ENS name resolves to null',
                    'tx.to',
                    t
                  ),
                e
              );
            })
          )),
          e.to.catch((t) => {}));
        const r = null != e.maxFeePerGas || null != e.maxPriorityFeePerGas;
        if (
          (null == e.gasPrice || (2 !== e.type && !r)
            ? (0 !== e.type && 1 !== e.type) ||
              !r ||
              Dn.throwArgumentError(
                'pre-eip-1559 transaction do not support maxFeePerGas/maxPriorityFeePerGas',
                'transaction',
                t
              )
            : Dn.throwArgumentError(
                'eip-1559 transaction do not support gasPrice',
                'transaction',
                t
              ),
          (2 !== e.type && null != e.type) ||
            null == e.maxFeePerGas ||
            null == e.maxPriorityFeePerGas)
        )
          if (0 === e.type || 1 === e.type)
            null == e.gasPrice && (e.gasPrice = this.getGasPrice());
          else {
            const t = yield this.getFeeData();
            if (null == e.type)
              if (null != t.maxFeePerGas && null != t.maxPriorityFeePerGas)
                if (((e.type = 2), null != e.gasPrice)) {
                  const t = e.gasPrice;
                  delete e.gasPrice,
                    (e.maxFeePerGas = t),
                    (e.maxPriorityFeePerGas = t);
                } else
                  null == e.maxFeePerGas && (e.maxFeePerGas = t.maxFeePerGas),
                    null == e.maxPriorityFeePerGas &&
                      (e.maxPriorityFeePerGas = t.maxPriorityFeePerGas);
              else
                null != t.gasPrice
                  ? (r &&
                      Dn.throwError(
                        'network does not support EIP-1559',
                        f.errors.UNSUPPORTED_OPERATION,
                        { operation: 'populateTransaction' }
                      ),
                    null == e.gasPrice && (e.gasPrice = t.gasPrice),
                    (e.type = 0))
                  : Dn.throwError(
                      'failed to get consistent fee data',
                      f.errors.UNSUPPORTED_OPERATION,
                      { operation: 'signer.getFeeData' }
                    );
            else
              2 === e.type &&
                (null == e.maxFeePerGas && (e.maxFeePerGas = t.maxFeePerGas),
                null == e.maxPriorityFeePerGas &&
                  (e.maxPriorityFeePerGas = t.maxPriorityFeePerGas));
          }
        else e.type = 2;
        return (
          null == e.nonce && (e.nonce = this.getTransactionCount('pending')),
          null == e.gasLimit &&
            (e.gasLimit = this.estimateGas(e).catch((t) => {
              if (Un.indexOf(t.code) >= 0) throw t;
              return Dn.throwError(
                'cannot estimate gas; transaction may fail or may require manual gas limit',
                f.errors.UNPREDICTABLE_GAS_LIMIT,
                { error: t, tx: e }
              );
            })),
          null == e.chainId
            ? (e.chainId = this.getChainId())
            : (e.chainId = Promise.all([
                Promise.resolve(e.chainId),
                this.getChainId(),
              ]).then(
                (e) => (
                  0 !== e[1] &&
                    e[0] !== e[1] &&
                    Dn.throwArgumentError(
                      'chainId address mismatch',
                      'transaction',
                      t
                    ),
                  e[0]
                )
              )),
          yield ut(e)
        );
      });
    }
    _checkProvider(t) {
      this.provider ||
        Dn.throwError('missing provider', f.errors.UNSUPPORTED_OPERATION, {
          operation: t || '_checkProvider',
        });
    }
    static isSigner(t) {
      return !(!t || !t._isSigner);
    }
  }
  class zn extends jn {
    constructor(t, e) {
      super(), st(this, 'address', t), st(this, 'provider', e || null);
    }
    getAddress() {
      return Promise.resolve(this.address);
    }
    _fail(t, e) {
      return Promise.resolve().then(() => {
        Dn.throwError(t, f.errors.UNSUPPORTED_OPERATION, { operation: e });
      });
    }
    signMessage(t) {
      return this._fail('VoidSigner cannot sign messages', 'signMessage');
    }
    signTransaction(t) {
      return this._fail(
        'VoidSigner cannot sign transactions',
        'signTransaction'
      );
    }
    _signTypedData(t, e, r) {
      return this._fail('VoidSigner cannot sign typed data', 'signTypedData');
    }
    connect(t) {
      return new zn(this.address, t);
    }
  }
  var Gn = { exports: {} };
  !(function (e) {
    !(function (t, e) {
      function r(t, e) {
        if (!t) throw new Error(e || 'Assertion failed');
      }
      function n(t, e) {
        t.super_ = e;
        var r = function () {};
        (r.prototype = e.prototype),
          (t.prototype = new r()),
          (t.prototype.constructor = t);
      }
      function i(t, e, r) {
        if (i.isBN(t)) return t;
        (this.negative = 0),
          (this.words = null),
          (this.length = 0),
          (this.red = null),
          null !== t &&
            (('le' !== e && 'be' !== e) || ((r = e), (e = 10)),
            this._init(t || 0, e || 10, r || 'be'));
      }
      var o;
      'object' == typeof t ? (t.exports = i) : (e.BN = i),
        (i.BN = i),
        (i.wordSize = 26);
      try {
        o =
          'undefined' != typeof window && void 0 !== window.Buffer
            ? window.Buffer
            : require('buffer').Buffer;
      } catch (t) {}
      function s(t, e) {
        var n = t.charCodeAt(e);
        return n >= 48 && n <= 57
          ? n - 48
          : n >= 65 && n <= 70
          ? n - 55
          : n >= 97 && n <= 102
          ? n - 87
          : void r(!1, 'Invalid character in ' + t);
      }
      function a(t, e, r) {
        var n = s(t, r);
        return r - 1 >= e && (n |= s(t, r - 1) << 4), n;
      }
      function u(t, e, n, i) {
        for (var o = 0, s = 0, a = Math.min(t.length, n), u = e; u < a; u++) {
          var h = t.charCodeAt(u) - 48;
          (o *= i),
            (s = h >= 49 ? h - 49 + 10 : h >= 17 ? h - 17 + 10 : h),
            r(h >= 0 && s < i, 'Invalid character'),
            (o += s);
        }
        return o;
      }
      function h(t, e) {
        (t.words = e.words),
          (t.length = e.length),
          (t.negative = e.negative),
          (t.red = e.red);
      }
      if (
        ((i.isBN = function (t) {
          return (
            t instanceof i ||
            (null !== t &&
              'object' == typeof t &&
              t.constructor.wordSize === i.wordSize &&
              Array.isArray(t.words))
          );
        }),
        (i.max = function (t, e) {
          return t.cmp(e) > 0 ? t : e;
        }),
        (i.min = function (t, e) {
          return t.cmp(e) < 0 ? t : e;
        }),
        (i.prototype._init = function (t, e, n) {
          if ('number' == typeof t) return this._initNumber(t, e, n);
          if ('object' == typeof t) return this._initArray(t, e, n);
          'hex' === e && (e = 16), r(e === (0 | e) && e >= 2 && e <= 36);
          var i = 0;
          '-' === (t = t.toString().replace(/\s+/g, ''))[0] &&
            (i++, (this.negative = 1)),
            i < t.length &&
              (16 === e
                ? this._parseHex(t, i, n)
                : (this._parseBase(t, e, i),
                  'le' === n && this._initArray(this.toArray(), e, n)));
        }),
        (i.prototype._initNumber = function (t, e, n) {
          t < 0 && ((this.negative = 1), (t = -t)),
            t < 67108864
              ? ((this.words = [67108863 & t]), (this.length = 1))
              : t < 4503599627370496
              ? ((this.words = [67108863 & t, (t / 67108864) & 67108863]),
                (this.length = 2))
              : (r(t < 9007199254740992),
                (this.words = [67108863 & t, (t / 67108864) & 67108863, 1]),
                (this.length = 3)),
            'le' === n && this._initArray(this.toArray(), e, n);
        }),
        (i.prototype._initArray = function (t, e, n) {
          if ((r('number' == typeof t.length), t.length <= 0))
            return (this.words = [0]), (this.length = 1), this;
          (this.length = Math.ceil(t.length / 3)),
            (this.words = new Array(this.length));
          for (var i = 0; i < this.length; i++) this.words[i] = 0;
          var o,
            s,
            a = 0;
          if ('be' === n)
            for (i = t.length - 1, o = 0; i >= 0; i -= 3)
              (s = t[i] | (t[i - 1] << 8) | (t[i - 2] << 16)),
                (this.words[o] |= (s << a) & 67108863),
                (this.words[o + 1] = (s >>> (26 - a)) & 67108863),
                (a += 24) >= 26 && ((a -= 26), o++);
          else if ('le' === n)
            for (i = 0, o = 0; i < t.length; i += 3)
              (s = t[i] | (t[i + 1] << 8) | (t[i + 2] << 16)),
                (this.words[o] |= (s << a) & 67108863),
                (this.words[o + 1] = (s >>> (26 - a)) & 67108863),
                (a += 24) >= 26 && ((a -= 26), o++);
          return this._strip();
        }),
        (i.prototype._parseHex = function (t, e, r) {
          (this.length = Math.ceil((t.length - e) / 6)),
            (this.words = new Array(this.length));
          for (var n = 0; n < this.length; n++) this.words[n] = 0;
          var i,
            o = 0,
            s = 0;
          if ('be' === r)
            for (n = t.length - 1; n >= e; n -= 2)
              (i = a(t, e, n) << o),
                (this.words[s] |= 67108863 & i),
                o >= 18
                  ? ((o -= 18), (s += 1), (this.words[s] |= i >>> 26))
                  : (o += 8);
          else
            for (n = (t.length - e) % 2 == 0 ? e + 1 : e; n < t.length; n += 2)
              (i = a(t, e, n) << o),
                (this.words[s] |= 67108863 & i),
                o >= 18
                  ? ((o -= 18), (s += 1), (this.words[s] |= i >>> 26))
                  : (o += 8);
          this._strip();
        }),
        (i.prototype._parseBase = function (t, e, r) {
          (this.words = [0]), (this.length = 1);
          for (var n = 0, i = 1; i <= 67108863; i *= e) n++;
          n--, (i = (i / e) | 0);
          for (
            var o = t.length - r,
              s = o % n,
              a = Math.min(o, o - s) + r,
              h = 0,
              l = r;
            l < a;
            l += n
          )
            (h = u(t, l, l + n, e)),
              this.imuln(i),
              this.words[0] + h < 67108864
                ? (this.words[0] += h)
                : this._iaddn(h);
          if (0 !== s) {
            var c = 1;
            for (h = u(t, l, t.length, e), l = 0; l < s; l++) c *= e;
            this.imuln(c),
              this.words[0] + h < 67108864
                ? (this.words[0] += h)
                : this._iaddn(h);
          }
          this._strip();
        }),
        (i.prototype.copy = function (t) {
          t.words = new Array(this.length);
          for (var e = 0; e < this.length; e++) t.words[e] = this.words[e];
          (t.length = this.length),
            (t.negative = this.negative),
            (t.red = this.red);
        }),
        (i.prototype._move = function (t) {
          h(t, this);
        }),
        (i.prototype.clone = function () {
          var t = new i(null);
          return this.copy(t), t;
        }),
        (i.prototype._expand = function (t) {
          for (; this.length < t; ) this.words[this.length++] = 0;
          return this;
        }),
        (i.prototype._strip = function () {
          for (; this.length > 1 && 0 === this.words[this.length - 1]; )
            this.length--;
          return this._normSign();
        }),
        (i.prototype._normSign = function () {
          return (
            1 === this.length && 0 === this.words[0] && (this.negative = 0),
            this
          );
        }),
        'undefined' != typeof Symbol && 'function' == typeof Symbol.for)
      )
        try {
          i.prototype[Symbol.for('nodejs.util.inspect.custom')] = l;
        } catch (t) {
          i.prototype.inspect = l;
        }
      else i.prototype.inspect = l;
      function l() {
        return (this.red ? '<BN-R: ' : '<BN: ') + this.toString(16) + '>';
      }
      var c = [
          '',
          '0',
          '00',
          '000',
          '0000',
          '00000',
          '000000',
          '0000000',
          '00000000',
          '000000000',
          '0000000000',
          '00000000000',
          '000000000000',
          '0000000000000',
          '00000000000000',
          '000000000000000',
          '0000000000000000',
          '00000000000000000',
          '000000000000000000',
          '0000000000000000000',
          '00000000000000000000',
          '000000000000000000000',
          '0000000000000000000000',
          '00000000000000000000000',
          '000000000000000000000000',
          '0000000000000000000000000',
        ],
        f = [
          0, 0, 25, 16, 12, 11, 10, 9, 8, 8, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5,
          5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
        ],
        d = [
          0, 0, 33554432, 43046721, 16777216, 48828125, 60466176, 40353607,
          16777216, 43046721, 1e7, 19487171, 35831808, 62748517, 7529536,
          11390625, 16777216, 24137569, 34012224, 47045881, 64e6, 4084101,
          5153632, 6436343, 7962624, 9765625, 11881376, 14348907, 17210368,
          20511149, 243e5, 28629151, 33554432, 39135393, 45435424, 52521875,
          60466176,
        ];
      (i.prototype.toString = function (t, e) {
        var n;
        if (((e = 0 | e || 1), 16 === (t = t || 10) || 'hex' === t)) {
          n = '';
          for (var i = 0, o = 0, s = 0; s < this.length; s++) {
            var a = this.words[s],
              u = (16777215 & ((a << i) | o)).toString(16);
            (o = (a >>> (24 - i)) & 16777215),
              (i += 2) >= 26 && ((i -= 26), s--),
              (n =
                0 !== o || s !== this.length - 1
                  ? c[6 - u.length] + u + n
                  : u + n);
          }
          for (0 !== o && (n = o.toString(16) + n); n.length % e != 0; )
            n = '0' + n;
          return 0 !== this.negative && (n = '-' + n), n;
        }
        if (t === (0 | t) && t >= 2 && t <= 36) {
          var h = f[t],
            l = d[t];
          n = '';
          var m = this.clone();
          for (m.negative = 0; !m.isZero(); ) {
            var p = m.modrn(l).toString(t);
            n = (m = m.idivn(l)).isZero() ? p + n : c[h - p.length] + p + n;
          }
          for (this.isZero() && (n = '0' + n); n.length % e != 0; ) n = '0' + n;
          return 0 !== this.negative && (n = '-' + n), n;
        }
        r(!1, 'Base should be between 2 and 36');
      }),
        (i.prototype.toNumber = function () {
          var t = this.words[0];
          return (
            2 === this.length
              ? (t += 67108864 * this.words[1])
              : 3 === this.length && 1 === this.words[2]
              ? (t += 4503599627370496 + 67108864 * this.words[1])
              : this.length > 2 &&
                r(!1, 'Number can only safely store up to 53 bits'),
            0 !== this.negative ? -t : t
          );
        }),
        (i.prototype.toJSON = function () {
          return this.toString(16, 2);
        }),
        o &&
          (i.prototype.toBuffer = function (t, e) {
            return this.toArrayLike(o, t, e);
          }),
        (i.prototype.toArray = function (t, e) {
          return this.toArrayLike(Array, t, e);
        });
      function m(t, e, r) {
        r.negative = e.negative ^ t.negative;
        var n = (t.length + e.length) | 0;
        (r.length = n), (n = (n - 1) | 0);
        var i = 0 | t.words[0],
          o = 0 | e.words[0],
          s = i * o,
          a = 67108863 & s,
          u = (s / 67108864) | 0;
        r.words[0] = a;
        for (var h = 1; h < n; h++) {
          for (
            var l = u >>> 26,
              c = 67108863 & u,
              f = Math.min(h, e.length - 1),
              d = Math.max(0, h - t.length + 1);
            d <= f;
            d++
          ) {
            var m = (h - d) | 0;
            (l +=
              ((s = (i = 0 | t.words[m]) * (o = 0 | e.words[d]) + c) /
                67108864) |
              0),
              (c = 67108863 & s);
          }
          (r.words[h] = 0 | c), (u = 0 | l);
        }
        return 0 !== u ? (r.words[h] = 0 | u) : r.length--, r._strip();
      }
      (i.prototype.toArrayLike = function (t, e, n) {
        this._strip();
        var i = this.byteLength(),
          o = n || Math.max(1, i);
        r(i <= o, 'byte array longer than desired length'),
          r(o > 0, 'Requested array length <= 0');
        var s = (function (t, e) {
          return t.allocUnsafe ? t.allocUnsafe(e) : new t(e);
        })(t, o);
        return this['_toArrayLike' + ('le' === e ? 'LE' : 'BE')](s, i), s;
      }),
        (i.prototype._toArrayLikeLE = function (t, e) {
          for (var r = 0, n = 0, i = 0, o = 0; i < this.length; i++) {
            var s = (this.words[i] << o) | n;
            (t[r++] = 255 & s),
              r < t.length && (t[r++] = (s >> 8) & 255),
              r < t.length && (t[r++] = (s >> 16) & 255),
              6 === o
                ? (r < t.length && (t[r++] = (s >> 24) & 255), (n = 0), (o = 0))
                : ((n = s >>> 24), (o += 2));
          }
          if (r < t.length) for (t[r++] = n; r < t.length; ) t[r++] = 0;
        }),
        (i.prototype._toArrayLikeBE = function (t, e) {
          for (
            var r = t.length - 1, n = 0, i = 0, o = 0;
            i < this.length;
            i++
          ) {
            var s = (this.words[i] << o) | n;
            (t[r--] = 255 & s),
              r >= 0 && (t[r--] = (s >> 8) & 255),
              r >= 0 && (t[r--] = (s >> 16) & 255),
              6 === o
                ? (r >= 0 && (t[r--] = (s >> 24) & 255), (n = 0), (o = 0))
                : ((n = s >>> 24), (o += 2));
          }
          if (r >= 0) for (t[r--] = n; r >= 0; ) t[r--] = 0;
        }),
        Math.clz32
          ? (i.prototype._countBits = function (t) {
              return 32 - Math.clz32(t);
            })
          : (i.prototype._countBits = function (t) {
              var e = t,
                r = 0;
              return (
                e >= 4096 && ((r += 13), (e >>>= 13)),
                e >= 64 && ((r += 7), (e >>>= 7)),
                e >= 8 && ((r += 4), (e >>>= 4)),
                e >= 2 && ((r += 2), (e >>>= 2)),
                r + e
              );
            }),
        (i.prototype._zeroBits = function (t) {
          if (0 === t) return 26;
          var e = t,
            r = 0;
          return (
            8191 & e || ((r += 13), (e >>>= 13)),
            127 & e || ((r += 7), (e >>>= 7)),
            15 & e || ((r += 4), (e >>>= 4)),
            3 & e || ((r += 2), (e >>>= 2)),
            1 & e || r++,
            r
          );
        }),
        (i.prototype.bitLength = function () {
          var t = this.words[this.length - 1],
            e = this._countBits(t);
          return 26 * (this.length - 1) + e;
        }),
        (i.prototype.zeroBits = function () {
          if (this.isZero()) return 0;
          for (var t = 0, e = 0; e < this.length; e++) {
            var r = this._zeroBits(this.words[e]);
            if (((t += r), 26 !== r)) break;
          }
          return t;
        }),
        (i.prototype.byteLength = function () {
          return Math.ceil(this.bitLength() / 8);
        }),
        (i.prototype.toTwos = function (t) {
          return 0 !== this.negative
            ? this.abs().inotn(t).iaddn(1)
            : this.clone();
        }),
        (i.prototype.fromTwos = function (t) {
          return this.testn(t - 1)
            ? this.notn(t).iaddn(1).ineg()
            : this.clone();
        }),
        (i.prototype.isNeg = function () {
          return 0 !== this.negative;
        }),
        (i.prototype.neg = function () {
          return this.clone().ineg();
        }),
        (i.prototype.ineg = function () {
          return this.isZero() || (this.negative ^= 1), this;
        }),
        (i.prototype.iuor = function (t) {
          for (; this.length < t.length; ) this.words[this.length++] = 0;
          for (var e = 0; e < t.length; e++)
            this.words[e] = this.words[e] | t.words[e];
          return this._strip();
        }),
        (i.prototype.ior = function (t) {
          return r(!(this.negative | t.negative)), this.iuor(t);
        }),
        (i.prototype.or = function (t) {
          return this.length > t.length
            ? this.clone().ior(t)
            : t.clone().ior(this);
        }),
        (i.prototype.uor = function (t) {
          return this.length > t.length
            ? this.clone().iuor(t)
            : t.clone().iuor(this);
        }),
        (i.prototype.iuand = function (t) {
          var e;
          e = this.length > t.length ? t : this;
          for (var r = 0; r < e.length; r++)
            this.words[r] = this.words[r] & t.words[r];
          return (this.length = e.length), this._strip();
        }),
        (i.prototype.iand = function (t) {
          return r(!(this.negative | t.negative)), this.iuand(t);
        }),
        (i.prototype.and = function (t) {
          return this.length > t.length
            ? this.clone().iand(t)
            : t.clone().iand(this);
        }),
        (i.prototype.uand = function (t) {
          return this.length > t.length
            ? this.clone().iuand(t)
            : t.clone().iuand(this);
        }),
        (i.prototype.iuxor = function (t) {
          var e, r;
          this.length > t.length
            ? ((e = this), (r = t))
            : ((e = t), (r = this));
          for (var n = 0; n < r.length; n++)
            this.words[n] = e.words[n] ^ r.words[n];
          if (this !== e) for (; n < e.length; n++) this.words[n] = e.words[n];
          return (this.length = e.length), this._strip();
        }),
        (i.prototype.ixor = function (t) {
          return r(!(this.negative | t.negative)), this.iuxor(t);
        }),
        (i.prototype.xor = function (t) {
          return this.length > t.length
            ? this.clone().ixor(t)
            : t.clone().ixor(this);
        }),
        (i.prototype.uxor = function (t) {
          return this.length > t.length
            ? this.clone().iuxor(t)
            : t.clone().iuxor(this);
        }),
        (i.prototype.inotn = function (t) {
          r('number' == typeof t && t >= 0);
          var e = 0 | Math.ceil(t / 26),
            n = t % 26;
          this._expand(e), n > 0 && e--;
          for (var i = 0; i < e; i++) this.words[i] = 67108863 & ~this.words[i];
          return (
            n > 0 && (this.words[i] = ~this.words[i] & (67108863 >> (26 - n))),
            this._strip()
          );
        }),
        (i.prototype.notn = function (t) {
          return this.clone().inotn(t);
        }),
        (i.prototype.setn = function (t, e) {
          r('number' == typeof t && t >= 0);
          var n = (t / 26) | 0,
            i = t % 26;
          return (
            this._expand(n + 1),
            (this.words[n] = e
              ? this.words[n] | (1 << i)
              : this.words[n] & ~(1 << i)),
            this._strip()
          );
        }),
        (i.prototype.iadd = function (t) {
          var e, r, n;
          if (0 !== this.negative && 0 === t.negative)
            return (
              (this.negative = 0),
              (e = this.isub(t)),
              (this.negative ^= 1),
              this._normSign()
            );
          if (0 === this.negative && 0 !== t.negative)
            return (
              (t.negative = 0),
              (e = this.isub(t)),
              (t.negative = 1),
              e._normSign()
            );
          this.length > t.length
            ? ((r = this), (n = t))
            : ((r = t), (n = this));
          for (var i = 0, o = 0; o < n.length; o++)
            (e = (0 | r.words[o]) + (0 | n.words[o]) + i),
              (this.words[o] = 67108863 & e),
              (i = e >>> 26);
          for (; 0 !== i && o < r.length; o++)
            (e = (0 | r.words[o]) + i),
              (this.words[o] = 67108863 & e),
              (i = e >>> 26);
          if (((this.length = r.length), 0 !== i))
            (this.words[this.length] = i), this.length++;
          else if (r !== this)
            for (; o < r.length; o++) this.words[o] = r.words[o];
          return this;
        }),
        (i.prototype.add = function (t) {
          var e;
          return 0 !== t.negative && 0 === this.negative
            ? ((t.negative = 0), (e = this.sub(t)), (t.negative ^= 1), e)
            : 0 === t.negative && 0 !== this.negative
            ? ((this.negative = 0), (e = t.sub(this)), (this.negative = 1), e)
            : this.length > t.length
            ? this.clone().iadd(t)
            : t.clone().iadd(this);
        }),
        (i.prototype.isub = function (t) {
          if (0 !== t.negative) {
            t.negative = 0;
            var e = this.iadd(t);
            return (t.negative = 1), e._normSign();
          }
          if (0 !== this.negative)
            return (
              (this.negative = 0),
              this.iadd(t),
              (this.negative = 1),
              this._normSign()
            );
          var r,
            n,
            i = this.cmp(t);
          if (0 === i)
            return (
              (this.negative = 0), (this.length = 1), (this.words[0] = 0), this
            );
          i > 0 ? ((r = this), (n = t)) : ((r = t), (n = this));
          for (var o = 0, s = 0; s < n.length; s++)
            (o = (e = (0 | r.words[s]) - (0 | n.words[s]) + o) >> 26),
              (this.words[s] = 67108863 & e);
          for (; 0 !== o && s < r.length; s++)
            (o = (e = (0 | r.words[s]) + o) >> 26),
              (this.words[s] = 67108863 & e);
          if (0 === o && s < r.length && r !== this)
            for (; s < r.length; s++) this.words[s] = r.words[s];
          return (
            (this.length = Math.max(this.length, s)),
            r !== this && (this.negative = 1),
            this._strip()
          );
        }),
        (i.prototype.sub = function (t) {
          return this.clone().isub(t);
        });
      var p = function (t, e, r) {
        var n,
          i,
          o,
          s = t.words,
          a = e.words,
          u = r.words,
          h = 0,
          l = 0 | s[0],
          c = 8191 & l,
          f = l >>> 13,
          d = 0 | s[1],
          m = 8191 & d,
          p = d >>> 13,
          g = 0 | s[2],
          y = 8191 & g,
          v = g >>> 13,
          w = 0 | s[3],
          b = 8191 & w,
          A = w >>> 13,
          E = 0 | s[4],
          M = 8191 & E,
          _ = E >>> 13,
          k = 0 | s[5],
          S = 8191 & k,
          N = k >>> 13,
          P = 0 | s[6],
          x = 8191 & P,
          T = P >>> 13,
          C = 0 | s[7],
          R = 8191 & C,
          I = C >>> 13,
          B = 0 | s[8],
          O = 8191 & B,
          F = B >>> 13,
          D = 0 | s[9],
          L = 8191 & D,
          U = D >>> 13,
          j = 0 | a[0],
          z = 8191 & j,
          G = j >>> 13,
          q = 0 | a[1],
          K = 8191 & q,
          H = q >>> 13,
          V = 0 | a[2],
          J = 8191 & V,
          W = V >>> 13,
          Q = 0 | a[3],
          Y = 8191 & Q,
          Z = Q >>> 13,
          X = 0 | a[4],
          $ = 8191 & X,
          tt = X >>> 13,
          et = 0 | a[5],
          rt = 8191 & et,
          nt = et >>> 13,
          it = 0 | a[6],
          ot = 8191 & it,
          st = it >>> 13,
          at = 0 | a[7],
          ut = 8191 & at,
          ht = at >>> 13,
          lt = 0 | a[8],
          ct = 8191 & lt,
          ft = lt >>> 13,
          dt = 0 | a[9],
          mt = 8191 & dt,
          pt = dt >>> 13;
        (r.negative = t.negative ^ e.negative), (r.length = 19);
        var gt =
          (((h + (n = Math.imul(c, z))) | 0) +
            ((8191 & (i = ((i = Math.imul(c, G)) + Math.imul(f, z)) | 0)) <<
              13)) |
          0;
        (h = ((((o = Math.imul(f, G)) + (i >>> 13)) | 0) + (gt >>> 26)) | 0),
          (gt &= 67108863),
          (n = Math.imul(m, z)),
          (i = ((i = Math.imul(m, G)) + Math.imul(p, z)) | 0),
          (o = Math.imul(p, G));
        var yt =
          (((h + (n = (n + Math.imul(c, K)) | 0)) | 0) +
            ((8191 &
              (i = ((i = (i + Math.imul(c, H)) | 0) + Math.imul(f, K)) | 0)) <<
              13)) |
          0;
        (h =
          ((((o = (o + Math.imul(f, H)) | 0) + (i >>> 13)) | 0) + (yt >>> 26)) |
          0),
          (yt &= 67108863),
          (n = Math.imul(y, z)),
          (i = ((i = Math.imul(y, G)) + Math.imul(v, z)) | 0),
          (o = Math.imul(v, G)),
          (n = (n + Math.imul(m, K)) | 0),
          (i = ((i = (i + Math.imul(m, H)) | 0) + Math.imul(p, K)) | 0),
          (o = (o + Math.imul(p, H)) | 0);
        var vt =
          (((h + (n = (n + Math.imul(c, J)) | 0)) | 0) +
            ((8191 &
              (i = ((i = (i + Math.imul(c, W)) | 0) + Math.imul(f, J)) | 0)) <<
              13)) |
          0;
        (h =
          ((((o = (o + Math.imul(f, W)) | 0) + (i >>> 13)) | 0) + (vt >>> 26)) |
          0),
          (vt &= 67108863),
          (n = Math.imul(b, z)),
          (i = ((i = Math.imul(b, G)) + Math.imul(A, z)) | 0),
          (o = Math.imul(A, G)),
          (n = (n + Math.imul(y, K)) | 0),
          (i = ((i = (i + Math.imul(y, H)) | 0) + Math.imul(v, K)) | 0),
          (o = (o + Math.imul(v, H)) | 0),
          (n = (n + Math.imul(m, J)) | 0),
          (i = ((i = (i + Math.imul(m, W)) | 0) + Math.imul(p, J)) | 0),
          (o = (o + Math.imul(p, W)) | 0);
        var wt =
          (((h + (n = (n + Math.imul(c, Y)) | 0)) | 0) +
            ((8191 &
              (i = ((i = (i + Math.imul(c, Z)) | 0) + Math.imul(f, Y)) | 0)) <<
              13)) |
          0;
        (h =
          ((((o = (o + Math.imul(f, Z)) | 0) + (i >>> 13)) | 0) + (wt >>> 26)) |
          0),
          (wt &= 67108863),
          (n = Math.imul(M, z)),
          (i = ((i = Math.imul(M, G)) + Math.imul(_, z)) | 0),
          (o = Math.imul(_, G)),
          (n = (n + Math.imul(b, K)) | 0),
          (i = ((i = (i + Math.imul(b, H)) | 0) + Math.imul(A, K)) | 0),
          (o = (o + Math.imul(A, H)) | 0),
          (n = (n + Math.imul(y, J)) | 0),
          (i = ((i = (i + Math.imul(y, W)) | 0) + Math.imul(v, J)) | 0),
          (o = (o + Math.imul(v, W)) | 0),
          (n = (n + Math.imul(m, Y)) | 0),
          (i = ((i = (i + Math.imul(m, Z)) | 0) + Math.imul(p, Y)) | 0),
          (o = (o + Math.imul(p, Z)) | 0);
        var bt =
          (((h + (n = (n + Math.imul(c, $)) | 0)) | 0) +
            ((8191 &
              (i = ((i = (i + Math.imul(c, tt)) | 0) + Math.imul(f, $)) | 0)) <<
              13)) |
          0;
        (h =
          ((((o = (o + Math.imul(f, tt)) | 0) + (i >>> 13)) | 0) +
            (bt >>> 26)) |
          0),
          (bt &= 67108863),
          (n = Math.imul(S, z)),
          (i = ((i = Math.imul(S, G)) + Math.imul(N, z)) | 0),
          (o = Math.imul(N, G)),
          (n = (n + Math.imul(M, K)) | 0),
          (i = ((i = (i + Math.imul(M, H)) | 0) + Math.imul(_, K)) | 0),
          (o = (o + Math.imul(_, H)) | 0),
          (n = (n + Math.imul(b, J)) | 0),
          (i = ((i = (i + Math.imul(b, W)) | 0) + Math.imul(A, J)) | 0),
          (o = (o + Math.imul(A, W)) | 0),
          (n = (n + Math.imul(y, Y)) | 0),
          (i = ((i = (i + Math.imul(y, Z)) | 0) + Math.imul(v, Y)) | 0),
          (o = (o + Math.imul(v, Z)) | 0),
          (n = (n + Math.imul(m, $)) | 0),
          (i = ((i = (i + Math.imul(m, tt)) | 0) + Math.imul(p, $)) | 0),
          (o = (o + Math.imul(p, tt)) | 0);
        var At =
          (((h + (n = (n + Math.imul(c, rt)) | 0)) | 0) +
            ((8191 &
              (i =
                ((i = (i + Math.imul(c, nt)) | 0) + Math.imul(f, rt)) | 0)) <<
              13)) |
          0;
        (h =
          ((((o = (o + Math.imul(f, nt)) | 0) + (i >>> 13)) | 0) +
            (At >>> 26)) |
          0),
          (At &= 67108863),
          (n = Math.imul(x, z)),
          (i = ((i = Math.imul(x, G)) + Math.imul(T, z)) | 0),
          (o = Math.imul(T, G)),
          (n = (n + Math.imul(S, K)) | 0),
          (i = ((i = (i + Math.imul(S, H)) | 0) + Math.imul(N, K)) | 0),
          (o = (o + Math.imul(N, H)) | 0),
          (n = (n + Math.imul(M, J)) | 0),
          (i = ((i = (i + Math.imul(M, W)) | 0) + Math.imul(_, J)) | 0),
          (o = (o + Math.imul(_, W)) | 0),
          (n = (n + Math.imul(b, Y)) | 0),
          (i = ((i = (i + Math.imul(b, Z)) | 0) + Math.imul(A, Y)) | 0),
          (o = (o + Math.imul(A, Z)) | 0),
          (n = (n + Math.imul(y, $)) | 0),
          (i = ((i = (i + Math.imul(y, tt)) | 0) + Math.imul(v, $)) | 0),
          (o = (o + Math.imul(v, tt)) | 0),
          (n = (n + Math.imul(m, rt)) | 0),
          (i = ((i = (i + Math.imul(m, nt)) | 0) + Math.imul(p, rt)) | 0),
          (o = (o + Math.imul(p, nt)) | 0);
        var Et =
          (((h + (n = (n + Math.imul(c, ot)) | 0)) | 0) +
            ((8191 &
              (i =
                ((i = (i + Math.imul(c, st)) | 0) + Math.imul(f, ot)) | 0)) <<
              13)) |
          0;
        (h =
          ((((o = (o + Math.imul(f, st)) | 0) + (i >>> 13)) | 0) +
            (Et >>> 26)) |
          0),
          (Et &= 67108863),
          (n = Math.imul(R, z)),
          (i = ((i = Math.imul(R, G)) + Math.imul(I, z)) | 0),
          (o = Math.imul(I, G)),
          (n = (n + Math.imul(x, K)) | 0),
          (i = ((i = (i + Math.imul(x, H)) | 0) + Math.imul(T, K)) | 0),
          (o = (o + Math.imul(T, H)) | 0),
          (n = (n + Math.imul(S, J)) | 0),
          (i = ((i = (i + Math.imul(S, W)) | 0) + Math.imul(N, J)) | 0),
          (o = (o + Math.imul(N, W)) | 0),
          (n = (n + Math.imul(M, Y)) | 0),
          (i = ((i = (i + Math.imul(M, Z)) | 0) + Math.imul(_, Y)) | 0),
          (o = (o + Math.imul(_, Z)) | 0),
          (n = (n + Math.imul(b, $)) | 0),
          (i = ((i = (i + Math.imul(b, tt)) | 0) + Math.imul(A, $)) | 0),
          (o = (o + Math.imul(A, tt)) | 0),
          (n = (n + Math.imul(y, rt)) | 0),
          (i = ((i = (i + Math.imul(y, nt)) | 0) + Math.imul(v, rt)) | 0),
          (o = (o + Math.imul(v, nt)) | 0),
          (n = (n + Math.imul(m, ot)) | 0),
          (i = ((i = (i + Math.imul(m, st)) | 0) + Math.imul(p, ot)) | 0),
          (o = (o + Math.imul(p, st)) | 0);
        var Mt =
          (((h + (n = (n + Math.imul(c, ut)) | 0)) | 0) +
            ((8191 &
              (i =
                ((i = (i + Math.imul(c, ht)) | 0) + Math.imul(f, ut)) | 0)) <<
              13)) |
          0;
        (h =
          ((((o = (o + Math.imul(f, ht)) | 0) + (i >>> 13)) | 0) +
            (Mt >>> 26)) |
          0),
          (Mt &= 67108863),
          (n = Math.imul(O, z)),
          (i = ((i = Math.imul(O, G)) + Math.imul(F, z)) | 0),
          (o = Math.imul(F, G)),
          (n = (n + Math.imul(R, K)) | 0),
          (i = ((i = (i + Math.imul(R, H)) | 0) + Math.imul(I, K)) | 0),
          (o = (o + Math.imul(I, H)) | 0),
          (n = (n + Math.imul(x, J)) | 0),
          (i = ((i = (i + Math.imul(x, W)) | 0) + Math.imul(T, J)) | 0),
          (o = (o + Math.imul(T, W)) | 0),
          (n = (n + Math.imul(S, Y)) | 0),
          (i = ((i = (i + Math.imul(S, Z)) | 0) + Math.imul(N, Y)) | 0),
          (o = (o + Math.imul(N, Z)) | 0),
          (n = (n + Math.imul(M, $)) | 0),
          (i = ((i = (i + Math.imul(M, tt)) | 0) + Math.imul(_, $)) | 0),
          (o = (o + Math.imul(_, tt)) | 0),
          (n = (n + Math.imul(b, rt)) | 0),
          (i = ((i = (i + Math.imul(b, nt)) | 0) + Math.imul(A, rt)) | 0),
          (o = (o + Math.imul(A, nt)) | 0),
          (n = (n + Math.imul(y, ot)) | 0),
          (i = ((i = (i + Math.imul(y, st)) | 0) + Math.imul(v, ot)) | 0),
          (o = (o + Math.imul(v, st)) | 0),
          (n = (n + Math.imul(m, ut)) | 0),
          (i = ((i = (i + Math.imul(m, ht)) | 0) + Math.imul(p, ut)) | 0),
          (o = (o + Math.imul(p, ht)) | 0);
        var _t =
          (((h + (n = (n + Math.imul(c, ct)) | 0)) | 0) +
            ((8191 &
              (i =
                ((i = (i + Math.imul(c, ft)) | 0) + Math.imul(f, ct)) | 0)) <<
              13)) |
          0;
        (h =
          ((((o = (o + Math.imul(f, ft)) | 0) + (i >>> 13)) | 0) +
            (_t >>> 26)) |
          0),
          (_t &= 67108863),
          (n = Math.imul(L, z)),
          (i = ((i = Math.imul(L, G)) + Math.imul(U, z)) | 0),
          (o = Math.imul(U, G)),
          (n = (n + Math.imul(O, K)) | 0),
          (i = ((i = (i + Math.imul(O, H)) | 0) + Math.imul(F, K)) | 0),
          (o = (o + Math.imul(F, H)) | 0),
          (n = (n + Math.imul(R, J)) | 0),
          (i = ((i = (i + Math.imul(R, W)) | 0) + Math.imul(I, J)) | 0),
          (o = (o + Math.imul(I, W)) | 0),
          (n = (n + Math.imul(x, Y)) | 0),
          (i = ((i = (i + Math.imul(x, Z)) | 0) + Math.imul(T, Y)) | 0),
          (o = (o + Math.imul(T, Z)) | 0),
          (n = (n + Math.imul(S, $)) | 0),
          (i = ((i = (i + Math.imul(S, tt)) | 0) + Math.imul(N, $)) | 0),
          (o = (o + Math.imul(N, tt)) | 0),
          (n = (n + Math.imul(M, rt)) | 0),
          (i = ((i = (i + Math.imul(M, nt)) | 0) + Math.imul(_, rt)) | 0),
          (o = (o + Math.imul(_, nt)) | 0),
          (n = (n + Math.imul(b, ot)) | 0),
          (i = ((i = (i + Math.imul(b, st)) | 0) + Math.imul(A, ot)) | 0),
          (o = (o + Math.imul(A, st)) | 0),
          (n = (n + Math.imul(y, ut)) | 0),
          (i = ((i = (i + Math.imul(y, ht)) | 0) + Math.imul(v, ut)) | 0),
          (o = (o + Math.imul(v, ht)) | 0),
          (n = (n + Math.imul(m, ct)) | 0),
          (i = ((i = (i + Math.imul(m, ft)) | 0) + Math.imul(p, ct)) | 0),
          (o = (o + Math.imul(p, ft)) | 0);
        var kt =
          (((h + (n = (n + Math.imul(c, mt)) | 0)) | 0) +
            ((8191 &
              (i =
                ((i = (i + Math.imul(c, pt)) | 0) + Math.imul(f, mt)) | 0)) <<
              13)) |
          0;
        (h =
          ((((o = (o + Math.imul(f, pt)) | 0) + (i >>> 13)) | 0) +
            (kt >>> 26)) |
          0),
          (kt &= 67108863),
          (n = Math.imul(L, K)),
          (i = ((i = Math.imul(L, H)) + Math.imul(U, K)) | 0),
          (o = Math.imul(U, H)),
          (n = (n + Math.imul(O, J)) | 0),
          (i = ((i = (i + Math.imul(O, W)) | 0) + Math.imul(F, J)) | 0),
          (o = (o + Math.imul(F, W)) | 0),
          (n = (n + Math.imul(R, Y)) | 0),
          (i = ((i = (i + Math.imul(R, Z)) | 0) + Math.imul(I, Y)) | 0),
          (o = (o + Math.imul(I, Z)) | 0),
          (n = (n + Math.imul(x, $)) | 0),
          (i = ((i = (i + Math.imul(x, tt)) | 0) + Math.imul(T, $)) | 0),
          (o = (o + Math.imul(T, tt)) | 0),
          (n = (n + Math.imul(S, rt)) | 0),
          (i = ((i = (i + Math.imul(S, nt)) | 0) + Math.imul(N, rt)) | 0),
          (o = (o + Math.imul(N, nt)) | 0),
          (n = (n + Math.imul(M, ot)) | 0),
          (i = ((i = (i + Math.imul(M, st)) | 0) + Math.imul(_, ot)) | 0),
          (o = (o + Math.imul(_, st)) | 0),
          (n = (n + Math.imul(b, ut)) | 0),
          (i = ((i = (i + Math.imul(b, ht)) | 0) + Math.imul(A, ut)) | 0),
          (o = (o + Math.imul(A, ht)) | 0),
          (n = (n + Math.imul(y, ct)) | 0),
          (i = ((i = (i + Math.imul(y, ft)) | 0) + Math.imul(v, ct)) | 0),
          (o = (o + Math.imul(v, ft)) | 0);
        var St =
          (((h + (n = (n + Math.imul(m, mt)) | 0)) | 0) +
            ((8191 &
              (i =
                ((i = (i + Math.imul(m, pt)) | 0) + Math.imul(p, mt)) | 0)) <<
              13)) |
          0;
        (h =
          ((((o = (o + Math.imul(p, pt)) | 0) + (i >>> 13)) | 0) +
            (St >>> 26)) |
          0),
          (St &= 67108863),
          (n = Math.imul(L, J)),
          (i = ((i = Math.imul(L, W)) + Math.imul(U, J)) | 0),
          (o = Math.imul(U, W)),
          (n = (n + Math.imul(O, Y)) | 0),
          (i = ((i = (i + Math.imul(O, Z)) | 0) + Math.imul(F, Y)) | 0),
          (o = (o + Math.imul(F, Z)) | 0),
          (n = (n + Math.imul(R, $)) | 0),
          (i = ((i = (i + Math.imul(R, tt)) | 0) + Math.imul(I, $)) | 0),
          (o = (o + Math.imul(I, tt)) | 0),
          (n = (n + Math.imul(x, rt)) | 0),
          (i = ((i = (i + Math.imul(x, nt)) | 0) + Math.imul(T, rt)) | 0),
          (o = (o + Math.imul(T, nt)) | 0),
          (n = (n + Math.imul(S, ot)) | 0),
          (i = ((i = (i + Math.imul(S, st)) | 0) + Math.imul(N, ot)) | 0),
          (o = (o + Math.imul(N, st)) | 0),
          (n = (n + Math.imul(M, ut)) | 0),
          (i = ((i = (i + Math.imul(M, ht)) | 0) + Math.imul(_, ut)) | 0),
          (o = (o + Math.imul(_, ht)) | 0),
          (n = (n + Math.imul(b, ct)) | 0),
          (i = ((i = (i + Math.imul(b, ft)) | 0) + Math.imul(A, ct)) | 0),
          (o = (o + Math.imul(A, ft)) | 0);
        var Nt =
          (((h + (n = (n + Math.imul(y, mt)) | 0)) | 0) +
            ((8191 &
              (i =
                ((i = (i + Math.imul(y, pt)) | 0) + Math.imul(v, mt)) | 0)) <<
              13)) |
          0;
        (h =
          ((((o = (o + Math.imul(v, pt)) | 0) + (i >>> 13)) | 0) +
            (Nt >>> 26)) |
          0),
          (Nt &= 67108863),
          (n = Math.imul(L, Y)),
          (i = ((i = Math.imul(L, Z)) + Math.imul(U, Y)) | 0),
          (o = Math.imul(U, Z)),
          (n = (n + Math.imul(O, $)) | 0),
          (i = ((i = (i + Math.imul(O, tt)) | 0) + Math.imul(F, $)) | 0),
          (o = (o + Math.imul(F, tt)) | 0),
          (n = (n + Math.imul(R, rt)) | 0),
          (i = ((i = (i + Math.imul(R, nt)) | 0) + Math.imul(I, rt)) | 0),
          (o = (o + Math.imul(I, nt)) | 0),
          (n = (n + Math.imul(x, ot)) | 0),
          (i = ((i = (i + Math.imul(x, st)) | 0) + Math.imul(T, ot)) | 0),
          (o = (o + Math.imul(T, st)) | 0),
          (n = (n + Math.imul(S, ut)) | 0),
          (i = ((i = (i + Math.imul(S, ht)) | 0) + Math.imul(N, ut)) | 0),
          (o = (o + Math.imul(N, ht)) | 0),
          (n = (n + Math.imul(M, ct)) | 0),
          (i = ((i = (i + Math.imul(M, ft)) | 0) + Math.imul(_, ct)) | 0),
          (o = (o + Math.imul(_, ft)) | 0);
        var Pt =
          (((h + (n = (n + Math.imul(b, mt)) | 0)) | 0) +
            ((8191 &
              (i =
                ((i = (i + Math.imul(b, pt)) | 0) + Math.imul(A, mt)) | 0)) <<
              13)) |
          0;
        (h =
          ((((o = (o + Math.imul(A, pt)) | 0) + (i >>> 13)) | 0) +
            (Pt >>> 26)) |
          0),
          (Pt &= 67108863),
          (n = Math.imul(L, $)),
          (i = ((i = Math.imul(L, tt)) + Math.imul(U, $)) | 0),
          (o = Math.imul(U, tt)),
          (n = (n + Math.imul(O, rt)) | 0),
          (i = ((i = (i + Math.imul(O, nt)) | 0) + Math.imul(F, rt)) | 0),
          (o = (o + Math.imul(F, nt)) | 0),
          (n = (n + Math.imul(R, ot)) | 0),
          (i = ((i = (i + Math.imul(R, st)) | 0) + Math.imul(I, ot)) | 0),
          (o = (o + Math.imul(I, st)) | 0),
          (n = (n + Math.imul(x, ut)) | 0),
          (i = ((i = (i + Math.imul(x, ht)) | 0) + Math.imul(T, ut)) | 0),
          (o = (o + Math.imul(T, ht)) | 0),
          (n = (n + Math.imul(S, ct)) | 0),
          (i = ((i = (i + Math.imul(S, ft)) | 0) + Math.imul(N, ct)) | 0),
          (o = (o + Math.imul(N, ft)) | 0);
        var xt =
          (((h + (n = (n + Math.imul(M, mt)) | 0)) | 0) +
            ((8191 &
              (i =
                ((i = (i + Math.imul(M, pt)) | 0) + Math.imul(_, mt)) | 0)) <<
              13)) |
          0;
        (h =
          ((((o = (o + Math.imul(_, pt)) | 0) + (i >>> 13)) | 0) +
            (xt >>> 26)) |
          0),
          (xt &= 67108863),
          (n = Math.imul(L, rt)),
          (i = ((i = Math.imul(L, nt)) + Math.imul(U, rt)) | 0),
          (o = Math.imul(U, nt)),
          (n = (n + Math.imul(O, ot)) | 0),
          (i = ((i = (i + Math.imul(O, st)) | 0) + Math.imul(F, ot)) | 0),
          (o = (o + Math.imul(F, st)) | 0),
          (n = (n + Math.imul(R, ut)) | 0),
          (i = ((i = (i + Math.imul(R, ht)) | 0) + Math.imul(I, ut)) | 0),
          (o = (o + Math.imul(I, ht)) | 0),
          (n = (n + Math.imul(x, ct)) | 0),
          (i = ((i = (i + Math.imul(x, ft)) | 0) + Math.imul(T, ct)) | 0),
          (o = (o + Math.imul(T, ft)) | 0);
        var Tt =
          (((h + (n = (n + Math.imul(S, mt)) | 0)) | 0) +
            ((8191 &
              (i =
                ((i = (i + Math.imul(S, pt)) | 0) + Math.imul(N, mt)) | 0)) <<
              13)) |
          0;
        (h =
          ((((o = (o + Math.imul(N, pt)) | 0) + (i >>> 13)) | 0) +
            (Tt >>> 26)) |
          0),
          (Tt &= 67108863),
          (n = Math.imul(L, ot)),
          (i = ((i = Math.imul(L, st)) + Math.imul(U, ot)) | 0),
          (o = Math.imul(U, st)),
          (n = (n + Math.imul(O, ut)) | 0),
          (i = ((i = (i + Math.imul(O, ht)) | 0) + Math.imul(F, ut)) | 0),
          (o = (o + Math.imul(F, ht)) | 0),
          (n = (n + Math.imul(R, ct)) | 0),
          (i = ((i = (i + Math.imul(R, ft)) | 0) + Math.imul(I, ct)) | 0),
          (o = (o + Math.imul(I, ft)) | 0);
        var Ct =
          (((h + (n = (n + Math.imul(x, mt)) | 0)) | 0) +
            ((8191 &
              (i =
                ((i = (i + Math.imul(x, pt)) | 0) + Math.imul(T, mt)) | 0)) <<
              13)) |
          0;
        (h =
          ((((o = (o + Math.imul(T, pt)) | 0) + (i >>> 13)) | 0) +
            (Ct >>> 26)) |
          0),
          (Ct &= 67108863),
          (n = Math.imul(L, ut)),
          (i = ((i = Math.imul(L, ht)) + Math.imul(U, ut)) | 0),
          (o = Math.imul(U, ht)),
          (n = (n + Math.imul(O, ct)) | 0),
          (i = ((i = (i + Math.imul(O, ft)) | 0) + Math.imul(F, ct)) | 0),
          (o = (o + Math.imul(F, ft)) | 0);
        var Rt =
          (((h + (n = (n + Math.imul(R, mt)) | 0)) | 0) +
            ((8191 &
              (i =
                ((i = (i + Math.imul(R, pt)) | 0) + Math.imul(I, mt)) | 0)) <<
              13)) |
          0;
        (h =
          ((((o = (o + Math.imul(I, pt)) | 0) + (i >>> 13)) | 0) +
            (Rt >>> 26)) |
          0),
          (Rt &= 67108863),
          (n = Math.imul(L, ct)),
          (i = ((i = Math.imul(L, ft)) + Math.imul(U, ct)) | 0),
          (o = Math.imul(U, ft));
        var It =
          (((h + (n = (n + Math.imul(O, mt)) | 0)) | 0) +
            ((8191 &
              (i =
                ((i = (i + Math.imul(O, pt)) | 0) + Math.imul(F, mt)) | 0)) <<
              13)) |
          0;
        (h =
          ((((o = (o + Math.imul(F, pt)) | 0) + (i >>> 13)) | 0) +
            (It >>> 26)) |
          0),
          (It &= 67108863);
        var Bt =
          (((h + (n = Math.imul(L, mt))) | 0) +
            ((8191 & (i = ((i = Math.imul(L, pt)) + Math.imul(U, mt)) | 0)) <<
              13)) |
          0;
        return (
          (h = ((((o = Math.imul(U, pt)) + (i >>> 13)) | 0) + (Bt >>> 26)) | 0),
          (Bt &= 67108863),
          (u[0] = gt),
          (u[1] = yt),
          (u[2] = vt),
          (u[3] = wt),
          (u[4] = bt),
          (u[5] = At),
          (u[6] = Et),
          (u[7] = Mt),
          (u[8] = _t),
          (u[9] = kt),
          (u[10] = St),
          (u[11] = Nt),
          (u[12] = Pt),
          (u[13] = xt),
          (u[14] = Tt),
          (u[15] = Ct),
          (u[16] = Rt),
          (u[17] = It),
          (u[18] = Bt),
          0 !== h && ((u[19] = h), r.length++),
          r
        );
      };
      function g(t, e, r) {
        (r.negative = e.negative ^ t.negative),
          (r.length = t.length + e.length);
        for (var n = 0, i = 0, o = 0; o < r.length - 1; o++) {
          var s = i;
          i = 0;
          for (
            var a = 67108863 & n,
              u = Math.min(o, e.length - 1),
              h = Math.max(0, o - t.length + 1);
            h <= u;
            h++
          ) {
            var l = o - h,
              c = (0 | t.words[l]) * (0 | e.words[h]),
              f = 67108863 & c;
            (a = 67108863 & (f = (f + a) | 0)),
              (i +=
                (s =
                  ((s = (s + ((c / 67108864) | 0)) | 0) + (f >>> 26)) | 0) >>>
                26),
              (s &= 67108863);
          }
          (r.words[o] = a), (n = s), (s = i);
        }
        return 0 !== n ? (r.words[o] = n) : r.length--, r._strip();
      }
      function y(t, e, r) {
        return g(t, e, r);
      }
      Math.imul || (p = m),
        (i.prototype.mulTo = function (t, e) {
          var r = this.length + t.length;
          return 10 === this.length && 10 === t.length
            ? p(this, t, e)
            : r < 63
            ? m(this, t, e)
            : r < 1024
            ? g(this, t, e)
            : y(this, t, e);
        }),
        (i.prototype.mul = function (t) {
          var e = new i(null);
          return (
            (e.words = new Array(this.length + t.length)), this.mulTo(t, e)
          );
        }),
        (i.prototype.mulf = function (t) {
          var e = new i(null);
          return (e.words = new Array(this.length + t.length)), y(this, t, e);
        }),
        (i.prototype.imul = function (t) {
          return this.clone().mulTo(t, this);
        }),
        (i.prototype.imuln = function (t) {
          var e = t < 0;
          e && (t = -t), r('number' == typeof t), r(t < 67108864);
          for (var n = 0, i = 0; i < this.length; i++) {
            var o = (0 | this.words[i]) * t,
              s = (67108863 & o) + (67108863 & n);
            (n >>= 26),
              (n += (o / 67108864) | 0),
              (n += s >>> 26),
              (this.words[i] = 67108863 & s);
          }
          return (
            0 !== n && ((this.words[i] = n), this.length++),
            e ? this.ineg() : this
          );
        }),
        (i.prototype.muln = function (t) {
          return this.clone().imuln(t);
        }),
        (i.prototype.sqr = function () {
          return this.mul(this);
        }),
        (i.prototype.isqr = function () {
          return this.imul(this.clone());
        }),
        (i.prototype.pow = function (t) {
          var e = (function (t) {
            for (var e = new Array(t.bitLength()), r = 0; r < e.length; r++) {
              var n = (r / 26) | 0,
                i = r % 26;
              e[r] = (t.words[n] >>> i) & 1;
            }
            return e;
          })(t);
          if (0 === e.length) return new i(1);
          for (
            var r = this, n = 0;
            n < e.length && 0 === e[n];
            n++, r = r.sqr()
          );
          if (++n < e.length)
            for (var o = r.sqr(); n < e.length; n++, o = o.sqr())
              0 !== e[n] && (r = r.mul(o));
          return r;
        }),
        (i.prototype.iushln = function (t) {
          r('number' == typeof t && t >= 0);
          var e,
            n = t % 26,
            i = (t - n) / 26,
            o = (67108863 >>> (26 - n)) << (26 - n);
          if (0 !== n) {
            var s = 0;
            for (e = 0; e < this.length; e++) {
              var a = this.words[e] & o,
                u = ((0 | this.words[e]) - a) << n;
              (this.words[e] = u | s), (s = a >>> (26 - n));
            }
            s && ((this.words[e] = s), this.length++);
          }
          if (0 !== i) {
            for (e = this.length - 1; e >= 0; e--)
              this.words[e + i] = this.words[e];
            for (e = 0; e < i; e++) this.words[e] = 0;
            this.length += i;
          }
          return this._strip();
        }),
        (i.prototype.ishln = function (t) {
          return r(0 === this.negative), this.iushln(t);
        }),
        (i.prototype.iushrn = function (t, e, n) {
          var i;
          r('number' == typeof t && t >= 0), (i = e ? (e - (e % 26)) / 26 : 0);
          var o = t % 26,
            s = Math.min((t - o) / 26, this.length),
            a = 67108863 ^ ((67108863 >>> o) << o),
            u = n;
          if (((i -= s), (i = Math.max(0, i)), u)) {
            for (var h = 0; h < s; h++) u.words[h] = this.words[h];
            u.length = s;
          }
          if (0 === s);
          else if (this.length > s)
            for (this.length -= s, h = 0; h < this.length; h++)
              this.words[h] = this.words[h + s];
          else (this.words[0] = 0), (this.length = 1);
          var l = 0;
          for (h = this.length - 1; h >= 0 && (0 !== l || h >= i); h--) {
            var c = 0 | this.words[h];
            (this.words[h] = (l << (26 - o)) | (c >>> o)), (l = c & a);
          }
          return (
            u && 0 !== l && (u.words[u.length++] = l),
            0 === this.length && ((this.words[0] = 0), (this.length = 1)),
            this._strip()
          );
        }),
        (i.prototype.ishrn = function (t, e, n) {
          return r(0 === this.negative), this.iushrn(t, e, n);
        }),
        (i.prototype.shln = function (t) {
          return this.clone().ishln(t);
        }),
        (i.prototype.ushln = function (t) {
          return this.clone().iushln(t);
        }),
        (i.prototype.shrn = function (t) {
          return this.clone().ishrn(t);
        }),
        (i.prototype.ushrn = function (t) {
          return this.clone().iushrn(t);
        }),
        (i.prototype.testn = function (t) {
          r('number' == typeof t && t >= 0);
          var e = t % 26,
            n = (t - e) / 26,
            i = 1 << e;
          return !(this.length <= n) && !!(this.words[n] & i);
        }),
        (i.prototype.imaskn = function (t) {
          r('number' == typeof t && t >= 0);
          var e = t % 26,
            n = (t - e) / 26;
          if (
            (r(0 === this.negative, 'imaskn works only with positive numbers'),
            this.length <= n)
          )
            return this;
          if (
            (0 !== e && n++, (this.length = Math.min(n, this.length)), 0 !== e)
          ) {
            var i = 67108863 ^ ((67108863 >>> e) << e);
            this.words[this.length - 1] &= i;
          }
          return this._strip();
        }),
        (i.prototype.maskn = function (t) {
          return this.clone().imaskn(t);
        }),
        (i.prototype.iaddn = function (t) {
          return (
            r('number' == typeof t),
            r(t < 67108864),
            t < 0
              ? this.isubn(-t)
              : 0 !== this.negative
              ? 1 === this.length && (0 | this.words[0]) <= t
                ? ((this.words[0] = t - (0 | this.words[0])),
                  (this.negative = 0),
                  this)
                : ((this.negative = 0),
                  this.isubn(t),
                  (this.negative = 1),
                  this)
              : this._iaddn(t)
          );
        }),
        (i.prototype._iaddn = function (t) {
          this.words[0] += t;
          for (var e = 0; e < this.length && this.words[e] >= 67108864; e++)
            (this.words[e] -= 67108864),
              e === this.length - 1
                ? (this.words[e + 1] = 1)
                : this.words[e + 1]++;
          return (this.length = Math.max(this.length, e + 1)), this;
        }),
        (i.prototype.isubn = function (t) {
          if ((r('number' == typeof t), r(t < 67108864), t < 0))
            return this.iaddn(-t);
          if (0 !== this.negative)
            return (
              (this.negative = 0), this.iaddn(t), (this.negative = 1), this
            );
          if (((this.words[0] -= t), 1 === this.length && this.words[0] < 0))
            (this.words[0] = -this.words[0]), (this.negative = 1);
          else
            for (var e = 0; e < this.length && this.words[e] < 0; e++)
              (this.words[e] += 67108864), (this.words[e + 1] -= 1);
          return this._strip();
        }),
        (i.prototype.addn = function (t) {
          return this.clone().iaddn(t);
        }),
        (i.prototype.subn = function (t) {
          return this.clone().isubn(t);
        }),
        (i.prototype.iabs = function () {
          return (this.negative = 0), this;
        }),
        (i.prototype.abs = function () {
          return this.clone().iabs();
        }),
        (i.prototype._ishlnsubmul = function (t, e, n) {
          var i,
            o,
            s = t.length + n;
          this._expand(s);
          var a = 0;
          for (i = 0; i < t.length; i++) {
            o = (0 | this.words[i + n]) + a;
            var u = (0 | t.words[i]) * e;
            (a = ((o -= 67108863 & u) >> 26) - ((u / 67108864) | 0)),
              (this.words[i + n] = 67108863 & o);
          }
          for (; i < this.length - n; i++)
            (a = (o = (0 | this.words[i + n]) + a) >> 26),
              (this.words[i + n] = 67108863 & o);
          if (0 === a) return this._strip();
          for (r(-1 === a), a = 0, i = 0; i < this.length; i++)
            (a = (o = -(0 | this.words[i]) + a) >> 26),
              (this.words[i] = 67108863 & o);
          return (this.negative = 1), this._strip();
        }),
        (i.prototype._wordDiv = function (t, e) {
          var r = (this.length, t.length),
            n = this.clone(),
            o = t,
            s = 0 | o.words[o.length - 1];
          0 !== (r = 26 - this._countBits(s)) &&
            ((o = o.ushln(r)), n.iushln(r), (s = 0 | o.words[o.length - 1]));
          var a,
            u = n.length - o.length;
          if ('mod' !== e) {
            ((a = new i(null)).length = u + 1), (a.words = new Array(a.length));
            for (var h = 0; h < a.length; h++) a.words[h] = 0;
          }
          var l = n.clone()._ishlnsubmul(o, 1, u);
          0 === l.negative && ((n = l), a && (a.words[u] = 1));
          for (var c = u - 1; c >= 0; c--) {
            var f =
              67108864 * (0 | n.words[o.length + c]) +
              (0 | n.words[o.length + c - 1]);
            for (
              f = Math.min((f / s) | 0, 67108863), n._ishlnsubmul(o, f, c);
              0 !== n.negative;

            )
              f--,
                (n.negative = 0),
                n._ishlnsubmul(o, 1, c),
                n.isZero() || (n.negative ^= 1);
            a && (a.words[c] = f);
          }
          return (
            a && a._strip(),
            n._strip(),
            'div' !== e && 0 !== r && n.iushrn(r),
            { div: a || null, mod: n }
          );
        }),
        (i.prototype.divmod = function (t, e, n) {
          return (
            r(!t.isZero()),
            this.isZero()
              ? { div: new i(0), mod: new i(0) }
              : 0 !== this.negative && 0 === t.negative
              ? ((a = this.neg().divmod(t, e)),
                'mod' !== e && (o = a.div.neg()),
                'div' !== e &&
                  ((s = a.mod.neg()), n && 0 !== s.negative && s.iadd(t)),
                { div: o, mod: s })
              : 0 === this.negative && 0 !== t.negative
              ? ((a = this.divmod(t.neg(), e)),
                'mod' !== e && (o = a.div.neg()),
                { div: o, mod: a.mod })
              : this.negative & t.negative
              ? ((a = this.neg().divmod(t.neg(), e)),
                'div' !== e &&
                  ((s = a.mod.neg()), n && 0 !== s.negative && s.isub(t)),
                { div: a.div, mod: s })
              : t.length > this.length || this.cmp(t) < 0
              ? { div: new i(0), mod: this }
              : 1 === t.length
              ? 'div' === e
                ? { div: this.divn(t.words[0]), mod: null }
                : 'mod' === e
                ? { div: null, mod: new i(this.modrn(t.words[0])) }
                : {
                    div: this.divn(t.words[0]),
                    mod: new i(this.modrn(t.words[0])),
                  }
              : this._wordDiv(t, e)
          );
          var o, s, a;
        }),
        (i.prototype.div = function (t) {
          return this.divmod(t, 'div', !1).div;
        }),
        (i.prototype.mod = function (t) {
          return this.divmod(t, 'mod', !1).mod;
        }),
        (i.prototype.umod = function (t) {
          return this.divmod(t, 'mod', !0).mod;
        }),
        (i.prototype.divRound = function (t) {
          var e = this.divmod(t);
          if (e.mod.isZero()) return e.div;
          var r = 0 !== e.div.negative ? e.mod.isub(t) : e.mod,
            n = t.ushrn(1),
            i = t.andln(1),
            o = r.cmp(n);
          return o < 0 || (1 === i && 0 === o)
            ? e.div
            : 0 !== e.div.negative
            ? e.div.isubn(1)
            : e.div.iaddn(1);
        }),
        (i.prototype.modrn = function (t) {
          var e = t < 0;
          e && (t = -t), r(t <= 67108863);
          for (var n = (1 << 26) % t, i = 0, o = this.length - 1; o >= 0; o--)
            i = (n * i + (0 | this.words[o])) % t;
          return e ? -i : i;
        }),
        (i.prototype.modn = function (t) {
          return this.modrn(t);
        }),
        (i.prototype.idivn = function (t) {
          var e = t < 0;
          e && (t = -t), r(t <= 67108863);
          for (var n = 0, i = this.length - 1; i >= 0; i--) {
            var o = (0 | this.words[i]) + 67108864 * n;
            (this.words[i] = (o / t) | 0), (n = o % t);
          }
          return this._strip(), e ? this.ineg() : this;
        }),
        (i.prototype.divn = function (t) {
          return this.clone().idivn(t);
        }),
        (i.prototype.egcd = function (t) {
          r(0 === t.negative), r(!t.isZero());
          var e = this,
            n = t.clone();
          e = 0 !== e.negative ? e.umod(t) : e.clone();
          for (
            var o = new i(1), s = new i(0), a = new i(0), u = new i(1), h = 0;
            e.isEven() && n.isEven();

          )
            e.iushrn(1), n.iushrn(1), ++h;
          for (var l = n.clone(), c = e.clone(); !e.isZero(); ) {
            for (var f = 0, d = 1; !(e.words[0] & d) && f < 26; ++f, d <<= 1);
            if (f > 0)
              for (e.iushrn(f); f-- > 0; )
                (o.isOdd() || s.isOdd()) && (o.iadd(l), s.isub(c)),
                  o.iushrn(1),
                  s.iushrn(1);
            for (var m = 0, p = 1; !(n.words[0] & p) && m < 26; ++m, p <<= 1);
            if (m > 0)
              for (n.iushrn(m); m-- > 0; )
                (a.isOdd() || u.isOdd()) && (a.iadd(l), u.isub(c)),
                  a.iushrn(1),
                  u.iushrn(1);
            e.cmp(n) >= 0
              ? (e.isub(n), o.isub(a), s.isub(u))
              : (n.isub(e), a.isub(o), u.isub(s));
          }
          return { a: a, b: u, gcd: n.iushln(h) };
        }),
        (i.prototype._invmp = function (t) {
          r(0 === t.negative), r(!t.isZero());
          var e = this,
            n = t.clone();
          e = 0 !== e.negative ? e.umod(t) : e.clone();
          for (
            var o, s = new i(1), a = new i(0), u = n.clone();
            e.cmpn(1) > 0 && n.cmpn(1) > 0;

          ) {
            for (var h = 0, l = 1; !(e.words[0] & l) && h < 26; ++h, l <<= 1);
            if (h > 0)
              for (e.iushrn(h); h-- > 0; ) s.isOdd() && s.iadd(u), s.iushrn(1);
            for (var c = 0, f = 1; !(n.words[0] & f) && c < 26; ++c, f <<= 1);
            if (c > 0)
              for (n.iushrn(c); c-- > 0; ) a.isOdd() && a.iadd(u), a.iushrn(1);
            e.cmp(n) >= 0 ? (e.isub(n), s.isub(a)) : (n.isub(e), a.isub(s));
          }
          return (o = 0 === e.cmpn(1) ? s : a).cmpn(0) < 0 && o.iadd(t), o;
        }),
        (i.prototype.gcd = function (t) {
          if (this.isZero()) return t.abs();
          if (t.isZero()) return this.abs();
          var e = this.clone(),
            r = t.clone();
          (e.negative = 0), (r.negative = 0);
          for (var n = 0; e.isEven() && r.isEven(); n++)
            e.iushrn(1), r.iushrn(1);
          for (;;) {
            for (; e.isEven(); ) e.iushrn(1);
            for (; r.isEven(); ) r.iushrn(1);
            var i = e.cmp(r);
            if (i < 0) {
              var o = e;
              (e = r), (r = o);
            } else if (0 === i || 0 === r.cmpn(1)) break;
            e.isub(r);
          }
          return r.iushln(n);
        }),
        (i.prototype.invm = function (t) {
          return this.egcd(t).a.umod(t);
        }),
        (i.prototype.isEven = function () {
          return !(1 & this.words[0]);
        }),
        (i.prototype.isOdd = function () {
          return !(1 & ~this.words[0]);
        }),
        (i.prototype.andln = function (t) {
          return this.words[0] & t;
        }),
        (i.prototype.bincn = function (t) {
          r('number' == typeof t);
          var e = t % 26,
            n = (t - e) / 26,
            i = 1 << e;
          if (this.length <= n)
            return this._expand(n + 1), (this.words[n] |= i), this;
          for (var o = i, s = n; 0 !== o && s < this.length; s++) {
            var a = 0 | this.words[s];
            (o = (a += o) >>> 26), (a &= 67108863), (this.words[s] = a);
          }
          return 0 !== o && ((this.words[s] = o), this.length++), this;
        }),
        (i.prototype.isZero = function () {
          return 1 === this.length && 0 === this.words[0];
        }),
        (i.prototype.cmpn = function (t) {
          var e,
            n = t < 0;
          if (0 !== this.negative && !n) return -1;
          if (0 === this.negative && n) return 1;
          if ((this._strip(), this.length > 1)) e = 1;
          else {
            n && (t = -t), r(t <= 67108863, 'Number is too big');
            var i = 0 | this.words[0];
            e = i === t ? 0 : i < t ? -1 : 1;
          }
          return 0 !== this.negative ? 0 | -e : e;
        }),
        (i.prototype.cmp = function (t) {
          if (0 !== this.negative && 0 === t.negative) return -1;
          if (0 === this.negative && 0 !== t.negative) return 1;
          var e = this.ucmp(t);
          return 0 !== this.negative ? 0 | -e : e;
        }),
        (i.prototype.ucmp = function (t) {
          if (this.length > t.length) return 1;
          if (this.length < t.length) return -1;
          for (var e = 0, r = this.length - 1; r >= 0; r--) {
            var n = 0 | this.words[r],
              i = 0 | t.words[r];
            if (n !== i) {
              n < i ? (e = -1) : n > i && (e = 1);
              break;
            }
          }
          return e;
        }),
        (i.prototype.gtn = function (t) {
          return 1 === this.cmpn(t);
        }),
        (i.prototype.gt = function (t) {
          return 1 === this.cmp(t);
        }),
        (i.prototype.gten = function (t) {
          return this.cmpn(t) >= 0;
        }),
        (i.prototype.gte = function (t) {
          return this.cmp(t) >= 0;
        }),
        (i.prototype.ltn = function (t) {
          return -1 === this.cmpn(t);
        }),
        (i.prototype.lt = function (t) {
          return -1 === this.cmp(t);
        }),
        (i.prototype.lten = function (t) {
          return this.cmpn(t) <= 0;
        }),
        (i.prototype.lte = function (t) {
          return this.cmp(t) <= 0;
        }),
        (i.prototype.eqn = function (t) {
          return 0 === this.cmpn(t);
        }),
        (i.prototype.eq = function (t) {
          return 0 === this.cmp(t);
        }),
        (i.red = function (t) {
          return new _(t);
        }),
        (i.prototype.toRed = function (t) {
          return (
            r(!this.red, 'Already a number in reduction context'),
            r(0 === this.negative, 'red works only with positives'),
            t.convertTo(this)._forceRed(t)
          );
        }),
        (i.prototype.fromRed = function () {
          return (
            r(this.red, 'fromRed works only with numbers in reduction context'),
            this.red.convertFrom(this)
          );
        }),
        (i.prototype._forceRed = function (t) {
          return (this.red = t), this;
        }),
        (i.prototype.forceRed = function (t) {
          return (
            r(!this.red, 'Already a number in reduction context'),
            this._forceRed(t)
          );
        }),
        (i.prototype.redAdd = function (t) {
          return (
            r(this.red, 'redAdd works only with red numbers'),
            this.red.add(this, t)
          );
        }),
        (i.prototype.redIAdd = function (t) {
          return (
            r(this.red, 'redIAdd works only with red numbers'),
            this.red.iadd(this, t)
          );
        }),
        (i.prototype.redSub = function (t) {
          return (
            r(this.red, 'redSub works only with red numbers'),
            this.red.sub(this, t)
          );
        }),
        (i.prototype.redISub = function (t) {
          return (
            r(this.red, 'redISub works only with red numbers'),
            this.red.isub(this, t)
          );
        }),
        (i.prototype.redShl = function (t) {
          return (
            r(this.red, 'redShl works only with red numbers'),
            this.red.shl(this, t)
          );
        }),
        (i.prototype.redMul = function (t) {
          return (
            r(this.red, 'redMul works only with red numbers'),
            this.red._verify2(this, t),
            this.red.mul(this, t)
          );
        }),
        (i.prototype.redIMul = function (t) {
          return (
            r(this.red, 'redMul works only with red numbers'),
            this.red._verify2(this, t),
            this.red.imul(this, t)
          );
        }),
        (i.prototype.redSqr = function () {
          return (
            r(this.red, 'redSqr works only with red numbers'),
            this.red._verify1(this),
            this.red.sqr(this)
          );
        }),
        (i.prototype.redISqr = function () {
          return (
            r(this.red, 'redISqr works only with red numbers'),
            this.red._verify1(this),
            this.red.isqr(this)
          );
        }),
        (i.prototype.redSqrt = function () {
          return (
            r(this.red, 'redSqrt works only with red numbers'),
            this.red._verify1(this),
            this.red.sqrt(this)
          );
        }),
        (i.prototype.redInvm = function () {
          return (
            r(this.red, 'redInvm works only with red numbers'),
            this.red._verify1(this),
            this.red.invm(this)
          );
        }),
        (i.prototype.redNeg = function () {
          return (
            r(this.red, 'redNeg works only with red numbers'),
            this.red._verify1(this),
            this.red.neg(this)
          );
        }),
        (i.prototype.redPow = function (t) {
          return (
            r(this.red && !t.red, 'redPow(normalNum)'),
            this.red._verify1(this),
            this.red.pow(this, t)
          );
        });
      var v = { k256: null, p224: null, p192: null, p25519: null };
      function w(t, e) {
        (this.name = t),
          (this.p = new i(e, 16)),
          (this.n = this.p.bitLength()),
          (this.k = new i(1).iushln(this.n).isub(this.p)),
          (this.tmp = this._tmp());
      }
      function b() {
        w.call(
          this,
          'k256',
          'ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff fffffffe fffffc2f'
        );
      }
      function A() {
        w.call(
          this,
          'p224',
          'ffffffff ffffffff ffffffff ffffffff 00000000 00000000 00000001'
        );
      }
      function E() {
        w.call(
          this,
          'p192',
          'ffffffff ffffffff ffffffff fffffffe ffffffff ffffffff'
        );
      }
      function M() {
        w.call(
          this,
          '25519',
          '7fffffffffffffff ffffffffffffffff ffffffffffffffff ffffffffffffffed'
        );
      }
      function _(t) {
        if ('string' == typeof t) {
          var e = i._prime(t);
          (this.m = e.p), (this.prime = e);
        } else r(t.gtn(1), 'modulus must be greater than 1'), (this.m = t), (this.prime = null);
      }
      function k(t) {
        _.call(this, t),
          (this.shift = this.m.bitLength()),
          this.shift % 26 != 0 && (this.shift += 26 - (this.shift % 26)),
          (this.r = new i(1).iushln(this.shift)),
          (this.r2 = this.imod(this.r.sqr())),
          (this.rinv = this.r._invmp(this.m)),
          (this.minv = this.rinv.mul(this.r).isubn(1).div(this.m)),
          (this.minv = this.minv.umod(this.r)),
          (this.minv = this.r.sub(this.minv));
      }
      (w.prototype._tmp = function () {
        var t = new i(null);
        return (t.words = new Array(Math.ceil(this.n / 13))), t;
      }),
        (w.prototype.ireduce = function (t) {
          var e,
            r = t;
          do {
            this.split(r, this.tmp),
              (e = (r = (r = this.imulK(r)).iadd(this.tmp)).bitLength());
          } while (e > this.n);
          var n = e < this.n ? -1 : r.ucmp(this.p);
          return (
            0 === n
              ? ((r.words[0] = 0), (r.length = 1))
              : n > 0
              ? r.isub(this.p)
              : void 0 !== r.strip
              ? r.strip()
              : r._strip(),
            r
          );
        }),
        (w.prototype.split = function (t, e) {
          t.iushrn(this.n, 0, e);
        }),
        (w.prototype.imulK = function (t) {
          return t.imul(this.k);
        }),
        n(b, w),
        (b.prototype.split = function (t, e) {
          for (var r = 4194303, n = Math.min(t.length, 9), i = 0; i < n; i++)
            e.words[i] = t.words[i];
          if (((e.length = n), t.length <= 9))
            return (t.words[0] = 0), void (t.length = 1);
          var o = t.words[9];
          for (e.words[e.length++] = o & r, i = 10; i < t.length; i++) {
            var s = 0 | t.words[i];
            (t.words[i - 10] = ((s & r) << 4) | (o >>> 22)), (o = s);
          }
          (o >>>= 22),
            (t.words[i - 10] = o),
            0 === o && t.length > 10 ? (t.length -= 10) : (t.length -= 9);
        }),
        (b.prototype.imulK = function (t) {
          (t.words[t.length] = 0), (t.words[t.length + 1] = 0), (t.length += 2);
          for (var e = 0, r = 0; r < t.length; r++) {
            var n = 0 | t.words[r];
            (e += 977 * n),
              (t.words[r] = 67108863 & e),
              (e = 64 * n + ((e / 67108864) | 0));
          }
          return (
            0 === t.words[t.length - 1] &&
              (t.length--, 0 === t.words[t.length - 1] && t.length--),
            t
          );
        }),
        n(A, w),
        n(E, w),
        n(M, w),
        (M.prototype.imulK = function (t) {
          for (var e = 0, r = 0; r < t.length; r++) {
            var n = 19 * (0 | t.words[r]) + e,
              i = 67108863 & n;
            (n >>>= 26), (t.words[r] = i), (e = n);
          }
          return 0 !== e && (t.words[t.length++] = e), t;
        }),
        (i._prime = function (t) {
          if (v[t]) return v[t];
          var e;
          if ('k256' === t) e = new b();
          else if ('p224' === t) e = new A();
          else if ('p192' === t) e = new E();
          else {
            if ('p25519' !== t) throw new Error('Unknown prime ' + t);
            e = new M();
          }
          return (v[t] = e), e;
        }),
        (_.prototype._verify1 = function (t) {
          r(0 === t.negative, 'red works only with positives'),
            r(t.red, 'red works only with red numbers');
        }),
        (_.prototype._verify2 = function (t, e) {
          r(!(t.negative | e.negative), 'red works only with positives'),
            r(t.red && t.red === e.red, 'red works only with red numbers');
        }),
        (_.prototype.imod = function (t) {
          return this.prime
            ? this.prime.ireduce(t)._forceRed(this)
            : (h(t, t.umod(this.m)._forceRed(this)), t);
        }),
        (_.prototype.neg = function (t) {
          return t.isZero() ? t.clone() : this.m.sub(t)._forceRed(this);
        }),
        (_.prototype.add = function (t, e) {
          this._verify2(t, e);
          var r = t.add(e);
          return r.cmp(this.m) >= 0 && r.isub(this.m), r._forceRed(this);
        }),
        (_.prototype.iadd = function (t, e) {
          this._verify2(t, e);
          var r = t.iadd(e);
          return r.cmp(this.m) >= 0 && r.isub(this.m), r;
        }),
        (_.prototype.sub = function (t, e) {
          this._verify2(t, e);
          var r = t.sub(e);
          return r.cmpn(0) < 0 && r.iadd(this.m), r._forceRed(this);
        }),
        (_.prototype.isub = function (t, e) {
          this._verify2(t, e);
          var r = t.isub(e);
          return r.cmpn(0) < 0 && r.iadd(this.m), r;
        }),
        (_.prototype.shl = function (t, e) {
          return this._verify1(t), this.imod(t.ushln(e));
        }),
        (_.prototype.imul = function (t, e) {
          return this._verify2(t, e), this.imod(t.imul(e));
        }),
        (_.prototype.mul = function (t, e) {
          return this._verify2(t, e), this.imod(t.mul(e));
        }),
        (_.prototype.isqr = function (t) {
          return this.imul(t, t.clone());
        }),
        (_.prototype.sqr = function (t) {
          return this.mul(t, t);
        }),
        (_.prototype.sqrt = function (t) {
          if (t.isZero()) return t.clone();
          var e = this.m.andln(3);
          if ((r(e % 2 == 1), 3 === e)) {
            var n = this.m.add(new i(1)).iushrn(2);
            return this.pow(t, n);
          }
          for (var o = this.m.subn(1), s = 0; !o.isZero() && 0 === o.andln(1); )
            s++, o.iushrn(1);
          r(!o.isZero());
          var a = new i(1).toRed(this),
            u = a.redNeg(),
            h = this.m.subn(1).iushrn(1),
            l = this.m.bitLength();
          for (l = new i(2 * l * l).toRed(this); 0 !== this.pow(l, h).cmp(u); )
            l.redIAdd(u);
          for (
            var c = this.pow(l, o),
              f = this.pow(t, o.addn(1).iushrn(1)),
              d = this.pow(t, o),
              m = s;
            0 !== d.cmp(a);

          ) {
            for (var p = d, g = 0; 0 !== p.cmp(a); g++) p = p.redSqr();
            r(g < m);
            var y = this.pow(c, new i(1).iushln(m - g - 1));
            (f = f.redMul(y)), (c = y.redSqr()), (d = d.redMul(c)), (m = g);
          }
          return f;
        }),
        (_.prototype.invm = function (t) {
          var e = t._invmp(this.m);
          return 0 !== e.negative
            ? ((e.negative = 0), this.imod(e).redNeg())
            : this.imod(e);
        }),
        (_.prototype.pow = function (t, e) {
          if (e.isZero()) return new i(1).toRed(this);
          if (0 === e.cmpn(1)) return t.clone();
          var r = new Array(16);
          (r[0] = new i(1).toRed(this)), (r[1] = t);
          for (var n = 2; n < r.length; n++) r[n] = this.mul(r[n - 1], t);
          var o = r[0],
            s = 0,
            a = 0,
            u = e.bitLength() % 26;
          for (0 === u && (u = 26), n = e.length - 1; n >= 0; n--) {
            for (var h = e.words[n], l = u - 1; l >= 0; l--) {
              var c = (h >> l) & 1;
              o !== r[0] && (o = this.sqr(o)),
                0 !== c || 0 !== s
                  ? ((s <<= 1),
                    (s |= c),
                    (4 === ++a || (0 === n && 0 === l)) &&
                      ((o = this.mul(o, r[s])), (a = 0), (s = 0)))
                  : (a = 0);
            }
            u = 26;
          }
          return o;
        }),
        (_.prototype.convertTo = function (t) {
          var e = t.umod(this.m);
          return e === t ? e.clone() : e;
        }),
        (_.prototype.convertFrom = function (t) {
          var e = t.clone();
          return (e.red = null), e;
        }),
        (i.mont = function (t) {
          return new k(t);
        }),
        n(k, _),
        (k.prototype.convertTo = function (t) {
          return this.imod(t.ushln(this.shift));
        }),
        (k.prototype.convertFrom = function (t) {
          var e = this.imod(t.mul(this.rinv));
          return (e.red = null), e;
        }),
        (k.prototype.imul = function (t, e) {
          if (t.isZero() || e.isZero())
            return (t.words[0] = 0), (t.length = 1), t;
          var r = t.imul(e),
            n = r
              .maskn(this.shift)
              .mul(this.minv)
              .imaskn(this.shift)
              .mul(this.m),
            i = r.isub(n).iushrn(this.shift),
            o = i;
          return (
            i.cmp(this.m) >= 0
              ? (o = i.isub(this.m))
              : i.cmpn(0) < 0 && (o = i.iadd(this.m)),
            o._forceRed(this)
          );
        }),
        (k.prototype.mul = function (t, e) {
          if (t.isZero() || e.isZero()) return new i(0)._forceRed(this);
          var r = t.mul(e),
            n = r
              .maskn(this.shift)
              .mul(this.minv)
              .imaskn(this.shift)
              .mul(this.m),
            o = r.isub(n).iushrn(this.shift),
            s = o;
          return (
            o.cmp(this.m) >= 0
              ? (s = o.isub(this.m))
              : o.cmpn(0) < 0 && (s = o.iadd(this.m)),
            s._forceRed(this)
          );
        }),
        (k.prototype.invm = function (t) {
          return this.imod(t._invmp(this.m).mul(this.r2))._forceRed(this);
        });
    })(e, t);
  })(Gn);
  var qn = Gn.exports,
    Kn = {},
    Hn = {},
    Vn = Jn;
  function Jn(t, e) {
    if (!t) throw new Error(e || 'Assertion failed');
  }
  Jn.equal = function (t, e, r) {
    if (t != e) throw new Error(r || 'Assertion failed: ' + t + ' != ' + e);
  };
  var Wn = { exports: {} },
    Qn = { exports: {} };
  'function' == typeof Object.create
    ? (Qn.exports = function (t, e) {
        e &&
          ((t.super_ = e),
          (t.prototype = Object.create(e.prototype, {
            constructor: {
              value: t,
              enumerable: !1,
              writable: !0,
              configurable: !0,
            },
          })));
      })
    : (Qn.exports = function (t, e) {
        if (e) {
          t.super_ = e;
          var r = function () {};
          (r.prototype = e.prototype),
            (t.prototype = new r()),
            (t.prototype.constructor = t);
        }
      });
  try {
    var Yn = require('util');
    if ('function' != typeof Yn.inherits) throw '';
    Wn.exports = Yn.inherits;
  } catch (t) {
    Wn.exports = Qn.exports;
  }
  var Zn = Vn,
    Xn = Wn.exports;
  function $n(t, e) {
    return (
      55296 == (64512 & t.charCodeAt(e)) &&
      !(e < 0 || e + 1 >= t.length) &&
      56320 == (64512 & t.charCodeAt(e + 1))
    );
  }
  function ti(t) {
    return (
      ((t >>> 24) |
        ((t >>> 8) & 65280) |
        ((t << 8) & 16711680) |
        ((255 & t) << 24)) >>>
      0
    );
  }
  function ei(t) {
    return 1 === t.length ? '0' + t : t;
  }
  function ri(t) {
    return 7 === t.length
      ? '0' + t
      : 6 === t.length
      ? '00' + t
      : 5 === t.length
      ? '000' + t
      : 4 === t.length
      ? '0000' + t
      : 3 === t.length
      ? '00000' + t
      : 2 === t.length
      ? '000000' + t
      : 1 === t.length
      ? '0000000' + t
      : t;
  }
  (Hn.inherits = Xn),
    (Hn.toArray = function (t, e) {
      if (Array.isArray(t)) return t.slice();
      if (!t) return [];
      var r = [];
      if ('string' == typeof t)
        if (e) {
          if ('hex' === e)
            for (
              (t = t.replace(/[^a-z0-9]+/gi, '')).length % 2 != 0 &&
                (t = '0' + t),
                i = 0;
              i < t.length;
              i += 2
            )
              r.push(parseInt(t[i] + t[i + 1], 16));
        } else
          for (var n = 0, i = 0; i < t.length; i++) {
            var o = t.charCodeAt(i);
            o < 128
              ? (r[n++] = o)
              : o < 2048
              ? ((r[n++] = (o >> 6) | 192), (r[n++] = (63 & o) | 128))
              : $n(t, i)
              ? ((o = 65536 + ((1023 & o) << 10) + (1023 & t.charCodeAt(++i))),
                (r[n++] = (o >> 18) | 240),
                (r[n++] = ((o >> 12) & 63) | 128),
                (r[n++] = ((o >> 6) & 63) | 128),
                (r[n++] = (63 & o) | 128))
              : ((r[n++] = (o >> 12) | 224),
                (r[n++] = ((o >> 6) & 63) | 128),
                (r[n++] = (63 & o) | 128));
          }
      else for (i = 0; i < t.length; i++) r[i] = 0 | t[i];
      return r;
    }),
    (Hn.toHex = function (t) {
      for (var e = '', r = 0; r < t.length; r++) e += ei(t[r].toString(16));
      return e;
    }),
    (Hn.htonl = ti),
    (Hn.toHex32 = function (t, e) {
      for (var r = '', n = 0; n < t.length; n++) {
        var i = t[n];
        'little' === e && (i = ti(i)), (r += ri(i.toString(16)));
      }
      return r;
    }),
    (Hn.zero2 = ei),
    (Hn.zero8 = ri),
    (Hn.join32 = function (t, e, r, n) {
      var i = r - e;
      Zn(i % 4 == 0);
      for (var o = new Array(i / 4), s = 0, a = e; s < o.length; s++, a += 4) {
        var u;
        (u =
          'big' === n
            ? (t[a] << 24) | (t[a + 1] << 16) | (t[a + 2] << 8) | t[a + 3]
            : (t[a + 3] << 24) | (t[a + 2] << 16) | (t[a + 1] << 8) | t[a]),
          (o[s] = u >>> 0);
      }
      return o;
    }),
    (Hn.split32 = function (t, e) {
      for (
        var r = new Array(4 * t.length), n = 0, i = 0;
        n < t.length;
        n++, i += 4
      ) {
        var o = t[n];
        'big' === e
          ? ((r[i] = o >>> 24),
            (r[i + 1] = (o >>> 16) & 255),
            (r[i + 2] = (o >>> 8) & 255),
            (r[i + 3] = 255 & o))
          : ((r[i + 3] = o >>> 24),
            (r[i + 2] = (o >>> 16) & 255),
            (r[i + 1] = (o >>> 8) & 255),
            (r[i] = 255 & o));
      }
      return r;
    }),
    (Hn.rotr32 = function (t, e) {
      return (t >>> e) | (t << (32 - e));
    }),
    (Hn.rotl32 = function (t, e) {
      return (t << e) | (t >>> (32 - e));
    }),
    (Hn.sum32 = function (t, e) {
      return (t + e) >>> 0;
    }),
    (Hn.sum32_3 = function (t, e, r) {
      return (t + e + r) >>> 0;
    }),
    (Hn.sum32_4 = function (t, e, r, n) {
      return (t + e + r + n) >>> 0;
    }),
    (Hn.sum32_5 = function (t, e, r, n, i) {
      return (t + e + r + n + i) >>> 0;
    }),
    (Hn.sum64 = function (t, e, r, n) {
      var i = t[e],
        o = (n + t[e + 1]) >>> 0,
        s = (o < n ? 1 : 0) + r + i;
      (t[e] = s >>> 0), (t[e + 1] = o);
    }),
    (Hn.sum64_hi = function (t, e, r, n) {
      return (((e + n) >>> 0 < e ? 1 : 0) + t + r) >>> 0;
    }),
    (Hn.sum64_lo = function (t, e, r, n) {
      return (e + n) >>> 0;
    }),
    (Hn.sum64_4_hi = function (t, e, r, n, i, o, s, a) {
      var u = 0,
        h = e;
      return (
        (u += (h = (h + n) >>> 0) < e ? 1 : 0),
        (u += (h = (h + o) >>> 0) < o ? 1 : 0),
        (t + r + i + s + (u += (h = (h + a) >>> 0) < a ? 1 : 0)) >>> 0
      );
    }),
    (Hn.sum64_4_lo = function (t, e, r, n, i, o, s, a) {
      return (e + n + o + a) >>> 0;
    }),
    (Hn.sum64_5_hi = function (t, e, r, n, i, o, s, a, u, h) {
      var l = 0,
        c = e;
      return (
        (l += (c = (c + n) >>> 0) < e ? 1 : 0),
        (l += (c = (c + o) >>> 0) < o ? 1 : 0),
        (l += (c = (c + a) >>> 0) < a ? 1 : 0),
        (t + r + i + s + u + (l += (c = (c + h) >>> 0) < h ? 1 : 0)) >>> 0
      );
    }),
    (Hn.sum64_5_lo = function (t, e, r, n, i, o, s, a, u, h) {
      return (e + n + o + a + h) >>> 0;
    }),
    (Hn.rotr64_hi = function (t, e, r) {
      return ((e << (32 - r)) | (t >>> r)) >>> 0;
    }),
    (Hn.rotr64_lo = function (t, e, r) {
      return ((t << (32 - r)) | (e >>> r)) >>> 0;
    }),
    (Hn.shr64_hi = function (t, e, r) {
      return t >>> r;
    }),
    (Hn.shr64_lo = function (t, e, r) {
      return ((t << (32 - r)) | (e >>> r)) >>> 0;
    });
  var ni = {},
    ii = Hn,
    oi = Vn;
  function si() {
    (this.pending = null),
      (this.pendingTotal = 0),
      (this.blockSize = this.constructor.blockSize),
      (this.outSize = this.constructor.outSize),
      (this.hmacStrength = this.constructor.hmacStrength),
      (this.padLength = this.constructor.padLength / 8),
      (this.endian = 'big'),
      (this._delta8 = this.blockSize / 8),
      (this._delta32 = this.blockSize / 32);
  }
  (ni.BlockHash = si),
    (si.prototype.update = function (t, e) {
      if (
        ((t = ii.toArray(t, e)),
        this.pending
          ? (this.pending = this.pending.concat(t))
          : (this.pending = t),
        (this.pendingTotal += t.length),
        this.pending.length >= this._delta8)
      ) {
        var r = (t = this.pending).length % this._delta8;
        (this.pending = t.slice(t.length - r, t.length)),
          0 === this.pending.length && (this.pending = null),
          (t = ii.join32(t, 0, t.length - r, this.endian));
        for (var n = 0; n < t.length; n += this._delta32)
          this._update(t, n, n + this._delta32);
      }
      return this;
    }),
    (si.prototype.digest = function (t) {
      return (
        this.update(this._pad()), oi(null === this.pending), this._digest(t)
      );
    }),
    (si.prototype._pad = function () {
      var t = this.pendingTotal,
        e = this._delta8,
        r = e - ((t + this.padLength) % e),
        n = new Array(r + this.padLength);
      n[0] = 128;
      for (var i = 1; i < r; i++) n[i] = 0;
      if (((t <<= 3), 'big' === this.endian)) {
        for (var o = 8; o < this.padLength; o++) n[i++] = 0;
        (n[i++] = 0),
          (n[i++] = 0),
          (n[i++] = 0),
          (n[i++] = 0),
          (n[i++] = (t >>> 24) & 255),
          (n[i++] = (t >>> 16) & 255),
          (n[i++] = (t >>> 8) & 255),
          (n[i++] = 255 & t);
      } else
        for (
          n[i++] = 255 & t,
            n[i++] = (t >>> 8) & 255,
            n[i++] = (t >>> 16) & 255,
            n[i++] = (t >>> 24) & 255,
            n[i++] = 0,
            n[i++] = 0,
            n[i++] = 0,
            n[i++] = 0,
            o = 8;
          o < this.padLength;
          o++
        )
          n[i++] = 0;
      return n;
    });
  var ai = {},
    ui = {},
    hi = Hn.rotr32;
  function li(t, e, r) {
    return (t & e) ^ (~t & r);
  }
  function ci(t, e, r) {
    return (t & e) ^ (t & r) ^ (e & r);
  }
  function fi(t, e, r) {
    return t ^ e ^ r;
  }
  (ui.ft_1 = function (t, e, r, n) {
    return 0 === t
      ? li(e, r, n)
      : 1 === t || 3 === t
      ? fi(e, r, n)
      : 2 === t
      ? ci(e, r, n)
      : void 0;
  }),
    (ui.ch32 = li),
    (ui.maj32 = ci),
    (ui.p32 = fi),
    (ui.s0_256 = function (t) {
      return hi(t, 2) ^ hi(t, 13) ^ hi(t, 22);
    }),
    (ui.s1_256 = function (t) {
      return hi(t, 6) ^ hi(t, 11) ^ hi(t, 25);
    }),
    (ui.g0_256 = function (t) {
      return hi(t, 7) ^ hi(t, 18) ^ (t >>> 3);
    }),
    (ui.g1_256 = function (t) {
      return hi(t, 17) ^ hi(t, 19) ^ (t >>> 10);
    });
  var di = Hn,
    mi = ni,
    pi = ui,
    gi = di.rotl32,
    yi = di.sum32,
    vi = di.sum32_5,
    wi = pi.ft_1,
    bi = mi.BlockHash,
    Ai = [1518500249, 1859775393, 2400959708, 3395469782];
  function Ei() {
    if (!(this instanceof Ei)) return new Ei();
    bi.call(this),
      (this.h = [1732584193, 4023233417, 2562383102, 271733878, 3285377520]),
      (this.W = new Array(80));
  }
  di.inherits(Ei, bi);
  var Mi = Ei;
  (Ei.blockSize = 512),
    (Ei.outSize = 160),
    (Ei.hmacStrength = 80),
    (Ei.padLength = 64),
    (Ei.prototype._update = function (t, e) {
      for (var r = this.W, n = 0; n < 16; n++) r[n] = t[e + n];
      for (; n < r.length; n++)
        r[n] = gi(r[n - 3] ^ r[n - 8] ^ r[n - 14] ^ r[n - 16], 1);
      var i = this.h[0],
        o = this.h[1],
        s = this.h[2],
        a = this.h[3],
        u = this.h[4];
      for (n = 0; n < r.length; n++) {
        var h = ~~(n / 20),
          l = vi(gi(i, 5), wi(h, o, s, a), u, r[n], Ai[h]);
        (u = a), (a = s), (s = gi(o, 30)), (o = i), (i = l);
      }
      (this.h[0] = yi(this.h[0], i)),
        (this.h[1] = yi(this.h[1], o)),
        (this.h[2] = yi(this.h[2], s)),
        (this.h[3] = yi(this.h[3], a)),
        (this.h[4] = yi(this.h[4], u));
    }),
    (Ei.prototype._digest = function (t) {
      return 'hex' === t
        ? di.toHex32(this.h, 'big')
        : di.split32(this.h, 'big');
    });
  var _i = Hn,
    ki = ni,
    Si = ui,
    Ni = Vn,
    Pi = _i.sum32,
    xi = _i.sum32_4,
    Ti = _i.sum32_5,
    Ci = Si.ch32,
    Ri = Si.maj32,
    Ii = Si.s0_256,
    Bi = Si.s1_256,
    Oi = Si.g0_256,
    Fi = Si.g1_256,
    Di = ki.BlockHash,
    Li = [
      1116352408, 1899447441, 3049323471, 3921009573, 961987163, 1508970993,
      2453635748, 2870763221, 3624381080, 310598401, 607225278, 1426881987,
      1925078388, 2162078206, 2614888103, 3248222580, 3835390401, 4022224774,
      264347078, 604807628, 770255983, 1249150122, 1555081692, 1996064986,
      2554220882, 2821834349, 2952996808, 3210313671, 3336571891, 3584528711,
      113926993, 338241895, 666307205, 773529912, 1294757372, 1396182291,
      1695183700, 1986661051, 2177026350, 2456956037, 2730485921, 2820302411,
      3259730800, 3345764771, 3516065817, 3600352804, 4094571909, 275423344,
      430227734, 506948616, 659060556, 883997877, 958139571, 1322822218,
      1537002063, 1747873779, 1955562222, 2024104815, 2227730452, 2361852424,
      2428436474, 2756734187, 3204031479, 3329325298,
    ];
  function Ui() {
    if (!(this instanceof Ui)) return new Ui();
    Di.call(this),
      (this.h = [
        1779033703, 3144134277, 1013904242, 2773480762, 1359893119, 2600822924,
        528734635, 1541459225,
      ]),
      (this.k = Li),
      (this.W = new Array(64));
  }
  _i.inherits(Ui, Di);
  var ji = Ui;
  (Ui.blockSize = 512),
    (Ui.outSize = 256),
    (Ui.hmacStrength = 192),
    (Ui.padLength = 64),
    (Ui.prototype._update = function (t, e) {
      for (var r = this.W, n = 0; n < 16; n++) r[n] = t[e + n];
      for (; n < r.length; n++)
        r[n] = xi(Fi(r[n - 2]), r[n - 7], Oi(r[n - 15]), r[n - 16]);
      var i = this.h[0],
        o = this.h[1],
        s = this.h[2],
        a = this.h[3],
        u = this.h[4],
        h = this.h[5],
        l = this.h[6],
        c = this.h[7];
      for (Ni(this.k.length === r.length), n = 0; n < r.length; n++) {
        var f = Ti(c, Bi(u), Ci(u, h, l), this.k[n], r[n]),
          d = Pi(Ii(i), Ri(i, o, s));
        (c = l),
          (l = h),
          (h = u),
          (u = Pi(a, f)),
          (a = s),
          (s = o),
          (o = i),
          (i = Pi(f, d));
      }
      (this.h[0] = Pi(this.h[0], i)),
        (this.h[1] = Pi(this.h[1], o)),
        (this.h[2] = Pi(this.h[2], s)),
        (this.h[3] = Pi(this.h[3], a)),
        (this.h[4] = Pi(this.h[4], u)),
        (this.h[5] = Pi(this.h[5], h)),
        (this.h[6] = Pi(this.h[6], l)),
        (this.h[7] = Pi(this.h[7], c));
    }),
    (Ui.prototype._digest = function (t) {
      return 'hex' === t
        ? _i.toHex32(this.h, 'big')
        : _i.split32(this.h, 'big');
    });
  var zi = Hn,
    Gi = ji;
  function qi() {
    if (!(this instanceof qi)) return new qi();
    Gi.call(this),
      (this.h = [
        3238371032, 914150663, 812702999, 4144912697, 4290775857, 1750603025,
        1694076839, 3204075428,
      ]);
  }
  zi.inherits(qi, Gi);
  var Ki = qi;
  (qi.blockSize = 512),
    (qi.outSize = 224),
    (qi.hmacStrength = 192),
    (qi.padLength = 64),
    (qi.prototype._digest = function (t) {
      return 'hex' === t
        ? zi.toHex32(this.h.slice(0, 7), 'big')
        : zi.split32(this.h.slice(0, 7), 'big');
    });
  var Hi = Hn,
    Vi = ni,
    Ji = Vn,
    Wi = Hi.rotr64_hi,
    Qi = Hi.rotr64_lo,
    Yi = Hi.shr64_hi,
    Zi = Hi.shr64_lo,
    Xi = Hi.sum64,
    $i = Hi.sum64_hi,
    to = Hi.sum64_lo,
    eo = Hi.sum64_4_hi,
    ro = Hi.sum64_4_lo,
    no = Hi.sum64_5_hi,
    io = Hi.sum64_5_lo,
    oo = Vi.BlockHash,
    so = [
      1116352408, 3609767458, 1899447441, 602891725, 3049323471, 3964484399,
      3921009573, 2173295548, 961987163, 4081628472, 1508970993, 3053834265,
      2453635748, 2937671579, 2870763221, 3664609560, 3624381080, 2734883394,
      310598401, 1164996542, 607225278, 1323610764, 1426881987, 3590304994,
      1925078388, 4068182383, 2162078206, 991336113, 2614888103, 633803317,
      3248222580, 3479774868, 3835390401, 2666613458, 4022224774, 944711139,
      264347078, 2341262773, 604807628, 2007800933, 770255983, 1495990901,
      1249150122, 1856431235, 1555081692, 3175218132, 1996064986, 2198950837,
      2554220882, 3999719339, 2821834349, 766784016, 2952996808, 2566594879,
      3210313671, 3203337956, 3336571891, 1034457026, 3584528711, 2466948901,
      113926993, 3758326383, 338241895, 168717936, 666307205, 1188179964,
      773529912, 1546045734, 1294757372, 1522805485, 1396182291, 2643833823,
      1695183700, 2343527390, 1986661051, 1014477480, 2177026350, 1206759142,
      2456956037, 344077627, 2730485921, 1290863460, 2820302411, 3158454273,
      3259730800, 3505952657, 3345764771, 106217008, 3516065817, 3606008344,
      3600352804, 1432725776, 4094571909, 1467031594, 275423344, 851169720,
      430227734, 3100823752, 506948616, 1363258195, 659060556, 3750685593,
      883997877, 3785050280, 958139571, 3318307427, 1322822218, 3812723403,
      1537002063, 2003034995, 1747873779, 3602036899, 1955562222, 1575990012,
      2024104815, 1125592928, 2227730452, 2716904306, 2361852424, 442776044,
      2428436474, 593698344, 2756734187, 3733110249, 3204031479, 2999351573,
      3329325298, 3815920427, 3391569614, 3928383900, 3515267271, 566280711,
      3940187606, 3454069534, 4118630271, 4000239992, 116418474, 1914138554,
      174292421, 2731055270, 289380356, 3203993006, 460393269, 320620315,
      685471733, 587496836, 852142971, 1086792851, 1017036298, 365543100,
      1126000580, 2618297676, 1288033470, 3409855158, 1501505948, 4234509866,
      1607167915, 987167468, 1816402316, 1246189591,
    ];
  function ao() {
    if (!(this instanceof ao)) return new ao();
    oo.call(this),
      (this.h = [
        1779033703, 4089235720, 3144134277, 2227873595, 1013904242, 4271175723,
        2773480762, 1595750129, 1359893119, 2917565137, 2600822924, 725511199,
        528734635, 4215389547, 1541459225, 327033209,
      ]),
      (this.k = so),
      (this.W = new Array(160));
  }
  Hi.inherits(ao, oo);
  var uo = ao;
  function ho(t, e, r, n, i) {
    var o = (t & r) ^ (~t & i);
    return o < 0 && (o += 4294967296), o;
  }
  function lo(t, e, r, n, i, o) {
    var s = (e & n) ^ (~e & o);
    return s < 0 && (s += 4294967296), s;
  }
  function co(t, e, r, n, i) {
    var o = (t & r) ^ (t & i) ^ (r & i);
    return o < 0 && (o += 4294967296), o;
  }
  function fo(t, e, r, n, i, o) {
    var s = (e & n) ^ (e & o) ^ (n & o);
    return s < 0 && (s += 4294967296), s;
  }
  function mo(t, e) {
    var r = Wi(t, e, 28) ^ Wi(e, t, 2) ^ Wi(e, t, 7);
    return r < 0 && (r += 4294967296), r;
  }
  function po(t, e) {
    var r = Qi(t, e, 28) ^ Qi(e, t, 2) ^ Qi(e, t, 7);
    return r < 0 && (r += 4294967296), r;
  }
  function go(t, e) {
    var r = Wi(t, e, 14) ^ Wi(t, e, 18) ^ Wi(e, t, 9);
    return r < 0 && (r += 4294967296), r;
  }
  function yo(t, e) {
    var r = Qi(t, e, 14) ^ Qi(t, e, 18) ^ Qi(e, t, 9);
    return r < 0 && (r += 4294967296), r;
  }
  function vo(t, e) {
    var r = Wi(t, e, 1) ^ Wi(t, e, 8) ^ Yi(t, e, 7);
    return r < 0 && (r += 4294967296), r;
  }
  function wo(t, e) {
    var r = Qi(t, e, 1) ^ Qi(t, e, 8) ^ Zi(t, e, 7);
    return r < 0 && (r += 4294967296), r;
  }
  function bo(t, e) {
    var r = Wi(t, e, 19) ^ Wi(e, t, 29) ^ Yi(t, e, 6);
    return r < 0 && (r += 4294967296), r;
  }
  function Ao(t, e) {
    var r = Qi(t, e, 19) ^ Qi(e, t, 29) ^ Zi(t, e, 6);
    return r < 0 && (r += 4294967296), r;
  }
  (ao.blockSize = 1024),
    (ao.outSize = 512),
    (ao.hmacStrength = 192),
    (ao.padLength = 128),
    (ao.prototype._prepareBlock = function (t, e) {
      for (var r = this.W, n = 0; n < 32; n++) r[n] = t[e + n];
      for (; n < r.length; n += 2) {
        var i = bo(r[n - 4], r[n - 3]),
          o = Ao(r[n - 4], r[n - 3]),
          s = r[n - 14],
          a = r[n - 13],
          u = vo(r[n - 30], r[n - 29]),
          h = wo(r[n - 30], r[n - 29]),
          l = r[n - 32],
          c = r[n - 31];
        (r[n] = eo(i, o, s, a, u, h, l, c)),
          (r[n + 1] = ro(i, o, s, a, u, h, l, c));
      }
    }),
    (ao.prototype._update = function (t, e) {
      this._prepareBlock(t, e);
      var r = this.W,
        n = this.h[0],
        i = this.h[1],
        o = this.h[2],
        s = this.h[3],
        a = this.h[4],
        u = this.h[5],
        h = this.h[6],
        l = this.h[7],
        c = this.h[8],
        f = this.h[9],
        d = this.h[10],
        m = this.h[11],
        p = this.h[12],
        g = this.h[13],
        y = this.h[14],
        v = this.h[15];
      Ji(this.k.length === r.length);
      for (var w = 0; w < r.length; w += 2) {
        var b = y,
          A = v,
          E = go(c, f),
          M = yo(c, f),
          _ = ho(c, f, d, m, p),
          k = lo(c, f, d, m, p, g),
          S = this.k[w],
          N = this.k[w + 1],
          P = r[w],
          x = r[w + 1],
          T = no(b, A, E, M, _, k, S, N, P, x),
          C = io(b, A, E, M, _, k, S, N, P, x);
        (b = mo(n, i)),
          (A = po(n, i)),
          (E = co(n, i, o, s, a)),
          (M = fo(n, i, o, s, a, u));
        var R = $i(b, A, E, M),
          I = to(b, A, E, M);
        (y = p),
          (v = g),
          (p = d),
          (g = m),
          (d = c),
          (m = f),
          (c = $i(h, l, T, C)),
          (f = to(l, l, T, C)),
          (h = a),
          (l = u),
          (a = o),
          (u = s),
          (o = n),
          (s = i),
          (n = $i(T, C, R, I)),
          (i = to(T, C, R, I));
      }
      Xi(this.h, 0, n, i),
        Xi(this.h, 2, o, s),
        Xi(this.h, 4, a, u),
        Xi(this.h, 6, h, l),
        Xi(this.h, 8, c, f),
        Xi(this.h, 10, d, m),
        Xi(this.h, 12, p, g),
        Xi(this.h, 14, y, v);
    }),
    (ao.prototype._digest = function (t) {
      return 'hex' === t
        ? Hi.toHex32(this.h, 'big')
        : Hi.split32(this.h, 'big');
    });
  var Eo = Hn,
    Mo = uo;
  function _o() {
    if (!(this instanceof _o)) return new _o();
    Mo.call(this),
      (this.h = [
        3418070365, 3238371032, 1654270250, 914150663, 2438529370, 812702999,
        355462360, 4144912697, 1731405415, 4290775857, 2394180231, 1750603025,
        3675008525, 1694076839, 1203062813, 3204075428,
      ]);
  }
  Eo.inherits(_o, Mo);
  var ko = _o;
  (_o.blockSize = 1024),
    (_o.outSize = 384),
    (_o.hmacStrength = 192),
    (_o.padLength = 128),
    (_o.prototype._digest = function (t) {
      return 'hex' === t
        ? Eo.toHex32(this.h.slice(0, 12), 'big')
        : Eo.split32(this.h.slice(0, 12), 'big');
    }),
    (ai.sha1 = Mi),
    (ai.sha224 = Ki),
    (ai.sha256 = ji),
    (ai.sha384 = ko),
    (ai.sha512 = uo);
  var So = {},
    No = Hn,
    Po = ni,
    xo = No.rotl32,
    To = No.sum32,
    Co = No.sum32_3,
    Ro = No.sum32_4,
    Io = Po.BlockHash;
  function Bo() {
    if (!(this instanceof Bo)) return new Bo();
    Io.call(this),
      (this.h = [1732584193, 4023233417, 2562383102, 271733878, 3285377520]),
      (this.endian = 'little');
  }
  function Oo(t, e, r, n) {
    return t <= 15
      ? e ^ r ^ n
      : t <= 31
      ? (e & r) | (~e & n)
      : t <= 47
      ? (e | ~r) ^ n
      : t <= 63
      ? (e & n) | (r & ~n)
      : e ^ (r | ~n);
  }
  function Fo(t) {
    return t <= 15
      ? 0
      : t <= 31
      ? 1518500249
      : t <= 47
      ? 1859775393
      : t <= 63
      ? 2400959708
      : 2840853838;
  }
  function Do(t) {
    return t <= 15
      ? 1352829926
      : t <= 31
      ? 1548603684
      : t <= 47
      ? 1836072691
      : t <= 63
      ? 2053994217
      : 0;
  }
  No.inherits(Bo, Io),
    (So.ripemd160 = Bo),
    (Bo.blockSize = 512),
    (Bo.outSize = 160),
    (Bo.hmacStrength = 192),
    (Bo.padLength = 64),
    (Bo.prototype._update = function (t, e) {
      for (
        var r = this.h[0],
          n = this.h[1],
          i = this.h[2],
          o = this.h[3],
          s = this.h[4],
          a = r,
          u = n,
          h = i,
          l = o,
          c = s,
          f = 0;
        f < 80;
        f++
      ) {
        var d = To(xo(Ro(r, Oo(f, n, i, o), t[Lo[f] + e], Fo(f)), jo[f]), s);
        (r = s),
          (s = o),
          (o = xo(i, 10)),
          (i = n),
          (n = d),
          (d = To(
            xo(Ro(a, Oo(79 - f, u, h, l), t[Uo[f] + e], Do(f)), zo[f]),
            c
          )),
          (a = c),
          (c = l),
          (l = xo(h, 10)),
          (h = u),
          (u = d);
      }
      (d = Co(this.h[1], i, l)),
        (this.h[1] = Co(this.h[2], o, c)),
        (this.h[2] = Co(this.h[3], s, a)),
        (this.h[3] = Co(this.h[4], r, u)),
        (this.h[4] = Co(this.h[0], n, h)),
        (this.h[0] = d);
    }),
    (Bo.prototype._digest = function (t) {
      return 'hex' === t
        ? No.toHex32(this.h, 'little')
        : No.split32(this.h, 'little');
    });
  var Lo = [
      0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 7, 4, 13, 1, 10, 6,
      15, 3, 12, 0, 9, 5, 2, 14, 11, 8, 3, 10, 14, 4, 9, 15, 8, 1, 2, 7, 0, 6,
      13, 11, 5, 12, 1, 9, 11, 10, 0, 8, 12, 4, 13, 3, 7, 15, 14, 5, 6, 2, 4, 0,
      5, 9, 7, 12, 2, 10, 14, 1, 3, 8, 11, 6, 15, 13,
    ],
    Uo = [
      5, 14, 7, 0, 9, 2, 11, 4, 13, 6, 15, 8, 1, 10, 3, 12, 6, 11, 3, 7, 0, 13,
      5, 10, 14, 15, 8, 12, 4, 9, 1, 2, 15, 5, 1, 3, 7, 14, 6, 9, 11, 8, 12, 2,
      10, 0, 4, 13, 8, 6, 4, 1, 3, 11, 15, 0, 5, 12, 2, 13, 9, 7, 10, 14, 12,
      15, 10, 4, 1, 5, 8, 7, 6, 2, 13, 14, 0, 3, 9, 11,
    ],
    jo = [
      11, 14, 15, 12, 5, 8, 7, 9, 11, 13, 14, 15, 6, 7, 9, 8, 7, 6, 8, 13, 11,
      9, 7, 15, 7, 12, 15, 9, 11, 7, 13, 12, 11, 13, 6, 7, 14, 9, 13, 15, 14, 8,
      13, 6, 5, 12, 7, 5, 11, 12, 14, 15, 14, 15, 9, 8, 9, 14, 5, 6, 8, 6, 5,
      12, 9, 15, 5, 11, 6, 8, 13, 12, 5, 12, 13, 14, 11, 8, 5, 6,
    ],
    zo = [
      8, 9, 9, 11, 13, 15, 15, 5, 7, 7, 8, 11, 14, 14, 12, 6, 9, 13, 15, 7, 12,
      8, 9, 11, 7, 7, 12, 7, 6, 15, 13, 11, 9, 7, 15, 11, 8, 6, 6, 14, 12, 13,
      5, 14, 13, 13, 7, 5, 15, 5, 8, 11, 14, 14, 6, 14, 6, 9, 12, 9, 12, 5, 15,
      8, 8, 5, 12, 9, 12, 5, 14, 6, 8, 13, 6, 5, 15, 13, 11, 11,
    ],
    Go = Hn,
    qo = Vn;
  function Ko(t, e, r) {
    if (!(this instanceof Ko)) return new Ko(t, e, r);
    (this.Hash = t),
      (this.blockSize = t.blockSize / 8),
      (this.outSize = t.outSize / 8),
      (this.inner = null),
      (this.outer = null),
      this._init(Go.toArray(e, r));
  }
  var Ho = Ko;
  function Vo(t, e, r) {
    return (
      (r = {
        path: e,
        exports: {},
        require: function (t, e) {
          return (function () {
            throw new Error(
              'Dynamic requires are not currently supported by @rollup/plugin-commonjs'
            );
          })(null == e && r.path);
        },
      }),
      t(r, r.exports),
      r.exports
    );
  }
  (Ko.prototype._init = function (t) {
    t.length > this.blockSize && (t = new this.Hash().update(t).digest()),
      qo(t.length <= this.blockSize);
    for (var e = t.length; e < this.blockSize; e++) t.push(0);
    for (e = 0; e < t.length; e++) t[e] ^= 54;
    for (this.inner = new this.Hash().update(t), e = 0; e < t.length; e++)
      t[e] ^= 106;
    this.outer = new this.Hash().update(t);
  }),
    (Ko.prototype.update = function (t, e) {
      return this.inner.update(t, e), this;
    }),
    (Ko.prototype.digest = function (t) {
      return this.outer.update(this.inner.digest()), this.outer.digest(t);
    }),
    (function (t) {
      var e = t;
      (e.utils = Hn),
        (e.common = ni),
        (e.sha = ai),
        (e.ripemd = So),
        (e.hmac = Ho),
        (e.sha1 = e.sha.sha1),
        (e.sha256 = e.sha.sha256),
        (e.sha224 = e.sha.sha224),
        (e.sha384 = e.sha.sha384),
        (e.sha512 = e.sha.sha512),
        (e.ripemd160 = e.ripemd.ripemd160);
    })(Kn);
  var Jo = Wo;
  function Wo(t, e) {
    if (!t) throw new Error(e || 'Assertion failed');
  }
  Wo.equal = function (t, e, r) {
    if (t != e) throw new Error(r || 'Assertion failed: ' + t + ' != ' + e);
  };
  var Qo = Vo(function (t, e) {
      var r = e;
      function n(t) {
        return 1 === t.length ? '0' + t : t;
      }
      function i(t) {
        for (var e = '', r = 0; r < t.length; r++) e += n(t[r].toString(16));
        return e;
      }
      (r.toArray = function (t, e) {
        if (Array.isArray(t)) return t.slice();
        if (!t) return [];
        var r = [];
        if ('string' != typeof t) {
          for (var n = 0; n < t.length; n++) r[n] = 0 | t[n];
          return r;
        }
        if ('hex' === e) {
          (t = t.replace(/[^a-z0-9]+/gi, '')).length % 2 != 0 && (t = '0' + t);
          for (n = 0; n < t.length; n += 2)
            r.push(parseInt(t[n] + t[n + 1], 16));
        } else
          for (n = 0; n < t.length; n++) {
            var i = t.charCodeAt(n),
              o = i >> 8,
              s = 255 & i;
            o ? r.push(o, s) : r.push(s);
          }
        return r;
      }),
        (r.zero2 = n),
        (r.toHex = i),
        (r.encode = function (t, e) {
          return 'hex' === e ? i(t) : t;
        });
    }),
    Yo = Vo(function (t, e) {
      var r = e;
      (r.assert = Jo),
        (r.toArray = Qo.toArray),
        (r.zero2 = Qo.zero2),
        (r.toHex = Qo.toHex),
        (r.encode = Qo.encode),
        (r.getNAF = function (t, e, r) {
          var n = new Array(Math.max(t.bitLength(), r) + 1);
          n.fill(0);
          for (var i = 1 << (e + 1), o = t.clone(), s = 0; s < n.length; s++) {
            var a,
              u = o.andln(i - 1);
            o.isOdd()
              ? ((a = u > (i >> 1) - 1 ? (i >> 1) - u : u), o.isubn(a))
              : (a = 0),
              (n[s] = a),
              o.iushrn(1);
          }
          return n;
        }),
        (r.getJSF = function (t, e) {
          var r = [[], []];
          (t = t.clone()), (e = e.clone());
          for (var n, i = 0, o = 0; t.cmpn(-i) > 0 || e.cmpn(-o) > 0; ) {
            var s,
              a,
              u = (t.andln(3) + i) & 3,
              h = (e.andln(3) + o) & 3;
            3 === u && (u = -1),
              3 === h && (h = -1),
              (s =
                1 & u
                  ? (3 !== (n = (t.andln(7) + i) & 7) && 5 !== n) || 2 !== h
                    ? u
                    : -u
                  : 0),
              r[0].push(s),
              (a =
                1 & h
                  ? (3 !== (n = (e.andln(7) + o) & 7) && 5 !== n) || 2 !== u
                    ? h
                    : -h
                  : 0),
              r[1].push(a),
              2 * i === s + 1 && (i = 1 - i),
              2 * o === a + 1 && (o = 1 - o),
              t.iushrn(1),
              e.iushrn(1);
          }
          return r;
        }),
        (r.cachedProperty = function (t, e, r) {
          var n = '_' + e;
          t.prototype[e] = function () {
            return void 0 !== this[n] ? this[n] : (this[n] = r.call(this));
          };
        }),
        (r.parseBytes = function (t) {
          return 'string' == typeof t ? r.toArray(t, 'hex') : t;
        }),
        (r.intFromLE = function (t) {
          return new qn(t, 'hex', 'le');
        });
    }),
    Zo = Yo.getNAF,
    Xo = Yo.getJSF,
    $o = Yo.assert;
  function ts(t, e) {
    (this.type = t),
      (this.p = new qn(e.p, 16)),
      (this.red = e.prime ? qn.red(e.prime) : qn.mont(this.p)),
      (this.zero = new qn(0).toRed(this.red)),
      (this.one = new qn(1).toRed(this.red)),
      (this.two = new qn(2).toRed(this.red)),
      (this.n = e.n && new qn(e.n, 16)),
      (this.g = e.g && this.pointFromJSON(e.g, e.gRed)),
      (this._wnafT1 = new Array(4)),
      (this._wnafT2 = new Array(4)),
      (this._wnafT3 = new Array(4)),
      (this._wnafT4 = new Array(4)),
      (this._bitLength = this.n ? this.n.bitLength() : 0);
    var r = this.n && this.p.div(this.n);
    !r || r.cmpn(100) > 0
      ? (this.redN = null)
      : ((this._maxwellTrick = !0), (this.redN = this.n.toRed(this.red)));
  }
  var es = ts;
  function rs(t, e) {
    (this.curve = t), (this.type = e), (this.precomputed = null);
  }
  (ts.prototype.point = function () {
    throw new Error('Not implemented');
  }),
    (ts.prototype.validate = function () {
      throw new Error('Not implemented');
    }),
    (ts.prototype._fixedNafMul = function (t, e) {
      $o(t.precomputed);
      var r = t._getDoubles(),
        n = Zo(e, 1, this._bitLength),
        i = (1 << (r.step + 1)) - (r.step % 2 == 0 ? 2 : 1);
      i /= 3;
      var o,
        s,
        a = [];
      for (o = 0; o < n.length; o += r.step) {
        s = 0;
        for (var u = o + r.step - 1; u >= o; u--) s = (s << 1) + n[u];
        a.push(s);
      }
      for (
        var h = this.jpoint(null, null, null),
          l = this.jpoint(null, null, null),
          c = i;
        c > 0;
        c--
      ) {
        for (o = 0; o < a.length; o++)
          (s = a[o]) === c
            ? (l = l.mixedAdd(r.points[o]))
            : s === -c && (l = l.mixedAdd(r.points[o].neg()));
        h = h.add(l);
      }
      return h.toP();
    }),
    (ts.prototype._wnafMul = function (t, e) {
      var r = 4,
        n = t._getNAFPoints(r);
      r = n.wnd;
      for (
        var i = n.points,
          o = Zo(e, r, this._bitLength),
          s = this.jpoint(null, null, null),
          a = o.length - 1;
        a >= 0;
        a--
      ) {
        for (var u = 0; a >= 0 && 0 === o[a]; a--) u++;
        if ((a >= 0 && u++, (s = s.dblp(u)), a < 0)) break;
        var h = o[a];
        $o(0 !== h),
          (s =
            'affine' === t.type
              ? h > 0
                ? s.mixedAdd(i[(h - 1) >> 1])
                : s.mixedAdd(i[(-h - 1) >> 1].neg())
              : h > 0
              ? s.add(i[(h - 1) >> 1])
              : s.add(i[(-h - 1) >> 1].neg()));
      }
      return 'affine' === t.type ? s.toP() : s;
    }),
    (ts.prototype._wnafMulAdd = function (t, e, r, n, i) {
      var o,
        s,
        a,
        u = this._wnafT1,
        h = this._wnafT2,
        l = this._wnafT3,
        c = 0;
      for (o = 0; o < n; o++) {
        var f = (a = e[o])._getNAFPoints(t);
        (u[o] = f.wnd), (h[o] = f.points);
      }
      for (o = n - 1; o >= 1; o -= 2) {
        var d = o - 1,
          m = o;
        if (1 === u[d] && 1 === u[m]) {
          var p = [e[d], null, null, e[m]];
          0 === e[d].y.cmp(e[m].y)
            ? ((p[1] = e[d].add(e[m])),
              (p[2] = e[d].toJ().mixedAdd(e[m].neg())))
            : 0 === e[d].y.cmp(e[m].y.redNeg())
            ? ((p[1] = e[d].toJ().mixedAdd(e[m])),
              (p[2] = e[d].add(e[m].neg())))
            : ((p[1] = e[d].toJ().mixedAdd(e[m])),
              (p[2] = e[d].toJ().mixedAdd(e[m].neg())));
          var g = [-3, -1, -5, -7, 0, 7, 5, 1, 3],
            y = Xo(r[d], r[m]);
          for (
            c = Math.max(y[0].length, c),
              l[d] = new Array(c),
              l[m] = new Array(c),
              s = 0;
            s < c;
            s++
          ) {
            var v = 0 | y[0][s],
              w = 0 | y[1][s];
            (l[d][s] = g[3 * (v + 1) + (w + 1)]), (l[m][s] = 0), (h[d] = p);
          }
        } else
          (l[d] = Zo(r[d], u[d], this._bitLength)),
            (l[m] = Zo(r[m], u[m], this._bitLength)),
            (c = Math.max(l[d].length, c)),
            (c = Math.max(l[m].length, c));
      }
      var b = this.jpoint(null, null, null),
        A = this._wnafT4;
      for (o = c; o >= 0; o--) {
        for (var E = 0; o >= 0; ) {
          var M = !0;
          for (s = 0; s < n; s++) (A[s] = 0 | l[s][o]), 0 !== A[s] && (M = !1);
          if (!M) break;
          E++, o--;
        }
        if ((o >= 0 && E++, (b = b.dblp(E)), o < 0)) break;
        for (s = 0; s < n; s++) {
          var _ = A[s];
          0 !== _ &&
            (_ > 0
              ? (a = h[s][(_ - 1) >> 1])
              : _ < 0 && (a = h[s][(-_ - 1) >> 1].neg()),
            (b = 'affine' === a.type ? b.mixedAdd(a) : b.add(a)));
        }
      }
      for (o = 0; o < n; o++) h[o] = null;
      return i ? b : b.toP();
    }),
    (ts.BasePoint = rs),
    (rs.prototype.eq = function () {
      throw new Error('Not implemented');
    }),
    (rs.prototype.validate = function () {
      return this.curve.validate(this);
    }),
    (ts.prototype.decodePoint = function (t, e) {
      t = Yo.toArray(t, e);
      var r = this.p.byteLength();
      if ((4 === t[0] || 6 === t[0] || 7 === t[0]) && t.length - 1 == 2 * r)
        return (
          6 === t[0]
            ? $o(t[t.length - 1] % 2 == 0)
            : 7 === t[0] && $o(t[t.length - 1] % 2 == 1),
          this.point(t.slice(1, 1 + r), t.slice(1 + r, 1 + 2 * r))
        );
      if ((2 === t[0] || 3 === t[0]) && t.length - 1 === r)
        return this.pointFromX(t.slice(1, 1 + r), 3 === t[0]);
      throw new Error('Unknown point format');
    }),
    (rs.prototype.encodeCompressed = function (t) {
      return this.encode(t, !0);
    }),
    (rs.prototype._encode = function (t) {
      var e = this.curve.p.byteLength(),
        r = this.getX().toArray('be', e);
      return t
        ? [this.getY().isEven() ? 2 : 3].concat(r)
        : [4].concat(r, this.getY().toArray('be', e));
    }),
    (rs.prototype.encode = function (t, e) {
      return Yo.encode(this._encode(e), t);
    }),
    (rs.prototype.precompute = function (t) {
      if (this.precomputed) return this;
      var e = { doubles: null, naf: null, beta: null };
      return (
        (e.naf = this._getNAFPoints(8)),
        (e.doubles = this._getDoubles(4, t)),
        (e.beta = this._getBeta()),
        (this.precomputed = e),
        this
      );
    }),
    (rs.prototype._hasDoubles = function (t) {
      if (!this.precomputed) return !1;
      var e = this.precomputed.doubles;
      return !!e && e.points.length >= Math.ceil((t.bitLength() + 1) / e.step);
    }),
    (rs.prototype._getDoubles = function (t, e) {
      if (this.precomputed && this.precomputed.doubles)
        return this.precomputed.doubles;
      for (var r = [this], n = this, i = 0; i < e; i += t) {
        for (var o = 0; o < t; o++) n = n.dbl();
        r.push(n);
      }
      return { step: t, points: r };
    }),
    (rs.prototype._getNAFPoints = function (t) {
      if (this.precomputed && this.precomputed.naf) return this.precomputed.naf;
      for (
        var e = [this],
          r = (1 << t) - 1,
          n = 1 === r ? null : this.dbl(),
          i = 1;
        i < r;
        i++
      )
        e[i] = e[i - 1].add(n);
      return { wnd: t, points: e };
    }),
    (rs.prototype._getBeta = function () {
      return null;
    }),
    (rs.prototype.dblp = function (t) {
      for (var e = this, r = 0; r < t; r++) e = e.dbl();
      return e;
    });
  var ns = Vo(function (t) {
      'function' == typeof Object.create
        ? (t.exports = function (t, e) {
            e &&
              ((t.super_ = e),
              (t.prototype = Object.create(e.prototype, {
                constructor: {
                  value: t,
                  enumerable: !1,
                  writable: !0,
                  configurable: !0,
                },
              })));
          })
        : (t.exports = function (t, e) {
            if (e) {
              t.super_ = e;
              var r = function () {};
              (r.prototype = e.prototype),
                (t.prototype = new r()),
                (t.prototype.constructor = t);
            }
          });
    }),
    is = Yo.assert;
  function os(t) {
    es.call(this, 'short', t),
      (this.a = new qn(t.a, 16).toRed(this.red)),
      (this.b = new qn(t.b, 16).toRed(this.red)),
      (this.tinv = this.two.redInvm()),
      (this.zeroA = 0 === this.a.fromRed().cmpn(0)),
      (this.threeA = 0 === this.a.fromRed().sub(this.p).cmpn(-3)),
      (this.endo = this._getEndomorphism(t)),
      (this._endoWnafT1 = new Array(4)),
      (this._endoWnafT2 = new Array(4));
  }
  ns(os, es);
  var ss = os;
  function as(t, e, r, n) {
    es.BasePoint.call(this, t, 'affine'),
      null === e && null === r
        ? ((this.x = null), (this.y = null), (this.inf = !0))
        : ((this.x = new qn(e, 16)),
          (this.y = new qn(r, 16)),
          n &&
            (this.x.forceRed(this.curve.red), this.y.forceRed(this.curve.red)),
          this.x.red || (this.x = this.x.toRed(this.curve.red)),
          this.y.red || (this.y = this.y.toRed(this.curve.red)),
          (this.inf = !1));
  }
  function us(t, e, r, n) {
    es.BasePoint.call(this, t, 'jacobian'),
      null === e && null === r && null === n
        ? ((this.x = this.curve.one),
          (this.y = this.curve.one),
          (this.z = new qn(0)))
        : ((this.x = new qn(e, 16)),
          (this.y = new qn(r, 16)),
          (this.z = new qn(n, 16))),
      this.x.red || (this.x = this.x.toRed(this.curve.red)),
      this.y.red || (this.y = this.y.toRed(this.curve.red)),
      this.z.red || (this.z = this.z.toRed(this.curve.red)),
      (this.zOne = this.z === this.curve.one);
  }
  (os.prototype._getEndomorphism = function (t) {
    if (this.zeroA && this.g && this.n && 1 === this.p.modn(3)) {
      var e, r;
      if (t.beta) e = new qn(t.beta, 16).toRed(this.red);
      else {
        var n = this._getEndoRoots(this.p);
        e = (e = n[0].cmp(n[1]) < 0 ? n[0] : n[1]).toRed(this.red);
      }
      if (t.lambda) r = new qn(t.lambda, 16);
      else {
        var i = this._getEndoRoots(this.n);
        0 === this.g.mul(i[0]).x.cmp(this.g.x.redMul(e))
          ? (r = i[0])
          : ((r = i[1]), is(0 === this.g.mul(r).x.cmp(this.g.x.redMul(e))));
      }
      return {
        beta: e,
        lambda: r,
        basis: t.basis
          ? t.basis.map(function (t) {
              return { a: new qn(t.a, 16), b: new qn(t.b, 16) };
            })
          : this._getEndoBasis(r),
      };
    }
  }),
    (os.prototype._getEndoRoots = function (t) {
      var e = t === this.p ? this.red : qn.mont(t),
        r = new qn(2).toRed(e).redInvm(),
        n = r.redNeg(),
        i = new qn(3).toRed(e).redNeg().redSqrt().redMul(r);
      return [n.redAdd(i).fromRed(), n.redSub(i).fromRed()];
    }),
    (os.prototype._getEndoBasis = function (t) {
      for (
        var e,
          r,
          n,
          i,
          o,
          s,
          a,
          u,
          h,
          l = this.n.ushrn(Math.floor(this.n.bitLength() / 2)),
          c = t,
          f = this.n.clone(),
          d = new qn(1),
          m = new qn(0),
          p = new qn(0),
          g = new qn(1),
          y = 0;
        0 !== c.cmpn(0);

      ) {
        var v = f.div(c);
        (u = f.sub(v.mul(c))), (h = p.sub(v.mul(d)));
        var w = g.sub(v.mul(m));
        if (!n && u.cmp(l) < 0) (e = a.neg()), (r = d), (n = u.neg()), (i = h);
        else if (n && 2 == ++y) break;
        (a = u), (f = c), (c = u), (p = d), (d = h), (g = m), (m = w);
      }
      (o = u.neg()), (s = h);
      var b = n.sqr().add(i.sqr());
      return (
        o.sqr().add(s.sqr()).cmp(b) >= 0 && ((o = e), (s = r)),
        n.negative && ((n = n.neg()), (i = i.neg())),
        o.negative && ((o = o.neg()), (s = s.neg())),
        [
          { a: n, b: i },
          { a: o, b: s },
        ]
      );
    }),
    (os.prototype._endoSplit = function (t) {
      var e = this.endo.basis,
        r = e[0],
        n = e[1],
        i = n.b.mul(t).divRound(this.n),
        o = r.b.neg().mul(t).divRound(this.n),
        s = i.mul(r.a),
        a = o.mul(n.a),
        u = i.mul(r.b),
        h = o.mul(n.b);
      return { k1: t.sub(s).sub(a), k2: u.add(h).neg() };
    }),
    (os.prototype.pointFromX = function (t, e) {
      (t = new qn(t, 16)).red || (t = t.toRed(this.red));
      var r = t.redSqr().redMul(t).redIAdd(t.redMul(this.a)).redIAdd(this.b),
        n = r.redSqrt();
      if (0 !== n.redSqr().redSub(r).cmp(this.zero))
        throw new Error('invalid point');
      var i = n.fromRed().isOdd();
      return ((e && !i) || (!e && i)) && (n = n.redNeg()), this.point(t, n);
    }),
    (os.prototype.validate = function (t) {
      if (t.inf) return !0;
      var e = t.x,
        r = t.y,
        n = this.a.redMul(e),
        i = e.redSqr().redMul(e).redIAdd(n).redIAdd(this.b);
      return 0 === r.redSqr().redISub(i).cmpn(0);
    }),
    (os.prototype._endoWnafMulAdd = function (t, e, r) {
      for (
        var n = this._endoWnafT1, i = this._endoWnafT2, o = 0;
        o < t.length;
        o++
      ) {
        var s = this._endoSplit(e[o]),
          a = t[o],
          u = a._getBeta();
        s.k1.negative && (s.k1.ineg(), (a = a.neg(!0))),
          s.k2.negative && (s.k2.ineg(), (u = u.neg(!0))),
          (n[2 * o] = a),
          (n[2 * o + 1] = u),
          (i[2 * o] = s.k1),
          (i[2 * o + 1] = s.k2);
      }
      for (var h = this._wnafMulAdd(1, n, i, 2 * o, r), l = 0; l < 2 * o; l++)
        (n[l] = null), (i[l] = null);
      return h;
    }),
    ns(as, es.BasePoint),
    (os.prototype.point = function (t, e, r) {
      return new as(this, t, e, r);
    }),
    (os.prototype.pointFromJSON = function (t, e) {
      return as.fromJSON(this, t, e);
    }),
    (as.prototype._getBeta = function () {
      if (this.curve.endo) {
        var t = this.precomputed;
        if (t && t.beta) return t.beta;
        var e = this.curve.point(this.x.redMul(this.curve.endo.beta), this.y);
        if (t) {
          var r = this.curve,
            n = function (t) {
              return r.point(t.x.redMul(r.endo.beta), t.y);
            };
          (t.beta = e),
            (e.precomputed = {
              beta: null,
              naf: t.naf && { wnd: t.naf.wnd, points: t.naf.points.map(n) },
              doubles: t.doubles && {
                step: t.doubles.step,
                points: t.doubles.points.map(n),
              },
            });
        }
        return e;
      }
    }),
    (as.prototype.toJSON = function () {
      return this.precomputed
        ? [
            this.x,
            this.y,
            this.precomputed && {
              doubles: this.precomputed.doubles && {
                step: this.precomputed.doubles.step,
                points: this.precomputed.doubles.points.slice(1),
              },
              naf: this.precomputed.naf && {
                wnd: this.precomputed.naf.wnd,
                points: this.precomputed.naf.points.slice(1),
              },
            },
          ]
        : [this.x, this.y];
    }),
    (as.fromJSON = function (t, e, r) {
      'string' == typeof e && (e = JSON.parse(e));
      var n = t.point(e[0], e[1], r);
      if (!e[2]) return n;
      function i(e) {
        return t.point(e[0], e[1], r);
      }
      var o = e[2];
      return (
        (n.precomputed = {
          beta: null,
          doubles: o.doubles && {
            step: o.doubles.step,
            points: [n].concat(o.doubles.points.map(i)),
          },
          naf: o.naf && {
            wnd: o.naf.wnd,
            points: [n].concat(o.naf.points.map(i)),
          },
        }),
        n
      );
    }),
    (as.prototype.inspect = function () {
      return this.isInfinity()
        ? '<EC Point Infinity>'
        : '<EC Point x: ' +
            this.x.fromRed().toString(16, 2) +
            ' y: ' +
            this.y.fromRed().toString(16, 2) +
            '>';
    }),
    (as.prototype.isInfinity = function () {
      return this.inf;
    }),
    (as.prototype.add = function (t) {
      if (this.inf) return t;
      if (t.inf) return this;
      if (this.eq(t)) return this.dbl();
      if (this.neg().eq(t)) return this.curve.point(null, null);
      if (0 === this.x.cmp(t.x)) return this.curve.point(null, null);
      var e = this.y.redSub(t.y);
      0 !== e.cmpn(0) && (e = e.redMul(this.x.redSub(t.x).redInvm()));
      var r = e.redSqr().redISub(this.x).redISub(t.x),
        n = e.redMul(this.x.redSub(r)).redISub(this.y);
      return this.curve.point(r, n);
    }),
    (as.prototype.dbl = function () {
      if (this.inf) return this;
      var t = this.y.redAdd(this.y);
      if (0 === t.cmpn(0)) return this.curve.point(null, null);
      var e = this.curve.a,
        r = this.x.redSqr(),
        n = t.redInvm(),
        i = r.redAdd(r).redIAdd(r).redIAdd(e).redMul(n),
        o = i.redSqr().redISub(this.x.redAdd(this.x)),
        s = i.redMul(this.x.redSub(o)).redISub(this.y);
      return this.curve.point(o, s);
    }),
    (as.prototype.getX = function () {
      return this.x.fromRed();
    }),
    (as.prototype.getY = function () {
      return this.y.fromRed();
    }),
    (as.prototype.mul = function (t) {
      return (
        (t = new qn(t, 16)),
        this.isInfinity()
          ? this
          : this._hasDoubles(t)
          ? this.curve._fixedNafMul(this, t)
          : this.curve.endo
          ? this.curve._endoWnafMulAdd([this], [t])
          : this.curve._wnafMul(this, t)
      );
    }),
    (as.prototype.mulAdd = function (t, e, r) {
      var n = [this, e],
        i = [t, r];
      return this.curve.endo
        ? this.curve._endoWnafMulAdd(n, i)
        : this.curve._wnafMulAdd(1, n, i, 2);
    }),
    (as.prototype.jmulAdd = function (t, e, r) {
      var n = [this, e],
        i = [t, r];
      return this.curve.endo
        ? this.curve._endoWnafMulAdd(n, i, !0)
        : this.curve._wnafMulAdd(1, n, i, 2, !0);
    }),
    (as.prototype.eq = function (t) {
      return (
        this === t ||
        (this.inf === t.inf &&
          (this.inf || (0 === this.x.cmp(t.x) && 0 === this.y.cmp(t.y))))
      );
    }),
    (as.prototype.neg = function (t) {
      if (this.inf) return this;
      var e = this.curve.point(this.x, this.y.redNeg());
      if (t && this.precomputed) {
        var r = this.precomputed,
          n = function (t) {
            return t.neg();
          };
        e.precomputed = {
          naf: r.naf && { wnd: r.naf.wnd, points: r.naf.points.map(n) },
          doubles: r.doubles && {
            step: r.doubles.step,
            points: r.doubles.points.map(n),
          },
        };
      }
      return e;
    }),
    (as.prototype.toJ = function () {
      return this.inf
        ? this.curve.jpoint(null, null, null)
        : this.curve.jpoint(this.x, this.y, this.curve.one);
    }),
    ns(us, es.BasePoint),
    (os.prototype.jpoint = function (t, e, r) {
      return new us(this, t, e, r);
    }),
    (us.prototype.toP = function () {
      if (this.isInfinity()) return this.curve.point(null, null);
      var t = this.z.redInvm(),
        e = t.redSqr(),
        r = this.x.redMul(e),
        n = this.y.redMul(e).redMul(t);
      return this.curve.point(r, n);
    }),
    (us.prototype.neg = function () {
      return this.curve.jpoint(this.x, this.y.redNeg(), this.z);
    }),
    (us.prototype.add = function (t) {
      if (this.isInfinity()) return t;
      if (t.isInfinity()) return this;
      var e = t.z.redSqr(),
        r = this.z.redSqr(),
        n = this.x.redMul(e),
        i = t.x.redMul(r),
        o = this.y.redMul(e.redMul(t.z)),
        s = t.y.redMul(r.redMul(this.z)),
        a = n.redSub(i),
        u = o.redSub(s);
      if (0 === a.cmpn(0))
        return 0 !== u.cmpn(0)
          ? this.curve.jpoint(null, null, null)
          : this.dbl();
      var h = a.redSqr(),
        l = h.redMul(a),
        c = n.redMul(h),
        f = u.redSqr().redIAdd(l).redISub(c).redISub(c),
        d = u.redMul(c.redISub(f)).redISub(o.redMul(l)),
        m = this.z.redMul(t.z).redMul(a);
      return this.curve.jpoint(f, d, m);
    }),
    (us.prototype.mixedAdd = function (t) {
      if (this.isInfinity()) return t.toJ();
      if (t.isInfinity()) return this;
      var e = this.z.redSqr(),
        r = this.x,
        n = t.x.redMul(e),
        i = this.y,
        o = t.y.redMul(e).redMul(this.z),
        s = r.redSub(n),
        a = i.redSub(o);
      if (0 === s.cmpn(0))
        return 0 !== a.cmpn(0)
          ? this.curve.jpoint(null, null, null)
          : this.dbl();
      var u = s.redSqr(),
        h = u.redMul(s),
        l = r.redMul(u),
        c = a.redSqr().redIAdd(h).redISub(l).redISub(l),
        f = a.redMul(l.redISub(c)).redISub(i.redMul(h)),
        d = this.z.redMul(s);
      return this.curve.jpoint(c, f, d);
    }),
    (us.prototype.dblp = function (t) {
      if (0 === t) return this;
      if (this.isInfinity()) return this;
      if (!t) return this.dbl();
      var e;
      if (this.curve.zeroA || this.curve.threeA) {
        var r = this;
        for (e = 0; e < t; e++) r = r.dbl();
        return r;
      }
      var n = this.curve.a,
        i = this.curve.tinv,
        o = this.x,
        s = this.y,
        a = this.z,
        u = a.redSqr().redSqr(),
        h = s.redAdd(s);
      for (e = 0; e < t; e++) {
        var l = o.redSqr(),
          c = h.redSqr(),
          f = c.redSqr(),
          d = l.redAdd(l).redIAdd(l).redIAdd(n.redMul(u)),
          m = o.redMul(c),
          p = d.redSqr().redISub(m.redAdd(m)),
          g = m.redISub(p),
          y = d.redMul(g);
        y = y.redIAdd(y).redISub(f);
        var v = h.redMul(a);
        e + 1 < t && (u = u.redMul(f)), (o = p), (a = v), (h = y);
      }
      return this.curve.jpoint(o, h.redMul(i), a);
    }),
    (us.prototype.dbl = function () {
      return this.isInfinity()
        ? this
        : this.curve.zeroA
        ? this._zeroDbl()
        : this.curve.threeA
        ? this._threeDbl()
        : this._dbl();
    }),
    (us.prototype._zeroDbl = function () {
      var t, e, r;
      if (this.zOne) {
        var n = this.x.redSqr(),
          i = this.y.redSqr(),
          o = i.redSqr(),
          s = this.x.redAdd(i).redSqr().redISub(n).redISub(o);
        s = s.redIAdd(s);
        var a = n.redAdd(n).redIAdd(n),
          u = a.redSqr().redISub(s).redISub(s),
          h = o.redIAdd(o);
        (h = (h = h.redIAdd(h)).redIAdd(h)),
          (t = u),
          (e = a.redMul(s.redISub(u)).redISub(h)),
          (r = this.y.redAdd(this.y));
      } else {
        var l = this.x.redSqr(),
          c = this.y.redSqr(),
          f = c.redSqr(),
          d = this.x.redAdd(c).redSqr().redISub(l).redISub(f);
        d = d.redIAdd(d);
        var m = l.redAdd(l).redIAdd(l),
          p = m.redSqr(),
          g = f.redIAdd(f);
        (g = (g = g.redIAdd(g)).redIAdd(g)),
          (t = p.redISub(d).redISub(d)),
          (e = m.redMul(d.redISub(t)).redISub(g)),
          (r = (r = this.y.redMul(this.z)).redIAdd(r));
      }
      return this.curve.jpoint(t, e, r);
    }),
    (us.prototype._threeDbl = function () {
      var t, e, r;
      if (this.zOne) {
        var n = this.x.redSqr(),
          i = this.y.redSqr(),
          o = i.redSqr(),
          s = this.x.redAdd(i).redSqr().redISub(n).redISub(o);
        s = s.redIAdd(s);
        var a = n.redAdd(n).redIAdd(n).redIAdd(this.curve.a),
          u = a.redSqr().redISub(s).redISub(s);
        t = u;
        var h = o.redIAdd(o);
        (h = (h = h.redIAdd(h)).redIAdd(h)),
          (e = a.redMul(s.redISub(u)).redISub(h)),
          (r = this.y.redAdd(this.y));
      } else {
        var l = this.z.redSqr(),
          c = this.y.redSqr(),
          f = this.x.redMul(c),
          d = this.x.redSub(l).redMul(this.x.redAdd(l));
        d = d.redAdd(d).redIAdd(d);
        var m = f.redIAdd(f),
          p = (m = m.redIAdd(m)).redAdd(m);
        (t = d.redSqr().redISub(p)),
          (r = this.y.redAdd(this.z).redSqr().redISub(c).redISub(l));
        var g = c.redSqr();
        (g = (g = (g = g.redIAdd(g)).redIAdd(g)).redIAdd(g)),
          (e = d.redMul(m.redISub(t)).redISub(g));
      }
      return this.curve.jpoint(t, e, r);
    }),
    (us.prototype._dbl = function () {
      var t = this.curve.a,
        e = this.x,
        r = this.y,
        n = this.z,
        i = n.redSqr().redSqr(),
        o = e.redSqr(),
        s = r.redSqr(),
        a = o.redAdd(o).redIAdd(o).redIAdd(t.redMul(i)),
        u = e.redAdd(e),
        h = (u = u.redIAdd(u)).redMul(s),
        l = a.redSqr().redISub(h.redAdd(h)),
        c = h.redISub(l),
        f = s.redSqr();
      f = (f = (f = f.redIAdd(f)).redIAdd(f)).redIAdd(f);
      var d = a.redMul(c).redISub(f),
        m = r.redAdd(r).redMul(n);
      return this.curve.jpoint(l, d, m);
    }),
    (us.prototype.trpl = function () {
      if (!this.curve.zeroA) return this.dbl().add(this);
      var t = this.x.redSqr(),
        e = this.y.redSqr(),
        r = this.z.redSqr(),
        n = e.redSqr(),
        i = t.redAdd(t).redIAdd(t),
        o = i.redSqr(),
        s = this.x.redAdd(e).redSqr().redISub(t).redISub(n),
        a = (s = (s = (s = s.redIAdd(s)).redAdd(s).redIAdd(s)).redISub(
          o
        )).redSqr(),
        u = n.redIAdd(n);
      u = (u = (u = u.redIAdd(u)).redIAdd(u)).redIAdd(u);
      var h = i.redIAdd(s).redSqr().redISub(o).redISub(a).redISub(u),
        l = e.redMul(h);
      l = (l = l.redIAdd(l)).redIAdd(l);
      var c = this.x.redMul(a).redISub(l);
      c = (c = c.redIAdd(c)).redIAdd(c);
      var f = this.y.redMul(h.redMul(u.redISub(h)).redISub(s.redMul(a)));
      f = (f = (f = f.redIAdd(f)).redIAdd(f)).redIAdd(f);
      var d = this.z.redAdd(s).redSqr().redISub(r).redISub(a);
      return this.curve.jpoint(c, f, d);
    }),
    (us.prototype.mul = function (t, e) {
      return (t = new qn(t, e)), this.curve._wnafMul(this, t);
    }),
    (us.prototype.eq = function (t) {
      if ('affine' === t.type) return this.eq(t.toJ());
      if (this === t) return !0;
      var e = this.z.redSqr(),
        r = t.z.redSqr();
      if (0 !== this.x.redMul(r).redISub(t.x.redMul(e)).cmpn(0)) return !1;
      var n = e.redMul(this.z),
        i = r.redMul(t.z);
      return 0 === this.y.redMul(i).redISub(t.y.redMul(n)).cmpn(0);
    }),
    (us.prototype.eqXToP = function (t) {
      var e = this.z.redSqr(),
        r = t.toRed(this.curve.red).redMul(e);
      if (0 === this.x.cmp(r)) return !0;
      for (var n = t.clone(), i = this.curve.redN.redMul(e); ; ) {
        if ((n.iadd(this.curve.n), n.cmp(this.curve.p) >= 0)) return !1;
        if ((r.redIAdd(i), 0 === this.x.cmp(r))) return !0;
      }
    }),
    (us.prototype.inspect = function () {
      return this.isInfinity()
        ? '<EC JPoint Infinity>'
        : '<EC JPoint x: ' +
            this.x.toString(16, 2) +
            ' y: ' +
            this.y.toString(16, 2) +
            ' z: ' +
            this.z.toString(16, 2) +
            '>';
    }),
    (us.prototype.isInfinity = function () {
      return 0 === this.z.cmpn(0);
    });
  var hs = Vo(function (t, e) {
      var r = e;
      (r.base = es), (r.short = ss), (r.mont = null), (r.edwards = null);
    }),
    ls = Vo(function (t, e) {
      var r,
        n = e,
        i = Yo.assert;
      function o(t) {
        'short' === t.type
          ? (this.curve = new hs.short(t))
          : 'edwards' === t.type
          ? (this.curve = new hs.edwards(t))
          : (this.curve = new hs.mont(t)),
          (this.g = this.curve.g),
          (this.n = this.curve.n),
          (this.hash = t.hash),
          i(this.g.validate(), 'Invalid curve'),
          i(this.g.mul(this.n).isInfinity(), 'Invalid curve, G*N != O');
      }
      function s(t, e) {
        Object.defineProperty(n, t, {
          configurable: !0,
          enumerable: !0,
          get: function () {
            var r = new o(e);
            return (
              Object.defineProperty(n, t, {
                configurable: !0,
                enumerable: !0,
                value: r,
              }),
              r
            );
          },
        });
      }
      (n.PresetCurve = o),
        s('p192', {
          type: 'short',
          prime: 'p192',
          p: 'ffffffff ffffffff ffffffff fffffffe ffffffff ffffffff',
          a: 'ffffffff ffffffff ffffffff fffffffe ffffffff fffffffc',
          b: '64210519 e59c80e7 0fa7e9ab 72243049 feb8deec c146b9b1',
          n: 'ffffffff ffffffff ffffffff 99def836 146bc9b1 b4d22831',
          hash: Kn.sha256,
          gRed: !1,
          g: [
            '188da80e b03090f6 7cbf20eb 43a18800 f4ff0afd 82ff1012',
            '07192b95 ffc8da78 631011ed 6b24cdd5 73f977a1 1e794811',
          ],
        }),
        s('p224', {
          type: 'short',
          prime: 'p224',
          p: 'ffffffff ffffffff ffffffff ffffffff 00000000 00000000 00000001',
          a: 'ffffffff ffffffff ffffffff fffffffe ffffffff ffffffff fffffffe',
          b: 'b4050a85 0c04b3ab f5413256 5044b0b7 d7bfd8ba 270b3943 2355ffb4',
          n: 'ffffffff ffffffff ffffffff ffff16a2 e0b8f03e 13dd2945 5c5c2a3d',
          hash: Kn.sha256,
          gRed: !1,
          g: [
            'b70e0cbd 6bb4bf7f 321390b9 4a03c1d3 56c21122 343280d6 115c1d21',
            'bd376388 b5f723fb 4c22dfe6 cd4375a0 5a074764 44d58199 85007e34',
          ],
        }),
        s('p256', {
          type: 'short',
          prime: null,
          p: 'ffffffff 00000001 00000000 00000000 00000000 ffffffff ffffffff ffffffff',
          a: 'ffffffff 00000001 00000000 00000000 00000000 ffffffff ffffffff fffffffc',
          b: '5ac635d8 aa3a93e7 b3ebbd55 769886bc 651d06b0 cc53b0f6 3bce3c3e 27d2604b',
          n: 'ffffffff 00000000 ffffffff ffffffff bce6faad a7179e84 f3b9cac2 fc632551',
          hash: Kn.sha256,
          gRed: !1,
          g: [
            '6b17d1f2 e12c4247 f8bce6e5 63a440f2 77037d81 2deb33a0 f4a13945 d898c296',
            '4fe342e2 fe1a7f9b 8ee7eb4a 7c0f9e16 2bce3357 6b315ece cbb64068 37bf51f5',
          ],
        }),
        s('p384', {
          type: 'short',
          prime: null,
          p: 'ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff fffffffe ffffffff 00000000 00000000 ffffffff',
          a: 'ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff fffffffe ffffffff 00000000 00000000 fffffffc',
          b: 'b3312fa7 e23ee7e4 988e056b e3f82d19 181d9c6e fe814112 0314088f 5013875a c656398d 8a2ed19d 2a85c8ed d3ec2aef',
          n: 'ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff c7634d81 f4372ddf 581a0db2 48b0a77a ecec196a ccc52973',
          hash: Kn.sha384,
          gRed: !1,
          g: [
            'aa87ca22 be8b0537 8eb1c71e f320ad74 6e1d3b62 8ba79b98 59f741e0 82542a38 5502f25d bf55296c 3a545e38 72760ab7',
            '3617de4a 96262c6f 5d9e98bf 9292dc29 f8f41dbd 289a147c e9da3113 b5f0b8c0 0a60b1ce 1d7e819d 7a431d7c 90ea0e5f',
          ],
        }),
        s('p521', {
          type: 'short',
          prime: null,
          p: '000001ff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff',
          a: '000001ff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff fffffffc',
          b: '00000051 953eb961 8e1c9a1f 929a21a0 b68540ee a2da725b 99b315f3 b8b48991 8ef109e1 56193951 ec7e937b 1652c0bd 3bb1bf07 3573df88 3d2c34f1 ef451fd4 6b503f00',
          n: '000001ff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff fffffffa 51868783 bf2f966b 7fcc0148 f709a5d0 3bb5c9b8 899c47ae bb6fb71e 91386409',
          hash: Kn.sha512,
          gRed: !1,
          g: [
            '000000c6 858e06b7 0404e9cd 9e3ecb66 2395b442 9c648139 053fb521 f828af60 6b4d3dba a14b5e77 efe75928 fe1dc127 a2ffa8de 3348b3c1 856a429b f97e7e31 c2e5bd66',
            '00000118 39296a78 9a3bc004 5c8a5fb4 2c7d1bd9 98f54449 579b4468 17afbd17 273e662c 97ee7299 5ef42640 c550b901 3fad0761 353c7086 a272c240 88be9476 9fd16650',
          ],
        }),
        s('curve25519', {
          type: 'mont',
          prime: 'p25519',
          p: '7fffffffffffffff ffffffffffffffff ffffffffffffffff ffffffffffffffed',
          a: '76d06',
          b: '1',
          n: '1000000000000000 0000000000000000 14def9dea2f79cd6 5812631a5cf5d3ed',
          hash: Kn.sha256,
          gRed: !1,
          g: ['9'],
        }),
        s('ed25519', {
          type: 'edwards',
          prime: 'p25519',
          p: '7fffffffffffffff ffffffffffffffff ffffffffffffffff ffffffffffffffed',
          a: '-1',
          c: '1',
          d: '52036cee2b6ffe73 8cc740797779e898 00700a4d4141d8ab 75eb4dca135978a3',
          n: '1000000000000000 0000000000000000 14def9dea2f79cd6 5812631a5cf5d3ed',
          hash: Kn.sha256,
          gRed: !1,
          g: [
            '216936d3cd6e53fec0a4e231fdd6dc5c692cc7609525a7b2c9562d608f25d51a',
            '6666666666666666666666666666666666666666666666666666666666666658',
          ],
        });
      try {
        r = null.crash();
      } catch (t) {
        r = void 0;
      }
      s('secp256k1', {
        type: 'short',
        prime: 'k256',
        p: 'ffffffff ffffffff ffffffff ffffffff ffffffff ffffffff fffffffe fffffc2f',
        a: '0',
        b: '7',
        n: 'ffffffff ffffffff ffffffff fffffffe baaedce6 af48a03b bfd25e8c d0364141',
        h: '1',
        hash: Kn.sha256,
        beta: '7ae96a2b657c07106e64479eac3434e99cf0497512f58995c1396c28719501ee',
        lambda:
          '5363ad4cc05c30e0a5261c028812645a122e22ea20816678df02967c1b23bd72',
        basis: [
          {
            a: '3086d221a7d46bcde86c90e49284eb15',
            b: '-e4437ed6010e88286f547fa90abfe4c3',
          },
          {
            a: '114ca50f7a8e2f3f657c1108d9d44cfd8',
            b: '3086d221a7d46bcde86c90e49284eb15',
          },
        ],
        gRed: !1,
        g: [
          '79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798',
          '483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8',
          r,
        ],
      });
    });
  function cs(t) {
    if (!(this instanceof cs)) return new cs(t);
    (this.hash = t.hash),
      (this.predResist = !!t.predResist),
      (this.outLen = this.hash.outSize),
      (this.minEntropy = t.minEntropy || this.hash.hmacStrength),
      (this._reseed = null),
      (this.reseedInterval = null),
      (this.K = null),
      (this.V = null);
    var e = Qo.toArray(t.entropy, t.entropyEnc || 'hex'),
      r = Qo.toArray(t.nonce, t.nonceEnc || 'hex'),
      n = Qo.toArray(t.pers, t.persEnc || 'hex');
    Jo(
      e.length >= this.minEntropy / 8,
      'Not enough entropy. Minimum is: ' + this.minEntropy + ' bits'
    ),
      this._init(e, r, n);
  }
  var fs = cs;
  (cs.prototype._init = function (t, e, r) {
    var n = t.concat(e).concat(r);
    (this.K = new Array(this.outLen / 8)),
      (this.V = new Array(this.outLen / 8));
    for (var i = 0; i < this.V.length; i++) (this.K[i] = 0), (this.V[i] = 1);
    this._update(n),
      (this._reseed = 1),
      (this.reseedInterval = 281474976710656);
  }),
    (cs.prototype._hmac = function () {
      return new Kn.hmac(this.hash, this.K);
    }),
    (cs.prototype._update = function (t) {
      var e = this._hmac().update(this.V).update([0]);
      t && (e = e.update(t)),
        (this.K = e.digest()),
        (this.V = this._hmac().update(this.V).digest()),
        t &&
          ((this.K = this._hmac()
            .update(this.V)
            .update([1])
            .update(t)
            .digest()),
          (this.V = this._hmac().update(this.V).digest()));
    }),
    (cs.prototype.reseed = function (t, e, r, n) {
      'string' != typeof e && ((n = r), (r = e), (e = null)),
        (t = Qo.toArray(t, e)),
        (r = Qo.toArray(r, n)),
        Jo(
          t.length >= this.minEntropy / 8,
          'Not enough entropy. Minimum is: ' + this.minEntropy + ' bits'
        ),
        this._update(t.concat(r || [])),
        (this._reseed = 1);
    }),
    (cs.prototype.generate = function (t, e, r, n) {
      if (this._reseed > this.reseedInterval)
        throw new Error('Reseed is required');
      'string' != typeof e && ((n = r), (r = e), (e = null)),
        r && ((r = Qo.toArray(r, n || 'hex')), this._update(r));
      for (var i = []; i.length < t; )
        (this.V = this._hmac().update(this.V).digest()), (i = i.concat(this.V));
      var o = i.slice(0, t);
      return this._update(r), this._reseed++, Qo.encode(o, e);
    });
  var ds = Yo.assert;
  function ms(t, e) {
    (this.ec = t),
      (this.priv = null),
      (this.pub = null),
      e.priv && this._importPrivate(e.priv, e.privEnc),
      e.pub && this._importPublic(e.pub, e.pubEnc);
  }
  var ps = ms;
  (ms.fromPublic = function (t, e, r) {
    return e instanceof ms ? e : new ms(t, { pub: e, pubEnc: r });
  }),
    (ms.fromPrivate = function (t, e, r) {
      return e instanceof ms ? e : new ms(t, { priv: e, privEnc: r });
    }),
    (ms.prototype.validate = function () {
      var t = this.getPublic();
      return t.isInfinity()
        ? { result: !1, reason: 'Invalid public key' }
        : t.validate()
        ? t.mul(this.ec.curve.n).isInfinity()
          ? { result: !0, reason: null }
          : { result: !1, reason: 'Public key * N != O' }
        : { result: !1, reason: 'Public key is not a point' };
    }),
    (ms.prototype.getPublic = function (t, e) {
      return (
        'string' == typeof t && ((e = t), (t = null)),
        this.pub || (this.pub = this.ec.g.mul(this.priv)),
        e ? this.pub.encode(e, t) : this.pub
      );
    }),
    (ms.prototype.getPrivate = function (t) {
      return 'hex' === t ? this.priv.toString(16, 2) : this.priv;
    }),
    (ms.prototype._importPrivate = function (t, e) {
      (this.priv = new qn(t, e || 16)),
        (this.priv = this.priv.umod(this.ec.curve.n));
    }),
    (ms.prototype._importPublic = function (t, e) {
      if (t.x || t.y)
        return (
          'mont' === this.ec.curve.type
            ? ds(t.x, 'Need x coordinate')
            : ('short' !== this.ec.curve.type &&
                'edwards' !== this.ec.curve.type) ||
              ds(t.x && t.y, 'Need both x and y coordinate'),
          void (this.pub = this.ec.curve.point(t.x, t.y))
        );
      this.pub = this.ec.curve.decodePoint(t, e);
    }),
    (ms.prototype.derive = function (t) {
      return (
        t.validate() || ds(t.validate(), 'public point not validated'),
        t.mul(this.priv).getX()
      );
    }),
    (ms.prototype.sign = function (t, e, r) {
      return this.ec.sign(t, this, e, r);
    }),
    (ms.prototype.verify = function (t, e) {
      return this.ec.verify(t, e, this);
    }),
    (ms.prototype.inspect = function () {
      return (
        '<Key priv: ' +
        (this.priv && this.priv.toString(16, 2)) +
        ' pub: ' +
        (this.pub && this.pub.inspect()) +
        ' >'
      );
    });
  var gs = Yo.assert;
  function ys(t, e) {
    if (t instanceof ys) return t;
    this._importDER(t, e) ||
      (gs(t.r && t.s, 'Signature without r or s'),
      (this.r = new qn(t.r, 16)),
      (this.s = new qn(t.s, 16)),
      void 0 === t.recoveryParam
        ? (this.recoveryParam = null)
        : (this.recoveryParam = t.recoveryParam));
  }
  var vs = ys;
  function ws() {
    this.place = 0;
  }
  function bs(t, e) {
    var r = t[e.place++];
    if (!(128 & r)) return r;
    var n = 15 & r;
    if (0 === n || n > 4) return !1;
    for (var i = 0, o = 0, s = e.place; o < n; o++, s++)
      (i <<= 8), (i |= t[s]), (i >>>= 0);
    return !(i <= 127) && ((e.place = s), i);
  }
  function As(t) {
    for (var e = 0, r = t.length - 1; !t[e] && !(128 & t[e + 1]) && e < r; )
      e++;
    return 0 === e ? t : t.slice(e);
  }
  function Es(t, e) {
    if (e < 128) t.push(e);
    else {
      var r = 1 + ((Math.log(e) / Math.LN2) >>> 3);
      for (t.push(128 | r); --r; ) t.push((e >>> (r << 3)) & 255);
      t.push(e);
    }
  }
  (ys.prototype._importDER = function (t, e) {
    t = Yo.toArray(t, e);
    var r = new ws();
    if (48 !== t[r.place++]) return !1;
    var n = bs(t, r);
    if (!1 === n) return !1;
    if (n + r.place !== t.length) return !1;
    if (2 !== t[r.place++]) return !1;
    var i = bs(t, r);
    if (!1 === i) return !1;
    var o = t.slice(r.place, i + r.place);
    if (((r.place += i), 2 !== t[r.place++])) return !1;
    var s = bs(t, r);
    if (!1 === s) return !1;
    if (t.length !== s + r.place) return !1;
    var a = t.slice(r.place, s + r.place);
    if (0 === o[0]) {
      if (!(128 & o[1])) return !1;
      o = o.slice(1);
    }
    if (0 === a[0]) {
      if (!(128 & a[1])) return !1;
      a = a.slice(1);
    }
    return (
      (this.r = new qn(o)),
      (this.s = new qn(a)),
      (this.recoveryParam = null),
      !0
    );
  }),
    (ys.prototype.toDER = function (t) {
      var e = this.r.toArray(),
        r = this.s.toArray();
      for (
        128 & e[0] && (e = [0].concat(e)),
          128 & r[0] && (r = [0].concat(r)),
          e = As(e),
          r = As(r);
        !(r[0] || 128 & r[1]);

      )
        r = r.slice(1);
      var n = [2];
      Es(n, e.length), (n = n.concat(e)).push(2), Es(n, r.length);
      var i = n.concat(r),
        o = [48];
      return Es(o, i.length), (o = o.concat(i)), Yo.encode(o, t);
    });
  var Ms = function () {
      throw new Error('unsupported');
    },
    _s = Yo.assert;
  function ks(t) {
    if (!(this instanceof ks)) return new ks(t);
    'string' == typeof t &&
      (_s(Object.prototype.hasOwnProperty.call(ls, t), 'Unknown curve ' + t),
      (t = ls[t])),
      t instanceof ls.PresetCurve && (t = { curve: t }),
      (this.curve = t.curve.curve),
      (this.n = this.curve.n),
      (this.nh = this.n.ushrn(1)),
      (this.g = this.curve.g),
      (this.g = t.curve.g),
      this.g.precompute(t.curve.n.bitLength() + 1),
      (this.hash = t.hash || t.curve.hash);
  }
  var Ss = ks;
  (ks.prototype.keyPair = function (t) {
    return new ps(this, t);
  }),
    (ks.prototype.keyFromPrivate = function (t, e) {
      return ps.fromPrivate(this, t, e);
    }),
    (ks.prototype.keyFromPublic = function (t, e) {
      return ps.fromPublic(this, t, e);
    }),
    (ks.prototype.genKeyPair = function (t) {
      t || (t = {});
      for (
        var e = new fs({
            hash: this.hash,
            pers: t.pers,
            persEnc: t.persEnc || 'utf8',
            entropy: t.entropy || Ms(this.hash.hmacStrength),
            entropyEnc: (t.entropy && t.entropyEnc) || 'utf8',
            nonce: this.n.toArray(),
          }),
          r = this.n.byteLength(),
          n = this.n.sub(new qn(2));
        ;

      ) {
        var i = new qn(e.generate(r));
        if (!(i.cmp(n) > 0)) return i.iaddn(1), this.keyFromPrivate(i);
      }
    }),
    (ks.prototype._truncateToN = function (t, e) {
      var r = 8 * t.byteLength() - this.n.bitLength();
      return (
        r > 0 && (t = t.ushrn(r)), !e && t.cmp(this.n) >= 0 ? t.sub(this.n) : t
      );
    }),
    (ks.prototype.sign = function (t, e, r, n) {
      'object' == typeof r && ((n = r), (r = null)),
        n || (n = {}),
        (e = this.keyFromPrivate(e, r)),
        (t = this._truncateToN(new qn(t, 16)));
      for (
        var i = this.n.byteLength(),
          o = e.getPrivate().toArray('be', i),
          s = t.toArray('be', i),
          a = new fs({
            hash: this.hash,
            entropy: o,
            nonce: s,
            pers: n.pers,
            persEnc: n.persEnc || 'utf8',
          }),
          u = this.n.sub(new qn(1)),
          h = 0;
        ;
        h++
      ) {
        var l = n.k ? n.k(h) : new qn(a.generate(this.n.byteLength()));
        if (!((l = this._truncateToN(l, !0)).cmpn(1) <= 0 || l.cmp(u) >= 0)) {
          var c = this.g.mul(l);
          if (!c.isInfinity()) {
            var f = c.getX(),
              d = f.umod(this.n);
            if (0 !== d.cmpn(0)) {
              var m = l.invm(this.n).mul(d.mul(e.getPrivate()).iadd(t));
              if (0 !== (m = m.umod(this.n)).cmpn(0)) {
                var p = (c.getY().isOdd() ? 1 : 0) | (0 !== f.cmp(d) ? 2 : 0);
                return (
                  n.canonical &&
                    m.cmp(this.nh) > 0 &&
                    ((m = this.n.sub(m)), (p ^= 1)),
                  new vs({ r: d, s: m, recoveryParam: p })
                );
              }
            }
          }
        }
      }
    }),
    (ks.prototype.verify = function (t, e, r, n) {
      (t = this._truncateToN(new qn(t, 16))), (r = this.keyFromPublic(r, n));
      var i = (e = new vs(e, 'hex')).r,
        o = e.s;
      if (i.cmpn(1) < 0 || i.cmp(this.n) >= 0) return !1;
      if (o.cmpn(1) < 0 || o.cmp(this.n) >= 0) return !1;
      var s,
        a = o.invm(this.n),
        u = a.mul(t).umod(this.n),
        h = a.mul(i).umod(this.n);
      return this.curve._maxwellTrick
        ? !(s = this.g.jmulAdd(u, r.getPublic(), h)).isInfinity() && s.eqXToP(i)
        : !(s = this.g.mulAdd(u, r.getPublic(), h)).isInfinity() &&
            0 === s.getX().umod(this.n).cmp(i);
    }),
    (ks.prototype.recoverPubKey = function (t, e, r, n) {
      _s((3 & r) === r, 'The recovery param is more than two bits'),
        (e = new vs(e, n));
      var i = this.n,
        o = new qn(t),
        s = e.r,
        a = e.s,
        u = 1 & r,
        h = r >> 1;
      if (s.cmp(this.curve.p.umod(this.curve.n)) >= 0 && h)
        throw new Error('Unable to find sencond key candinate');
      s = h
        ? this.curve.pointFromX(s.add(this.curve.n), u)
        : this.curve.pointFromX(s, u);
      var l = e.r.invm(i),
        c = i.sub(o).mul(l).umod(i),
        f = a.mul(l).umod(i);
      return this.g.mulAdd(c, s, f);
    }),
    (ks.prototype.getKeyRecoveryParam = function (t, e, r, n) {
      if (null !== (e = new vs(e, n)).recoveryParam) return e.recoveryParam;
      for (var i = 0; i < 4; i++) {
        var o;
        try {
          o = this.recoverPubKey(t, e, i);
        } catch (t) {
          continue;
        }
        if (o.eq(r)) return i;
      }
      throw new Error('Unable to find valid recovery factor');
    });
  var Ns = Vo(function (t, e) {
    var r = e;
    (r.version = '6.5.4'),
      (r.utils = Yo),
      (r.rand = function () {
        throw new Error('unsupported');
      }),
      (r.curve = hs),
      (r.curves = ls),
      (r.ec = Ss),
      (r.eddsa = null);
  }).ec;
  const Ps = new f('signing-key/5.7.0');
  let xs = null;
  function Ts() {
    return xs || (xs = new Ns('secp256k1')), xs;
  }
  class Cs {
    constructor(t) {
      st(this, 'curve', 'secp256k1'),
        st(this, 'privateKey', k(t)),
        32 !== S(this.privateKey) &&
          Ps.throwArgumentError(
            'invalid private key',
            'privateKey',
            '[[ REDACTED ]]'
          );
      const e = Ts().keyFromPrivate(w(this.privateKey));
      st(this, 'publicKey', '0x' + e.getPublic(!1, 'hex')),
        st(this, 'compressedPublicKey', '0x' + e.getPublic(!0, 'hex')),
        st(this, '_isSigningKey', !0);
    }
    _addPoint(t) {
      const e = Ts().keyFromPublic(w(this.publicKey)),
        r = Ts().keyFromPublic(w(t));
      return '0x' + e.pub.add(r.pub).encodeCompressed('hex');
    }
    signDigest(t) {
      const e = Ts().keyFromPrivate(w(this.privateKey)),
        r = w(t);
      32 !== r.length &&
        Ps.throwArgumentError('bad digest length', 'digest', t);
      const n = e.sign(r, { canonical: !0 });
      return R({
        recoveryParam: n.recoveryParam,
        r: C('0x' + n.r.toString(16), 32),
        s: C('0x' + n.s.toString(16), 32),
      });
    }
    computeSharedSecret(t) {
      const e = Ts().keyFromPrivate(w(this.privateKey)),
        r = Ts().keyFromPublic(w(Is(t)));
      return C('0x' + e.derive(r.getPublic()).toString(16), 32);
    }
    static isSigningKey(t) {
      return !(!t || !t._isSigningKey);
    }
  }
  function Rs(t, e) {
    const r = R(e),
      n = { r: w(r.r), s: w(r.s) };
    return (
      '0x' + Ts().recoverPubKey(w(t), n, r.recoveryParam).encode('hex', !1)
    );
  }
  function Is(t, e) {
    const r = w(t);
    if (32 === r.length) {
      const t = new Cs(r);
      return e
        ? '0x' + Ts().keyFromPrivate(r).getPublic(!0, 'hex')
        : t.publicKey;
    }
    return 33 === r.length
      ? e
        ? k(r)
        : '0x' + Ts().keyFromPublic(r).getPublic(!1, 'hex')
      : 65 === r.length
      ? e
        ? '0x' + Ts().keyFromPublic(r).getPublic(!0, 'hex')
        : k(r)
      : Ps.throwArgumentError(
          'invalid public or private key',
          'key',
          '[REDACTED]'
        );
  }
  const Bs = new f('transactions/5.7.0');
  var Os;
  function Fs(t) {
    return '0x' === t ? null : Re(t);
  }
  function Ds(t) {
    return '0x' === t ? Je : j.from(t);
  }
  !(function (t) {
    (t[(t.legacy = 0)] = 'legacy'),
      (t[(t.eip2930 = 1)] = 'eip2930'),
      (t[(t.eip1559 = 2)] = 'eip1559');
  })(Os || (Os = {}));
  const Ls = [
      { name: 'nonce', maxLength: 32, numeric: !0 },
      { name: 'gasPrice', maxLength: 32, numeric: !0 },
      { name: 'gasLimit', maxLength: 32, numeric: !0 },
      { name: 'to', length: 20 },
      { name: 'value', maxLength: 32, numeric: !0 },
      { name: 'data' },
    ],
    Us = {
      chainId: !0,
      data: !0,
      gasLimit: !0,
      gasPrice: !0,
      nonce: !0,
      to: !0,
      type: !0,
      value: !0,
    };
  function js(t) {
    return Re(N(ye(N(Is(t), 1)), 12));
  }
  function zs(t, e) {
    return js(Rs(w(t), e));
  }
  function Gs(t, e) {
    const r = A(j.from(t).toHexString());
    return (
      r.length > 32 &&
        Bs.throwArgumentError('invalid length for ' + e, 'transaction:' + e, t),
      r
    );
  }
  function qs(t, e) {
    return {
      address: Re(t),
      storageKeys: (e || []).map(
        (e, r) => (
          32 !== S(e) &&
            Bs.throwArgumentError(
              'invalid access list storageKey',
              `accessList[${t}:${r}]`,
              e
            ),
          e.toLowerCase()
        )
      ),
    };
  }
  function Ks(t) {
    if (Array.isArray(t))
      return t.map((t, e) =>
        Array.isArray(t)
          ? (t.length > 2 &&
              Bs.throwArgumentError(
                'access list expected to be [ address, storageKeys[] ]',
                `value[${e}]`,
                t
              ),
            qs(t[0], t[1]))
          : qs(t.address, t.storageKeys)
      );
    const e = Object.keys(t).map((e) => {
      const r = t[e].reduce((t, e) => ((t[e] = !0), t), {});
      return qs(e, Object.keys(r).sort());
    });
    return e.sort((t, e) => t.address.localeCompare(e.address)), e;
  }
  function Hs(t) {
    return Ks(t).map((t) => [t.address, t.storageKeys]);
  }
  function Vs(t, e) {
    if (null != t.gasPrice) {
      const e = j.from(t.gasPrice),
        r = j.from(t.maxFeePerGas || 0);
      e.eq(r) ||
        Bs.throwArgumentError(
          'mismatch EIP-1559 gasPrice != maxFeePerGas',
          'tx',
          { gasPrice: e, maxFeePerGas: r }
        );
    }
    const r = [
      Gs(t.chainId || 0, 'chainId'),
      Gs(t.nonce || 0, 'nonce'),
      Gs(t.maxPriorityFeePerGas || 0, 'maxPriorityFeePerGas'),
      Gs(t.maxFeePerGas || 0, 'maxFeePerGas'),
      Gs(t.gasLimit || 0, 'gasLimit'),
      null != t.to ? Re(t.to) : '0x',
      Gs(t.value || 0, 'value'),
      t.data || '0x',
      Hs(t.accessList || []),
    ];
    if (e) {
      const t = R(e);
      r.push(Gs(t.recoveryParam, 'recoveryParam')),
        r.push(A(t.r)),
        r.push(A(t.s));
    }
    return P(['0x02', Ee(r)]);
  }
  function Js(t, e) {
    const r = [
      Gs(t.chainId || 0, 'chainId'),
      Gs(t.nonce || 0, 'nonce'),
      Gs(t.gasPrice || 0, 'gasPrice'),
      Gs(t.gasLimit || 0, 'gasLimit'),
      null != t.to ? Re(t.to) : '0x',
      Gs(t.value || 0, 'value'),
      t.data || '0x',
      Hs(t.accessList || []),
    ];
    if (e) {
      const t = R(e);
      r.push(Gs(t.recoveryParam, 'recoveryParam')),
        r.push(A(t.r)),
        r.push(A(t.s));
    }
    return P(['0x01', Ee(r)]);
  }
  function Ws(t, e) {
    if (null == t.type || 0 === t.type)
      return (
        null != t.accessList &&
          Bs.throwArgumentError(
            'untyped transactions do not support accessList; include type: 1',
            'transaction',
            t
          ),
        (function (t, e) {
          ht(t, Us);
          const r = [];
          Ls.forEach(function (e) {
            let n = t[e.name] || [];
            const i = {};
            e.numeric && (i.hexPad = 'left'),
              (n = w(k(n, i))),
              e.length &&
                n.length !== e.length &&
                n.length > 0 &&
                Bs.throwArgumentError(
                  'invalid length for ' + e.name,
                  'transaction:' + e.name,
                  n
                ),
              e.maxLength &&
                ((n = A(n)),
                n.length > e.maxLength &&
                  Bs.throwArgumentError(
                    'invalid length for ' + e.name,
                    'transaction:' + e.name,
                    n
                  )),
              r.push(k(n));
          });
          let n = 0;
          if (
            (null != t.chainId
              ? ((n = t.chainId),
                'number' != typeof n &&
                  Bs.throwArgumentError(
                    'invalid transaction.chainId',
                    'transaction',
                    t
                  ))
              : e && !g(e) && e.v > 28 && (n = Math.floor((e.v - 35) / 2)),
            0 !== n && (r.push(k(n)), r.push('0x'), r.push('0x')),
            !e)
          )
            return Ee(r);
          const i = R(e);
          let o = 27 + i.recoveryParam;
          return (
            0 !== n
              ? (r.pop(),
                r.pop(),
                r.pop(),
                (o += 2 * n + 8),
                i.v > 28 &&
                  i.v !== o &&
                  Bs.throwArgumentError(
                    'transaction.chainId/signature.v mismatch',
                    'signature',
                    e
                  ))
              : i.v !== o &&
                Bs.throwArgumentError(
                  'transaction.chainId/signature.v mismatch',
                  'signature',
                  e
                ),
            r.push(k(o)),
            r.push(A(w(i.r))),
            r.push(A(w(i.s))),
            Ee(r)
          );
        })(t, e)
      );
    switch (t.type) {
      case 1:
        return Js(t, e);
      case 2:
        return Vs(t, e);
    }
    return Bs.throwError(
      `unsupported transaction type: ${t.type}`,
      f.errors.UNSUPPORTED_OPERATION,
      { operation: 'serializeTransaction', transactionType: t.type }
    );
  }
  function Qs(t, e, r) {
    try {
      const r = Ds(e[0]).toNumber();
      if (0 !== r && 1 !== r) throw new Error('bad recid');
      t.v = r;
    } catch (t) {
      Bs.throwArgumentError('invalid v for transaction type: 1', 'v', e[0]);
    }
    (t.r = C(e[1], 32)), (t.s = C(e[2], 32));
    try {
      const e = ye(r(t));
      t.from = zs(e, { r: t.r, s: t.s, recoveryParam: t.v });
    } catch (t) {}
  }
  function Ys(t) {
    const e = w(t);
    if (e[0] > 127)
      return (function (t) {
        const e = ke(t);
        9 !== e.length &&
          6 !== e.length &&
          Bs.throwArgumentError('invalid raw transaction', 'rawTransaction', t);
        const r = {
          nonce: Ds(e[0]).toNumber(),
          gasPrice: Ds(e[1]),
          gasLimit: Ds(e[2]),
          to: Fs(e[3]),
          value: Ds(e[4]),
          data: e[5],
          chainId: 0,
        };
        if (6 === e.length) return r;
        try {
          r.v = j.from(e[6]).toNumber();
        } catch (t) {
          return r;
        }
        if (
          ((r.r = C(e[7], 32)),
          (r.s = C(e[8], 32)),
          j.from(r.r).isZero() && j.from(r.s).isZero())
        )
          (r.chainId = r.v), (r.v = 0);
        else {
          (r.chainId = Math.floor((r.v - 35) / 2)),
            r.chainId < 0 && (r.chainId = 0);
          let n = r.v - 27;
          const i = e.slice(0, 6);
          0 !== r.chainId &&
            (i.push(k(r.chainId)),
            i.push('0x'),
            i.push('0x'),
            (n -= 2 * r.chainId + 8));
          const o = ye(Ee(i));
          try {
            r.from = zs(o, { r: k(r.r), s: k(r.s), recoveryParam: n });
          } catch (t) {}
          r.hash = ye(t);
        }
        return (r.type = null), r;
      })(e);
    switch (e[0]) {
      case 1:
        return (function (t) {
          const e = ke(t.slice(1));
          8 !== e.length &&
            11 !== e.length &&
            Bs.throwArgumentError(
              'invalid component count for transaction type: 1',
              'payload',
              k(t)
            );
          const r = {
            type: 1,
            chainId: Ds(e[0]).toNumber(),
            nonce: Ds(e[1]).toNumber(),
            gasPrice: Ds(e[2]),
            gasLimit: Ds(e[3]),
            to: Fs(e[4]),
            value: Ds(e[5]),
            data: e[6],
            accessList: Ks(e[7]),
          };
          return 8 === e.length || ((r.hash = ye(t)), Qs(r, e.slice(8), Js)), r;
        })(e);
      case 2:
        return (function (t) {
          const e = ke(t.slice(1));
          9 !== e.length &&
            12 !== e.length &&
            Bs.throwArgumentError(
              'invalid component count for transaction type: 2',
              'payload',
              k(t)
            );
          const r = Ds(e[2]),
            n = Ds(e[3]),
            i = {
              type: 2,
              chainId: Ds(e[0]).toNumber(),
              nonce: Ds(e[1]).toNumber(),
              maxPriorityFeePerGas: r,
              maxFeePerGas: n,
              gasPrice: null,
              gasLimit: Ds(e[4]),
              to: Fs(e[5]),
              value: Ds(e[6]),
              data: e[7],
              accessList: Ks(e[8]),
            };
          return 9 === e.length || ((i.hash = ye(t)), Qs(i, e.slice(9), Vs)), i;
        })(e);
    }
    return Bs.throwError(
      `unsupported transaction type: ${e[0]}`,
      f.errors.UNSUPPORTED_OPERATION,
      { operation: 'parseTransaction', transactionType: e[0] }
    );
  }
  var Zs = function (t, e, r, n) {
    return new (r || (r = Promise))(function (i, o) {
      function s(t) {
        try {
          u(n.next(t));
        } catch (t) {
          o(t);
        }
      }
      function a(t) {
        try {
          u(n.throw(t));
        } catch (t) {
          o(t);
        }
      }
      function u(t) {
        var e;
        t.done
          ? i(t.value)
          : ((e = t.value),
            e instanceof r
              ? e
              : new r(function (t) {
                  t(e);
                })).then(s, a);
      }
      u((n = n.apply(t, e || [])).next());
    });
  };
  const Xs = new f('contracts/5.7.0'),
    $s = {
      chainId: !0,
      data: !0,
      from: !0,
      gasLimit: !0,
      gasPrice: !0,
      nonce: !0,
      to: !0,
      value: !0,
      type: !0,
      accessList: !0,
      maxFeePerGas: !0,
      maxPriorityFeePerGas: !0,
      customData: !0,
      ccipReadEnabled: !0,
    };
  function ta(t, e) {
    return Zs(this, void 0, void 0, function* () {
      const r = yield e;
      'string' != typeof r &&
        Xs.throwArgumentError('invalid address or ENS name', 'name', r);
      try {
        return Re(r);
      } catch (t) {}
      t ||
        Xs.throwError(
          'a provider or signer is needed to resolve ENS names',
          f.errors.UNSUPPORTED_OPERATION,
          { operation: 'resolveName' }
        );
      const n = yield t.resolveName(r);
      return (
        null == n &&
          Xs.throwArgumentError(
            'resolver or addr is not configured for ENS name',
            'name',
            r
          ),
        n
      );
    });
  }
  function ea(t, e, r) {
    return Zs(this, void 0, void 0, function* () {
      return Array.isArray(r)
        ? yield Promise.all(
            r.map((r, n) => ea(t, Array.isArray(e) ? e[n] : e[r.name], r))
          )
        : 'address' === r.type
        ? yield ta(t, e)
        : 'tuple' === r.type
        ? yield ea(t, e, r.components)
        : 'array' === r.baseType
        ? Array.isArray(e)
          ? yield Promise.all(e.map((e) => ea(t, e, r.arrayChildren)))
          : Promise.reject(
              Xs.makeError(
                'invalid value for array',
                f.errors.INVALID_ARGUMENT,
                { argument: 'value', value: e }
              )
            )
        : e;
    });
  }
  function ra(t, e, r) {
    return Zs(this, void 0, void 0, function* () {
      let n = {};
      r.length === e.inputs.length + 1 &&
        'object' == typeof r[r.length - 1] &&
        (n = lt(r.pop())),
        Xs.checkArgumentCount(r.length, e.inputs.length, 'passed to contract'),
        t.signer
          ? n.from
            ? (n.from = ut({
                override: ta(t.signer, n.from),
                signer: t.signer.getAddress(),
              }).then((t) =>
                Zs(this, void 0, void 0, function* () {
                  return (
                    Re(t.signer) !== t.override &&
                      Xs.throwError(
                        'Contract with a Signer cannot override from',
                        f.errors.UNSUPPORTED_OPERATION,
                        { operation: 'overrides.from' }
                      ),
                    t.override
                  );
                })
              ))
            : (n.from = t.signer.getAddress())
          : n.from && (n.from = ta(t.provider, n.from));
      const i = yield ut({
          args: ea(t.signer || t.provider, r, e.inputs),
          address: t.resolvedAddress,
          overrides: ut(n) || {},
        }),
        o = t.interface.encodeFunctionData(e, i.args),
        s = { data: o, to: i.address },
        a = i.overrides;
      if (
        (null != a.nonce && (s.nonce = j.from(a.nonce).toNumber()),
        null != a.gasLimit && (s.gasLimit = j.from(a.gasLimit)),
        null != a.gasPrice && (s.gasPrice = j.from(a.gasPrice)),
        null != a.maxFeePerGas && (s.maxFeePerGas = j.from(a.maxFeePerGas)),
        null != a.maxPriorityFeePerGas &&
          (s.maxPriorityFeePerGas = j.from(a.maxPriorityFeePerGas)),
        null != a.from && (s.from = a.from),
        null != a.type && (s.type = a.type),
        null != a.accessList && (s.accessList = Ks(a.accessList)),
        null == s.gasLimit && null != e.gas)
      ) {
        let t = 21e3;
        const r = w(o);
        for (let e = 0; e < r.length; e++) (t += 4), r[e] && (t += 64);
        s.gasLimit = j.from(e.gas).add(t);
      }
      if (a.value) {
        const t = j.from(a.value);
        t.isZero() ||
          e.payable ||
          Xs.throwError(
            'non-payable method cannot override value',
            f.errors.UNSUPPORTED_OPERATION,
            { operation: 'overrides.value', value: n.value }
          ),
          (s.value = t);
      }
      a.customData && (s.customData = lt(a.customData)),
        a.ccipReadEnabled && (s.ccipReadEnabled = !!a.ccipReadEnabled),
        delete n.nonce,
        delete n.gasLimit,
        delete n.gasPrice,
        delete n.from,
        delete n.value,
        delete n.type,
        delete n.accessList,
        delete n.maxFeePerGas,
        delete n.maxPriorityFeePerGas,
        delete n.customData,
        delete n.ccipReadEnabled;
      const u = Object.keys(n).filter((t) => null != n[t]);
      return (
        u.length &&
          Xs.throwError(
            `cannot override ${u.map((t) => JSON.stringify(t)).join(',')}`,
            f.errors.UNSUPPORTED_OPERATION,
            { operation: 'overrides', overrides: u }
          ),
        s
      );
    });
  }
  function na(t, e) {
    const r = e.wait.bind(e);
    e.wait = (e) =>
      r(e).then(
        (e) => (
          (e.events = e.logs.map((r) => {
            let n = mt(r),
              i = null;
            try {
              i = t.interface.parseLog(r);
            } catch (t) {}
            return (
              i &&
                ((n.args = i.args),
                (n.decode = (e, r) =>
                  t.interface.decodeEventLog(i.eventFragment, e, r)),
                (n.event = i.name),
                (n.eventSignature = i.signature)),
              (n.removeListener = () => t.provider),
              (n.getBlock = () => t.provider.getBlock(e.blockHash)),
              (n.getTransaction = () =>
                t.provider.getTransaction(e.transactionHash)),
              (n.getTransactionReceipt = () => Promise.resolve(e)),
              n
            );
          })),
          e
        )
      );
  }
  function ia(t, e, r) {
    const n = t.signer || t.provider;
    return function (...i) {
      return Zs(this, void 0, void 0, function* () {
        let o;
        if (
          i.length === e.inputs.length + 1 &&
          'object' == typeof i[i.length - 1]
        ) {
          const t = lt(i.pop());
          null != t.blockTag && (o = yield t.blockTag),
            delete t.blockTag,
            i.push(t);
        }
        null != t.deployTransaction && (yield t._deployed(o));
        const s = yield ra(t, e, i),
          a = yield n.call(s, o);
        try {
          let n = t.interface.decodeFunctionResult(e, a);
          return r && 1 === e.outputs.length && (n = n[0]), n;
        } catch (e) {
          throw (
            (e.code === f.errors.CALL_EXCEPTION &&
              ((e.address = t.address), (e.args = i), (e.transaction = s)),
            e)
          );
        }
      });
    };
  }
  function oa(t, e, r) {
    return e.constant
      ? ia(t, e, r)
      : (function (t, e) {
          return function (...r) {
            return Zs(this, void 0, void 0, function* () {
              t.signer ||
                Xs.throwError(
                  'sending a transaction requires a signer',
                  f.errors.UNSUPPORTED_OPERATION,
                  { operation: 'sendTransaction' }
                ),
                null != t.deployTransaction && (yield t._deployed());
              const n = yield ra(t, e, r),
                i = yield t.signer.sendTransaction(n);
              return na(t, i), i;
            });
          };
        })(t, e);
  }
  function sa(t) {
    return !t.address || (null != t.topics && 0 !== t.topics.length)
      ? (t.address || '*') +
          '@' +
          (t.topics
            ? t.topics
                .map((t) => (Array.isArray(t) ? t.join('|') : t))
                .join(':')
            : '')
      : '*';
  }
  class aa {
    constructor(t, e) {
      st(this, 'tag', t), st(this, 'filter', e), (this._listeners = []);
    }
    addListener(t, e) {
      this._listeners.push({ listener: t, once: e });
    }
    removeListener(t) {
      let e = !1;
      this._listeners = this._listeners.filter(
        (r) => !(!e && r.listener === t) || ((e = !0), !1)
      );
    }
    removeAllListeners() {
      this._listeners = [];
    }
    listeners() {
      return this._listeners.map((t) => t.listener);
    }
    listenerCount() {
      return this._listeners.length;
    }
    run(t) {
      const e = this.listenerCount();
      return (
        (this._listeners = this._listeners.filter((e) => {
          const r = t.slice();
          return (
            setTimeout(() => {
              e.listener.apply(this, r);
            }, 0),
            !e.once
          );
        })),
        e
      );
    }
    prepareEvent(t) {}
    getEmit(t) {
      return [t];
    }
  }
  class ua extends aa {
    constructor() {
      super('error', null);
    }
  }
  class ha extends aa {
    constructor(t, e, r, n) {
      const i = { address: t };
      let o = e.getEventTopic(r);
      n
        ? (o !== n[0] && Xs.throwArgumentError('topic mismatch', 'topics', n),
          (i.topics = n.slice()))
        : (i.topics = [o]),
        super(sa(i), i),
        st(this, 'address', t),
        st(this, 'interface', e),
        st(this, 'fragment', r);
    }
    prepareEvent(t) {
      super.prepareEvent(t),
        (t.event = this.fragment.name),
        (t.eventSignature = this.fragment.format()),
        (t.decode = (t, e) =>
          this.interface.decodeEventLog(this.fragment, t, e));
      try {
        t.args = this.interface.decodeEventLog(this.fragment, t.data, t.topics);
      } catch (e) {
        (t.args = null), (t.decodeError = e);
      }
    }
    getEmit(t) {
      const e = (function (t) {
        const e = [],
          r = function (t, n) {
            if (Array.isArray(n))
              for (let i in n) {
                const o = t.slice();
                o.push(i);
                try {
                  r(o, n[i]);
                } catch (t) {
                  e.push({ path: o, error: t });
                }
              }
          };
        return r([], t), e;
      })(t.args);
      if (e.length) throw e[0].error;
      const r = (t.args || []).slice();
      return r.push(t), r;
    }
  }
  class la extends aa {
    constructor(t, e) {
      super('*', { address: t }),
        st(this, 'address', t),
        st(this, 'interface', e);
    }
    prepareEvent(t) {
      super.prepareEvent(t);
      try {
        const e = this.interface.parseLog(t);
        (t.event = e.name),
          (t.eventSignature = e.signature),
          (t.decode = (t, r) =>
            this.interface.decodeEventLog(e.eventFragment, t, r)),
          (t.args = e.args);
      } catch (t) {}
    }
  }
  class ca {
    constructor(t, e, r) {
      st(this, 'interface', at(new.target, 'getInterface')(e)),
        null == r
          ? (st(this, 'provider', null), st(this, 'signer', null))
          : jn.isSigner(r)
          ? (st(this, 'provider', r.provider || null), st(this, 'signer', r))
          : On.isProvider(r)
          ? (st(this, 'provider', r), st(this, 'signer', null))
          : Xs.throwArgumentError(
              'invalid signer or provider',
              'signerOrProvider',
              r
            ),
        st(this, 'callStatic', {}),
        st(this, 'estimateGas', {}),
        st(this, 'functions', {}),
        st(this, 'populateTransaction', {}),
        st(this, 'filters', {});
      {
        const t = {};
        Object.keys(this.interface.events).forEach((e) => {
          const r = this.interface.events[e];
          st(this.filters, e, (...t) => ({
            address: this.address,
            topics: this.interface.encodeFilterTopics(r, t),
          })),
            t[r.name] || (t[r.name] = []),
            t[r.name].push(e);
        }),
          Object.keys(t).forEach((e) => {
            const r = t[e];
            1 === r.length
              ? st(this.filters, e, this.filters[r[0]])
              : Xs.warn(`Duplicate definition of ${e} (${r.join(', ')})`);
          });
      }
      if (
        (st(this, '_runningEvents', {}),
        st(this, '_wrappedEmits', {}),
        null == t &&
          Xs.throwArgumentError(
            'invalid contract address or ENS name',
            'addressOrName',
            t
          ),
        st(this, 'address', t),
        this.provider)
      )
        st(this, 'resolvedAddress', ta(this.provider, t));
      else
        try {
          st(this, 'resolvedAddress', Promise.resolve(Re(t)));
        } catch (t) {
          Xs.throwError(
            'provider is required to use ENS name as contract address',
            f.errors.UNSUPPORTED_OPERATION,
            { operation: 'new Contract' }
          );
        }
      this.resolvedAddress.catch((t) => {});
      const n = {},
        i = {};
      Object.keys(this.interface.functions).forEach((t) => {
        const e = this.interface.functions[t];
        if (i[t]) Xs.warn(`Duplicate ABI entry for ${JSON.stringify(t)}`);
        else {
          i[t] = !0;
          {
            const r = e.name;
            n[`%${r}`] || (n[`%${r}`] = []), n[`%${r}`].push(t);
          }
          null == this[t] && st(this, t, oa(this, e, !0)),
            null == this.functions[t] && st(this.functions, t, oa(this, e, !1)),
            null == this.callStatic[t] &&
              st(this.callStatic, t, ia(this, e, !0)),
            null == this.populateTransaction[t] &&
              st(
                this.populateTransaction,
                t,
                (function (t, e) {
                  return function (...r) {
                    return ra(t, e, r);
                  };
                })(this, e)
              ),
            null == this.estimateGas[t] &&
              st(
                this.estimateGas,
                t,
                (function (t, e) {
                  const r = t.signer || t.provider;
                  return function (...n) {
                    return Zs(this, void 0, void 0, function* () {
                      r ||
                        Xs.throwError(
                          'estimate require a provider or signer',
                          f.errors.UNSUPPORTED_OPERATION,
                          { operation: 'estimateGas' }
                        );
                      const i = yield ra(t, e, n);
                      return yield r.estimateGas(i);
                    });
                  };
                })(this, e)
              );
        }
      }),
        Object.keys(n).forEach((t) => {
          const e = n[t];
          if (e.length > 1) return;
          t = t.substring(1);
          const r = e[0];
          try {
            null == this[t] && st(this, t, this[r]);
          } catch (t) {}
          null == this.functions[t] && st(this.functions, t, this.functions[r]),
            null == this.callStatic[t] &&
              st(this.callStatic, t, this.callStatic[r]),
            null == this.populateTransaction[t] &&
              st(this.populateTransaction, t, this.populateTransaction[r]),
            null == this.estimateGas[t] &&
              st(this.estimateGas, t, this.estimateGas[r]);
        });
    }
    static getContractAddress(t) {
      return Ie(t);
    }
    static getInterface(t) {
      return Cn.isInterface(t) ? t : new Cn(t);
    }
    deployed() {
      return this._deployed();
    }
    _deployed(t) {
      return (
        this._deployedPromise ||
          (this.deployTransaction
            ? (this._deployedPromise = this.deployTransaction
                .wait()
                .then(() => this))
            : (this._deployedPromise = this.provider
                .getCode(this.address, t)
                .then(
                  (t) => (
                    '0x' === t &&
                      Xs.throwError(
                        'contract not deployed',
                        f.errors.UNSUPPORTED_OPERATION,
                        {
                          contractAddress: this.address,
                          operation: 'getDeployed',
                        }
                      ),
                    this
                  )
                ))),
        this._deployedPromise
      );
    }
    fallback(t) {
      this.signer ||
        Xs.throwError(
          'sending a transactions require a signer',
          f.errors.UNSUPPORTED_OPERATION,
          { operation: 'sendTransaction(fallback)' }
        );
      const e = lt(t || {});
      return (
        ['from', 'to'].forEach(function (t) {
          null != e[t] &&
            Xs.throwError(
              'cannot override ' + t,
              f.errors.UNSUPPORTED_OPERATION,
              { operation: t }
            );
        }),
        (e.to = this.resolvedAddress),
        this.deployed().then(() => this.signer.sendTransaction(e))
      );
    }
    connect(t) {
      'string' == typeof t && (t = new zn(t, this.provider));
      const e = new this.constructor(this.address, this.interface, t);
      return (
        this.deployTransaction &&
          st(e, 'deployTransaction', this.deployTransaction),
        e
      );
    }
    attach(t) {
      return new this.constructor(
        t,
        this.interface,
        this.signer || this.provider
      );
    }
    static isIndexed(t) {
      return Pn.isIndexed(t);
    }
    _normalizeRunningEvent(t) {
      return this._runningEvents[t.tag] ? this._runningEvents[t.tag] : t;
    }
    _getRunningEvent(t) {
      if ('string' == typeof t) {
        if ('error' === t) return this._normalizeRunningEvent(new ua());
        if ('event' === t)
          return this._normalizeRunningEvent(new aa('event', null));
        if ('*' === t)
          return this._normalizeRunningEvent(
            new la(this.address, this.interface)
          );
        const e = this.interface.getEvent(t);
        return this._normalizeRunningEvent(
          new ha(this.address, this.interface, e)
        );
      }
      if (t.topics && t.topics.length > 0) {
        try {
          const e = t.topics[0];
          if ('string' != typeof e) throw new Error('invalid topic');
          const r = this.interface.getEvent(e);
          return this._normalizeRunningEvent(
            new ha(this.address, this.interface, r, t.topics)
          );
        } catch (t) {}
        const e = { address: this.address, topics: t.topics };
        return this._normalizeRunningEvent(new aa(sa(e), e));
      }
      return this._normalizeRunningEvent(new la(this.address, this.interface));
    }
    _checkRunningEvents(t) {
      if (0 === t.listenerCount()) {
        delete this._runningEvents[t.tag];
        const e = this._wrappedEmits[t.tag];
        e &&
          t.filter &&
          (this.provider.off(t.filter, e), delete this._wrappedEmits[t.tag]);
      }
    }
    _wrapEvent(t, e, r) {
      const n = mt(e);
      return (
        (n.removeListener = () => {
          r && (t.removeListener(r), this._checkRunningEvents(t));
        }),
        (n.getBlock = () => this.provider.getBlock(e.blockHash)),
        (n.getTransaction = () =>
          this.provider.getTransaction(e.transactionHash)),
        (n.getTransactionReceipt = () =>
          this.provider.getTransactionReceipt(e.transactionHash)),
        t.prepareEvent(n),
        n
      );
    }
    _addEventListener(t, e, r) {
      if (
        (this.provider ||
          Xs.throwError(
            'events require a provider or a signer with a provider',
            f.errors.UNSUPPORTED_OPERATION,
            { operation: 'once' }
          ),
        t.addListener(e, r),
        (this._runningEvents[t.tag] = t),
        !this._wrappedEmits[t.tag])
      ) {
        const r = (r) => {
          let n = this._wrapEvent(t, r, e);
          if (null == n.decodeError)
            try {
              const e = t.getEmit(n);
              this.emit(t.filter, ...e);
            } catch (t) {
              n.decodeError = t.error;
            }
          null != t.filter && this.emit('event', n),
            null != n.decodeError && this.emit('error', n.decodeError, n);
        };
        (this._wrappedEmits[t.tag] = r),
          null != t.filter && this.provider.on(t.filter, r);
      }
    }
    queryFilter(t, e, r) {
      const n = this._getRunningEvent(t),
        i = lt(n.filter);
      return (
        'string' == typeof e && M(e, 32)
          ? (null != r &&
              Xs.throwArgumentError(
                'cannot specify toBlock with blockhash',
                'toBlock',
                r
              ),
            (i.blockHash = e))
          : ((i.fromBlock = null != e ? e : 0),
            (i.toBlock = null != r ? r : 'latest')),
        this.provider
          .getLogs(i)
          .then((t) => t.map((t) => this._wrapEvent(n, t, null)))
      );
    }
    on(t, e) {
      return this._addEventListener(this._getRunningEvent(t), e, !1), this;
    }
    once(t, e) {
      return this._addEventListener(this._getRunningEvent(t), e, !0), this;
    }
    emit(t, ...e) {
      if (!this.provider) return !1;
      const r = this._getRunningEvent(t),
        n = r.run(e) > 0;
      return this._checkRunningEvents(r), n;
    }
    listenerCount(t) {
      return this.provider
        ? null == t
          ? Object.keys(this._runningEvents).reduce(
              (t, e) => t + this._runningEvents[e].listenerCount(),
              0
            )
          : this._getRunningEvent(t).listenerCount()
        : 0;
    }
    listeners(t) {
      if (!this.provider) return [];
      if (null == t) {
        const t = [];
        for (let e in this._runningEvents)
          this._runningEvents[e].listeners().forEach((e) => {
            t.push(e);
          });
        return t;
      }
      return this._getRunningEvent(t).listeners();
    }
    removeAllListeners(t) {
      if (!this.provider) return this;
      if (null == t) {
        for (const t in this._runningEvents) {
          const e = this._runningEvents[t];
          e.removeAllListeners(), this._checkRunningEvents(e);
        }
        return this;
      }
      const e = this._getRunningEvent(t);
      return e.removeAllListeners(), this._checkRunningEvents(e), this;
    }
    off(t, e) {
      if (!this.provider) return this;
      const r = this._getRunningEvent(t);
      return r.removeListener(e), this._checkRunningEvents(r), this;
    }
    removeListener(t, e) {
      return this.off(t, e);
    }
  }
  class fa extends ca {}
  class da {
    constructor(t, e, r) {
      let n = null;
      (n =
        'string' == typeof e
          ? e
          : v(e)
          ? k(e)
          : e && 'string' == typeof e.object
          ? e.object
          : '!'),
        '0x' !== n.substring(0, 2) && (n = '0x' + n),
        (!M(n) || n.length % 2) &&
          Xs.throwArgumentError('invalid bytecode', 'bytecode', e),
        r &&
          !jn.isSigner(r) &&
          Xs.throwArgumentError('invalid signer', 'signer', r),
        st(this, 'bytecode', n),
        st(this, 'interface', at(new.target, 'getInterface')(t)),
        st(this, 'signer', r || null);
    }
    getDeployTransaction(...t) {
      let e = {};
      if (
        t.length === this.interface.deploy.inputs.length + 1 &&
        'object' == typeof t[t.length - 1]
      ) {
        e = lt(t.pop());
        for (const t in e)
          if (!$s[t]) throw new Error('unknown transaction override ' + t);
      }
      if (
        (['data', 'from', 'to'].forEach((t) => {
          null != e[t] &&
            Xs.throwError(
              'cannot override ' + t,
              f.errors.UNSUPPORTED_OPERATION,
              { operation: t }
            );
        }),
        e.value)
      ) {
        j.from(e.value).isZero() ||
          this.interface.deploy.payable ||
          Xs.throwError(
            'non-payable constructor cannot override value',
            f.errors.UNSUPPORTED_OPERATION,
            { operation: 'overrides.value', value: e.value }
          );
      }
      return (
        Xs.checkArgumentCount(
          t.length,
          this.interface.deploy.inputs.length,
          ' in Contract constructor'
        ),
        (e.data = k(b([this.bytecode, this.interface.encodeDeploy(t)]))),
        e
      );
    }
    deploy(...t) {
      return Zs(this, void 0, void 0, function* () {
        let e = {};
        t.length === this.interface.deploy.inputs.length + 1 && (e = t.pop()),
          Xs.checkArgumentCount(
            t.length,
            this.interface.deploy.inputs.length,
            ' in Contract constructor'
          );
        const r = yield ea(this.signer, t, this.interface.deploy.inputs);
        r.push(e);
        const n = this.getDeployTransaction(...r),
          i = yield this.signer.sendTransaction(n),
          o = at(this.constructor, 'getContractAddress')(i),
          s = at(this.constructor, 'getContract')(
            o,
            this.interface,
            this.signer
          );
        return na(s, i), st(s, 'deployTransaction', i), s;
      });
    }
    attach(t) {
      return this.constructor.getContract(t, this.interface, this.signer);
    }
    connect(t) {
      return new this.constructor(this.interface, this.bytecode, t);
    }
    static fromSolidity(t, e) {
      null == t &&
        Xs.throwError('missing compiler output', f.errors.MISSING_ARGUMENT, {
          argument: 'compilerOutput',
        }),
        'string' == typeof t && (t = JSON.parse(t));
      const r = t.abi;
      let n = null;
      return (
        t.bytecode
          ? (n = t.bytecode)
          : t.evm && t.evm.bytecode && (n = t.evm.bytecode),
        new this(r, n, e)
      );
    }
    static getInterface(t) {
      return fa.getInterface(t);
    }
    static getContractAddress(t) {
      return Ie(t);
    }
    static getContract(t, e, r) {
      return new fa(t, e, r);
    }
  }
  class ma {
    constructor(t) {
      st(this, 'alphabet', t),
        st(this, 'base', t.length),
        st(this, '_alphabetMap', {}),
        st(this, '_leader', t.charAt(0));
      for (let e = 0; e < t.length; e++) this._alphabetMap[t.charAt(e)] = e;
    }
    encode(t) {
      let e = w(t);
      if (0 === e.length) return '';
      let r = [0];
      for (let t = 0; t < e.length; ++t) {
        let n = e[t];
        for (let t = 0; t < r.length; ++t)
          (n += r[t] << 8), (r[t] = n % this.base), (n = (n / this.base) | 0);
        for (; n > 0; ) r.push(n % this.base), (n = (n / this.base) | 0);
      }
      let n = '';
      for (let t = 0; 0 === e[t] && t < e.length - 1; ++t) n += this._leader;
      for (let t = r.length - 1; t >= 0; --t) n += this.alphabet[r[t]];
      return n;
    }
    decode(t) {
      if ('string' != typeof t) throw new TypeError('Expected String');
      let e = [];
      if (0 === t.length) return new Uint8Array(e);
      e.push(0);
      for (let r = 0; r < t.length; r++) {
        let n = this._alphabetMap[t[r]];
        if (void 0 === n)
          throw new Error('Non-base' + this.base + ' character');
        let i = n;
        for (let t = 0; t < e.length; ++t)
          (i += e[t] * this.base), (e[t] = 255 & i), (i >>= 8);
        for (; i > 0; ) e.push(255 & i), (i >>= 8);
      }
      for (let r = 0; t[r] === this._leader && r < t.length - 1; ++r) e.push(0);
      return w(new Uint8Array(e.reverse()));
    }
  }
  new ma('abcdefghijklmnopqrstuvwxyz234567');
  const pa = new ma(
    '123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz'
  );
  var ga;
  !(function (t) {
    (t.sha256 = 'sha256'), (t.sha512 = 'sha512');
  })(ga || (ga = {}));
  const ya = new f('sha2/5.7.0');
  function va(t) {
    return '0x' + Kn.ripemd160().update(w(t)).digest('hex');
  }
  function wa(t) {
    return '0x' + Kn.sha256().update(w(t)).digest('hex');
  }
  function ba(t, e, r) {
    return (
      ga[t] ||
        ya.throwError(
          'unsupported algorithm ' + t,
          f.errors.UNSUPPORTED_OPERATION,
          { operation: 'hmac', algorithm: t }
        ),
      '0x' + Kn.hmac(Kn[t], w(e)).update(w(r)).digest('hex')
    );
  }
  function Aa(t, e, r, n, i) {
    let o;
    (t = w(t)), (e = w(e));
    let s = 1;
    const a = new Uint8Array(n),
      u = new Uint8Array(e.length + 4);
    let h, l;
    u.set(e);
    for (let c = 1; c <= s; c++) {
      (u[e.length] = (c >> 24) & 255),
        (u[e.length + 1] = (c >> 16) & 255),
        (u[e.length + 2] = (c >> 8) & 255),
        (u[e.length + 3] = 255 & c);
      let f = w(ba(i, t, u));
      o ||
        ((o = f.length),
        (l = new Uint8Array(o)),
        (s = Math.ceil(n / o)),
        (h = n - (s - 1) * o)),
        l.set(f);
      for (let e = 1; e < r; e++) {
        f = w(ba(i, t, f));
        for (let t = 0; t < o; t++) l[t] ^= f[t];
      }
      const d = (c - 1) * o,
        m = c === s ? h : o;
      a.set(w(l).slice(0, m), d);
    }
    return k(a);
  }
  const Ea = new f('wordlists/5.7.0');
  class Ma {
    constructor(t) {
      Ea.checkAbstract(new.target, Ma), st(this, 'locale', t);
    }
    split(t) {
      return t.toLowerCase().split(/ +/g);
    }
    join(t) {
      return t.join(' ');
    }
    static check(t) {
      const e = [];
      for (let r = 0; r < 2048; r++) {
        const n = t.getWord(r);
        if (r !== t.getWordIndex(n)) return '0x';
        e.push(n);
      }
      return Tr(e.join('\n') + '\n');
    }
    static register(t, e) {
      e || (e = t.locale);
    }
  }
  let _a = null;
  function ka(t) {
    if (
      null == _a &&
      ((_a =
        'AbandonAbilityAbleAboutAboveAbsentAbsorbAbstractAbsurdAbuseAccessAccidentAccountAccuseAchieveAcidAcousticAcquireAcrossActActionActorActressActualAdaptAddAddictAddressAdjustAdmitAdultAdvanceAdviceAerobicAffairAffordAfraidAgainAgeAgentAgreeAheadAimAirAirportAisleAlarmAlbumAlcoholAlertAlienAllAlleyAllowAlmostAloneAlphaAlreadyAlsoAlterAlwaysAmateurAmazingAmongAmountAmusedAnalystAnchorAncientAngerAngleAngryAnimalAnkleAnnounceAnnualAnotherAnswerAntennaAntiqueAnxietyAnyApartApologyAppearAppleApproveAprilArchArcticAreaArenaArgueArmArmedArmorArmyAroundArrangeArrestArriveArrowArtArtefactArtistArtworkAskAspectAssaultAssetAssistAssumeAsthmaAthleteAtomAttackAttendAttitudeAttractAuctionAuditAugustAuntAuthorAutoAutumnAverageAvocadoAvoidAwakeAwareAwayAwesomeAwfulAwkwardAxisBabyBachelorBaconBadgeBagBalanceBalconyBallBambooBananaBannerBarBarelyBargainBarrelBaseBasicBasketBattleBeachBeanBeautyBecauseBecomeBeefBeforeBeginBehaveBehindBelieveBelowBeltBenchBenefitBestBetrayBetterBetweenBeyondBicycleBidBikeBindBiologyBirdBirthBitterBlackBladeBlameBlanketBlastBleakBlessBlindBloodBlossomBlouseBlueBlurBlushBoardBoatBodyBoilBombBoneBonusBookBoostBorderBoringBorrowBossBottomBounceBoxBoyBracketBrainBrandBrassBraveBreadBreezeBrickBridgeBriefBrightBringBriskBroccoliBrokenBronzeBroomBrotherBrownBrushBubbleBuddyBudgetBuffaloBuildBulbBulkBulletBundleBunkerBurdenBurgerBurstBusBusinessBusyButterBuyerBuzzCabbageCabinCableCactusCageCakeCallCalmCameraCampCanCanalCancelCandyCannonCanoeCanvasCanyonCapableCapitalCaptainCarCarbonCardCargoCarpetCarryCartCaseCashCasinoCastleCasualCatCatalogCatchCategoryCattleCaughtCauseCautionCaveCeilingCeleryCementCensusCenturyCerealCertainChairChalkChampionChangeChaosChapterChargeChaseChatCheapCheckCheeseChefCherryChestChickenChiefChildChimneyChoiceChooseChronicChuckleChunkChurnCigarCinnamonCircleCitizenCityCivilClaimClapClarifyClawClayCleanClerkCleverClickClientCliffClimbClinicClipClockClogCloseClothCloudClownClubClumpClusterClutchCoachCoastCoconutCodeCoffeeCoilCoinCollectColorColumnCombineComeComfortComicCommonCompanyConcertConductConfirmCongressConnectConsiderControlConvinceCookCoolCopperCopyCoralCoreCornCorrectCostCottonCouchCountryCoupleCourseCousinCoverCoyoteCrackCradleCraftCramCraneCrashCraterCrawlCrazyCreamCreditCreekCrewCricketCrimeCrispCriticCropCrossCrouchCrowdCrucialCruelCruiseCrumbleCrunchCrushCryCrystalCubeCultureCupCupboardCuriousCurrentCurtainCurveCushionCustomCuteCycleDadDamageDampDanceDangerDaringDashDaughterDawnDayDealDebateDebrisDecadeDecemberDecideDeclineDecorateDecreaseDeerDefenseDefineDefyDegreeDelayDeliverDemandDemiseDenialDentistDenyDepartDependDepositDepthDeputyDeriveDescribeDesertDesignDeskDespairDestroyDetailDetectDevelopDeviceDevoteDiagramDialDiamondDiaryDiceDieselDietDifferDigitalDignityDilemmaDinnerDinosaurDirectDirtDisagreeDiscoverDiseaseDishDismissDisorderDisplayDistanceDivertDivideDivorceDizzyDoctorDocumentDogDollDolphinDomainDonateDonkeyDonorDoorDoseDoubleDoveDraftDragonDramaDrasticDrawDreamDressDriftDrillDrinkDripDriveDropDrumDryDuckDumbDuneDuringDustDutchDutyDwarfDynamicEagerEagleEarlyEarnEarthEasilyEastEasyEchoEcologyEconomyEdgeEditEducateEffortEggEightEitherElbowElderElectricElegantElementElephantElevatorEliteElseEmbarkEmbodyEmbraceEmergeEmotionEmployEmpowerEmptyEnableEnactEndEndlessEndorseEnemyEnergyEnforceEngageEngineEnhanceEnjoyEnlistEnoughEnrichEnrollEnsureEnterEntireEntryEnvelopeEpisodeEqualEquipEraEraseErodeErosionErrorEruptEscapeEssayEssenceEstateEternalEthicsEvidenceEvilEvokeEvolveExactExampleExcessExchangeExciteExcludeExcuseExecuteExerciseExhaustExhibitExileExistExitExoticExpandExpectExpireExplainExposeExpressExtendExtraEyeEyebrowFabricFaceFacultyFadeFaintFaithFallFalseFameFamilyFamousFanFancyFantasyFarmFashionFatFatalFatherFatigueFaultFavoriteFeatureFebruaryFederalFeeFeedFeelFemaleFenceFestivalFetchFeverFewFiberFictionFieldFigureFileFilmFilterFinalFindFineFingerFinishFireFirmFirstFiscalFishFitFitnessFixFlagFlameFlashFlatFlavorFleeFlightFlipFloatFlockFloorFlowerFluidFlushFlyFoamFocusFogFoilFoldFollowFoodFootForceForestForgetForkFortuneForumForwardFossilFosterFoundFoxFragileFrameFrequentFreshFriendFringeFrogFrontFrostFrownFrozenFruitFuelFunFunnyFurnaceFuryFutureGadgetGainGalaxyGalleryGameGapGarageGarbageGardenGarlicGarmentGasGaspGateGatherGaugeGazeGeneralGeniusGenreGentleGenuineGestureGhostGiantGiftGiggleGingerGiraffeGirlGiveGladGlanceGlareGlassGlideGlimpseGlobeGloomGloryGloveGlowGlueGoatGoddessGoldGoodGooseGorillaGospelGossipGovernGownGrabGraceGrainGrantGrapeGrassGravityGreatGreenGridGriefGritGroceryGroupGrowGruntGuardGuessGuideGuiltGuitarGunGymHabitHairHalfHammerHamsterHandHappyHarborHardHarshHarvestHatHaveHawkHazardHeadHealthHeartHeavyHedgehogHeightHelloHelmetHelpHenHeroHiddenHighHillHintHipHireHistoryHobbyHockeyHoldHoleHolidayHollowHomeHoneyHoodHopeHornHorrorHorseHospitalHostHotelHourHoverHubHugeHumanHumbleHumorHundredHungryHuntHurdleHurryHurtHusbandHybridIceIconIdeaIdentifyIdleIgnoreIllIllegalIllnessImageImitateImmenseImmuneImpactImposeImproveImpulseInchIncludeIncomeIncreaseIndexIndicateIndoorIndustryInfantInflictInformInhaleInheritInitialInjectInjuryInmateInnerInnocentInputInquiryInsaneInsectInsideInspireInstallIntactInterestIntoInvestInviteInvolveIronIslandIsolateIssueItemIvoryJacketJaguarJarJazzJealousJeansJellyJewelJobJoinJokeJourneyJoyJudgeJuiceJumpJungleJuniorJunkJustKangarooKeenKeepKetchupKeyKickKidKidneyKindKingdomKissKitKitchenKiteKittenKiwiKneeKnifeKnockKnowLabLabelLaborLadderLadyLakeLampLanguageLaptopLargeLaterLatinLaughLaundryLavaLawLawnLawsuitLayerLazyLeaderLeafLearnLeaveLectureLeftLegLegalLegendLeisureLemonLendLengthLensLeopardLessonLetterLevelLiarLibertyLibraryLicenseLifeLiftLightLikeLimbLimitLinkLionLiquidListLittleLiveLizardLoadLoanLobsterLocalLockLogicLonelyLongLoopLotteryLoudLoungeLoveLoyalLuckyLuggageLumberLunarLunchLuxuryLyricsMachineMadMagicMagnetMaidMailMainMajorMakeMammalManManageMandateMangoMansionManualMapleMarbleMarchMarginMarineMarketMarriageMaskMassMasterMatchMaterialMathMatrixMatterMaximumMazeMeadowMeanMeasureMeatMechanicMedalMediaMelodyMeltMemberMemoryMentionMenuMercyMergeMeritMerryMeshMessageMetalMethodMiddleMidnightMilkMillionMimicMindMinimumMinorMinuteMiracleMirrorMiseryMissMistakeMixMixedMixtureMobileModelModifyMomMomentMonitorMonkeyMonsterMonthMoonMoralMoreMorningMosquitoMotherMotionMotorMountainMouseMoveMovieMuchMuffinMuleMultiplyMuscleMuseumMushroomMusicMustMutualMyselfMysteryMythNaiveNameNapkinNarrowNastyNationNatureNearNeckNeedNegativeNeglectNeitherNephewNerveNestNetNetworkNeutralNeverNewsNextNiceNightNobleNoiseNomineeNoodleNormalNorthNoseNotableNoteNothingNoticeNovelNowNuclearNumberNurseNutOakObeyObjectObligeObscureObserveObtainObviousOccurOceanOctoberOdorOffOfferOfficeOftenOilOkayOldOliveOlympicOmitOnceOneOnionOnlineOnlyOpenOperaOpinionOpposeOptionOrangeOrbitOrchardOrderOrdinaryOrganOrientOriginalOrphanOstrichOtherOutdoorOuterOutputOutsideOvalOvenOverOwnOwnerOxygenOysterOzonePactPaddlePagePairPalacePalmPandaPanelPanicPantherPaperParadeParentParkParrotPartyPassPatchPathPatientPatrolPatternPausePavePaymentPeacePeanutPearPeasantPelicanPenPenaltyPencilPeoplePepperPerfectPermitPersonPetPhonePhotoPhrasePhysicalPianoPicnicPicturePiecePigPigeonPillPilotPinkPioneerPipePistolPitchPizzaPlacePlanetPlasticPlatePlayPleasePledgePluckPlugPlungePoemPoetPointPolarPolePolicePondPonyPoolPopularPortionPositionPossiblePostPotatoPotteryPovertyPowderPowerPracticePraisePredictPreferPreparePresentPrettyPreventPricePridePrimaryPrintPriorityPrisonPrivatePrizeProblemProcessProduceProfitProgramProjectPromoteProofPropertyProsperProtectProudProvidePublicPuddingPullPulpPulsePumpkinPunchPupilPuppyPurchasePurityPurposePursePushPutPuzzlePyramidQualityQuantumQuarterQuestionQuickQuitQuizQuoteRabbitRaccoonRaceRackRadarRadioRailRainRaiseRallyRampRanchRandomRangeRapidRareRateRatherRavenRawRazorReadyRealReasonRebelRebuildRecallReceiveRecipeRecordRecycleReduceReflectReformRefuseRegionRegretRegularRejectRelaxReleaseReliefRelyRemainRememberRemindRemoveRenderRenewRentReopenRepairRepeatReplaceReportRequireRescueResembleResistResourceResponseResultRetireRetreatReturnReunionRevealReviewRewardRhythmRibRibbonRiceRichRideRidgeRifleRightRigidRingRiotRippleRiskRitualRivalRiverRoadRoastRobotRobustRocketRomanceRoofRookieRoomRoseRotateRoughRoundRouteRoyalRubberRudeRugRuleRunRunwayRuralSadSaddleSadnessSafeSailSaladSalmonSalonSaltSaluteSameSampleSandSatisfySatoshiSauceSausageSaveSayScaleScanScareScatterSceneSchemeSchoolScienceScissorsScorpionScoutScrapScreenScriptScrubSeaSearchSeasonSeatSecondSecretSectionSecuritySeedSeekSegmentSelectSellSeminarSeniorSenseSentenceSeriesServiceSessionSettleSetupSevenShadowShaftShallowShareShedShellSheriffShieldShiftShineShipShiverShockShoeShootShopShortShoulderShoveShrimpShrugShuffleShySiblingSickSideSiegeSightSignSilentSilkSillySilverSimilarSimpleSinceSingSirenSisterSituateSixSizeSkateSketchSkiSkillSkinSkirtSkullSlabSlamSleepSlenderSliceSlideSlightSlimSloganSlotSlowSlushSmallSmartSmileSmokeSmoothSnackSnakeSnapSniffSnowSoapSoccerSocialSockSodaSoftSolarSoldierSolidSolutionSolveSomeoneSongSoonSorrySortSoulSoundSoupSourceSouthSpaceSpareSpatialSpawnSpeakSpecialSpeedSpellSpendSphereSpiceSpiderSpikeSpinSpiritSplitSpoilSponsorSpoonSportSpotSpraySpreadSpringSpySquareSqueezeSquirrelStableStadiumStaffStageStairsStampStandStartStateStaySteakSteelStemStepStereoStickStillStingStockStomachStoneStoolStoryStoveStrategyStreetStrikeStrongStruggleStudentStuffStumbleStyleSubjectSubmitSubwaySuccessSuchSuddenSufferSugarSuggestSuitSummerSunSunnySunsetSuperSupplySupremeSureSurfaceSurgeSurpriseSurroundSurveySuspectSustainSwallowSwampSwapSwarmSwearSweetSwiftSwimSwingSwitchSwordSymbolSymptomSyrupSystemTableTackleTagTailTalentTalkTankTapeTargetTaskTasteTattooTaxiTeachTeamTellTenTenantTennisTentTermTestTextThankThatThemeThenTheoryThereTheyThingThisThoughtThreeThriveThrowThumbThunderTicketTideTigerTiltTimberTimeTinyTipTiredTissueTitleToastTobaccoTodayToddlerToeTogetherToiletTokenTomatoTomorrowToneTongueTonightToolToothTopTopicToppleTorchTornadoTortoiseTossTotalTouristTowardTowerTownToyTrackTradeTrafficTragicTrainTransferTrapTrashTravelTrayTreatTreeTrendTrialTribeTrickTriggerTrimTripTrophyTroubleTruckTrueTrulyTrumpetTrustTruthTryTubeTuitionTumbleTunaTunnelTurkeyTurnTurtleTwelveTwentyTwiceTwinTwistTwoTypeTypicalUglyUmbrellaUnableUnawareUncleUncoverUnderUndoUnfairUnfoldUnhappyUniformUniqueUnitUniverseUnknownUnlockUntilUnusualUnveilUpdateUpgradeUpholdUponUpperUpsetUrbanUrgeUsageUseUsedUsefulUselessUsualUtilityVacantVacuumVagueValidValleyValveVanVanishVaporVariousVastVaultVehicleVelvetVendorVentureVenueVerbVerifyVersionVeryVesselVeteranViableVibrantViciousVictoryVideoViewVillageVintageViolinVirtualVirusVisaVisitVisualVitalVividVocalVoiceVoidVolcanoVolumeVoteVoyageWageWagonWaitWalkWallWalnutWantWarfareWarmWarriorWashWaspWasteWaterWaveWayWealthWeaponWearWeaselWeatherWebWeddingWeekendWeirdWelcomeWestWetWhaleWhatWheatWheelWhenWhereWhipWhisperWideWidthWifeWildWillWinWindowWineWingWinkWinnerWinterWireWisdomWiseWishWitnessWolfWomanWonderWoodWoolWordWorkWorldWorryWorthWrapWreckWrestleWristWriteWrongYardYearYellowYouYoungYouthZebraZeroZoneZoo'
          .replace(/([A-Z])/g, ' $1')
          .toLowerCase()
          .substring(1)
          .split(' ')),
      '0x3c8acc1e7b08d8e76f9fda015ef48dc8c710a73cb7e0f77b2c18a9b5a7adde60' !==
        Ma.check(t))
    )
      throw ((_a = null), new Error('BIP39 Wordlist for en (English) FAILED'));
  }
  const Sa = new (class extends Ma {
    constructor() {
      super('en');
    }
    getWord(t) {
      return ka(this), _a[t];
    }
    getWordIndex(t) {
      return ka(this), _a.indexOf(t);
    }
  })();
  Ma.register(Sa);
  const Na = { en: Sa },
    Pa = new f('hdnode/5.7.0'),
    xa = j.from(
      '0xfffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141'
    ),
    Ta = hr('Bitcoin seed'),
    Ca = 2147483648;
  function Ra(t) {
    return ((1 << t) - 1) << (8 - t);
  }
  function Ia(t) {
    return C(k(t), 32);
  }
  function Ba(t) {
    return pa.encode(b([t, N(wa(wa(t)), 0, 4)]));
  }
  function Oa(t) {
    if (null == t) return Na.en;
    if ('string' == typeof t) {
      const e = Na[t];
      return (
        null == e && Pa.throwArgumentError('unknown locale', 'wordlist', t), e
      );
    }
    return t;
  }
  const Fa = {},
    Da = "m/44'/60'/0'/0/0";
  class La {
    constructor(t, e, r, n, i, o, s, a) {
      if (t !== Fa)
        throw new Error('HDNode constructor cannot be called directly');
      if (e) {
        const t = new Cs(e);
        st(this, 'privateKey', t.privateKey),
          st(this, 'publicKey', t.compressedPublicKey);
      } else st(this, 'privateKey', null), st(this, 'publicKey', k(r));
      st(this, 'parentFingerprint', n),
        st(this, 'fingerprint', N(va(wa(this.publicKey)), 0, 4)),
        st(this, 'address', js(this.publicKey)),
        st(this, 'chainCode', i),
        st(this, 'index', o),
        st(this, 'depth', s),
        null == a
          ? (st(this, 'mnemonic', null), st(this, 'path', null))
          : 'string' == typeof a
          ? (st(this, 'mnemonic', null), st(this, 'path', a))
          : (st(this, 'mnemonic', a), st(this, 'path', a.path));
    }
    get extendedKey() {
      if (this.depth >= 256) throw new Error('Depth too large!');
      return Ba(
        b([
          null != this.privateKey ? '0x0488ADE4' : '0x0488B21E',
          k(this.depth),
          this.parentFingerprint,
          C(k(this.index), 4),
          this.chainCode,
          null != this.privateKey
            ? b(['0x00', this.privateKey])
            : this.publicKey,
        ])
      );
    }
    neuter() {
      return new La(
        Fa,
        null,
        this.publicKey,
        this.parentFingerprint,
        this.chainCode,
        this.index,
        this.depth,
        this.path
      );
    }
    _derive(t) {
      if (t > 4294967295) throw new Error('invalid index - ' + String(t));
      let e = this.path;
      e && (e += '/' + (2147483647 & t));
      const r = new Uint8Array(37);
      if (t & Ca) {
        if (!this.privateKey)
          throw new Error('cannot derive child of neutered node');
        r.set(w(this.privateKey), 1), e && (e += "'");
      } else r.set(w(this.publicKey));
      for (let e = 24; e >= 0; e -= 8) r[33 + (e >> 3)] = (t >> (24 - e)) & 255;
      const n = w(ba(ga.sha512, this.chainCode, r)),
        i = n.slice(0, 32),
        o = n.slice(32);
      let s = null,
        a = null;
      if (this.privateKey) s = Ia(j.from(i).add(this.privateKey).mod(xa));
      else {
        a = new Cs(k(i))._addPoint(this.publicKey);
      }
      let u = e;
      const h = this.mnemonic;
      return (
        h &&
          (u = Object.freeze({
            phrase: h.phrase,
            path: e,
            locale: h.locale || 'en',
          })),
        new La(Fa, s, a, this.fingerprint, Ia(o), t, this.depth + 1, u)
      );
    }
    derivePath(t) {
      const e = t.split('/');
      if (0 === e.length || ('m' === e[0] && 0 !== this.depth))
        throw new Error('invalid path - ' + t);
      'm' === e[0] && e.shift();
      let r = this;
      for (let t = 0; t < e.length; t++) {
        const n = e[t];
        if (n.match(/^[0-9]+'$/)) {
          const t = parseInt(n.substring(0, n.length - 1));
          if (t >= Ca) throw new Error('invalid path index - ' + n);
          r = r._derive(Ca + t);
        } else {
          if (!n.match(/^[0-9]+$/))
            throw new Error('invalid path component - ' + n);
          {
            const t = parseInt(n);
            if (t >= Ca) throw new Error('invalid path index - ' + n);
            r = r._derive(t);
          }
        }
      }
      return r;
    }
    static _fromSeed(t, e) {
      const r = w(t);
      if (r.length < 16 || r.length > 64) throw new Error('invalid seed');
      const n = w(ba(ga.sha512, Ta, r));
      return new La(
        Fa,
        Ia(n.slice(0, 32)),
        null,
        '0x00000000',
        Ia(n.slice(32)),
        0,
        0,
        e
      );
    }
    static fromMnemonic(t, e, r) {
      return (
        (t = za(ja(t, (r = Oa(r))), r)),
        La._fromSeed(Ua(t, e), { phrase: t, path: 'm', locale: r.locale })
      );
    }
    static fromSeed(t) {
      return La._fromSeed(t, null);
    }
    static fromExtendedKey(t) {
      const e = pa.decode(t);
      (82 === e.length && Ba(e.slice(0, 78)) === t) ||
        Pa.throwArgumentError(
          'invalid extended key',
          'extendedKey',
          '[REDACTED]'
        );
      const r = e[4],
        n = k(e.slice(5, 9)),
        i = parseInt(k(e.slice(9, 13)).substring(2), 16),
        o = k(e.slice(13, 45)),
        s = e.slice(45, 78);
      switch (k(e.slice(0, 4))) {
        case '0x0488b21e':
        case '0x043587cf':
          return new La(Fa, null, k(s), n, o, i, r, null);
        case '0x0488ade4':
        case '0x04358394 ':
          if (0 !== s[0]) break;
          return new La(Fa, k(s.slice(1)), null, n, o, i, r, null);
      }
      return Pa.throwArgumentError(
        'invalid extended key',
        'extendedKey',
        '[REDACTED]'
      );
    }
  }
  function Ua(t, e) {
    e || (e = '');
    const r = hr('mnemonic' + e, ir.NFKD);
    return Aa(hr(t, ir.NFKD), r, 2048, 64, 'sha512');
  }
  function ja(t, e) {
    (e = Oa(e)), Pa.checkNormalize();
    const r = e.split(t);
    if (r.length % 3 != 0) throw new Error('invalid mnemonic');
    const n = w(new Uint8Array(Math.ceil((11 * r.length) / 8)));
    let i = 0;
    for (let t = 0; t < r.length; t++) {
      let o = e.getWordIndex(r[t].normalize('NFKD'));
      if (-1 === o) throw new Error('invalid mnemonic');
      for (let t = 0; t < 11; t++)
        o & (1 << (10 - t)) && (n[i >> 3] |= 1 << (7 - (i % 8))), i++;
    }
    const o = (32 * r.length) / 3,
      s = Ra(r.length / 3);
    if ((w(wa(n.slice(0, o / 8)))[0] & s) !== (n[n.length - 1] & s))
      throw new Error('invalid checksum');
    return k(n.slice(0, o / 8));
  }
  function za(t, e) {
    if (
      ((e = Oa(e)),
      (t = w(t)).length % 4 != 0 || t.length < 16 || t.length > 32)
    )
      throw new Error('invalid entropy');
    const r = [0];
    let n = 11;
    for (let e = 0; e < t.length; e++)
      n > 8
        ? ((r[r.length - 1] <<= 8), (r[r.length - 1] |= t[e]), (n -= 8))
        : ((r[r.length - 1] <<= n),
          (r[r.length - 1] |= t[e] >> (8 - n)),
          r.push(t[e] & ((1 << (8 - n)) - 1)),
          (n += 3));
    const i = t.length / 4,
      o = w(wa(t))[0] & Ra(i);
    return (
      (r[r.length - 1] <<= i),
      (r[r.length - 1] |= o >> (8 - i)),
      e.join(r.map((t) => e.getWord(t)))
    );
  }
  const Ga = new f('random/5.7.0');
  const qa = (function () {
    if ('undefined' != typeof self) return self;
    if ('undefined' != typeof window) return window;
    if (void 0 !== Kt) return Kt;
    throw new Error('unable to locate global object');
  })();
  let Ka = qa.crypto || qa.msCrypto;
  function Ha(t) {
    (t <= 0 || t > 1024 || t % 1 || t != t) &&
      Ga.throwArgumentError('invalid length', 'length', t);
    const e = new Uint8Array(t);
    return Ka.getRandomValues(e), w(e);
  }
  function Va(t) {
    for (let e = (t = t.slice()).length - 1; e > 0; e--) {
      const r = Math.floor(Math.random() * (e + 1)),
        n = t[e];
      (t[e] = t[r]), (t[r] = n);
    }
    return t;
  }
  (Ka && Ka.getRandomValues) ||
    (Ga.warn('WARNING: Missing strong random number source'),
    (Ka = {
      getRandomValues: function (t) {
        return Ga.throwError(
          'no secure random source avaialble',
          f.errors.UNSUPPORTED_OPERATION,
          { operation: 'crypto.getRandomValues' }
        );
      },
    }));
  var Ja = { exports: {} };
  !(function (t) {
    !(function () {
      function e(t) {
        return parseInt(t) === t;
      }
      function r(t) {
        if (!e(t.length)) return !1;
        for (var r = 0; r < t.length; r++)
          if (!e(t[r]) || t[r] < 0 || t[r] > 255) return !1;
        return !0;
      }
      function n(t, n) {
        if (t.buffer && ArrayBuffer.isView(t) && 'Uint8Array' === t.name)
          return (
            n && (t = t.slice ? t.slice() : Array.prototype.slice.call(t)), t
          );
        if (Array.isArray(t)) {
          if (!r(t)) throw new Error('Array contains invalid value: ' + t);
          return new Uint8Array(t);
        }
        if (e(t.length) && r(t)) return new Uint8Array(t);
        throw new Error('unsupported array-like object');
      }
      function i(t) {
        return new Uint8Array(t);
      }
      function o(t, e, r, n, i) {
        (null == n && null == i) ||
          (t = t.slice ? t.slice(n, i) : Array.prototype.slice.call(t, n, i)),
          e.set(t, r);
      }
      var s,
        a = {
          toBytes: function (t) {
            var e = [],
              r = 0;
            for (t = encodeURI(t); r < t.length; ) {
              var i = t.charCodeAt(r++);
              37 === i
                ? (e.push(parseInt(t.substr(r, 2), 16)), (r += 2))
                : e.push(i);
            }
            return n(e);
          },
          fromBytes: function (t) {
            for (var e = [], r = 0; r < t.length; ) {
              var n = t[r];
              n < 128
                ? (e.push(String.fromCharCode(n)), r++)
                : n > 191 && n < 224
                ? (e.push(
                    String.fromCharCode(((31 & n) << 6) | (63 & t[r + 1]))
                  ),
                  (r += 2))
                : (e.push(
                    String.fromCharCode(
                      ((15 & n) << 12) |
                        ((63 & t[r + 1]) << 6) |
                        (63 & t[r + 2])
                    )
                  ),
                  (r += 3));
            }
            return e.join('');
          },
        },
        u =
          ((s = '0123456789abcdef'),
          {
            toBytes: function (t) {
              for (var e = [], r = 0; r < t.length; r += 2)
                e.push(parseInt(t.substr(r, 2), 16));
              return e;
            },
            fromBytes: function (t) {
              for (var e = [], r = 0; r < t.length; r++) {
                var n = t[r];
                e.push(s[(240 & n) >> 4] + s[15 & n]);
              }
              return e.join('');
            },
          }),
        h = { 16: 10, 24: 12, 32: 14 },
        l = [
          1, 2, 4, 8, 16, 32, 64, 128, 27, 54, 108, 216, 171, 77, 154, 47, 94,
          188, 99, 198, 151, 53, 106, 212, 179, 125, 250, 239, 197, 145,
        ],
        c = [
          99, 124, 119, 123, 242, 107, 111, 197, 48, 1, 103, 43, 254, 215, 171,
          118, 202, 130, 201, 125, 250, 89, 71, 240, 173, 212, 162, 175, 156,
          164, 114, 192, 183, 253, 147, 38, 54, 63, 247, 204, 52, 165, 229, 241,
          113, 216, 49, 21, 4, 199, 35, 195, 24, 150, 5, 154, 7, 18, 128, 226,
          235, 39, 178, 117, 9, 131, 44, 26, 27, 110, 90, 160, 82, 59, 214, 179,
          41, 227, 47, 132, 83, 209, 0, 237, 32, 252, 177, 91, 106, 203, 190,
          57, 74, 76, 88, 207, 208, 239, 170, 251, 67, 77, 51, 133, 69, 249, 2,
          127, 80, 60, 159, 168, 81, 163, 64, 143, 146, 157, 56, 245, 188, 182,
          218, 33, 16, 255, 243, 210, 205, 12, 19, 236, 95, 151, 68, 23, 196,
          167, 126, 61, 100, 93, 25, 115, 96, 129, 79, 220, 34, 42, 144, 136,
          70, 238, 184, 20, 222, 94, 11, 219, 224, 50, 58, 10, 73, 6, 36, 92,
          194, 211, 172, 98, 145, 149, 228, 121, 231, 200, 55, 109, 141, 213,
          78, 169, 108, 86, 244, 234, 101, 122, 174, 8, 186, 120, 37, 46, 28,
          166, 180, 198, 232, 221, 116, 31, 75, 189, 139, 138, 112, 62, 181,
          102, 72, 3, 246, 14, 97, 53, 87, 185, 134, 193, 29, 158, 225, 248,
          152, 17, 105, 217, 142, 148, 155, 30, 135, 233, 206, 85, 40, 223, 140,
          161, 137, 13, 191, 230, 66, 104, 65, 153, 45, 15, 176, 84, 187, 22,
        ],
        f = [
          82, 9, 106, 213, 48, 54, 165, 56, 191, 64, 163, 158, 129, 243, 215,
          251, 124, 227, 57, 130, 155, 47, 255, 135, 52, 142, 67, 68, 196, 222,
          233, 203, 84, 123, 148, 50, 166, 194, 35, 61, 238, 76, 149, 11, 66,
          250, 195, 78, 8, 46, 161, 102, 40, 217, 36, 178, 118, 91, 162, 73,
          109, 139, 209, 37, 114, 248, 246, 100, 134, 104, 152, 22, 212, 164,
          92, 204, 93, 101, 182, 146, 108, 112, 72, 80, 253, 237, 185, 218, 94,
          21, 70, 87, 167, 141, 157, 132, 144, 216, 171, 0, 140, 188, 211, 10,
          247, 228, 88, 5, 184, 179, 69, 6, 208, 44, 30, 143, 202, 63, 15, 2,
          193, 175, 189, 3, 1, 19, 138, 107, 58, 145, 17, 65, 79, 103, 220, 234,
          151, 242, 207, 206, 240, 180, 230, 115, 150, 172, 116, 34, 231, 173,
          53, 133, 226, 249, 55, 232, 28, 117, 223, 110, 71, 241, 26, 113, 29,
          41, 197, 137, 111, 183, 98, 14, 170, 24, 190, 27, 252, 86, 62, 75,
          198, 210, 121, 32, 154, 219, 192, 254, 120, 205, 90, 244, 31, 221,
          168, 51, 136, 7, 199, 49, 177, 18, 16, 89, 39, 128, 236, 95, 96, 81,
          127, 169, 25, 181, 74, 13, 45, 229, 122, 159, 147, 201, 156, 239, 160,
          224, 59, 77, 174, 42, 245, 176, 200, 235, 187, 60, 131, 83, 153, 97,
          23, 43, 4, 126, 186, 119, 214, 38, 225, 105, 20, 99, 85, 33, 12, 125,
        ],
        d = [
          3328402341, 4168907908, 4000806809, 4135287693, 4294111757,
          3597364157, 3731845041, 2445657428, 1613770832, 33620227, 3462883241,
          1445669757, 3892248089, 3050821474, 1303096294, 3967186586,
          2412431941, 528646813, 2311702848, 4202528135, 4026202645, 2992200171,
          2387036105, 4226871307, 1101901292, 3017069671, 1604494077,
          1169141738, 597466303, 1403299063, 3832705686, 2613100635, 1974974402,
          3791519004, 1033081774, 1277568618, 1815492186, 2118074177,
          4126668546, 2211236943, 1748251740, 1369810420, 3521504564,
          4193382664, 3799085459, 2883115123, 1647391059, 706024767, 134480908,
          2512897874, 1176707941, 2646852446, 806885416, 932615841, 168101135,
          798661301, 235341577, 605164086, 461406363, 3756188221, 3454790438,
          1311188841, 2142417613, 3933566367, 302582043, 495158174, 1479289972,
          874125870, 907746093, 3698224818, 3025820398, 1537253627, 2756858614,
          1983593293, 3084310113, 2108928974, 1378429307, 3722699582,
          1580150641, 327451799, 2790478837, 3117535592, 0, 3253595436,
          1075847264, 3825007647, 2041688520, 3059440621, 3563743934,
          2378943302, 1740553945, 1916352843, 2487896798, 2555137236,
          2958579944, 2244988746, 3151024235, 3320835882, 1336584933,
          3992714006, 2252555205, 2588757463, 1714631509, 293963156, 2319795663,
          3925473552, 67240454, 4269768577, 2689618160, 2017213508, 631218106,
          1269344483, 2723238387, 1571005438, 2151694528, 93294474, 1066570413,
          563977660, 1882732616, 4059428100, 1673313503, 2008463041, 2950355573,
          1109467491, 537923632, 3858759450, 4260623118, 3218264685, 2177748300,
          403442708, 638784309, 3287084079, 3193921505, 899127202, 2286175436,
          773265209, 2479146071, 1437050866, 4236148354, 2050833735, 3362022572,
          3126681063, 840505643, 3866325909, 3227541664, 427917720, 2655997905,
          2749160575, 1143087718, 1412049534, 999329963, 193497219, 2353415882,
          3354324521, 1807268051, 672404540, 2816401017, 3160301282, 369822493,
          2916866934, 3688947771, 1681011286, 1949973070, 336202270, 2454276571,
          201721354, 1210328172, 3093060836, 2680341085, 3184776046, 1135389935,
          3294782118, 965841320, 831886756, 3554993207, 4068047243, 3588745010,
          2345191491, 1849112409, 3664604599, 26054028, 2983581028, 2622377682,
          1235855840, 3630984372, 2891339514, 4092916743, 3488279077,
          3395642799, 4101667470, 1202630377, 268961816, 1874508501, 4034427016,
          1243948399, 1546530418, 941366308, 1470539505, 1941222599, 2546386513,
          3421038627, 2715671932, 3899946140, 1042226977, 2521517021,
          1639824860, 227249030, 260737669, 3765465232, 2084453954, 1907733956,
          3429263018, 2420656344, 100860677, 4160157185, 470683154, 3261161891,
          1781871967, 2924959737, 1773779408, 394692241, 2579611992, 974986535,
          664706745, 3655459128, 3958962195, 731420851, 571543859, 3530123707,
          2849626480, 126783113, 865375399, 765172662, 1008606754, 361203602,
          3387549984, 2278477385, 2857719295, 1344809080, 2782912378, 59542671,
          1503764984, 160008576, 437062935, 1707065306, 3622233649, 2218934982,
          3496503480, 2185314755, 697932208, 1512910199, 504303377, 2075177163,
          2824099068, 1841019862, 739644986,
        ],
        m = [
          2781242211, 2230877308, 2582542199, 2381740923, 234877682, 3184946027,
          2984144751, 1418839493, 1348481072, 50462977, 2848876391, 2102799147,
          434634494, 1656084439, 3863849899, 2599188086, 1167051466, 2636087938,
          1082771913, 2281340285, 368048890, 3954334041, 3381544775, 201060592,
          3963727277, 1739838676, 4250903202, 3930435503, 3206782108,
          4149453988, 2531553906, 1536934080, 3262494647, 484572669, 2923271059,
          1783375398, 1517041206, 1098792767, 49674231, 1334037708, 1550332980,
          4098991525, 886171109, 150598129, 2481090929, 1940642008, 1398944049,
          1059722517, 201851908, 1385547719, 1699095331, 1587397571, 674240536,
          2704774806, 252314885, 3039795866, 151914247, 908333586, 2602270848,
          1038082786, 651029483, 1766729511, 3447698098, 2682942837, 454166793,
          2652734339, 1951935532, 775166490, 758520603, 3000790638, 4004797018,
          4217086112, 4137964114, 1299594043, 1639438038, 3464344499,
          2068982057, 1054729187, 1901997871, 2534638724, 4121318227,
          1757008337, 0, 750906861, 1614815264, 535035132, 3363418545,
          3988151131, 3201591914, 1183697867, 3647454910, 1265776953,
          3734260298, 3566750796, 3903871064, 1250283471, 1807470800, 717615087,
          3847203498, 384695291, 3313910595, 3617213773, 1432761139, 2484176261,
          3481945413, 283769337, 100925954, 2180939647, 4037038160, 1148730428,
          3123027871, 3813386408, 4087501137, 4267549603, 3229630528,
          2315620239, 2906624658, 3156319645, 1215313976, 82966005, 3747855548,
          3245848246, 1974459098, 1665278241, 807407632, 451280895, 251524083,
          1841287890, 1283575245, 337120268, 891687699, 801369324, 3787349855,
          2721421207, 3431482436, 959321879, 1469301956, 4065699751, 2197585534,
          1199193405, 2898814052, 3887750493, 724703513, 2514908019, 2696962144,
          2551808385, 3516813135, 2141445340, 1715741218, 2119445034,
          2872807568, 2198571144, 3398190662, 700968686, 3547052216, 1009259540,
          2041044702, 3803995742, 487983883, 1991105499, 1004265696, 1449407026,
          1316239930, 504629770, 3683797321, 168560134, 1816667172, 3837287516,
          1570751170, 1857934291, 4014189740, 2797888098, 2822345105,
          2754712981, 936633572, 2347923833, 852879335, 1133234376, 1500395319,
          3084545389, 2348912013, 1689376213, 3533459022, 3762923945,
          3034082412, 4205598294, 133428468, 634383082, 2949277029, 2398386810,
          3913789102, 403703816, 3580869306, 2297460856, 1867130149, 1918643758,
          607656988, 4049053350, 3346248884, 1368901318, 600565992, 2090982877,
          2632479860, 557719327, 3717614411, 3697393085, 2249034635, 2232388234,
          2430627952, 1115438654, 3295786421, 2865522278, 3633334344, 84280067,
          33027830, 303828494, 2747425121, 1600795957, 4188952407, 3496589753,
          2434238086, 1486471617, 658119965, 3106381470, 953803233, 334231800,
          3005978776, 857870609, 3151128937, 1890179545, 2298973838, 2805175444,
          3056442267, 574365214, 2450884487, 550103529, 1233637070, 4289353045,
          2018519080, 2057691103, 2399374476, 4166623649, 2148108681, 387583245,
          3664101311, 836232934, 3330556482, 3100665960, 3280093505, 2955516313,
          2002398509, 287182607, 3413881008, 4238890068, 3597515707, 975967766,
        ],
        p = [
          1671808611, 2089089148, 2006576759, 2072901243, 4061003762,
          1807603307, 1873927791, 3310653893, 810573872, 16974337, 1739181671,
          729634347, 4263110654, 3613570519, 2883997099, 1989864566, 3393556426,
          2191335298, 3376449993, 2106063485, 4195741690, 1508618841,
          1204391495, 4027317232, 2917941677, 3563566036, 2734514082,
          2951366063, 2629772188, 2767672228, 1922491506, 3227229120,
          3082974647, 4246528509, 2477669779, 644500518, 911895606, 1061256767,
          4144166391, 3427763148, 878471220, 2784252325, 3845444069, 4043897329,
          1905517169, 3631459288, 827548209, 356461077, 67897348, 3344078279,
          593839651, 3277757891, 405286936, 2527147926, 84871685, 2595565466,
          118033927, 305538066, 2157648768, 3795705826, 3945188843, 661212711,
          2999812018, 1973414517, 152769033, 2208177539, 745822252, 439235610,
          455947803, 1857215598, 1525593178, 2700827552, 1391895634, 994932283,
          3596728278, 3016654259, 695947817, 3812548067, 795958831, 2224493444,
          1408607827, 3513301457, 0, 3979133421, 543178784, 4229948412,
          2982705585, 1542305371, 1790891114, 3410398667, 3201918910, 961245753,
          1256100938, 1289001036, 1491644504, 3477767631, 3496721360,
          4012557807, 2867154858, 4212583931, 1137018435, 1305975373, 861234739,
          2241073541, 1171229253, 4178635257, 33948674, 2139225727, 1357946960,
          1011120188, 2679776671, 2833468328, 1374921297, 2751356323,
          1086357568, 2408187279, 2460827538, 2646352285, 944271416, 4110742005,
          3168756668, 3066132406, 3665145818, 560153121, 271589392, 4279952895,
          4077846003, 3530407890, 3444343245, 202643468, 322250259, 3962553324,
          1608629855, 2543990167, 1154254916, 389623319, 3294073796, 2817676711,
          2122513534, 1028094525, 1689045092, 1575467613, 422261273, 1939203699,
          1621147744, 2174228865, 1339137615, 3699352540, 577127458, 712922154,
          2427141008, 2290289544, 1187679302, 3995715566, 3100863416, 339486740,
          3732514782, 1591917662, 186455563, 3681988059, 3762019296, 844522546,
          978220090, 169743370, 1239126601, 101321734, 611076132, 1558493276,
          3260915650, 3547250131, 2901361580, 1655096418, 2443721105,
          2510565781, 3828863972, 2039214713, 3878868455, 3359869896, 928607799,
          1840765549, 2374762893, 3580146133, 1322425422, 2850048425,
          1823791212, 1459268694, 4094161908, 3928346602, 1706019429,
          2056189050, 2934523822, 135794696, 3134549946, 2022240376, 628050469,
          779246638, 472135708, 2800834470, 3032970164, 3327236038, 3894660072,
          3715932637, 1956440180, 522272287, 1272813131, 3185336765, 2340818315,
          2323976074, 1888542832, 1044544574, 3049550261, 1722469478,
          1222152264, 50660867, 4127324150, 236067854, 1638122081, 895445557,
          1475980887, 3117443513, 2257655686, 3243809217, 489110045, 2662934430,
          3778599393, 4162055160, 2561878936, 288563729, 1773916777, 3648039385,
          2391345038, 2493985684, 2612407707, 505560094, 2274497927, 3911240169,
          3460925390, 1442818645, 678973480, 3749357023, 2358182796, 2717407649,
          2306869641, 219617805, 3218761151, 3862026214, 1120306242, 1756942440,
          1103331905, 2578459033, 762796589, 252780047, 2966125488, 1425844308,
          3151392187, 372911126,
        ],
        g = [
          1667474886, 2088535288, 2004326894, 2071694838, 4075949567,
          1802223062, 1869591006, 3318043793, 808472672, 16843522, 1734846926,
          724270422, 4278065639, 3621216949, 2880169549, 1987484396, 3402253711,
          2189597983, 3385409673, 2105378810, 4210693615, 1499065266,
          1195886990, 4042263547, 2913856577, 3570689971, 2728590687,
          2947541573, 2627518243, 2762274643, 1920112356, 3233831835,
          3082273397, 4261223649, 2475929149, 640051788, 909531756, 1061110142,
          4160160501, 3435941763, 875846760, 2779116625, 3857003729, 4059105529,
          1903268834, 3638064043, 825316194, 353713962, 67374088, 3351728789,
          589522246, 3284360861, 404236336, 2526454071, 84217610, 2593830191,
          117901582, 303183396, 2155911963, 3806477791, 3958056653, 656894286,
          2998062463, 1970642922, 151591698, 2206440989, 741110872, 437923380,
          454765878, 1852748508, 1515908788, 2694904667, 1381168804, 993742198,
          3604373943, 3014905469, 690584402, 3823320797, 791638366, 2223281939,
          1398011302, 3520161977, 0, 3991743681, 538992704, 4244381667,
          2981218425, 1532751286, 1785380564, 3419096717, 3200178535, 960056178,
          1246420628, 1280103576, 1482221744, 3486468741, 3503319995,
          4025428677, 2863326543, 4227536621, 1128514950, 1296947098, 859002214,
          2240123921, 1162203018, 4193849577, 33687044, 2139062782, 1347481760,
          1010582648, 2678045221, 2829640523, 1364325282, 2745433693,
          1077985408, 2408548869, 2459086143, 2644360225, 943212656, 4126475505,
          3166494563, 3065430391, 3671750063, 555836226, 269496352, 4294908645,
          4092792573, 3537006015, 3452783745, 202118168, 320025894, 3974901699,
          1600119230, 2543297077, 1145359496, 387397934, 3301201811, 2812801621,
          2122220284, 1027426170, 1684319432, 1566435258, 421079858, 1936954854,
          1616945344, 2172753945, 1330631070, 3705438115, 572679748, 707427924,
          2425400123, 2290647819, 1179044492, 4008585671, 3099120491, 336870440,
          3739122087, 1583276732, 185277718, 3688593069, 3772791771, 842159716,
          976899700, 168435220, 1229577106, 101059084, 606366792, 1549591736,
          3267517855, 3553849021, 2897014595, 1650632388, 2442242105,
          2509612081, 3840161747, 2038008818, 3890688725, 3368567691, 926374254,
          1835907034, 2374863873, 3587531953, 1313788572, 2846482505,
          1819063512, 1448540844, 4109633523, 3941213647, 1701162954,
          2054852340, 2930698567, 134748176, 3132806511, 2021165296, 623210314,
          774795868, 471606328, 2795958615, 3031746419, 3334885783, 3907527627,
          3722280097, 1953799400, 522133822, 1263263126, 3183336545, 2341176845,
          2324333839, 1886425312, 1044267644, 3048588401, 1718004428,
          1212733584, 50529542, 4143317495, 235803164, 1633788866, 892690282,
          1465383342, 3115962473, 2256965911, 3250673817, 488449850, 2661202215,
          3789633753, 4177007595, 2560144171, 286339874, 1768537042, 3654906025,
          2391705863, 2492770099, 2610673197, 505291324, 2273808917, 3924369609,
          3469625735, 1431699370, 673740880, 3755965093, 2358021891, 2711746649,
          2307489801, 218961690, 3217021541, 3873845719, 1111672452, 1751693520,
          1094828930, 2576986153, 757954394, 252645662, 2964376443, 1414855848,
          3149649517, 370555436,
        ],
        y = [
          1374988112, 2118214995, 437757123, 975658646, 1001089995, 530400753,
          2902087851, 1273168787, 540080725, 2910219766, 2295101073, 4110568485,
          1340463100, 3307916247, 641025152, 3043140495, 3736164937, 632953703,
          1172967064, 1576976609, 3274667266, 2169303058, 2370213795,
          1809054150, 59727847, 361929877, 3211623147, 2505202138, 3569255213,
          1484005843, 1239443753, 2395588676, 1975683434, 4102977912,
          2572697195, 666464733, 3202437046, 4035489047, 3374361702, 2110667444,
          1675577880, 3843699074, 2538681184, 1649639237, 2976151520,
          3144396420, 4269907996, 4178062228, 1883793496, 2403728665,
          2497604743, 1383856311, 2876494627, 1917518562, 3810496343,
          1716890410, 3001755655, 800440835, 2261089178, 3543599269, 807962610,
          599762354, 33778362, 3977675356, 2328828971, 2809771154, 4077384432,
          1315562145, 1708848333, 101039829, 3509871135, 3299278474, 875451293,
          2733856160, 92987698, 2767645557, 193195065, 1080094634, 1584504582,
          3178106961, 1042385657, 2531067453, 3711829422, 1306967366,
          2438237621, 1908694277, 67556463, 1615861247, 429456164, 3602770327,
          2302690252, 1742315127, 2968011453, 126454664, 3877198648, 2043211483,
          2709260871, 2084704233, 4169408201, 0, 159417987, 841739592,
          504459436, 1817866830, 4245618683, 260388950, 1034867998, 908933415,
          168810852, 1750902305, 2606453969, 607530554, 202008497, 2472011535,
          3035535058, 463180190, 2160117071, 1641816226, 1517767529, 470948374,
          3801332234, 3231722213, 1008918595, 303765277, 235474187, 4069246893,
          766945465, 337553864, 1475418501, 2943682380, 4003061179, 2743034109,
          4144047775, 1551037884, 1147550661, 1543208500, 2336434550,
          3408119516, 3069049960, 3102011747, 3610369226, 1113818384, 328671808,
          2227573024, 2236228733, 3535486456, 2935566865, 3341394285, 496906059,
          3702665459, 226906860, 2009195472, 733156972, 2842737049, 294930682,
          1206477858, 2835123396, 2700099354, 1451044056, 573804783, 2269728455,
          3644379585, 2362090238, 2564033334, 2801107407, 2776292904,
          3669462566, 1068351396, 742039012, 1350078989, 1784663195, 1417561698,
          4136440770, 2430122216, 775550814, 2193862645, 2673705150, 1775276924,
          1876241833, 3475313331, 3366754619, 270040487, 3902563182, 3678124923,
          3441850377, 1851332852, 3969562369, 2203032232, 3868552805,
          2868897406, 566021896, 4011190502, 3135740889, 1248802510, 3936291284,
          699432150, 832877231, 708780849, 3332740144, 899835584, 1951317047,
          4236429990, 3767586992, 866637845, 4043610186, 1106041591, 2144161806,
          395441711, 1984812685, 1139781709, 3433712980, 3835036895, 2664543715,
          1282050075, 3240894392, 1181045119, 2640243204, 25965917, 4203181171,
          4211818798, 3009879386, 2463879762, 3910161971, 1842759443,
          2597806476, 933301370, 1509430414, 3943906441, 3467192302, 3076639029,
          3776767469, 2051518780, 2631065433, 1441952575, 404016761, 1942435775,
          1408749034, 1610459739, 3745345300, 2017778566, 3400528769,
          3110650942, 941896748, 3265478751, 371049330, 3168937228, 675039627,
          4279080257, 967311729, 135050206, 3635733660, 1683407248, 2076935265,
          3576870512, 1215061108, 3501741890,
        ],
        v = [
          1347548327, 1400783205, 3273267108, 2520393566, 3409685355,
          4045380933, 2880240216, 2471224067, 1428173050, 4138563181,
          2441661558, 636813900, 4233094615, 3620022987, 2149987652, 2411029155,
          1239331162, 1730525723, 2554718734, 3781033664, 46346101, 310463728,
          2743944855, 3328955385, 3875770207, 2501218972, 3955191162,
          3667219033, 768917123, 3545789473, 692707433, 1150208456, 1786102409,
          2029293177, 1805211710, 3710368113, 3065962831, 401639597, 1724457132,
          3028143674, 409198410, 2196052529, 1620529459, 1164071807, 3769721975,
          2226875310, 486441376, 2499348523, 1483753576, 428819965, 2274680428,
          3075636216, 598438867, 3799141122, 1474502543, 711349675, 129166120,
          53458370, 2592523643, 2782082824, 4063242375, 2988687269, 3120694122,
          1559041666, 730517276, 2460449204, 4042459122, 2706270690, 3446004468,
          3573941694, 533804130, 2328143614, 2637442643, 2695033685, 839224033,
          1973745387, 957055980, 2856345839, 106852767, 1371368976, 4181598602,
          1033297158, 2933734917, 1179510461, 3046200461, 91341917, 1862534868,
          4284502037, 605657339, 2547432937, 3431546947, 2003294622, 3182487618,
          2282195339, 954669403, 3682191598, 1201765386, 3917234703, 3388507166,
          0, 2198438022, 1211247597, 2887651696, 1315723890, 4227665663,
          1443857720, 507358933, 657861945, 1678381017, 560487590, 3516619604,
          975451694, 2970356327, 261314535, 3535072918, 2652609425, 1333838021,
          2724322336, 1767536459, 370938394, 182621114, 3854606378, 1128014560,
          487725847, 185469197, 2918353863, 3106780840, 3356761769, 2237133081,
          1286567175, 3152976349, 4255350624, 2683765030, 3160175349,
          3309594171, 878443390, 1988838185, 3704300486, 1756818940, 1673061617,
          3403100636, 272786309, 1075025698, 545572369, 2105887268, 4174560061,
          296679730, 1841768865, 1260232239, 4091327024, 3960309330, 3497509347,
          1814803222, 2578018489, 4195456072, 575138148, 3299409036, 446754879,
          3629546796, 4011996048, 3347532110, 3252238545, 4270639778, 915985419,
          3483825537, 681933534, 651868046, 2755636671, 3828103837, 223377554,
          2607439820, 1649704518, 3270937875, 3901806776, 1580087799,
          4118987695, 3198115200, 2087309459, 2842678573, 3016697106,
          1003007129, 2802849917, 1860738147, 2077965243, 164439672, 4100872472,
          32283319, 2827177882, 1709610350, 2125135846, 136428751, 3874428392,
          3652904859, 3460984630, 3572145929, 3593056380, 2939266226, 824852259,
          818324884, 3224740454, 930369212, 2801566410, 2967507152, 355706840,
          1257309336, 4148292826, 243256656, 790073846, 2373340630, 1296297904,
          1422699085, 3756299780, 3818836405, 457992840, 3099667487, 2135319889,
          77422314, 1560382517, 1945798516, 788204353, 1521706781, 1385356242,
          870912086, 325965383, 2358957921, 2050466060, 2388260884, 2313884476,
          4006521127, 901210569, 3990953189, 1014646705, 1503449823, 1062597235,
          2031621326, 3212035895, 3931371469, 1533017514, 350174575, 2256028891,
          2177544179, 1052338372, 741876788, 1606591296, 1914052035, 213705253,
          2334669897, 1107234197, 1899603969, 3725069491, 2631447780,
          2422494913, 1635502980, 1893020342, 1950903388, 1120974935,
        ],
        w = [
          2807058932, 1699970625, 2764249623, 1586903591, 1808481195,
          1173430173, 1487645946, 59984867, 4199882800, 1844882806, 1989249228,
          1277555970, 3623636965, 3419915562, 1149249077, 2744104290,
          1514790577, 459744698, 244860394, 3235995134, 1963115311, 4027744588,
          2544078150, 4190530515, 1608975247, 2627016082, 2062270317,
          1507497298, 2200818878, 567498868, 1764313568, 3359936201, 2305455554,
          2037970062, 1047239e3, 1910319033, 1337376481, 2904027272, 2892417312,
          984907214, 1243112415, 830661914, 861968209, 2135253587, 2011214180,
          2927934315, 2686254721, 731183368, 1750626376, 4246310725, 1820824798,
          4172763771, 3542330227, 48394827, 2404901663, 2871682645, 671593195,
          3254988725, 2073724613, 145085239, 2280796200, 2779915199, 1790575107,
          2187128086, 472615631, 3029510009, 4075877127, 3802222185, 4107101658,
          3201631749, 1646252340, 4270507174, 1402811438, 1436590835,
          3778151818, 3950355702, 3963161475, 4020912224, 2667994737, 273792366,
          2331590177, 104699613, 95345982, 3175501286, 2377486676, 1560637892,
          3564045318, 369057872, 4213447064, 3919042237, 1137477952, 2658625497,
          1119727848, 2340947849, 1530455833, 4007360968, 172466556, 266959938,
          516552836, 0, 2256734592, 3980931627, 1890328081, 1917742170,
          4294704398, 945164165, 3575528878, 958871085, 3647212047, 2787207260,
          1423022939, 775562294, 1739656202, 3876557655, 2530391278, 2443058075,
          3310321856, 547512796, 1265195639, 437656594, 3121275539, 719700128,
          3762502690, 387781147, 218828297, 3350065803, 2830708150, 2848461854,
          428169201, 122466165, 3720081049, 1627235199, 648017665, 4122762354,
          1002783846, 2117360635, 695634755, 3336358691, 4234721005, 4049844452,
          3704280881, 2232435299, 574624663, 287343814, 612205898, 1039717051,
          840019705, 2708326185, 793451934, 821288114, 1391201670, 3822090177,
          376187827, 3113855344, 1224348052, 1679968233, 2361698556, 1058709744,
          752375421, 2431590963, 1321699145, 3519142200, 2734591178, 188127444,
          2177869557, 3727205754, 2384911031, 3215212461, 2648976442,
          2450346104, 3432737375, 1180849278, 331544205, 3102249176, 4150144569,
          2952102595, 2159976285, 2474404304, 766078933, 313773861, 2570832044,
          2108100632, 1668212892, 3145456443, 2013908262, 418672217, 3070356634,
          2594734927, 1852171925, 3867060991, 3473416636, 3907448597,
          2614737639, 919489135, 164948639, 2094410160, 2997825956, 590424639,
          2486224549, 1723872674, 3157750862, 3399941250, 3501252752,
          3625268135, 2555048196, 3673637356, 1343127501, 4130281361,
          3599595085, 2957853679, 1297403050, 81781910, 3051593425, 2283490410,
          532201772, 1367295589, 3926170974, 895287692, 1953757831, 1093597963,
          492483431, 3528626907, 1446242576, 1192455638, 1636604631, 209336225,
          344873464, 1015671571, 669961897, 3375740769, 3857572124, 2973530695,
          3747192018, 1933530610, 3464042516, 935293895, 3454686199, 2858115069,
          1863638845, 3683022916, 4085369519, 3292445032, 875313188, 1080017571,
          3279033885, 621591778, 1233856572, 2504130317, 24197544, 3017672716,
          3835484340, 3247465558, 2220981195, 3060847922, 1551124588,
          1463996600,
        ],
        b = [
          4104605777, 1097159550, 396673818, 660510266, 2875968315, 2638606623,
          4200115116, 3808662347, 821712160, 1986918061, 3430322568, 38544885,
          3856137295, 718002117, 893681702, 1654886325, 2975484382, 3122358053,
          3926825029, 4274053469, 796197571, 1290801793, 1184342925, 3556361835,
          2405426947, 2459735317, 1836772287, 1381620373, 3196267988,
          1948373848, 3764988233, 3385345166, 3263785589, 2390325492,
          1480485785, 3111247143, 3780097726, 2293045232, 548169417, 3459953789,
          3746175075, 439452389, 1362321559, 1400849762, 1685577905, 1806599355,
          2174754046, 137073913, 1214797936, 1174215055, 3731654548, 2079897426,
          1943217067, 1258480242, 529487843, 1437280870, 3945269170, 3049390895,
          3313212038, 923313619, 679998e3, 3215307299, 57326082, 377642221,
          3474729866, 2041877159, 133361907, 1776460110, 3673476453, 96392454,
          878845905, 2801699524, 777231668, 4082475170, 2330014213, 4142626212,
          2213296395, 1626319424, 1906247262, 1846563261, 562755902, 3708173718,
          1040559837, 3871163981, 1418573201, 3294430577, 114585348, 1343618912,
          2566595609, 3186202582, 1078185097, 3651041127, 3896688048,
          2307622919, 425408743, 3371096953, 2081048481, 1108339068, 2216610296,
          0, 2156299017, 736970802, 292596766, 1517440620, 251657213,
          2235061775, 2933202493, 758720310, 265905162, 1554391400, 1532285339,
          908999204, 174567692, 1474760595, 4002861748, 2610011675, 3234156416,
          3693126241, 2001430874, 303699484, 2478443234, 2687165888, 585122620,
          454499602, 151849742, 2345119218, 3064510765, 514443284, 4044981591,
          1963412655, 2581445614, 2137062819, 19308535, 1928707164, 1715193156,
          4219352155, 1126790795, 600235211, 3992742070, 3841024952, 836553431,
          1669664834, 2535604243, 3323011204, 1243905413, 3141400786,
          4180808110, 698445255, 2653899549, 2989552604, 2253581325, 3252932727,
          3004591147, 1891211689, 2487810577, 3915653703, 4237083816,
          4030667424, 2100090966, 865136418, 1229899655, 953270745, 3399679628,
          3557504664, 4118925222, 2061379749, 3079546586, 2915017791, 983426092,
          2022837584, 1607244650, 2118541908, 2366882550, 3635996816, 972512814,
          3283088770, 1568718495, 3499326569, 3576539503, 621982671, 2895723464,
          410887952, 2623762152, 1002142683, 645401037, 1494807662, 2595684844,
          1335535747, 2507040230, 4293295786, 3167684641, 367585007, 3885750714,
          1865862730, 2668221674, 2960971305, 2763173681, 1059270954,
          2777952454, 2724642869, 1320957812, 2194319100, 2429595872,
          2815956275, 77089521, 3973773121, 3444575871, 2448830231, 1305906550,
          4021308739, 2857194700, 2516901860, 3518358430, 1787304780, 740276417,
          1699839814, 1592394909, 2352307457, 2272556026, 188821243, 1729977011,
          3687994002, 274084841, 3594982253, 3613494426, 2701949495, 4162096729,
          322734571, 2837966542, 1640576439, 484830689, 1202797690, 3537852828,
          4067639125, 349075736, 3342319475, 4157467219, 4255800159, 1030690015,
          1155237496, 2951971274, 1757691577, 607398968, 2738905026, 499347990,
          3794078908, 1011452712, 227885567, 2818666809, 213114376, 3034881240,
          1455525988, 3414450555, 850817237, 1817998408, 3092726480,
        ],
        A = [
          0, 235474187, 470948374, 303765277, 941896748, 908933415, 607530554,
          708780849, 1883793496, 2118214995, 1817866830, 1649639237, 1215061108,
          1181045119, 1417561698, 1517767529, 3767586992, 4003061179,
          4236429990, 4069246893, 3635733660, 3602770327, 3299278474,
          3400528769, 2430122216, 2664543715, 2362090238, 2193862645,
          2835123396, 2801107407, 3035535058, 3135740889, 3678124923,
          3576870512, 3341394285, 3374361702, 3810496343, 3977675356,
          4279080257, 4043610186, 2876494627, 2776292904, 3076639029,
          3110650942, 2472011535, 2640243204, 2403728665, 2169303058,
          1001089995, 899835584, 666464733, 699432150, 59727847, 226906860,
          530400753, 294930682, 1273168787, 1172967064, 1475418501, 1509430414,
          1942435775, 2110667444, 1876241833, 1641816226, 2910219766,
          2743034109, 2976151520, 3211623147, 2505202138, 2606453969,
          2302690252, 2269728455, 3711829422, 3543599269, 3240894392,
          3475313331, 3843699074, 3943906441, 4178062228, 4144047775,
          1306967366, 1139781709, 1374988112, 1610459739, 1975683434,
          2076935265, 1775276924, 1742315127, 1034867998, 866637845, 566021896,
          800440835, 92987698, 193195065, 429456164, 395441711, 1984812685,
          2017778566, 1784663195, 1683407248, 1315562145, 1080094634,
          1383856311, 1551037884, 101039829, 135050206, 437757123, 337553864,
          1042385657, 807962610, 573804783, 742039012, 2531067453, 2564033334,
          2328828971, 2227573024, 2935566865, 2700099354, 3001755655,
          3168937228, 3868552805, 3902563182, 4203181171, 4102977912,
          3736164937, 3501741890, 3265478751, 3433712980, 1106041591,
          1340463100, 1576976609, 1408749034, 2043211483, 2009195472,
          1708848333, 1809054150, 832877231, 1068351396, 766945465, 599762354,
          159417987, 126454664, 361929877, 463180190, 2709260871, 2943682380,
          3178106961, 3009879386, 2572697195, 2538681184, 2236228733,
          2336434550, 3509871135, 3745345300, 3441850377, 3274667266,
          3910161971, 3877198648, 4110568485, 4211818798, 2597806476,
          2497604743, 2261089178, 2295101073, 2733856160, 2902087851,
          3202437046, 2968011453, 3936291284, 3835036895, 4136440770,
          4169408201, 3535486456, 3702665459, 3467192302, 3231722213,
          2051518780, 1951317047, 1716890410, 1750902305, 1113818384,
          1282050075, 1584504582, 1350078989, 168810852, 67556463, 371049330,
          404016761, 841739592, 1008918595, 775550814, 540080725, 3969562369,
          3801332234, 4035489047, 4269907996, 3569255213, 3669462566,
          3366754619, 3332740144, 2631065433, 2463879762, 2160117071,
          2395588676, 2767645557, 2868897406, 3102011747, 3069049960, 202008497,
          33778362, 270040487, 504459436, 875451293, 975658646, 675039627,
          641025152, 2084704233, 1917518562, 1615861247, 1851332852, 1147550661,
          1248802510, 1484005843, 1451044056, 933301370, 967311729, 733156972,
          632953703, 260388950, 25965917, 328671808, 496906059, 1206477858,
          1239443753, 1543208500, 1441952575, 2144161806, 1908694277,
          1675577880, 1842759443, 3610369226, 3644379585, 3408119516,
          3307916247, 4011190502, 3776767469, 4077384432, 4245618683,
          2809771154, 2842737049, 3144396420, 3043140495, 2673705150,
          2438237621, 2203032232, 2370213795,
        ],
        E = [
          0, 185469197, 370938394, 487725847, 741876788, 657861945, 975451694,
          824852259, 1483753576, 1400783205, 1315723890, 1164071807, 1950903388,
          2135319889, 1649704518, 1767536459, 2967507152, 3152976349,
          2801566410, 2918353863, 2631447780, 2547432937, 2328143614,
          2177544179, 3901806776, 3818836405, 4270639778, 4118987695,
          3299409036, 3483825537, 3535072918, 3652904859, 2077965243,
          1893020342, 1841768865, 1724457132, 1474502543, 1559041666,
          1107234197, 1257309336, 598438867, 681933534, 901210569, 1052338372,
          261314535, 77422314, 428819965, 310463728, 3409685355, 3224740454,
          3710368113, 3593056380, 3875770207, 3960309330, 4045380933,
          4195456072, 2471224067, 2554718734, 2237133081, 2388260884,
          3212035895, 3028143674, 2842678573, 2724322336, 4138563181,
          4255350624, 3769721975, 3955191162, 3667219033, 3516619604,
          3431546947, 3347532110, 2933734917, 2782082824, 3099667487,
          3016697106, 2196052529, 2313884476, 2499348523, 2683765030,
          1179510461, 1296297904, 1347548327, 1533017514, 1786102409,
          1635502980, 2087309459, 2003294622, 507358933, 355706840, 136428751,
          53458370, 839224033, 957055980, 605657339, 790073846, 2373340630,
          2256028891, 2607439820, 2422494913, 2706270690, 2856345839,
          3075636216, 3160175349, 3573941694, 3725069491, 3273267108,
          3356761769, 4181598602, 4063242375, 4011996048, 3828103837,
          1033297158, 915985419, 730517276, 545572369, 296679730, 446754879,
          129166120, 213705253, 1709610350, 1860738147, 1945798516, 2029293177,
          1239331162, 1120974935, 1606591296, 1422699085, 4148292826,
          4233094615, 3781033664, 3931371469, 3682191598, 3497509347,
          3446004468, 3328955385, 2939266226, 2755636671, 3106780840,
          2988687269, 2198438022, 2282195339, 2501218972, 2652609425,
          1201765386, 1286567175, 1371368976, 1521706781, 1805211710,
          1620529459, 2105887268, 1988838185, 533804130, 350174575, 164439672,
          46346101, 870912086, 954669403, 636813900, 788204353, 2358957921,
          2274680428, 2592523643, 2441661558, 2695033685, 2880240216,
          3065962831, 3182487618, 3572145929, 3756299780, 3270937875,
          3388507166, 4174560061, 4091327024, 4006521127, 3854606378,
          1014646705, 930369212, 711349675, 560487590, 272786309, 457992840,
          106852767, 223377554, 1678381017, 1862534868, 1914052035, 2031621326,
          1211247597, 1128014560, 1580087799, 1428173050, 32283319, 182621114,
          401639597, 486441376, 768917123, 651868046, 1003007129, 818324884,
          1503449823, 1385356242, 1333838021, 1150208456, 1973745387,
          2125135846, 1673061617, 1756818940, 2970356327, 3120694122,
          2802849917, 2887651696, 2637442643, 2520393566, 2334669897,
          2149987652, 3917234703, 3799141122, 4284502037, 4100872472,
          3309594171, 3460984630, 3545789473, 3629546796, 2050466060,
          1899603969, 1814803222, 1730525723, 1443857720, 1560382517,
          1075025698, 1260232239, 575138148, 692707433, 878443390, 1062597235,
          243256656, 91341917, 409198410, 325965383, 3403100636, 3252238545,
          3704300486, 3620022987, 3874428392, 3990953189, 4042459122,
          4227665663, 2460449204, 2578018489, 2226875310, 2411029155,
          3198115200, 3046200461, 2827177882, 2743944855,
        ],
        M = [
          0, 218828297, 437656594, 387781147, 875313188, 958871085, 775562294,
          590424639, 1750626376, 1699970625, 1917742170, 2135253587, 1551124588,
          1367295589, 1180849278, 1265195639, 3501252752, 3720081049,
          3399941250, 3350065803, 3835484340, 3919042237, 4270507174,
          4085369519, 3102249176, 3051593425, 2734591178, 2952102595,
          2361698556, 2177869557, 2530391278, 2614737639, 3145456443,
          3060847922, 2708326185, 2892417312, 2404901663, 2187128086,
          2504130317, 2555048196, 3542330227, 3727205754, 3375740769,
          3292445032, 3876557655, 3926170974, 4246310725, 4027744588,
          1808481195, 1723872674, 1910319033, 2094410160, 1608975247,
          1391201670, 1173430173, 1224348052, 59984867, 244860394, 428169201,
          344873464, 935293895, 984907214, 766078933, 547512796, 1844882806,
          1627235199, 2011214180, 2062270317, 1507497298, 1423022939,
          1137477952, 1321699145, 95345982, 145085239, 532201772, 313773861,
          830661914, 1015671571, 731183368, 648017665, 3175501286, 2957853679,
          2807058932, 2858115069, 2305455554, 2220981195, 2474404304,
          2658625497, 3575528878, 3625268135, 3473416636, 3254988725,
          3778151818, 3963161475, 4213447064, 4130281361, 3599595085,
          3683022916, 3432737375, 3247465558, 3802222185, 4020912224,
          4172763771, 4122762354, 3201631749, 3017672716, 2764249623,
          2848461854, 2331590177, 2280796200, 2431590963, 2648976442, 104699613,
          188127444, 472615631, 287343814, 840019705, 1058709744, 671593195,
          621591778, 1852171925, 1668212892, 1953757831, 2037970062, 1514790577,
          1463996600, 1080017571, 1297403050, 3673637356, 3623636965,
          3235995134, 3454686199, 4007360968, 3822090177, 4107101658,
          4190530515, 2997825956, 3215212461, 2830708150, 2779915199,
          2256734592, 2340947849, 2627016082, 2443058075, 172466556, 122466165,
          273792366, 492483431, 1047239e3, 861968209, 612205898, 695634755,
          1646252340, 1863638845, 2013908262, 1963115311, 1446242576,
          1530455833, 1277555970, 1093597963, 1636604631, 1820824798,
          2073724613, 1989249228, 1436590835, 1487645946, 1337376481,
          1119727848, 164948639, 81781910, 331544205, 516552836, 1039717051,
          821288114, 669961897, 719700128, 2973530695, 3157750862, 2871682645,
          2787207260, 2232435299, 2283490410, 2667994737, 2450346104,
          3647212047, 3564045318, 3279033885, 3464042516, 3980931627,
          3762502690, 4150144569, 4199882800, 3070356634, 3121275539,
          2904027272, 2686254721, 2200818878, 2384911031, 2570832044,
          2486224549, 3747192018, 3528626907, 3310321856, 3359936201,
          3950355702, 3867060991, 4049844452, 4234721005, 1739656202,
          1790575107, 2108100632, 1890328081, 1402811438, 1586903591,
          1233856572, 1149249077, 266959938, 48394827, 369057872, 418672217,
          1002783846, 919489135, 567498868, 752375421, 209336225, 24197544,
          376187827, 459744698, 945164165, 895287692, 574624663, 793451934,
          1679968233, 1764313568, 2117360635, 1933530610, 1343127501,
          1560637892, 1243112415, 1192455638, 3704280881, 3519142200,
          3336358691, 3419915562, 3907448597, 3857572124, 4075877127,
          4294704398, 3029510009, 3113855344, 2927934315, 2744104290,
          2159976285, 2377486676, 2594734927, 2544078150,
        ],
        _ = [
          0, 151849742, 303699484, 454499602, 607398968, 758720310, 908999204,
          1059270954, 1214797936, 1097159550, 1517440620, 1400849762,
          1817998408, 1699839814, 2118541908, 2001430874, 2429595872,
          2581445614, 2194319100, 2345119218, 3034881240, 3186202582,
          2801699524, 2951971274, 3635996816, 3518358430, 3399679628,
          3283088770, 4237083816, 4118925222, 4002861748, 3885750714,
          1002142683, 850817237, 698445255, 548169417, 529487843, 377642221,
          227885567, 77089521, 1943217067, 2061379749, 1640576439, 1757691577,
          1474760595, 1592394909, 1174215055, 1290801793, 2875968315,
          2724642869, 3111247143, 2960971305, 2405426947, 2253581325,
          2638606623, 2487810577, 3808662347, 3926825029, 4044981591,
          4162096729, 3342319475, 3459953789, 3576539503, 3693126241,
          1986918061, 2137062819, 1685577905, 1836772287, 1381620373,
          1532285339, 1078185097, 1229899655, 1040559837, 923313619, 740276417,
          621982671, 439452389, 322734571, 137073913, 19308535, 3871163981,
          4021308739, 4104605777, 4255800159, 3263785589, 3414450555,
          3499326569, 3651041127, 2933202493, 2815956275, 3167684641,
          3049390895, 2330014213, 2213296395, 2566595609, 2448830231,
          1305906550, 1155237496, 1607244650, 1455525988, 1776460110,
          1626319424, 2079897426, 1928707164, 96392454, 213114376, 396673818,
          514443284, 562755902, 679998e3, 865136418, 983426092, 3708173718,
          3557504664, 3474729866, 3323011204, 4180808110, 4030667424,
          3945269170, 3794078908, 2507040230, 2623762152, 2272556026,
          2390325492, 2975484382, 3092726480, 2738905026, 2857194700,
          3973773121, 3856137295, 4274053469, 4157467219, 3371096953,
          3252932727, 3673476453, 3556361835, 2763173681, 2915017791,
          3064510765, 3215307299, 2156299017, 2307622919, 2459735317,
          2610011675, 2081048481, 1963412655, 1846563261, 1729977011,
          1480485785, 1362321559, 1243905413, 1126790795, 878845905, 1030690015,
          645401037, 796197571, 274084841, 425408743, 38544885, 188821243,
          3613494426, 3731654548, 3313212038, 3430322568, 4082475170,
          4200115116, 3780097726, 3896688048, 2668221674, 2516901860,
          2366882550, 2216610296, 3141400786, 2989552604, 2837966542,
          2687165888, 1202797690, 1320957812, 1437280870, 1554391400,
          1669664834, 1787304780, 1906247262, 2022837584, 265905162, 114585348,
          499347990, 349075736, 736970802, 585122620, 972512814, 821712160,
          2595684844, 2478443234, 2293045232, 2174754046, 3196267988,
          3079546586, 2895723464, 2777952454, 3537852828, 3687994002,
          3234156416, 3385345166, 4142626212, 4293295786, 3841024952,
          3992742070, 174567692, 57326082, 410887952, 292596766, 777231668,
          660510266, 1011452712, 893681702, 1108339068, 1258480242, 1343618912,
          1494807662, 1715193156, 1865862730, 1948373848, 2100090966,
          2701949495, 2818666809, 3004591147, 3122358053, 2235061775,
          2352307457, 2535604243, 2653899549, 3915653703, 3764988233,
          4219352155, 4067639125, 3444575871, 3294430577, 3746175075,
          3594982253, 836553431, 953270745, 600235211, 718002117, 367585007,
          484830689, 133361907, 251657213, 2041877159, 1891211689, 1806599355,
          1654886325, 1568718495, 1418573201, 1335535747, 1184342925,
        ];
      function k(t) {
        for (var e = [], r = 0; r < t.length; r += 4)
          e.push((t[r] << 24) | (t[r + 1] << 16) | (t[r + 2] << 8) | t[r + 3]);
        return e;
      }
      var S = function (t) {
        if (!(this instanceof S))
          throw Error('AES must be instanitated with `new`');
        Object.defineProperty(this, 'key', { value: n(t, !0) }),
          this._prepare();
      };
      (S.prototype._prepare = function () {
        var t = h[this.key.length];
        if (null == t)
          throw new Error('invalid key size (must be 16, 24 or 32 bytes)');
        (this._Ke = []), (this._Kd = []);
        for (var e = 0; e <= t; e++)
          this._Ke.push([0, 0, 0, 0]), this._Kd.push([0, 0, 0, 0]);
        var r,
          n = 4 * (t + 1),
          i = this.key.length / 4,
          o = k(this.key);
        for (e = 0; e < i; e++)
          (r = e >> 2),
            (this._Ke[r][e % 4] = o[e]),
            (this._Kd[t - r][e % 4] = o[e]);
        for (var s, a = 0, u = i; u < n; ) {
          if (
            ((s = o[i - 1]),
            (o[0] ^=
              (c[(s >> 16) & 255] << 24) ^
              (c[(s >> 8) & 255] << 16) ^
              (c[255 & s] << 8) ^
              c[(s >> 24) & 255] ^
              (l[a] << 24)),
            (a += 1),
            8 != i)
          )
            for (e = 1; e < i; e++) o[e] ^= o[e - 1];
          else {
            for (e = 1; e < i / 2; e++) o[e] ^= o[e - 1];
            (s = o[i / 2 - 1]),
              (o[i / 2] ^=
                c[255 & s] ^
                (c[(s >> 8) & 255] << 8) ^
                (c[(s >> 16) & 255] << 16) ^
                (c[(s >> 24) & 255] << 24));
            for (e = i / 2 + 1; e < i; e++) o[e] ^= o[e - 1];
          }
          for (e = 0; e < i && u < n; )
            (f = u >> 2),
              (d = u % 4),
              (this._Ke[f][d] = o[e]),
              (this._Kd[t - f][d] = o[e++]),
              u++;
        }
        for (var f = 1; f < t; f++)
          for (var d = 0; d < 4; d++)
            (s = this._Kd[f][d]),
              (this._Kd[f][d] =
                A[(s >> 24) & 255] ^
                E[(s >> 16) & 255] ^
                M[(s >> 8) & 255] ^
                _[255 & s]);
      }),
        (S.prototype.encrypt = function (t) {
          if (16 != t.length)
            throw new Error('invalid plaintext size (must be 16 bytes)');
          for (
            var e = this._Ke.length - 1, r = [0, 0, 0, 0], n = k(t), o = 0;
            o < 4;
            o++
          )
            n[o] ^= this._Ke[0][o];
          for (var s = 1; s < e; s++) {
            for (o = 0; o < 4; o++)
              r[o] =
                d[(n[o] >> 24) & 255] ^
                m[(n[(o + 1) % 4] >> 16) & 255] ^
                p[(n[(o + 2) % 4] >> 8) & 255] ^
                g[255 & n[(o + 3) % 4]] ^
                this._Ke[s][o];
            n = r.slice();
          }
          var a,
            u = i(16);
          for (o = 0; o < 4; o++)
            (a = this._Ke[e][o]),
              (u[4 * o] = 255 & (c[(n[o] >> 24) & 255] ^ (a >> 24))),
              (u[4 * o + 1] =
                255 & (c[(n[(o + 1) % 4] >> 16) & 255] ^ (a >> 16))),
              (u[4 * o + 2] =
                255 & (c[(n[(o + 2) % 4] >> 8) & 255] ^ (a >> 8))),
              (u[4 * o + 3] = 255 & (c[255 & n[(o + 3) % 4]] ^ a));
          return u;
        }),
        (S.prototype.decrypt = function (t) {
          if (16 != t.length)
            throw new Error('invalid ciphertext size (must be 16 bytes)');
          for (
            var e = this._Kd.length - 1, r = [0, 0, 0, 0], n = k(t), o = 0;
            o < 4;
            o++
          )
            n[o] ^= this._Kd[0][o];
          for (var s = 1; s < e; s++) {
            for (o = 0; o < 4; o++)
              r[o] =
                y[(n[o] >> 24) & 255] ^
                v[(n[(o + 3) % 4] >> 16) & 255] ^
                w[(n[(o + 2) % 4] >> 8) & 255] ^
                b[255 & n[(o + 1) % 4]] ^
                this._Kd[s][o];
            n = r.slice();
          }
          var a,
            u = i(16);
          for (o = 0; o < 4; o++)
            (a = this._Kd[e][o]),
              (u[4 * o] = 255 & (f[(n[o] >> 24) & 255] ^ (a >> 24))),
              (u[4 * o + 1] =
                255 & (f[(n[(o + 3) % 4] >> 16) & 255] ^ (a >> 16))),
              (u[4 * o + 2] =
                255 & (f[(n[(o + 2) % 4] >> 8) & 255] ^ (a >> 8))),
              (u[4 * o + 3] = 255 & (f[255 & n[(o + 1) % 4]] ^ a));
          return u;
        });
      var N = function (t) {
        if (!(this instanceof N))
          throw Error('AES must be instanitated with `new`');
        (this.description = 'Electronic Code Block'),
          (this.name = 'ecb'),
          (this._aes = new S(t));
      };
      (N.prototype.encrypt = function (t) {
        if ((t = n(t)).length % 16 != 0)
          throw new Error(
            'invalid plaintext size (must be multiple of 16 bytes)'
          );
        for (var e = i(t.length), r = i(16), s = 0; s < t.length; s += 16)
          o(t, r, 0, s, s + 16), o((r = this._aes.encrypt(r)), e, s);
        return e;
      }),
        (N.prototype.decrypt = function (t) {
          if ((t = n(t)).length % 16 != 0)
            throw new Error(
              'invalid ciphertext size (must be multiple of 16 bytes)'
            );
          for (var e = i(t.length), r = i(16), s = 0; s < t.length; s += 16)
            o(t, r, 0, s, s + 16), o((r = this._aes.decrypt(r)), e, s);
          return e;
        });
      var P = function (t, e) {
        if (!(this instanceof P))
          throw Error('AES must be instanitated with `new`');
        if (
          ((this.description = 'Cipher Block Chaining'), (this.name = 'cbc'), e)
        ) {
          if (16 != e.length)
            throw new Error(
              'invalid initialation vector size (must be 16 bytes)'
            );
        } else e = i(16);
        (this._lastCipherblock = n(e, !0)), (this._aes = new S(t));
      };
      (P.prototype.encrypt = function (t) {
        if ((t = n(t)).length % 16 != 0)
          throw new Error(
            'invalid plaintext size (must be multiple of 16 bytes)'
          );
        for (var e = i(t.length), r = i(16), s = 0; s < t.length; s += 16) {
          o(t, r, 0, s, s + 16);
          for (var a = 0; a < 16; a++) r[a] ^= this._lastCipherblock[a];
          (this._lastCipherblock = this._aes.encrypt(r)),
            o(this._lastCipherblock, e, s);
        }
        return e;
      }),
        (P.prototype.decrypt = function (t) {
          if ((t = n(t)).length % 16 != 0)
            throw new Error(
              'invalid ciphertext size (must be multiple of 16 bytes)'
            );
          for (var e = i(t.length), r = i(16), s = 0; s < t.length; s += 16) {
            o(t, r, 0, s, s + 16), (r = this._aes.decrypt(r));
            for (var a = 0; a < 16; a++)
              e[s + a] = r[a] ^ this._lastCipherblock[a];
            o(t, this._lastCipherblock, 0, s, s + 16);
          }
          return e;
        });
      var x = function (t, e, r) {
        if (!(this instanceof x))
          throw Error('AES must be instanitated with `new`');
        if (((this.description = 'Cipher Feedback'), (this.name = 'cfb'), e)) {
          if (16 != e.length)
            throw new Error(
              'invalid initialation vector size (must be 16 size)'
            );
        } else e = i(16);
        r || (r = 1),
          (this.segmentSize = r),
          (this._shiftRegister = n(e, !0)),
          (this._aes = new S(t));
      };
      (x.prototype.encrypt = function (t) {
        if (t.length % this.segmentSize != 0)
          throw new Error('invalid plaintext size (must be segmentSize bytes)');
        for (var e, r = n(t, !0), i = 0; i < r.length; i += this.segmentSize) {
          e = this._aes.encrypt(this._shiftRegister);
          for (var s = 0; s < this.segmentSize; s++) r[i + s] ^= e[s];
          o(this._shiftRegister, this._shiftRegister, 0, this.segmentSize),
            o(
              r,
              this._shiftRegister,
              16 - this.segmentSize,
              i,
              i + this.segmentSize
            );
        }
        return r;
      }),
        (x.prototype.decrypt = function (t) {
          if (t.length % this.segmentSize != 0)
            throw new Error(
              'invalid ciphertext size (must be segmentSize bytes)'
            );
          for (
            var e, r = n(t, !0), i = 0;
            i < r.length;
            i += this.segmentSize
          ) {
            e = this._aes.encrypt(this._shiftRegister);
            for (var s = 0; s < this.segmentSize; s++) r[i + s] ^= e[s];
            o(this._shiftRegister, this._shiftRegister, 0, this.segmentSize),
              o(
                t,
                this._shiftRegister,
                16 - this.segmentSize,
                i,
                i + this.segmentSize
              );
          }
          return r;
        });
      var T = function (t, e) {
        if (!(this instanceof T))
          throw Error('AES must be instanitated with `new`');
        if (((this.description = 'Output Feedback'), (this.name = 'ofb'), e)) {
          if (16 != e.length)
            throw new Error(
              'invalid initialation vector size (must be 16 bytes)'
            );
        } else e = i(16);
        (this._lastPrecipher = n(e, !0)),
          (this._lastPrecipherIndex = 16),
          (this._aes = new S(t));
      };
      (T.prototype.encrypt = function (t) {
        for (var e = n(t, !0), r = 0; r < e.length; r++)
          16 === this._lastPrecipherIndex &&
            ((this._lastPrecipher = this._aes.encrypt(this._lastPrecipher)),
            (this._lastPrecipherIndex = 0)),
            (e[r] ^= this._lastPrecipher[this._lastPrecipherIndex++]);
        return e;
      }),
        (T.prototype.decrypt = T.prototype.encrypt);
      var C = function (t) {
        if (!(this instanceof C))
          throw Error('Counter must be instanitated with `new`');
        0 === t || t || (t = 1),
          'number' == typeof t
            ? ((this._counter = i(16)), this.setValue(t))
            : this.setBytes(t);
      };
      (C.prototype.setValue = function (t) {
        if ('number' != typeof t || parseInt(t) != t)
          throw new Error('invalid counter value (must be an integer)');
        for (var e = 15; e >= 0; --e) (this._counter[e] = t % 256), (t >>= 8);
      }),
        (C.prototype.setBytes = function (t) {
          if (16 != (t = n(t, !0)).length)
            throw new Error('invalid counter bytes size (must be 16 bytes)');
          this._counter = t;
        }),
        (C.prototype.increment = function () {
          for (var t = 15; t >= 0; t--) {
            if (255 !== this._counter[t]) {
              this._counter[t]++;
              break;
            }
            this._counter[t] = 0;
          }
        });
      var R = function (t, e) {
        if (!(this instanceof R))
          throw Error('AES must be instanitated with `new`');
        (this.description = 'Counter'),
          (this.name = 'ctr'),
          e instanceof C || (e = new C(e)),
          (this._counter = e),
          (this._remainingCounter = null),
          (this._remainingCounterIndex = 16),
          (this._aes = new S(t));
      };
      (R.prototype.encrypt = function (t) {
        for (var e = n(t, !0), r = 0; r < e.length; r++)
          16 === this._remainingCounterIndex &&
            ((this._remainingCounter = this._aes.encrypt(
              this._counter._counter
            )),
            (this._remainingCounterIndex = 0),
            this._counter.increment()),
            (e[r] ^= this._remainingCounter[this._remainingCounterIndex++]);
        return e;
      }),
        (R.prototype.decrypt = R.prototype.encrypt);
      var I = {
        AES: S,
        Counter: C,
        ModeOfOperation: { ecb: N, cbc: P, cfb: x, ofb: T, ctr: R },
        utils: { hex: u, utf8: a },
        padding: {
          pkcs7: {
            pad: function (t) {
              var e = 16 - ((t = n(t, !0)).length % 16),
                r = i(t.length + e);
              o(t, r);
              for (var s = t.length; s < r.length; s++) r[s] = e;
              return r;
            },
            strip: function (t) {
              if ((t = n(t, !0)).length < 16)
                throw new Error('PKCS#7 invalid length');
              var e = t[t.length - 1];
              if (e > 16) throw new Error('PKCS#7 padding byte out of range');
              for (var r = t.length - e, s = 0; s < e; s++)
                if (t[r + s] !== e)
                  throw new Error('PKCS#7 invalid padding byte');
              var a = i(r);
              return o(t, a, 0, 0, r), a;
            },
          },
        },
        _arrayTest: { coerceArray: n, createArray: i, copyArray: o },
      };
      t.exports = I;
    })();
  })(Ja);
  var Wa = Ja.exports;
  const Qa = 'json-wallets/5.7.0';
  function Ya(t) {
    return (
      'string' == typeof t && '0x' !== t.substring(0, 2) && (t = '0x' + t), w(t)
    );
  }
  function Za(t, e) {
    for (t = String(t); t.length < e; ) t = '0' + t;
    return t;
  }
  function Xa(t) {
    return 'string' == typeof t ? hr(t, ir.NFKC) : w(t);
  }
  function $a(t, e) {
    let r = t;
    const n = e.toLowerCase().split('/');
    for (let t = 0; t < n.length; t++) {
      let e = null;
      for (const i in r)
        if (i.toLowerCase() === n[t]) {
          e = r[i];
          break;
        }
      if (null === e) return null;
      r = e;
    }
    return r;
  }
  function tu(t) {
    const e = w(t);
    (e[6] = (15 & e[6]) | 64), (e[8] = (63 & e[8]) | 128);
    const r = k(e);
    return [
      r.substring(2, 10),
      r.substring(10, 14),
      r.substring(14, 18),
      r.substring(18, 22),
      r.substring(22, 34),
    ].join('-');
  }
  const eu = new f(Qa);
  class ru extends pt {
    isCrowdsaleAccount(t) {
      return !(!t || !t._isCrowdsaleAccount);
    }
  }
  function nu(t, e) {
    const r = JSON.parse(t);
    e = Xa(e);
    const n = Re($a(r, 'ethaddr')),
      i = Ya($a(r, 'encseed'));
    (i && i.length % 16 == 0) ||
      eu.throwArgumentError('invalid encseed', 'json', t);
    const o = w(Aa(e, e, 2e3, 32, 'sha256')).slice(0, 16),
      s = i.slice(0, 16),
      a = i.slice(16),
      u = new Wa.ModeOfOperation.cbc(o, s),
      h = Wa.padding.pkcs7.strip(w(u.decrypt(a)));
    let l = '';
    for (let t = 0; t < h.length; t++) l += String.fromCharCode(h[t]);
    const c = ye(hr(l));
    return new ru({ _isCrowdsaleAccount: !0, address: n, privateKey: c });
  }
  function iu(t) {
    let e = null;
    try {
      e = JSON.parse(t);
    } catch (t) {
      return !1;
    }
    return e.encseed && e.ethaddr;
  }
  function ou(t) {
    let e = null;
    try {
      e = JSON.parse(t);
    } catch (t) {
      return !1;
    }
    return !(
      !e.version ||
      parseInt(e.version) !== e.version ||
      3 !== parseInt(e.version)
    );
  }
  var su = { exports: {} };
  !(function (t) {
    !(function () {
      const e = 2147483647;
      function r(t) {
        const e = new Uint32Array([
          1116352408, 1899447441, 3049323471, 3921009573, 961987163, 1508970993,
          2453635748, 2870763221, 3624381080, 310598401, 607225278, 1426881987,
          1925078388, 2162078206, 2614888103, 3248222580, 3835390401,
          4022224774, 264347078, 604807628, 770255983, 1249150122, 1555081692,
          1996064986, 2554220882, 2821834349, 2952996808, 3210313671,
          3336571891, 3584528711, 113926993, 338241895, 666307205, 773529912,
          1294757372, 1396182291, 1695183700, 1986661051, 2177026350,
          2456956037, 2730485921, 2820302411, 3259730800, 3345764771,
          3516065817, 3600352804, 4094571909, 275423344, 430227734, 506948616,
          659060556, 883997877, 958139571, 1322822218, 1537002063, 1747873779,
          1955562222, 2024104815, 2227730452, 2361852424, 2428436474,
          2756734187, 3204031479, 3329325298,
        ]);
        let r = 1779033703,
          n = 3144134277,
          i = 1013904242,
          o = 2773480762,
          s = 1359893119,
          a = 2600822924,
          u = 528734635,
          h = 1541459225;
        const l = new Uint32Array(64);
        function c(t) {
          let c = 0,
            f = t.length;
          for (; f >= 64; ) {
            let d,
              m,
              p,
              g,
              y,
              v = r,
              w = n,
              b = i,
              A = o,
              E = s,
              M = a,
              _ = u,
              k = h;
            for (m = 0; m < 16; m++)
              (p = c + 4 * m),
                (l[m] =
                  ((255 & t[p]) << 24) |
                  ((255 & t[p + 1]) << 16) |
                  ((255 & t[p + 2]) << 8) |
                  (255 & t[p + 3]));
            for (m = 16; m < 64; m++)
              (d = l[m - 2]),
                (g =
                  ((d >>> 17) | (d << 15)) ^
                  ((d >>> 19) | (d << 13)) ^
                  (d >>> 10)),
                (d = l[m - 15]),
                (y =
                  ((d >>> 7) | (d << 25)) ^
                  ((d >>> 18) | (d << 14)) ^
                  (d >>> 3)),
                (l[m] = (((g + l[m - 7]) | 0) + ((y + l[m - 16]) | 0)) | 0);
            for (m = 0; m < 64; m++)
              (g =
                ((((((E >>> 6) | (E << 26)) ^
                  ((E >>> 11) | (E << 21)) ^
                  ((E >>> 25) | (E << 7))) +
                  ((E & M) ^ (~E & _))) |
                  0) +
                  ((k + ((e[m] + l[m]) | 0)) | 0)) |
                0),
                (y =
                  ((((v >>> 2) | (v << 30)) ^
                    ((v >>> 13) | (v << 19)) ^
                    ((v >>> 22) | (v << 10))) +
                    ((v & w) ^ (v & b) ^ (w & b))) |
                  0),
                (k = _),
                (_ = M),
                (M = E),
                (E = (A + g) | 0),
                (A = b),
                (b = w),
                (w = v),
                (v = (g + y) | 0);
            (r = (r + v) | 0),
              (n = (n + w) | 0),
              (i = (i + b) | 0),
              (o = (o + A) | 0),
              (s = (s + E) | 0),
              (a = (a + M) | 0),
              (u = (u + _) | 0),
              (h = (h + k) | 0),
              (c += 64),
              (f -= 64);
          }
        }
        c(t);
        let f,
          d = t.length % 64,
          m = (t.length / 536870912) | 0,
          p = t.length << 3,
          g = d < 56 ? 56 : 120,
          y = t.slice(t.length - d, t.length);
        for (y.push(128), f = d + 1; f < g; f++) y.push(0);
        return (
          y.push((m >>> 24) & 255),
          y.push((m >>> 16) & 255),
          y.push((m >>> 8) & 255),
          y.push((m >>> 0) & 255),
          y.push((p >>> 24) & 255),
          y.push((p >>> 16) & 255),
          y.push((p >>> 8) & 255),
          y.push((p >>> 0) & 255),
          c(y),
          [
            (r >>> 24) & 255,
            (r >>> 16) & 255,
            (r >>> 8) & 255,
            (r >>> 0) & 255,
            (n >>> 24) & 255,
            (n >>> 16) & 255,
            (n >>> 8) & 255,
            (n >>> 0) & 255,
            (i >>> 24) & 255,
            (i >>> 16) & 255,
            (i >>> 8) & 255,
            (i >>> 0) & 255,
            (o >>> 24) & 255,
            (o >>> 16) & 255,
            (o >>> 8) & 255,
            (o >>> 0) & 255,
            (s >>> 24) & 255,
            (s >>> 16) & 255,
            (s >>> 8) & 255,
            (s >>> 0) & 255,
            (a >>> 24) & 255,
            (a >>> 16) & 255,
            (a >>> 8) & 255,
            (a >>> 0) & 255,
            (u >>> 24) & 255,
            (u >>> 16) & 255,
            (u >>> 8) & 255,
            (u >>> 0) & 255,
            (h >>> 24) & 255,
            (h >>> 16) & 255,
            (h >>> 8) & 255,
            (h >>> 0) & 255,
          ]
        );
      }
      function n(t, e, n) {
        t = t.length <= 64 ? t : r(t);
        const i = 64 + e.length + 4,
          o = new Array(i),
          s = new Array(64);
        let a,
          u = [];
        for (a = 0; a < 64; a++) o[a] = 54;
        for (a = 0; a < t.length; a++) o[a] ^= t[a];
        for (a = 0; a < e.length; a++) o[64 + a] = e[a];
        for (a = i - 4; a < i; a++) o[a] = 0;
        for (a = 0; a < 64; a++) s[a] = 92;
        for (a = 0; a < t.length; a++) s[a] ^= t[a];
        function h() {
          for (let t = i - 1; t >= i - 4; t--) {
            if ((o[t]++, o[t] <= 255)) return;
            o[t] = 0;
          }
        }
        for (; n >= 32; ) h(), (u = u.concat(r(s.concat(r(o))))), (n -= 32);
        return n > 0 && (h(), (u = u.concat(r(s.concat(r(o))).slice(0, n)))), u;
      }
      function i(t, e, r, n, i) {
        let o;
        for (u(t, 16 * (2 * r - 1), i, 0, 16), o = 0; o < 2 * r; o++)
          a(t, 16 * o, i, 16), s(i, n), u(i, 0, t, e + 16 * o, 16);
        for (o = 0; o < r; o++) u(t, e + 2 * o * 16, t, 16 * o, 16);
        for (o = 0; o < r; o++) u(t, e + 16 * (2 * o + 1), t, 16 * (o + r), 16);
      }
      function o(t, e) {
        return (t << e) | (t >>> (32 - e));
      }
      function s(t, e) {
        u(t, 0, e, 0, 16);
        for (let t = 8; t > 0; t -= 2)
          (e[4] ^= o(e[0] + e[12], 7)),
            (e[8] ^= o(e[4] + e[0], 9)),
            (e[12] ^= o(e[8] + e[4], 13)),
            (e[0] ^= o(e[12] + e[8], 18)),
            (e[9] ^= o(e[5] + e[1], 7)),
            (e[13] ^= o(e[9] + e[5], 9)),
            (e[1] ^= o(e[13] + e[9], 13)),
            (e[5] ^= o(e[1] + e[13], 18)),
            (e[14] ^= o(e[10] + e[6], 7)),
            (e[2] ^= o(e[14] + e[10], 9)),
            (e[6] ^= o(e[2] + e[14], 13)),
            (e[10] ^= o(e[6] + e[2], 18)),
            (e[3] ^= o(e[15] + e[11], 7)),
            (e[7] ^= o(e[3] + e[15], 9)),
            (e[11] ^= o(e[7] + e[3], 13)),
            (e[15] ^= o(e[11] + e[7], 18)),
            (e[1] ^= o(e[0] + e[3], 7)),
            (e[2] ^= o(e[1] + e[0], 9)),
            (e[3] ^= o(e[2] + e[1], 13)),
            (e[0] ^= o(e[3] + e[2], 18)),
            (e[6] ^= o(e[5] + e[4], 7)),
            (e[7] ^= o(e[6] + e[5], 9)),
            (e[4] ^= o(e[7] + e[6], 13)),
            (e[5] ^= o(e[4] + e[7], 18)),
            (e[11] ^= o(e[10] + e[9], 7)),
            (e[8] ^= o(e[11] + e[10], 9)),
            (e[9] ^= o(e[8] + e[11], 13)),
            (e[10] ^= o(e[9] + e[8], 18)),
            (e[12] ^= o(e[15] + e[14], 7)),
            (e[13] ^= o(e[12] + e[15], 9)),
            (e[14] ^= o(e[13] + e[12], 13)),
            (e[15] ^= o(e[14] + e[13], 18));
        for (let r = 0; r < 16; ++r) t[r] += e[r];
      }
      function a(t, e, r, n) {
        for (let i = 0; i < n; i++) r[i] ^= t[e + i];
      }
      function u(t, e, r, n, i) {
        for (; i--; ) r[n++] = t[e++];
      }
      function h(t) {
        if (!t || 'number' != typeof t.length) return !1;
        for (let e = 0; e < t.length; e++) {
          const r = t[e];
          if ('number' != typeof r || r % 1 || r < 0 || r >= 256) return !1;
        }
        return !0;
      }
      function l(t, e) {
        if ('number' != typeof t || t % 1) throw new Error('invalid ' + e);
        return t;
      }
      function c(t, r, o, s, c, f, d) {
        if (
          ((o = l(o, 'N')),
          (s = l(s, 'r')),
          (c = l(c, 'p')),
          (f = l(f, 'dkLen')),
          0 === o || o & (o - 1))
        )
          throw new Error('N must be power of 2');
        if (o > e / 128 / s) throw new Error('N too large');
        if (s > e / 128 / c) throw new Error('r too large');
        if (!h(t)) throw new Error('password must be an array or buffer');
        if (((t = Array.prototype.slice.call(t)), !h(r)))
          throw new Error('salt must be an array or buffer');
        r = Array.prototype.slice.call(r);
        let m = n(t, r, 128 * c * s);
        const p = new Uint32Array(32 * c * s);
        for (let t = 0; t < p.length; t++) {
          const e = 4 * t;
          p[t] =
            ((255 & m[e + 3]) << 24) |
            ((255 & m[e + 2]) << 16) |
            ((255 & m[e + 1]) << 8) |
            (255 & m[e + 0]);
        }
        const g = new Uint32Array(64 * s),
          y = new Uint32Array(32 * s * o),
          v = 32 * s,
          w = new Uint32Array(16),
          b = new Uint32Array(16),
          A = c * o * 2;
        let E,
          M,
          _ = 0,
          k = null,
          S = !1,
          N = 0,
          P = 0;
        const x = d ? parseInt(1e3 / s) : 4294967295,
          T = 'undefined' != typeof setImmediate ? setImmediate : setTimeout,
          C = function () {
            if (S) return d(new Error('cancelled'), _ / A);
            let e;
            switch (N) {
              case 0:
                (M = 32 * P * s), u(p, M, g, 0, v), (N = 1), (E = 0);
              case 1:
                (e = o - E), e > x && (e = x);
                for (let t = 0; t < e; t++)
                  u(g, 0, y, (E + t) * v, v), i(g, v, s, w, b);
                if (((E += e), (_ += e), d)) {
                  const t = parseInt((1e3 * _) / A);
                  if (t !== k) {
                    if (((S = d(null, _ / A)), S)) break;
                    k = t;
                  }
                }
                if (E < o) break;
                (E = 0), (N = 2);
              case 2:
                (e = o - E), e > x && (e = x);
                for (let t = 0; t < e; t++) {
                  const t = g[16 * (2 * s - 1)] & (o - 1);
                  a(y, t * v, g, v), i(g, v, s, w, b);
                }
                if (((E += e), (_ += e), d)) {
                  const t = parseInt((1e3 * _) / A);
                  if (t !== k) {
                    if (((S = d(null, _ / A)), S)) break;
                    k = t;
                  }
                }
                if (E < o) break;
                if ((u(g, 0, p, M, v), P++, P < c)) {
                  N = 0;
                  break;
                }
                m = [];
                for (let t = 0; t < p.length; t++)
                  m.push(255 & p[t]),
                    m.push((p[t] >> 8) & 255),
                    m.push((p[t] >> 16) & 255),
                    m.push((p[t] >> 24) & 255);
                const r = n(t, m, f);
                return d && d(null, 1, r), r;
            }
            d && T(C);
          };
        if (!d)
          for (;;) {
            const t = C();
            if (null != t) return t;
          }
        C();
      }
      const f = {
        scrypt: function (t, e, r, n, i, o, s) {
          return new Promise(function (a, u) {
            let h = 0;
            s && s(0),
              c(t, e, r, n, i, o, function (t, e, r) {
                if (t) u(t);
                else if (r) s && 1 !== h && s(1), a(new Uint8Array(r));
                else if (s && e !== h) return (h = e), s(e);
              });
          });
        },
        syncScrypt: function (t, e, r, n, i, o) {
          return new Uint8Array(c(t, e, r, n, i, o));
        },
      };
      t.exports = f;
    })();
  })(su);
  var au = su.exports,
    uu = function (t, e, r, n) {
      return new (r || (r = Promise))(function (i, o) {
        function s(t) {
          try {
            u(n.next(t));
          } catch (t) {
            o(t);
          }
        }
        function a(t) {
          try {
            u(n.throw(t));
          } catch (t) {
            o(t);
          }
        }
        function u(t) {
          var e;
          t.done
            ? i(t.value)
            : ((e = t.value),
              e instanceof r
                ? e
                : new r(function (t) {
                    t(e);
                  })).then(s, a);
        }
        u((n = n.apply(t, e || [])).next());
      });
    };
  const hu = new f(Qa);
  function lu(t) {
    return null != t && t.mnemonic && t.mnemonic.phrase;
  }
  class cu extends pt {
    isKeystoreAccount(t) {
      return !(!t || !t._isKeystoreAccount);
    }
  }
  function fu(t, e) {
    const r = Ya($a(t, 'crypto/ciphertext'));
    if (
      k(ye(b([e.slice(16, 32), r]))).substring(2) !==
      $a(t, 'crypto/mac').toLowerCase()
    )
      throw new Error('invalid password');
    const n = (function (t, e, r) {
      if ('aes-128-ctr' === $a(t, 'crypto/cipher')) {
        const n = Ya($a(t, 'crypto/cipherparams/iv')),
          i = new Wa.Counter(n);
        return w(new Wa.ModeOfOperation.ctr(e, i).decrypt(r));
      }
      return null;
    })(t, e.slice(0, 16), r);
    n ||
      hu.throwError('unsupported cipher', f.errors.UNSUPPORTED_OPERATION, {
        operation: 'decrypt',
      });
    const i = e.slice(32, 64),
      o = js(n);
    if (t.address) {
      let e = t.address.toLowerCase();
      if (('0x' !== e.substring(0, 2) && (e = '0x' + e), Re(e) !== o))
        throw new Error('address mismatch');
    }
    const s = { _isKeystoreAccount: !0, address: o, privateKey: k(n) };
    if ('0.1' === $a(t, 'x-ethers/version')) {
      const e = Ya($a(t, 'x-ethers/mnemonicCiphertext')),
        r = Ya($a(t, 'x-ethers/mnemonicCounter')),
        n = new Wa.Counter(r),
        o = new Wa.ModeOfOperation.ctr(i, n),
        a = $a(t, 'x-ethers/path') || Da,
        u = $a(t, 'x-ethers/locale') || 'en',
        h = w(o.decrypt(e));
      try {
        const t = za(h, u),
          e = La.fromMnemonic(t, null, u).derivePath(a);
        if (e.privateKey != s.privateKey) throw new Error('mnemonic mismatch');
        s.mnemonic = e.mnemonic;
      } catch (t) {
        if (t.code !== f.errors.INVALID_ARGUMENT || 'wordlist' !== t.argument)
          throw t;
      }
    }
    return new cu(s);
  }
  function du(t, e, r, n, i) {
    return w(Aa(t, e, r, n, i));
  }
  function mu(t, e, r, n, i) {
    return Promise.resolve(du(t, e, r, n, i));
  }
  function pu(t, e, r, n, i) {
    const o = Xa(e),
      s = $a(t, 'crypto/kdf');
    if (s && 'string' == typeof s) {
      const e = function (t, e) {
        return hu.throwArgumentError(
          'invalid key-derivation function parameters',
          t,
          e
        );
      };
      if ('scrypt' === s.toLowerCase()) {
        const r = Ya($a(t, 'crypto/kdfparams/salt')),
          a = parseInt($a(t, 'crypto/kdfparams/n')),
          u = parseInt($a(t, 'crypto/kdfparams/r')),
          h = parseInt($a(t, 'crypto/kdfparams/p'));
        (a && u && h) || e('kdf', s), a & (a - 1) && e('N', a);
        const l = parseInt($a(t, 'crypto/kdfparams/dklen'));
        return 32 !== l && e('dklen', l), n(o, r, a, u, h, 64, i);
      }
      if ('pbkdf2' === s.toLowerCase()) {
        const n = Ya($a(t, 'crypto/kdfparams/salt'));
        let i = null;
        const s = $a(t, 'crypto/kdfparams/prf');
        'hmac-sha256' === s
          ? (i = 'sha256')
          : 'hmac-sha512' === s
          ? (i = 'sha512')
          : e('prf', s);
        const a = parseInt($a(t, 'crypto/kdfparams/c')),
          u = parseInt($a(t, 'crypto/kdfparams/dklen'));
        return 32 !== u && e('dklen', u), r(o, n, a, u, i);
      }
    }
    return hu.throwArgumentError(
      'unsupported key-derivation function',
      'kdf',
      s
    );
  }
  function gu(t, e, r) {
    if (iu(t)) {
      r && r(0);
      const n = nu(t, e);
      return r && r(1), Promise.resolve(n);
    }
    return ou(t)
      ? (function (t, e, r) {
          return uu(this, void 0, void 0, function* () {
            const n = JSON.parse(t);
            return fu(n, yield pu(n, e, mu, au.scrypt, r));
          });
        })(t, e, r)
      : Promise.reject(new Error('invalid JSON wallet'));
  }
  function yu(t, e) {
    if (iu(t)) return nu(t, e);
    if (ou(t))
      return (function (t, e) {
        const r = JSON.parse(t);
        return fu(r, pu(r, e, du, au.syncScrypt));
      })(t, e);
    throw new Error('invalid JSON wallet');
  }
  var vu = function (t, e, r, n) {
    return new (r || (r = Promise))(function (i, o) {
      function s(t) {
        try {
          u(n.next(t));
        } catch (t) {
          o(t);
        }
      }
      function a(t) {
        try {
          u(n.throw(t));
        } catch (t) {
          o(t);
        }
      }
      function u(t) {
        var e;
        t.done
          ? i(t.value)
          : ((e = t.value),
            e instanceof r
              ? e
              : new r(function (t) {
                  t(e);
                })).then(s, a);
      }
      u((n = n.apply(t, e || [])).next());
    });
  };
  const wu = new f('wallet/5.7.0');
  class bu extends jn {
    constructor(t, e) {
      if (
        (super(), null != (r = t) && M(r.privateKey, 32) && null != r.address)
      ) {
        const e = new Cs(t.privateKey);
        if (
          (st(this, '_signingKey', () => e),
          st(this, 'address', js(this.publicKey)),
          this.address !== Re(t.address) &&
            wu.throwArgumentError(
              'privateKey/address mismatch',
              'privateKey',
              '[REDACTED]'
            ),
          (function (t) {
            const e = t.mnemonic;
            return e && e.phrase;
          })(t))
        ) {
          const e = t.mnemonic;
          st(this, '_mnemonic', () => ({
            phrase: e.phrase,
            path: e.path || Da,
            locale: e.locale || 'en',
          }));
          const r = this.mnemonic;
          js(
            La.fromMnemonic(r.phrase, null, r.locale).derivePath(r.path)
              .privateKey
          ) !== this.address &&
            wu.throwArgumentError(
              'mnemonic/address mismatch',
              'privateKey',
              '[REDACTED]'
            );
        } else st(this, '_mnemonic', () => null);
      } else {
        if (Cs.isSigningKey(t))
          'secp256k1' !== t.curve &&
            wu.throwArgumentError(
              'unsupported curve; must be secp256k1',
              'privateKey',
              '[REDACTED]'
            ),
            st(this, '_signingKey', () => t);
        else {
          'string' == typeof t &&
            t.match(/^[0-9a-f]*$/i) &&
            64 === t.length &&
            (t = '0x' + t);
          const e = new Cs(t);
          st(this, '_signingKey', () => e);
        }
        st(this, '_mnemonic', () => null),
          st(this, 'address', js(this.publicKey));
      }
      var r;
      e &&
        !On.isProvider(e) &&
        wu.throwArgumentError('invalid provider', 'provider', e),
        st(this, 'provider', e || null);
    }
    get mnemonic() {
      return this._mnemonic();
    }
    get privateKey() {
      return this._signingKey().privateKey;
    }
    get publicKey() {
      return this._signingKey().publicKey;
    }
    getAddress() {
      return Promise.resolve(this.address);
    }
    connect(t) {
      return new bu(this, t);
    }
    signTransaction(t) {
      return ut(t).then((e) => {
        null != e.from &&
          (Re(e.from) !== this.address &&
            wu.throwArgumentError(
              'transaction from address mismatch',
              'transaction.from',
              t.from
            ),
          delete e.from);
        const r = this._signingKey().signDigest(ye(Ws(e)));
        return Ws(e, r);
      });
    }
    signMessage(t) {
      return vu(this, void 0, void 0, function* () {
        return I(this._signingKey().signDigest(an(t)));
      });
    }
    _signTypedData(t, e, r) {
      return vu(this, void 0, void 0, function* () {
        const n = yield Mn.resolveNames(
          t,
          e,
          r,
          (t) => (
            null == this.provider &&
              wu.throwError(
                'cannot resolve ENS names without a provider',
                f.errors.UNSUPPORTED_OPERATION,
                { operation: 'resolveName', value: t }
              ),
            this.provider.resolveName(t)
          )
        );
        return I(this._signingKey().signDigest(Mn.hash(n.domain, e, n.value)));
      });
    }
    encrypt(t, e, r) {
      if (
        ('function' != typeof e || r || ((r = e), (e = {})),
        r && 'function' != typeof r)
      )
        throw new Error('invalid callback');
      return (
        e || (e = {}),
        (function (t, e, r, n) {
          try {
            if (Re(t.address) !== js(t.privateKey))
              throw new Error('address/privateKey mismatch');
            if (lu(t)) {
              const e = t.mnemonic;
              if (
                La.fromMnemonic(e.phrase, null, e.locale).derivePath(
                  e.path || Da
                ).privateKey != t.privateKey
              )
                throw new Error('mnemonic mismatch');
            }
          } catch (t) {
            return Promise.reject(t);
          }
          'function' != typeof r || n || ((n = r), (r = {})), r || (r = {});
          const i = w(t.privateKey),
            o = Xa(e);
          let s = null,
            a = null,
            u = null;
          if (lu(t)) {
            const e = t.mnemonic;
            (s = w(ja(e.phrase, e.locale || 'en'))),
              (a = e.path || Da),
              (u = e.locale || 'en');
          }
          let h = r.client;
          h || (h = 'ethers.js');
          let l = null;
          l = r.salt ? w(r.salt) : Ha(32);
          let c = null;
          if (r.iv) {
            if (((c = w(r.iv)), 16 !== c.length)) throw new Error('invalid iv');
          } else c = Ha(16);
          let f = null;
          if (r.uuid) {
            if (((f = w(r.uuid)), 16 !== f.length))
              throw new Error('invalid uuid');
          } else f = Ha(16);
          let d = 1 << 17,
            m = 8,
            p = 1;
          return (
            r.scrypt &&
              (r.scrypt.N && (d = r.scrypt.N),
              r.scrypt.r && (m = r.scrypt.r),
              r.scrypt.p && (p = r.scrypt.p)),
            au.scrypt(o, l, d, m, p, 64, n).then((e) => {
              const r = (e = w(e)).slice(0, 16),
                n = e.slice(16, 32),
                o = e.slice(32, 64),
                g = new Wa.Counter(c),
                y = w(new Wa.ModeOfOperation.ctr(r, g).encrypt(i)),
                v = ye(b([n, y])),
                A = {
                  address: t.address.substring(2).toLowerCase(),
                  id: tu(f),
                  version: 3,
                  crypto: {
                    cipher: 'aes-128-ctr',
                    cipherparams: { iv: k(c).substring(2) },
                    ciphertext: k(y).substring(2),
                    kdf: 'scrypt',
                    kdfparams: {
                      salt: k(l).substring(2),
                      n: d,
                      dklen: 32,
                      p: p,
                      r: m,
                    },
                    mac: v.substring(2),
                  },
                };
              if (s) {
                const t = Ha(16),
                  e = new Wa.Counter(t),
                  r = w(new Wa.ModeOfOperation.ctr(o, e).encrypt(s)),
                  n = new Date(),
                  i =
                    n.getUTCFullYear() +
                    '-' +
                    Za(n.getUTCMonth() + 1, 2) +
                    '-' +
                    Za(n.getUTCDate(), 2) +
                    'T' +
                    Za(n.getUTCHours(), 2) +
                    '-' +
                    Za(n.getUTCMinutes(), 2) +
                    '-' +
                    Za(n.getUTCSeconds(), 2) +
                    '.0Z';
                A['x-ethers'] = {
                  client: h,
                  gethFilename: 'UTC--' + i + '--' + A.address,
                  mnemonicCounter: k(t).substring(2),
                  mnemonicCiphertext: k(r).substring(2),
                  path: a,
                  locale: u,
                  version: '0.1',
                };
              }
              return JSON.stringify(A);
            })
          );
        })(this, t, e, r)
      );
    }
    static createRandom(t) {
      let e = Ha(16);
      t || (t = {}),
        t.extraEntropy && (e = w(N(ye(b([e, t.extraEntropy])), 0, 16)));
      const r = za(e, t.locale);
      return bu.fromMnemonic(r, t.path, t.locale);
    }
    static fromEncryptedJson(t, e, r) {
      return gu(t, e, r).then((t) => new bu(t));
    }
    static fromEncryptedJsonSync(t, e) {
      return new bu(yu(t, e));
    }
    static fromMnemonic(t, e, r) {
      return e || (e = Da), new bu(La.fromMnemonic(t, null, r).derivePath(e));
    }
  }
  const Au = new f('networks/5.7.1');
  function Eu(t) {
    const e = function (e, r) {
      null == r && (r = {});
      const n = [];
      if (e.InfuraProvider && '-' !== r.infura)
        try {
          n.push(new e.InfuraProvider(t, r.infura));
        } catch (t) {}
      if (e.EtherscanProvider && '-' !== r.etherscan)
        try {
          n.push(new e.EtherscanProvider(t, r.etherscan));
        } catch (t) {}
      if (e.AlchemyProvider && '-' !== r.alchemy)
        try {
          n.push(new e.AlchemyProvider(t, r.alchemy));
        } catch (t) {}
      if (e.PocketProvider && '-' !== r.pocket) {
        const i = ['goerli', 'ropsten', 'rinkeby', 'sepolia'];
        try {
          const o = new e.PocketProvider(t, r.pocket);
          o.network && -1 === i.indexOf(o.network.name) && n.push(o);
        } catch (t) {}
      }
      if (e.CloudflareProvider && '-' !== r.cloudflare)
        try {
          n.push(new e.CloudflareProvider(t));
        } catch (t) {}
      if (e.AnkrProvider && '-' !== r.ankr)
        try {
          const i = ['ropsten'],
            o = new e.AnkrProvider(t, r.ankr);
          o.network && -1 === i.indexOf(o.network.name) && n.push(o);
        } catch (t) {}
      if (0 === n.length) return null;
      if (e.FallbackProvider) {
        let i = 1;
        return (
          null != r.quorum ? (i = r.quorum) : 'homestead' === t && (i = 2),
          new e.FallbackProvider(n, i)
        );
      }
      return n[0];
    };
    return (
      (e.renetwork = function (t) {
        return Eu(t);
      }),
      e
    );
  }
  function Mu(t, e) {
    const r = function (r, n) {
      return r.JsonRpcProvider ? new r.JsonRpcProvider(t, e) : null;
    };
    return (
      (r.renetwork = function (e) {
        return Mu(t, e);
      }),
      r
    );
  }
  const _u = {
      chainId: 1,
      ensAddress: '0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e',
      name: 'homestead',
      _defaultProvider: Eu('homestead'),
    },
    ku = {
      chainId: 3,
      ensAddress: '0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e',
      name: 'ropsten',
      _defaultProvider: Eu('ropsten'),
    },
    Su = {
      chainId: 63,
      name: 'classicMordor',
      _defaultProvider: Mu(
        'https://www.ethercluster.com/mordor',
        'classicMordor'
      ),
    },
    Nu = {
      unspecified: { chainId: 0, name: 'unspecified' },
      homestead: _u,
      mainnet: _u,
      morden: { chainId: 2, name: 'morden' },
      ropsten: ku,
      testnet: ku,
      rinkeby: {
        chainId: 4,
        ensAddress: '0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e',
        name: 'rinkeby',
        _defaultProvider: Eu('rinkeby'),
      },
      kovan: { chainId: 42, name: 'kovan', _defaultProvider: Eu('kovan') },
      goerli: {
        chainId: 5,
        ensAddress: '0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e',
        name: 'goerli',
        _defaultProvider: Eu('goerli'),
      },
      kintsugi: { chainId: 1337702, name: 'kintsugi' },
      sepolia: {
        chainId: 11155111,
        name: 'sepolia',
        _defaultProvider: Eu('sepolia'),
      },
      classic: {
        chainId: 61,
        name: 'classic',
        _defaultProvider: Mu('https://www.ethercluster.com/etc', 'classic'),
      },
      classicMorden: { chainId: 62, name: 'classicMorden' },
      classicMordor: Su,
      classicTestnet: Su,
      classicKotti: {
        chainId: 6,
        name: 'classicKotti',
        _defaultProvider: Mu(
          'https://www.ethercluster.com/kotti',
          'classicKotti'
        ),
      },
      xdai: { chainId: 100, name: 'xdai' },
      matic: { chainId: 137, name: 'matic', _defaultProvider: Eu('matic') },
      maticmum: { chainId: 80001, name: 'maticmum' },
      optimism: {
        chainId: 10,
        name: 'optimism',
        _defaultProvider: Eu('optimism'),
      },
      'optimism-kovan': { chainId: 69, name: 'optimism-kovan' },
      'optimism-goerli': { chainId: 420, name: 'optimism-goerli' },
      arbitrum: { chainId: 42161, name: 'arbitrum' },
      'arbitrum-rinkeby': { chainId: 421611, name: 'arbitrum-rinkeby' },
      'arbitrum-goerli': { chainId: 421613, name: 'arbitrum-goerli' },
      bnb: { chainId: 56, name: 'bnb' },
      bnbt: { chainId: 97, name: 'bnbt' },
    };
  function Pu(t) {
    if (null == t) return null;
    if ('number' == typeof t) {
      for (const e in Nu) {
        const r = Nu[e];
        if (r.chainId === t)
          return {
            name: r.name,
            chainId: r.chainId,
            ensAddress: r.ensAddress || null,
            _defaultProvider: r._defaultProvider || null,
          };
      }
      return { chainId: t, name: 'unknown' };
    }
    if ('string' == typeof t) {
      const e = Nu[t];
      return null == e
        ? null
        : {
            name: e.name,
            chainId: e.chainId,
            ensAddress: e.ensAddress,
            _defaultProvider: e._defaultProvider || null,
          };
    }
    const e = Nu[t.name];
    if (!e)
      return (
        'number' != typeof t.chainId &&
          Au.throwArgumentError('invalid network chainId', 'network', t),
        t
      );
    0 !== t.chainId &&
      t.chainId !== e.chainId &&
      Au.throwArgumentError('network chainId mismatch', 'network', t);
    let r = t._defaultProvider || null;
    var n;
    return (
      null == r &&
        e._defaultProvider &&
        (r =
          (n = e._defaultProvider) && 'function' == typeof n.renetwork
            ? e._defaultProvider.renetwork(t)
            : e._defaultProvider),
      {
        name: t.name,
        chainId: e.chainId,
        ensAddress: t.ensAddress || e.ensAddress || null,
        _defaultProvider: r,
      }
    );
  }
  var xu = function (t, e, r, n) {
    return new (r || (r = Promise))(function (i, o) {
      function s(t) {
        try {
          u(n.next(t));
        } catch (t) {
          o(t);
        }
      }
      function a(t) {
        try {
          u(n.throw(t));
        } catch (t) {
          o(t);
        }
      }
      function u(t) {
        var e;
        t.done
          ? i(t.value)
          : ((e = t.value),
            e instanceof r
              ? e
              : new r(function (t) {
                  t(e);
                })).then(s, a);
      }
      u((n = n.apply(t, e || [])).next());
    });
  };
  function Tu(t, e) {
    return xu(this, void 0, void 0, function* () {
      null == e && (e = {});
      const r = {
        method: e.method || 'GET',
        headers: e.headers || {},
        body: e.body || void 0,
      };
      if (
        (!0 !== e.skipFetchSetup &&
          ((r.mode = 'cors'),
          (r.cache = 'no-cache'),
          (r.credentials = 'same-origin'),
          (r.redirect = 'follow'),
          (r.referrer = 'client')),
        null != e.fetchOptions)
      ) {
        const t = e.fetchOptions;
        t.mode && (r.mode = t.mode),
          t.cache && (r.cache = t.cache),
          t.credentials && (r.credentials = t.credentials),
          t.redirect && (r.redirect = t.redirect),
          t.referrer && (r.referrer = t.referrer);
      }
      const n = yield fetch(t, r),
        i = yield n.arrayBuffer(),
        o = {};
      return (
        n.headers.forEach
          ? n.headers.forEach((t, e) => {
              o[e.toLowerCase()] = t;
            })
          : n.headers.keys().forEach((t) => {
              o[t.toLowerCase()] = n.headers.get(t);
            }),
        {
          headers: o,
          statusCode: n.status,
          statusMessage: n.statusText,
          body: w(new Uint8Array(i)),
        }
      );
    });
  }
  var Cu = function (t, e, r, n) {
    return new (r || (r = Promise))(function (i, o) {
      function s(t) {
        try {
          u(n.next(t));
        } catch (t) {
          o(t);
        }
      }
      function a(t) {
        try {
          u(n.throw(t));
        } catch (t) {
          o(t);
        }
      }
      function u(t) {
        var e;
        t.done
          ? i(t.value)
          : ((e = t.value),
            e instanceof r
              ? e
              : new r(function (t) {
                  t(e);
                })).then(s, a);
      }
      u((n = n.apply(t, e || [])).next());
    });
  };
  const Ru = new f('web/5.7.1');
  function Iu(t) {
    return new Promise((e) => {
      setTimeout(e, t);
    });
  }
  function Bu(t, e) {
    if (null == t) return null;
    if ('string' == typeof t) return t;
    if (g(t)) {
      if (
        e &&
        ('text' === e.split('/')[0] ||
          'application/json' === e.split(';')[0].trim())
      )
        try {
          return fr(t);
        } catch (t) {}
      return k(t);
    }
    return t;
  }
  function Ou(t, e, r) {
    const n =
      'object' == typeof t && null != t.throttleLimit ? t.throttleLimit : 12;
    Ru.assertArgument(
      n > 0 && n % 1 == 0,
      'invalid connection throttle limit',
      'connection.throttleLimit',
      n
    );
    const i = 'object' == typeof t ? t.throttleCallback : null,
      o =
        'object' == typeof t && 'number' == typeof t.throttleSlotInterval
          ? t.throttleSlotInterval
          : 100;
    Ru.assertArgument(
      o > 0 && o % 1 == 0,
      'invalid connection throttle slot interval',
      'connection.throttleSlotInterval',
      o
    );
    const s = 'object' == typeof t && !!t.errorPassThrough,
      a = {};
    let u = null;
    const h = { method: 'GET' };
    let l = !1,
      c = 12e4;
    if ('string' == typeof t) u = t;
    else if ('object' == typeof t) {
      if (
        ((null != t && null != t.url) ||
          Ru.throwArgumentError('missing URL', 'connection.url', t),
        (u = t.url),
        'number' == typeof t.timeout && t.timeout > 0 && (c = t.timeout),
        t.headers)
      )
        for (const e in t.headers)
          (a[e.toLowerCase()] = { key: e, value: String(t.headers[e]) }),
            ['if-none-match', 'if-modified-since'].indexOf(e.toLowerCase()) >=
              0 && (l = !0);
      if (
        ((h.allowGzip = !!t.allowGzip), null != t.user && null != t.password)
      ) {
        'https:' !== u.substring(0, 6) &&
          !0 !== t.allowInsecureAuthentication &&
          Ru.throwError(
            'basic authentication requires a secure https url',
            f.errors.INVALID_ARGUMENT,
            { argument: 'url', url: u, user: t.user, password: '[REDACTED]' }
          );
        const e = t.user + ':' + t.password;
        a.authorization = { key: 'Authorization', value: 'Basic ' + Ir(hr(e)) };
      }
      null != t.skipFetchSetup && (h.skipFetchSetup = !!t.skipFetchSetup),
        null != t.fetchOptions && (h.fetchOptions = lt(t.fetchOptions));
    }
    const d = new RegExp('^data:([^;:]*)?(;base64)?,(.*)$', 'i'),
      m = u ? u.match(d) : null;
    if (m)
      try {
        const t = {
          statusCode: 200,
          statusMessage: 'OK',
          headers: { 'content-type': m[1] || 'text/plain' },
          body: m[2]
            ? Rr(m[3])
            : ((p = m[3]),
              hr(
                p.replace(/%([0-9a-f][0-9a-f])/gi, (t, e) =>
                  String.fromCharCode(parseInt(e, 16))
                )
              )),
        };
        let e = t.body;
        return r && (e = r(t.body, t)), Promise.resolve(e);
      } catch (t) {
        Ru.throwError('processing response error', f.errors.SERVER_ERROR, {
          body: Bu(m[1], m[2]),
          error: t,
          requestBody: null,
          requestMethod: 'GET',
          url: u,
        });
      }
    var p;
    e &&
      ((h.method = 'POST'),
      (h.body = e),
      null == a['content-type'] &&
        (a['content-type'] = {
          key: 'Content-Type',
          value: 'application/octet-stream',
        }),
      null == a['content-length'] &&
        (a['content-length'] = {
          key: 'Content-Length',
          value: String(e.length),
        }));
    const g = {};
    Object.keys(a).forEach((t) => {
      const e = a[t];
      g[e.key] = e.value;
    }),
      (h.headers = g);
    const y = (function () {
        let t = null;
        return {
          promise: new Promise(function (e, r) {
            c &&
              (t = setTimeout(() => {
                null != t &&
                  ((t = null),
                  r(
                    Ru.makeError('timeout', f.errors.TIMEOUT, {
                      requestBody: Bu(h.body, g['content-type']),
                      requestMethod: h.method,
                      timeout: c,
                      url: u,
                    })
                  ));
              }, c));
          }),
          cancel: function () {
            null != t && (clearTimeout(t), (t = null));
          },
        };
      })(),
      v = (function () {
        return Cu(this, void 0, void 0, function* () {
          for (let t = 0; t < n; t++) {
            let e = null;
            try {
              if (((e = yield Tu(u, h)), t < n))
                if (301 === e.statusCode || 302 === e.statusCode) {
                  const t = e.headers.location || '';
                  if ('GET' === h.method && t.match(/^https:/)) {
                    u = e.headers.location;
                    continue;
                  }
                } else if (429 === e.statusCode) {
                  let r = !0;
                  if ((i && (r = yield i(t, u)), r)) {
                    let r = 0;
                    const n = e.headers['retry-after'];
                    (r =
                      'string' == typeof n && n.match(/^[1-9][0-9]*$/)
                        ? 1e3 * parseInt(n)
                        : o * parseInt(String(Math.random() * Math.pow(2, t)))),
                      yield Iu(r);
                    continue;
                  }
                }
            } catch (t) {
              (e = t.response),
                null == e &&
                  (y.cancel(),
                  Ru.throwError('missing response', f.errors.SERVER_ERROR, {
                    requestBody: Bu(h.body, g['content-type']),
                    requestMethod: h.method,
                    serverError: t,
                    url: u,
                  }));
            }
            let a = e.body;
            if (
              (l && 304 === e.statusCode
                ? (a = null)
                : !s &&
                  (e.statusCode < 200 || e.statusCode >= 300) &&
                  (y.cancel(),
                  Ru.throwError('bad response', f.errors.SERVER_ERROR, {
                    status: e.statusCode,
                    headers: e.headers,
                    body: Bu(a, e.headers ? e.headers['content-type'] : null),
                    requestBody: Bu(h.body, g['content-type']),
                    requestMethod: h.method,
                    url: u,
                  })),
              r)
            )
              try {
                const t = yield r(a, e);
                return y.cancel(), t;
              } catch (r) {
                if (r.throttleRetry && t < n) {
                  let e = !0;
                  if ((i && (e = yield i(t, u)), e)) {
                    const e =
                      o * parseInt(String(Math.random() * Math.pow(2, t)));
                    yield Iu(e);
                    continue;
                  }
                }
                y.cancel(),
                  Ru.throwError(
                    'processing response error',
                    f.errors.SERVER_ERROR,
                    {
                      body: Bu(a, e.headers ? e.headers['content-type'] : null),
                      error: r,
                      requestBody: Bu(h.body, g['content-type']),
                      requestMethod: h.method,
                      url: u,
                    }
                  );
              }
            return y.cancel(), a;
          }
          return Ru.throwError('failed response', f.errors.SERVER_ERROR, {
            requestBody: Bu(h.body, g['content-type']),
            requestMethod: h.method,
            url: u,
          });
        });
      })();
    return Promise.race([y.promise, v]);
  }
  function Fu(t, e, r) {
    let n = null;
    if (null != e) {
      n = hr(e);
      const r = 'string' == typeof t ? { url: t } : lt(t);
      if (r.headers) {
        0 !==
          Object.keys(r.headers).filter(
            (t) => 'content-type' === t.toLowerCase()
          ).length ||
          ((r.headers = lt(r.headers)),
          (r.headers['content-type'] = 'application/json'));
      } else r.headers = { 'content-type': 'application/json' };
      t = r;
    }
    return Ou(t, n, (t, e) => {
      let n = null;
      if (null != t)
        try {
          n = JSON.parse(fr(t));
        } catch (e) {
          Ru.throwError('invalid JSON', f.errors.SERVER_ERROR, {
            body: t,
            error: e,
          });
        }
      return r && (n = r(n, e)), n;
    });
  }
  function Du(t, e) {
    return (
      e || (e = {}),
      null == (e = lt(e)).floor && (e.floor = 0),
      null == e.ceiling && (e.ceiling = 1e4),
      null == e.interval && (e.interval = 250),
      new Promise(function (r, n) {
        let i = null,
          o = !1;
        const s = () => !o && ((o = !0), i && clearTimeout(i), !0);
        e.timeout &&
          (i = setTimeout(() => {
            s() && n(new Error('timeout'));
          }, e.timeout));
        const a = e.retryLimit;
        let u = 0;
        !(function i() {
          return t().then(
            function (t) {
              if (void 0 !== t) s() && r(t);
              else if (e.oncePoll) e.oncePoll.once('poll', i);
              else if (e.onceBlock) e.onceBlock.once('block', i);
              else if (!o) {
                if ((u++, u > a))
                  return void (s() && n(new Error('retry limit reached')));
                let t =
                  e.interval * parseInt(String(Math.random() * Math.pow(2, u)));
                t < e.floor && (t = e.floor),
                  t > e.ceiling && (t = e.ceiling),
                  setTimeout(i, t);
              }
              return null;
            },
            function (t) {
              s() && n(t);
            }
          );
        })();
      })
    );
  }
  for (
    var Lu = 'qpzry9x8gf2tvdw0s3jn54khce6mua7l', Uu = {}, ju = 0;
    ju < 32;
    ju++
  ) {
    var zu = Lu.charAt(ju);
    if (void 0 !== Uu[zu]) throw new TypeError(zu + ' is ambiguous');
    Uu[zu] = ju;
  }
  function Gu(t) {
    var e = t >> 25;
    return (
      ((33554431 & t) << 5) ^
      (996825010 & -(1 & e)) ^
      (642813549 & -((e >> 1) & 1)) ^
      (513874426 & -((e >> 2) & 1)) ^
      (1027748829 & -((e >> 3) & 1)) ^
      (705979059 & -((e >> 4) & 1))
    );
  }
  function qu(t) {
    for (var e = 1, r = 0; r < t.length; ++r) {
      var n = t.charCodeAt(r);
      if (n < 33 || n > 126) return 'Invalid prefix (' + t + ')';
      e = Gu(e) ^ (n >> 5);
    }
    for (e = Gu(e), r = 0; r < t.length; ++r) {
      var i = t.charCodeAt(r);
      e = Gu(e) ^ (31 & i);
    }
    return e;
  }
  function Ku(t, e) {
    if (((e = e || 90), t.length < 8)) return t + ' too short';
    if (t.length > e) return 'Exceeds length limit';
    var r = t.toLowerCase(),
      n = t.toUpperCase();
    if (t !== r && t !== n) return 'Mixed-case string ' + t;
    var i = (t = r).lastIndexOf('1');
    if (-1 === i) return 'No separator character for ' + t;
    if (0 === i) return 'Missing prefix for ' + t;
    var o = t.slice(0, i),
      s = t.slice(i + 1);
    if (s.length < 6) return 'Data too short';
    var a = qu(o);
    if ('string' == typeof a) return a;
    for (var u = [], h = 0; h < s.length; ++h) {
      var l = s.charAt(h),
        c = Uu[l];
      if (void 0 === c) return 'Unknown character ' + l;
      (a = Gu(a) ^ c), h + 6 >= s.length || u.push(c);
    }
    return 1 !== a ? 'Invalid checksum for ' + t : { prefix: o, words: u };
  }
  function Hu(t, e, r, n) {
    for (var i = 0, o = 0, s = (1 << r) - 1, a = [], u = 0; u < t.length; ++u)
      for (i = (i << e) | t[u], o += e; o >= r; )
        (o -= r), a.push((i >> o) & s);
    if (n) o > 0 && a.push((i << (r - o)) & s);
    else {
      if (o >= e) return 'Excess padding';
      if ((i << (r - o)) & s) return 'Non-zero padding';
    }
    return a;
  }
  var Vu = {
      decodeUnsafe: function () {
        var t = Ku.apply(null, arguments);
        if ('object' == typeof t) return t;
      },
      decode: function (t) {
        var e = Ku.apply(null, arguments);
        if ('object' == typeof e) return e;
        throw new Error(e);
      },
      encode: function (t, e, r) {
        if (((r = r || 90), t.length + 7 + e.length > r))
          throw new TypeError('Exceeds length limit');
        var n = qu((t = t.toLowerCase()));
        if ('string' == typeof n) throw new Error(n);
        for (var i = t + '1', o = 0; o < e.length; ++o) {
          var s = e[o];
          if (s >> 5) throw new Error('Non 5-bit word');
          (n = Gu(n) ^ s), (i += Lu.charAt(s));
        }
        for (o = 0; o < 6; ++o) n = Gu(n);
        for (n ^= 1, o = 0; o < 6; ++o) {
          i += Lu.charAt((n >> (5 * (5 - o))) & 31);
        }
        return i;
      },
      toWordsUnsafe: function (t) {
        var e = Hu(t, 8, 5, !0);
        if (Array.isArray(e)) return e;
      },
      toWords: function (t) {
        var e = Hu(t, 8, 5, !0);
        if (Array.isArray(e)) return e;
        throw new Error(e);
      },
      fromWordsUnsafe: function (t) {
        var e = Hu(t, 5, 8, !1);
        if (Array.isArray(e)) return e;
      },
      fromWords: function (t) {
        var e = Hu(t, 5, 8, !1);
        if (Array.isArray(e)) return e;
        throw new Error(e);
      },
    },
    Ju = Vu;
  const Wu = 'providers/5.7.2',
    Qu = new f(Wu);
  class Yu {
    constructor() {
      this.formats = this.getDefaultFormats();
    }
    getDefaultFormats() {
      const t = {},
        e = this.address.bind(this),
        r = this.bigNumber.bind(this),
        n = this.blockTag.bind(this),
        i = this.data.bind(this),
        o = this.hash.bind(this),
        s = this.hex.bind(this),
        a = this.number.bind(this),
        u = this.type.bind(this);
      return (
        (t.transaction = {
          hash: o,
          type: u,
          accessList: Yu.allowNull(this.accessList.bind(this), null),
          blockHash: Yu.allowNull(o, null),
          blockNumber: Yu.allowNull(a, null),
          transactionIndex: Yu.allowNull(a, null),
          confirmations: Yu.allowNull(a, null),
          from: e,
          gasPrice: Yu.allowNull(r),
          maxPriorityFeePerGas: Yu.allowNull(r),
          maxFeePerGas: Yu.allowNull(r),
          gasLimit: r,
          to: Yu.allowNull(e, null),
          value: r,
          nonce: a,
          data: i,
          r: Yu.allowNull(this.uint256),
          s: Yu.allowNull(this.uint256),
          v: Yu.allowNull(a),
          creates: Yu.allowNull(e, null),
          raw: Yu.allowNull(i),
        }),
        (t.transactionRequest = {
          from: Yu.allowNull(e),
          nonce: Yu.allowNull(a),
          gasLimit: Yu.allowNull(r),
          gasPrice: Yu.allowNull(r),
          maxPriorityFeePerGas: Yu.allowNull(r),
          maxFeePerGas: Yu.allowNull(r),
          to: Yu.allowNull(e),
          value: Yu.allowNull(r),
          data: Yu.allowNull((t) => this.data(t, !0)),
          type: Yu.allowNull(a),
          accessList: Yu.allowNull(this.accessList.bind(this), null),
        }),
        (t.receiptLog = {
          transactionIndex: a,
          blockNumber: a,
          transactionHash: o,
          address: e,
          topics: Yu.arrayOf(o),
          data: i,
          logIndex: a,
          blockHash: o,
        }),
        (t.receipt = {
          to: Yu.allowNull(this.address, null),
          from: Yu.allowNull(this.address, null),
          contractAddress: Yu.allowNull(e, null),
          transactionIndex: a,
          root: Yu.allowNull(s),
          gasUsed: r,
          logsBloom: Yu.allowNull(i),
          blockHash: o,
          transactionHash: o,
          logs: Yu.arrayOf(this.receiptLog.bind(this)),
          blockNumber: a,
          confirmations: Yu.allowNull(a, null),
          cumulativeGasUsed: r,
          effectiveGasPrice: Yu.allowNull(r),
          status: Yu.allowNull(a),
          type: u,
        }),
        (t.block = {
          hash: Yu.allowNull(o),
          parentHash: o,
          number: a,
          timestamp: a,
          nonce: Yu.allowNull(s),
          difficulty: this.difficulty.bind(this),
          gasLimit: r,
          gasUsed: r,
          miner: Yu.allowNull(e),
          extraData: i,
          transactions: Yu.allowNull(Yu.arrayOf(o)),
          baseFeePerGas: Yu.allowNull(r),
        }),
        (t.blockWithTransactions = lt(t.block)),
        (t.blockWithTransactions.transactions = Yu.allowNull(
          Yu.arrayOf(this.transactionResponse.bind(this))
        )),
        (t.filter = {
          fromBlock: Yu.allowNull(n, void 0),
          toBlock: Yu.allowNull(n, void 0),
          blockHash: Yu.allowNull(o, void 0),
          address: Yu.allowNull(e, void 0),
          topics: Yu.allowNull(this.topics.bind(this), void 0),
        }),
        (t.filterLog = {
          blockNumber: Yu.allowNull(a),
          blockHash: Yu.allowNull(o),
          transactionIndex: a,
          removed: Yu.allowNull(this.boolean.bind(this)),
          address: e,
          data: Yu.allowFalsish(i, '0x'),
          topics: Yu.arrayOf(o),
          transactionHash: o,
          logIndex: a,
        }),
        t
      );
    }
    accessList(t) {
      return Ks(t || []);
    }
    number(t) {
      return '0x' === t ? 0 : j.from(t).toNumber();
    }
    type(t) {
      return '0x' === t || null == t ? 0 : j.from(t).toNumber();
    }
    bigNumber(t) {
      return j.from(t);
    }
    boolean(t) {
      if ('boolean' == typeof t) return t;
      if ('string' == typeof t) {
        if ('true' === (t = t.toLowerCase())) return !0;
        if ('false' === t) return !1;
      }
      throw new Error('invalid boolean - ' + t);
    }
    hex(t, e) {
      return 'string' == typeof t &&
        (e || '0x' === t.substring(0, 2) || (t = '0x' + t), M(t))
        ? t.toLowerCase()
        : Qu.throwArgumentError('invalid hash', 'value', t);
    }
    data(t, e) {
      const r = this.hex(t, e);
      if (r.length % 2 != 0) throw new Error('invalid data; odd-length - ' + t);
      return r;
    }
    address(t) {
      return Re(t);
    }
    callAddress(t) {
      if (!M(t, 32)) return null;
      const e = Re(N(t, 12));
      return e === He ? null : e;
    }
    contractAddress(t) {
      return Ie(t);
    }
    blockTag(t) {
      if (null == t) return 'latest';
      if ('earliest' === t) return '0x0';
      switch (t) {
        case 'earliest':
          return '0x0';
        case 'latest':
        case 'pending':
        case 'safe':
        case 'finalized':
          return t;
      }
      if ('number' == typeof t || M(t)) return x(t);
      throw new Error('invalid blockTag');
    }
    hash(t, e) {
      const r = this.hex(t, e);
      return 32 !== S(r)
        ? Qu.throwArgumentError('invalid hash', 'value', t)
        : r;
    }
    difficulty(t) {
      if (null == t) return null;
      const e = j.from(t);
      try {
        return e.toNumber();
      } catch (t) {}
      return null;
    }
    uint256(t) {
      if (!M(t)) throw new Error('invalid uint256');
      return C(t, 32);
    }
    _block(t, e) {
      null != t.author && null == t.miner && (t.miner = t.author);
      const r = null != t._difficulty ? t._difficulty : t.difficulty,
        n = Yu.check(e, t);
      return (n._difficulty = null == r ? null : j.from(r)), n;
    }
    block(t) {
      return this._block(t, this.formats.block);
    }
    blockWithTransactions(t) {
      return this._block(t, this.formats.blockWithTransactions);
    }
    transactionRequest(t) {
      return Yu.check(this.formats.transactionRequest, t);
    }
    transactionResponse(t) {
      null != t.gas && null == t.gasLimit && (t.gasLimit = t.gas),
        t.to &&
          j.from(t.to).isZero() &&
          (t.to = '0x0000000000000000000000000000000000000000'),
        null != t.input && null == t.data && (t.data = t.input),
        null == t.to &&
          null == t.creates &&
          (t.creates = this.contractAddress(t)),
        (1 !== t.type && 2 !== t.type) ||
          null != t.accessList ||
          (t.accessList = []);
      const e = Yu.check(this.formats.transaction, t);
      if (null != t.chainId) {
        let r = t.chainId;
        M(r) && (r = j.from(r).toNumber()), (e.chainId = r);
      } else {
        let r = t.networkId;
        null == r && null == e.v && (r = t.chainId),
          M(r) && (r = j.from(r).toNumber()),
          'number' != typeof r &&
            null != e.v &&
            ((r = (e.v - 35) / 2), r < 0 && (r = 0), (r = parseInt(r))),
          'number' != typeof r && (r = 0),
          (e.chainId = r);
      }
      return (
        e.blockHash &&
          'x' === e.blockHash.replace(/0/g, '') &&
          (e.blockHash = null),
        e
      );
    }
    transaction(t) {
      return Ys(t);
    }
    receiptLog(t) {
      return Yu.check(this.formats.receiptLog, t);
    }
    receipt(t) {
      const e = Yu.check(this.formats.receipt, t);
      if (null != e.root)
        if (e.root.length <= 4) {
          const t = j.from(e.root).toNumber();
          0 === t || 1 === t
            ? (null != e.status &&
                e.status !== t &&
                Qu.throwArgumentError(
                  'alt-root-status/status mismatch',
                  'value',
                  { root: e.root, status: e.status }
                ),
              (e.status = t),
              delete e.root)
            : Qu.throwArgumentError(
                'invalid alt-root-status',
                'value.root',
                e.root
              );
        } else
          66 !== e.root.length &&
            Qu.throwArgumentError('invalid root hash', 'value.root', e.root);
      return null != e.status && (e.byzantium = !0), e;
    }
    topics(t) {
      return Array.isArray(t)
        ? t.map((t) => this.topics(t))
        : null != t
        ? this.hash(t, !0)
        : null;
    }
    filter(t) {
      return Yu.check(this.formats.filter, t);
    }
    filterLog(t) {
      return Yu.check(this.formats.filterLog, t);
    }
    static check(t, e) {
      const r = {};
      for (const n in t)
        try {
          const i = t[n](e[n]);
          void 0 !== i && (r[n] = i);
        } catch (t) {
          throw ((t.checkKey = n), (t.checkValue = e[n]), t);
        }
      return r;
    }
    static allowNull(t, e) {
      return function (r) {
        return null == r ? e : t(r);
      };
    }
    static allowFalsish(t, e) {
      return function (r) {
        return r ? t(r) : e;
      };
    }
    static arrayOf(t) {
      return function (e) {
        if (!Array.isArray(e)) throw new Error('not an array');
        const r = [];
        return (
          e.forEach(function (e) {
            r.push(t(e));
          }),
          r
        );
      };
    }
  }
  function Zu(t) {
    return t && 'function' == typeof t.isCommunityResource;
  }
  function Xu(t) {
    return Zu(t) && t.isCommunityResource();
  }
  let $u = !1;
  function th() {
    $u ||
      (($u = !0),
      console.log('========= NOTICE ========='),
      console.log('Request-Rate Exceeded  (this message will not be repeated)'),
      console.log(''),
      console.log(
        'The default API keys for each service are provided as a highly-throttled,'
      ),
      console.log(
        'community resource for low-traffic projects and early prototyping.'
      ),
      console.log(''),
      console.log(
        'While your application will continue to function, we highly recommended'
      ),
      console.log(
        'signing up for your own API keys to improve performance, increase your'
      ),
      console.log(
        'request rate/limit and enable other perks, such as metrics and advanced APIs.'
      ),
      console.log(''),
      console.log('For more details: https://docs.ethers.io/api-keys/'),
      console.log('=========================='));
  }
  var eh = function (t, e, r, n) {
    return new (r || (r = Promise))(function (i, o) {
      function s(t) {
        try {
          u(n.next(t));
        } catch (t) {
          o(t);
        }
      }
      function a(t) {
        try {
          u(n.throw(t));
        } catch (t) {
          o(t);
        }
      }
      function u(t) {
        var e;
        t.done
          ? i(t.value)
          : ((e = t.value),
            e instanceof r
              ? e
              : new r(function (t) {
                  t(e);
                })).then(s, a);
      }
      u((n = n.apply(t, e || [])).next());
    });
  };
  const rh = new f(Wu);
  function nh(t) {
    return null == t
      ? 'null'
      : (32 !== S(t) && rh.throwArgumentError('invalid topic', 'topic', t),
        t.toLowerCase());
  }
  function ih(t) {
    for (t = t.slice(); t.length > 0 && null == t[t.length - 1]; ) t.pop();
    return t
      .map((t) => {
        if (Array.isArray(t)) {
          const e = {};
          t.forEach((t) => {
            e[nh(t)] = !0;
          });
          const r = Object.keys(e);
          return r.sort(), r.join('|');
        }
        return nh(t);
      })
      .join('&');
  }
  function oh(t) {
    if ('string' == typeof t) {
      if (32 === S((t = t.toLowerCase()))) return 'tx:' + t;
      if (-1 === t.indexOf(':')) return t;
    } else {
      if (Array.isArray(t)) return 'filter:*:' + ih(t);
      if (Bn.isForkEvent(t))
        throw (rh.warn('not implemented'), new Error('not implemented'));
      if (t && 'object' == typeof t)
        return 'filter:' + (t.address || '*') + ':' + ih(t.topics || []);
    }
    throw new Error('invalid event - ' + t);
  }
  function sh() {
    return new Date().getTime();
  }
  function ah(t) {
    return new Promise((e) => {
      setTimeout(e, t);
    });
  }
  const uh = ['block', 'network', 'pending', 'poll'];
  class hh {
    constructor(t, e, r) {
      st(this, 'tag', t),
        st(this, 'listener', e),
        st(this, 'once', r),
        (this._lastBlockNumber = -2),
        (this._inflight = !1);
    }
    get event() {
      switch (this.type) {
        case 'tx':
          return this.hash;
        case 'filter':
          return this.filter;
      }
      return this.tag;
    }
    get type() {
      return this.tag.split(':')[0];
    }
    get hash() {
      const t = this.tag.split(':');
      return 'tx' !== t[0] ? null : t[1];
    }
    get filter() {
      const t = this.tag.split(':');
      if ('filter' !== t[0]) return null;
      const e = t[1],
        r =
          '' === (n = t[2])
            ? []
            : n.split(/&/g).map((t) => {
                if ('' === t) return [];
                const e = t.split('|').map((t) => ('null' === t ? null : t));
                return 1 === e.length ? e[0] : e;
              });
      var n;
      const i = {};
      return (
        r.length > 0 && (i.topics = r), e && '*' !== e && (i.address = e), i
      );
    }
    pollable() {
      return this.tag.indexOf(':') >= 0 || uh.indexOf(this.tag) >= 0;
    }
  }
  const lh = {
    0: { symbol: 'btc', p2pkh: 0, p2sh: 5, prefix: 'bc' },
    2: { symbol: 'ltc', p2pkh: 48, p2sh: 50, prefix: 'ltc' },
    3: { symbol: 'doge', p2pkh: 30, p2sh: 22 },
    60: { symbol: 'eth', ilk: 'eth' },
    61: { symbol: 'etc', ilk: 'eth' },
    700: { symbol: 'xdai', ilk: 'eth' },
  };
  function ch(t) {
    return C(j.from(t).toHexString(), 32);
  }
  function fh(t) {
    return pa.encode(b([t, N(wa(wa(t)), 0, 4)]));
  }
  const dh = new RegExp('^(ipfs)://(.*)$', 'i'),
    mh = [
      new RegExp('^(https)://(.*)$', 'i'),
      new RegExp('^(data):(.*)$', 'i'),
      dh,
      new RegExp('^eip155:[0-9]+/(erc[0-9]+):(.*)$', 'i'),
    ];
  function ph(t, e) {
    try {
      return fr(gh(t, e));
    } catch (t) {}
    return null;
  }
  function gh(t, e) {
    if ('0x' === t) return null;
    const r = j.from(N(t, e, e + 32)).toNumber(),
      n = j.from(N(t, r, r + 32)).toNumber();
    return N(t, r + 32, r + 32 + n);
  }
  function yh(t) {
    return (
      t.match(/^ipfs:\/\/ipfs\//i)
        ? (t = t.substring(12))
        : t.match(/^ipfs:\/\//i)
        ? (t = t.substring(7))
        : rh.throwArgumentError('unsupported IPFS format', 'link', t),
      `https://gateway.ipfs.io/ipfs/${t}`
    );
  }
  function vh(t) {
    const e = w(t);
    if (e.length > 32) throw new Error('internal; should not happen');
    const r = new Uint8Array(32);
    return r.set(e, 32 - e.length), r;
  }
  function wh(t) {
    if (t.length % 32 == 0) return t;
    const e = new Uint8Array(32 * Math.ceil(t.length / 32));
    return e.set(t), e;
  }
  function bh(t) {
    const e = [];
    let r = 0;
    for (let n = 0; n < t.length; n++) e.push(null), (r += 32);
    for (let n = 0; n < t.length; n++) {
      const i = w(t[n]);
      (e[n] = vh(r)),
        e.push(vh(i.length)),
        e.push(wh(i)),
        (r += 32 + 32 * Math.ceil(i.length / 32));
    }
    return P(e);
  }
  class Ah {
    constructor(t, e, r, n) {
      st(this, 'provider', t),
        st(this, 'name', r),
        st(this, 'address', t.formatter.address(e)),
        st(this, '_resolvedAddress', n);
    }
    supportsWildcard() {
      return (
        this._supportsEip2544 ||
          (this._supportsEip2544 = this.provider
            .call({
              to: this.address,
              data: '0x01ffc9a79061b92300000000000000000000000000000000000000000000000000000000',
            })
            .then((t) => j.from(t).eq(1))
            .catch((t) => {
              if (t.code === f.errors.CALL_EXCEPTION) return !1;
              throw ((this._supportsEip2544 = null), t);
            })),
        this._supportsEip2544
      );
    }
    _fetch(t, e) {
      return eh(this, void 0, void 0, function* () {
        const r = {
          to: this.address,
          ccipReadEnabled: !0,
          data: P([t, on(this.name), e || '0x']),
        };
        let n = !1;
        (yield this.supportsWildcard()) &&
          ((n = !0), (r.data = P(['0x9061b923', bh([sn(this.name), r.data])])));
        try {
          let t = yield this.provider.call(r);
          return (
            w(t).length % 32 == 4 &&
              rh.throwError('resolver threw error', f.errors.CALL_EXCEPTION, {
                transaction: r,
                data: t,
              }),
            n && (t = gh(t, 0)),
            t
          );
        } catch (t) {
          if (t.code === f.errors.CALL_EXCEPTION) return null;
          throw t;
        }
      });
    }
    _fetchBytes(t, e) {
      return eh(this, void 0, void 0, function* () {
        const r = yield this._fetch(t, e);
        return null != r ? gh(r, 0) : null;
      });
    }
    _getAddress(t, e) {
      const r = lh[String(t)];
      if (
        (null == r &&
          rh.throwError(
            `unsupported coin type: ${t}`,
            f.errors.UNSUPPORTED_OPERATION,
            { operation: `getAddress(${t})` }
          ),
        'eth' === r.ilk)
      )
        return this.provider.formatter.address(e);
      const n = w(e);
      if (null != r.p2pkh) {
        const t = e.match(/^0x76a9([0-9a-f][0-9a-f])([0-9a-f]*)88ac$/);
        if (t) {
          const e = parseInt(t[1], 16);
          if (t[2].length === 2 * e && e >= 1 && e <= 75)
            return fh(b([[r.p2pkh], '0x' + t[2]]));
        }
      }
      if (null != r.p2sh) {
        const t = e.match(/^0xa9([0-9a-f][0-9a-f])([0-9a-f]*)87$/);
        if (t) {
          const e = parseInt(t[1], 16);
          if (t[2].length === 2 * e && e >= 1 && e <= 75)
            return fh(b([[r.p2sh], '0x' + t[2]]));
        }
      }
      if (null != r.prefix) {
        const t = n[1];
        let e = n[0];
        if (
          (0 === e ? 20 !== t && 32 !== t && (e = -1) : (e = -1),
          e >= 0 && n.length === 2 + t && t >= 1 && t <= 75)
        ) {
          const t = Ju.toWords(n.slice(2));
          return t.unshift(e), Ju.encode(r.prefix, t);
        }
      }
      return null;
    }
    getAddress(t) {
      return eh(this, void 0, void 0, function* () {
        if ((null == t && (t = 60), 60 === t))
          try {
            const t = yield this._fetch('0x3b3b57de');
            return '0x' === t || t === tr
              ? null
              : this.provider.formatter.callAddress(t);
          } catch (t) {
            if (t.code === f.errors.CALL_EXCEPTION) return null;
            throw t;
          }
        const e = yield this._fetchBytes('0xf1cb7e06', ch(t));
        if (null == e || '0x' === e) return null;
        const r = this._getAddress(t, e);
        return (
          null == r &&
            rh.throwError(
              'invalid or unsupported coin data',
              f.errors.UNSUPPORTED_OPERATION,
              { operation: `getAddress(${t})`, coinType: t, data: e }
            ),
          r
        );
      });
    }
    getAvatar() {
      return eh(this, void 0, void 0, function* () {
        const t = [{ type: 'name', content: this.name }];
        try {
          const e = yield this.getText('avatar');
          if (null == e) return null;
          for (let r = 0; r < mh.length; r++) {
            const n = e.match(mh[r]);
            if (null == n) continue;
            const i = n[1].toLowerCase();
            switch (i) {
              case 'https':
                return (
                  t.push({ type: 'url', content: e }), { linkage: t, url: e }
                );
              case 'data':
                return (
                  t.push({ type: 'data', content: e }), { linkage: t, url: e }
                );
              case 'ipfs':
                return (
                  t.push({ type: 'ipfs', content: e }),
                  { linkage: t, url: yh(e) }
                );
              case 'erc721':
              case 'erc1155': {
                const r = 'erc721' === i ? '0xc87b56dd' : '0x0e89341c';
                t.push({ type: i, content: e });
                const o = this._resolvedAddress || (yield this.getAddress()),
                  s = (n[2] || '').split('/');
                if (2 !== s.length) return null;
                const a = yield this.provider.formatter.address(s[0]),
                  u = C(j.from(s[1]).toHexString(), 32);
                if ('erc721' === i) {
                  const e = this.provider.formatter.callAddress(
                    yield this.provider.call({
                      to: a,
                      data: P(['0x6352211e', u]),
                    })
                  );
                  if (o !== e) return null;
                  t.push({ type: 'owner', content: e });
                } else if ('erc1155' === i) {
                  const e = j.from(
                    yield this.provider.call({
                      to: a,
                      data: P(['0x00fdd58e', C(o, 32), u]),
                    })
                  );
                  if (e.isZero()) return null;
                  t.push({ type: 'balance', content: e.toString() });
                }
                const h = {
                  to: this.provider.formatter.address(s[0]),
                  data: P([r, u]),
                };
                let l = ph(yield this.provider.call(h), 0);
                if (null == l) return null;
                t.push({ type: 'metadata-url-base', content: l }),
                  'erc1155' === i &&
                    ((l = l.replace('{id}', u.substring(2))),
                    t.push({ type: 'metadata-url-expanded', content: l })),
                  l.match(/^ipfs:/i) && (l = yh(l)),
                  t.push({ type: 'metadata-url', content: l });
                const c = yield Fu(l);
                if (!c) return null;
                t.push({ type: 'metadata', content: JSON.stringify(c) });
                let f = c.image;
                if ('string' != typeof f) return null;
                if (f.match(/^(https:\/\/|data:)/i));
                else {
                  if (null == f.match(dh)) return null;
                  t.push({ type: 'url-ipfs', content: f }), (f = yh(f));
                }
                return (
                  t.push({ type: 'url', content: f }), { linkage: t, url: f }
                );
              }
            }
          }
        } catch (t) {}
        return null;
      });
    }
    getContentHash() {
      return eh(this, void 0, void 0, function* () {
        const t = yield this._fetchBytes('0xbc1c58d1');
        if (null == t || '0x' === t) return null;
        const e = t.match(
          /^0xe3010170(([0-9a-f][0-9a-f])([0-9a-f][0-9a-f])([0-9a-f]*))$/
        );
        if (e) {
          const t = parseInt(e[3], 16);
          if (e[4].length === 2 * t) return 'ipfs://' + pa.encode('0x' + e[1]);
        }
        const r = t.match(
          /^0xe5010172(([0-9a-f][0-9a-f])([0-9a-f][0-9a-f])([0-9a-f]*))$/
        );
        if (r) {
          const t = parseInt(r[3], 16);
          if (r[4].length === 2 * t) return 'ipns://' + pa.encode('0x' + r[1]);
        }
        const n = t.match(/^0xe40101fa011b20([0-9a-f]*)$/);
        if (n && 64 === n[1].length) return 'bzz://' + n[1];
        const i = t.match(/^0x90b2c605([0-9a-f]*)$/);
        if (i && 68 === i[1].length) {
          const t = { '=': '', '+': '-', '/': '_' };
          return 'sia://' + Ir('0x' + i[1]).replace(/[=+\/]/g, (e) => t[e]);
        }
        return rh.throwError(
          'invalid or unsupported content hash data',
          f.errors.UNSUPPORTED_OPERATION,
          { operation: 'getContentHash()', data: t }
        );
      });
    }
    getText(t) {
      return eh(this, void 0, void 0, function* () {
        let e = hr(t);
        (e = b([ch(64), ch(e.length), e])),
          e.length % 32 != 0 && (e = b([e, C('0x', 32 - (t.length % 32))]));
        const r = yield this._fetchBytes('0x59d1d43c', k(e));
        return null == r || '0x' === r ? null : fr(r);
      });
    }
  }
  let Eh = null,
    Mh = 1;
  class _h extends On {
    constructor(t) {
      if (
        (super(),
        (this._events = []),
        (this._emitted = { block: -2 }),
        (this.disableCcipRead = !1),
        (this.formatter = new.target.getFormatter()),
        st(this, 'anyNetwork', 'any' === t),
        this.anyNetwork && (t = this.detectNetwork()),
        t instanceof Promise)
      )
        (this._networkPromise = t),
          t.catch((t) => {}),
          this._ready().catch((t) => {});
      else {
        const e = at(new.target, 'getNetwork')(t);
        e
          ? (st(this, '_network', e), this.emit('network', e, null))
          : rh.throwArgumentError('invalid network', 'network', t);
      }
      (this._maxInternalBlockNumber = -1024),
        (this._lastBlockNumber = -2),
        (this._maxFilterBlockRange = 10),
        (this._pollingInterval = 4e3),
        (this._fastQueryDate = 0);
    }
    _ready() {
      return eh(this, void 0, void 0, function* () {
        if (null == this._network) {
          let t = null;
          if (this._networkPromise)
            try {
              t = yield this._networkPromise;
            } catch (t) {}
          null == t && (t = yield this.detectNetwork()),
            t ||
              rh.throwError('no network detected', f.errors.UNKNOWN_ERROR, {}),
            null == this._network &&
              (this.anyNetwork ? (this._network = t) : st(this, '_network', t),
              this.emit('network', t, null));
        }
        return this._network;
      });
    }
    get ready() {
      return Du(() =>
        this._ready().then(
          (t) => t,
          (t) => {
            if (t.code !== f.errors.NETWORK_ERROR || 'noNetwork' !== t.event)
              throw t;
          }
        )
      );
    }
    static getFormatter() {
      return null == Eh && (Eh = new Yu()), Eh;
    }
    static getNetwork(t) {
      return Pu(null == t ? 'homestead' : t);
    }
    ccipReadFetch(t, e, r) {
      return eh(this, void 0, void 0, function* () {
        if (this.disableCcipRead || 0 === r.length) return null;
        const n = t.to.toLowerCase(),
          i = e.toLowerCase(),
          o = [];
        for (let t = 0; t < r.length; t++) {
          const e = r[t],
            s = e.replace('{sender}', n).replace('{data}', i),
            a =
              e.indexOf('{data}') >= 0
                ? null
                : JSON.stringify({ data: i, sender: n }),
            u = yield Fu(
              { url: s, errorPassThrough: !0 },
              a,
              (t, e) => ((t.status = e.statusCode), t)
            );
          if (u.data) return u.data;
          const h = u.message || 'unknown error';
          if (u.status >= 400 && u.status < 500)
            return rh.throwError(
              `response not found during CCIP fetch: ${h}`,
              f.errors.SERVER_ERROR,
              { url: e, errorMessage: h }
            );
          o.push(h);
        }
        return rh.throwError(
          `error encountered during CCIP fetch: ${o
            .map((t) => JSON.stringify(t))
            .join(', ')}`,
          f.errors.SERVER_ERROR,
          { urls: r, errorMessages: o }
        );
      });
    }
    _getInternalBlockNumber(t) {
      return eh(this, void 0, void 0, function* () {
        if ((yield this._ready(), t > 0))
          for (; this._internalBlockNumber; ) {
            const e = this._internalBlockNumber;
            try {
              const r = yield e;
              if (sh() - r.respTime <= t) return r.blockNumber;
              break;
            } catch (t) {
              if (this._internalBlockNumber === e) break;
            }
          }
        const e = sh(),
          r = ut({
            blockNumber: this.perform('getBlockNumber', {}),
            networkError: this.getNetwork().then(
              (t) => null,
              (t) => t
            ),
          }).then(({ blockNumber: t, networkError: n }) => {
            if (n)
              throw (
                (this._internalBlockNumber === r &&
                  (this._internalBlockNumber = null),
                n)
              );
            const i = sh();
            return (
              (t = j.from(t).toNumber()) < this._maxInternalBlockNumber &&
                (t = this._maxInternalBlockNumber),
              (this._maxInternalBlockNumber = t),
              this._setFastBlockNumber(t),
              { blockNumber: t, reqTime: e, respTime: i }
            );
          });
        return (
          (this._internalBlockNumber = r),
          r.catch((t) => {
            this._internalBlockNumber === r &&
              (this._internalBlockNumber = null);
          }),
          (yield r).blockNumber
        );
      });
    }
    poll() {
      return eh(this, void 0, void 0, function* () {
        const t = Mh++,
          e = [];
        let r = null;
        try {
          r = yield this._getInternalBlockNumber(
            100 + this.pollingInterval / 2
          );
        } catch (t) {
          return void this.emit('error', t);
        }
        if (
          (this._setFastBlockNumber(r),
          this.emit('poll', t, r),
          r !== this._lastBlockNumber)
        ) {
          if (
            (-2 === this._emitted.block && (this._emitted.block = r - 1),
            Math.abs(this._emitted.block - r) > 1e3)
          )
            rh.warn(
              `network block skew detected; skipping block events (emitted=${this._emitted.block} blockNumber${r})`
            ),
              this.emit(
                'error',
                rh.makeError(
                  'network block skew detected',
                  f.errors.NETWORK_ERROR,
                  {
                    blockNumber: r,
                    event: 'blockSkew',
                    previousBlockNumber: this._emitted.block,
                  }
                )
              ),
              this.emit('block', r);
          else
            for (let t = this._emitted.block + 1; t <= r; t++)
              this.emit('block', t);
          this._emitted.block !== r &&
            ((this._emitted.block = r),
            Object.keys(this._emitted).forEach((t) => {
              if ('block' === t) return;
              const e = this._emitted[t];
              'pending' !== e && r - e > 12 && delete this._emitted[t];
            })),
            -2 === this._lastBlockNumber && (this._lastBlockNumber = r - 1),
            this._events.forEach((t) => {
              switch (t.type) {
                case 'tx': {
                  const r = t.hash;
                  let n = this.getTransactionReceipt(r)
                    .then((t) =>
                      t && null != t.blockNumber
                        ? ((this._emitted['t:' + r] = t.blockNumber),
                          this.emit(r, t),
                          null)
                        : null
                    )
                    .catch((t) => {
                      this.emit('error', t);
                    });
                  e.push(n);
                  break;
                }
                case 'filter':
                  if (!t._inflight) {
                    (t._inflight = !0),
                      -2 === t._lastBlockNumber && (t._lastBlockNumber = r - 1);
                    const n = t.filter;
                    (n.fromBlock = t._lastBlockNumber + 1), (n.toBlock = r);
                    const i = n.toBlock - this._maxFilterBlockRange;
                    i > n.fromBlock && (n.fromBlock = i),
                      n.fromBlock < 0 && (n.fromBlock = 0);
                    const o = this.getLogs(n)
                      .then((e) => {
                        (t._inflight = !1),
                          0 !== e.length &&
                            e.forEach((e) => {
                              e.blockNumber > t._lastBlockNumber &&
                                (t._lastBlockNumber = e.blockNumber),
                                (this._emitted['b:' + e.blockHash] =
                                  e.blockNumber),
                                (this._emitted['t:' + e.transactionHash] =
                                  e.blockNumber),
                                this.emit(n, e);
                            });
                      })
                      .catch((e) => {
                        this.emit('error', e), (t._inflight = !1);
                      });
                    e.push(o);
                  }
              }
            }),
            (this._lastBlockNumber = r),
            Promise.all(e)
              .then(() => {
                this.emit('didPoll', t);
              })
              .catch((t) => {
                this.emit('error', t);
              });
        } else this.emit('didPoll', t);
      });
    }
    resetEventsBlock(t) {
      (this._lastBlockNumber = t - 1), this.polling && this.poll();
    }
    get network() {
      return this._network;
    }
    detectNetwork() {
      return eh(this, void 0, void 0, function* () {
        return rh.throwError(
          'provider does not support network detection',
          f.errors.UNSUPPORTED_OPERATION,
          { operation: 'provider.detectNetwork' }
        );
      });
    }
    getNetwork() {
      return eh(this, void 0, void 0, function* () {
        const t = yield this._ready(),
          e = yield this.detectNetwork();
        if (t.chainId !== e.chainId) {
          if (this.anyNetwork)
            return (
              (this._network = e),
              (this._lastBlockNumber = -2),
              (this._fastBlockNumber = null),
              (this._fastBlockNumberPromise = null),
              (this._fastQueryDate = 0),
              (this._emitted.block = -2),
              (this._maxInternalBlockNumber = -1024),
              (this._internalBlockNumber = null),
              this.emit('network', e, t),
              yield ah(0),
              this._network
            );
          const r = rh.makeError(
            'underlying network changed',
            f.errors.NETWORK_ERROR,
            { event: 'changed', network: t, detectedNetwork: e }
          );
          throw (this.emit('error', r), r);
        }
        return t;
      });
    }
    get blockNumber() {
      return (
        this._getInternalBlockNumber(100 + this.pollingInterval / 2).then(
          (t) => {
            this._setFastBlockNumber(t);
          },
          (t) => {}
        ),
        null != this._fastBlockNumber ? this._fastBlockNumber : -1
      );
    }
    get polling() {
      return null != this._poller;
    }
    set polling(t) {
      t && !this._poller
        ? ((this._poller = setInterval(() => {
            this.poll();
          }, this.pollingInterval)),
          this._bootstrapPoll ||
            (this._bootstrapPoll = setTimeout(() => {
              this.poll(),
                (this._bootstrapPoll = setTimeout(() => {
                  this._poller || this.poll(), (this._bootstrapPoll = null);
                }, this.pollingInterval));
            }, 0)))
        : !t &&
          this._poller &&
          (clearInterval(this._poller), (this._poller = null));
    }
    get pollingInterval() {
      return this._pollingInterval;
    }
    set pollingInterval(t) {
      if ('number' != typeof t || t <= 0 || parseInt(String(t)) != t)
        throw new Error('invalid polling interval');
      (this._pollingInterval = t),
        this._poller &&
          (clearInterval(this._poller),
          (this._poller = setInterval(() => {
            this.poll();
          }, this._pollingInterval)));
    }
    _getFastBlockNumber() {
      const t = sh();
      return (
        t - this._fastQueryDate > 2 * this._pollingInterval &&
          ((this._fastQueryDate = t),
          (this._fastBlockNumberPromise = this.getBlockNumber().then(
            (t) => (
              (null == this._fastBlockNumber || t > this._fastBlockNumber) &&
                (this._fastBlockNumber = t),
              this._fastBlockNumber
            )
          ))),
        this._fastBlockNumberPromise
      );
    }
    _setFastBlockNumber(t) {
      (null != this._fastBlockNumber && t < this._fastBlockNumber) ||
        ((this._fastQueryDate = sh()),
        (null == this._fastBlockNumber || t > this._fastBlockNumber) &&
          ((this._fastBlockNumber = t),
          (this._fastBlockNumberPromise = Promise.resolve(t))));
    }
    waitForTransaction(t, e, r) {
      return eh(this, void 0, void 0, function* () {
        return this._waitForTransaction(t, null == e ? 1 : e, r || 0, null);
      });
    }
    _waitForTransaction(t, e, r, n) {
      return eh(this, void 0, void 0, function* () {
        const i = yield this.getTransactionReceipt(t);
        return (i ? i.confirmations : 0) >= e
          ? i
          : new Promise((i, o) => {
              const s = [];
              let a = !1;
              const u = function () {
                  return (
                    !!a ||
                    ((a = !0),
                    s.forEach((t) => {
                      t();
                    }),
                    !1)
                  );
                },
                h = (t) => {
                  t.confirmations < e || u() || i(t);
                };
              if (
                (this.on(t, h),
                s.push(() => {
                  this.removeListener(t, h);
                }),
                n)
              ) {
                let r = n.startBlock,
                  i = null;
                const h = (s) =>
                  eh(this, void 0, void 0, function* () {
                    a ||
                      (yield ah(1e3),
                      this.getTransactionCount(n.from).then(
                        (l) =>
                          eh(this, void 0, void 0, function* () {
                            if (!a) {
                              if (l <= n.nonce) r = s;
                              else {
                                {
                                  const e = yield this.getTransaction(t);
                                  if (e && null != e.blockNumber) return;
                                }
                                for (
                                  null == i &&
                                  ((i = r - 3),
                                  i < n.startBlock && (i = n.startBlock));
                                  i <= s;

                                ) {
                                  if (a) return;
                                  const r = yield this.getBlockWithTransactions(
                                    i
                                  );
                                  for (
                                    let i = 0;
                                    i < r.transactions.length;
                                    i++
                                  ) {
                                    const s = r.transactions[i];
                                    if (s.hash === t) return;
                                    if (
                                      s.from === n.from &&
                                      s.nonce === n.nonce
                                    ) {
                                      if (a) return;
                                      const r = yield this.waitForTransaction(
                                        s.hash,
                                        e
                                      );
                                      if (u()) return;
                                      let i = 'replaced';
                                      return (
                                        s.data === n.data &&
                                        s.to === n.to &&
                                        s.value.eq(n.value)
                                          ? (i = 'repriced')
                                          : '0x' === s.data &&
                                            s.from === s.to &&
                                            s.value.isZero() &&
                                            (i = 'cancelled'),
                                        void o(
                                          rh.makeError(
                                            'transaction was replaced',
                                            f.errors.TRANSACTION_REPLACED,
                                            {
                                              cancelled:
                                                'replaced' === i ||
                                                'cancelled' === i,
                                              reason: i,
                                              replacement:
                                                this._wrapTransaction(s),
                                              hash: t,
                                              receipt: r,
                                            }
                                          )
                                        )
                                      );
                                    }
                                  }
                                  i++;
                                }
                              }
                              a || this.once('block', h);
                            }
                          }),
                        (t) => {
                          a || this.once('block', h);
                        }
                      ));
                  });
                if (a) return;
                this.once('block', h),
                  s.push(() => {
                    this.removeListener('block', h);
                  });
              }
              if ('number' == typeof r && r > 0) {
                const t = setTimeout(() => {
                  u() ||
                    o(
                      rh.makeError('timeout exceeded', f.errors.TIMEOUT, {
                        timeout: r,
                      })
                    );
                }, r);
                t.unref && t.unref(),
                  s.push(() => {
                    clearTimeout(t);
                  });
              }
            });
      });
    }
    getBlockNumber() {
      return eh(this, void 0, void 0, function* () {
        return this._getInternalBlockNumber(0);
      });
    }
    getGasPrice() {
      return eh(this, void 0, void 0, function* () {
        yield this.getNetwork();
        const t = yield this.perform('getGasPrice', {});
        try {
          return j.from(t);
        } catch (e) {
          return rh.throwError(
            'bad result from backend',
            f.errors.SERVER_ERROR,
            { method: 'getGasPrice', result: t, error: e }
          );
        }
      });
    }
    getBalance(t, e) {
      return eh(this, void 0, void 0, function* () {
        yield this.getNetwork();
        const r = yield ut({
            address: this._getAddress(t),
            blockTag: this._getBlockTag(e),
          }),
          n = yield this.perform('getBalance', r);
        try {
          return j.from(n);
        } catch (t) {
          return rh.throwError(
            'bad result from backend',
            f.errors.SERVER_ERROR,
            { method: 'getBalance', params: r, result: n, error: t }
          );
        }
      });
    }
    getTransactionCount(t, e) {
      return eh(this, void 0, void 0, function* () {
        yield this.getNetwork();
        const r = yield ut({
            address: this._getAddress(t),
            blockTag: this._getBlockTag(e),
          }),
          n = yield this.perform('getTransactionCount', r);
        try {
          return j.from(n).toNumber();
        } catch (t) {
          return rh.throwError(
            'bad result from backend',
            f.errors.SERVER_ERROR,
            { method: 'getTransactionCount', params: r, result: n, error: t }
          );
        }
      });
    }
    getCode(t, e) {
      return eh(this, void 0, void 0, function* () {
        yield this.getNetwork();
        const r = yield ut({
            address: this._getAddress(t),
            blockTag: this._getBlockTag(e),
          }),
          n = yield this.perform('getCode', r);
        try {
          return k(n);
        } catch (t) {
          return rh.throwError(
            'bad result from backend',
            f.errors.SERVER_ERROR,
            { method: 'getCode', params: r, result: n, error: t }
          );
        }
      });
    }
    getStorageAt(t, e, r) {
      return eh(this, void 0, void 0, function* () {
        yield this.getNetwork();
        const n = yield ut({
            address: this._getAddress(t),
            blockTag: this._getBlockTag(r),
            position: Promise.resolve(e).then((t) => x(t)),
          }),
          i = yield this.perform('getStorageAt', n);
        try {
          return k(i);
        } catch (t) {
          return rh.throwError(
            'bad result from backend',
            f.errors.SERVER_ERROR,
            { method: 'getStorageAt', params: n, result: i, error: t }
          );
        }
      });
    }
    _wrapTransaction(t, e, r) {
      if (null != e && 32 !== S(e))
        throw new Error('invalid response - sendTransaction');
      const n = t;
      return (
        null != e &&
          t.hash !== e &&
          rh.throwError(
            'Transaction hash mismatch from Provider.sendTransaction.',
            f.errors.UNKNOWN_ERROR,
            { expectedHash: t.hash, returnedHash: e }
          ),
        (n.wait = (e, n) =>
          eh(this, void 0, void 0, function* () {
            let i;
            null == e && (e = 1),
              null == n && (n = 0),
              0 !== e &&
                null != r &&
                (i = {
                  data: t.data,
                  from: t.from,
                  nonce: t.nonce,
                  to: t.to,
                  value: t.value,
                  startBlock: r,
                });
            const o = yield this._waitForTransaction(t.hash, e, n, i);
            return null == o && 0 === e
              ? null
              : ((this._emitted['t:' + t.hash] = o.blockNumber),
                0 === o.status &&
                  rh.throwError('transaction failed', f.errors.CALL_EXCEPTION, {
                    transactionHash: t.hash,
                    transaction: t,
                    receipt: o,
                  }),
                o);
          })),
        n
      );
    }
    sendTransaction(t) {
      return eh(this, void 0, void 0, function* () {
        yield this.getNetwork();
        const e = yield Promise.resolve(t).then((t) => k(t)),
          r = this.formatter.transaction(t);
        null == r.confirmations && (r.confirmations = 0);
        const n = yield this._getInternalBlockNumber(
          100 + 2 * this.pollingInterval
        );
        try {
          const t = yield this.perform('sendTransaction', {
            signedTransaction: e,
          });
          return this._wrapTransaction(r, t, n);
        } catch (t) {
          throw ((t.transaction = r), (t.transactionHash = r.hash), t);
        }
      });
    }
    _getTransactionRequest(t) {
      return eh(this, void 0, void 0, function* () {
        const e = yield t,
          r = {};
        return (
          ['from', 'to'].forEach((t) => {
            null != e[t] &&
              (r[t] = Promise.resolve(e[t]).then((t) =>
                t ? this._getAddress(t) : null
              ));
          }),
          [
            'gasLimit',
            'gasPrice',
            'maxFeePerGas',
            'maxPriorityFeePerGas',
            'value',
          ].forEach((t) => {
            null != e[t] &&
              (r[t] = Promise.resolve(e[t]).then((t) =>
                t ? j.from(t) : null
              ));
          }),
          ['type'].forEach((t) => {
            null != e[t] &&
              (r[t] = Promise.resolve(e[t]).then((t) =>
                null != t ? t : null
              ));
          }),
          e.accessList &&
            (r.accessList = this.formatter.accessList(e.accessList)),
          ['data'].forEach((t) => {
            null != e[t] &&
              (r[t] = Promise.resolve(e[t]).then((t) => (t ? k(t) : null)));
          }),
          this.formatter.transactionRequest(yield ut(r))
        );
      });
    }
    _getFilter(t) {
      return eh(this, void 0, void 0, function* () {
        t = yield t;
        const e = {};
        return (
          null != t.address && (e.address = this._getAddress(t.address)),
          ['blockHash', 'topics'].forEach((r) => {
            null != t[r] && (e[r] = t[r]);
          }),
          ['fromBlock', 'toBlock'].forEach((r) => {
            null != t[r] && (e[r] = this._getBlockTag(t[r]));
          }),
          this.formatter.filter(yield ut(e))
        );
      });
    }
    _call(t, e, r) {
      return eh(this, void 0, void 0, function* () {
        r >= 10 &&
          rh.throwError(
            'CCIP read exceeded maximum redirections',
            f.errors.SERVER_ERROR,
            { redirects: r, transaction: t }
          );
        const n = t.to,
          i = yield this.perform('call', { transaction: t, blockTag: e });
        if (
          r >= 0 &&
          'latest' === e &&
          null != n &&
          '0x556f1830' === i.substring(0, 10) &&
          S(i) % 32 == 4
        )
          try {
            const o = N(i, 4),
              s = N(o, 0, 32);
            j.from(s).eq(n) ||
              rh.throwError(
                'CCIP Read sender did not match',
                f.errors.CALL_EXCEPTION,
                {
                  name: 'OffchainLookup',
                  signature:
                    'OffchainLookup(address,string[],bytes,bytes4,bytes)',
                  transaction: t,
                  data: i,
                }
              );
            const a = [],
              u = j.from(N(o, 32, 64)).toNumber(),
              h = j.from(N(o, u, u + 32)).toNumber(),
              l = N(o, u + 32);
            for (let e = 0; e < h; e++) {
              const r = ph(l, 32 * e);
              null == r &&
                rh.throwError(
                  'CCIP Read contained corrupt URL string',
                  f.errors.CALL_EXCEPTION,
                  {
                    name: 'OffchainLookup',
                    signature:
                      'OffchainLookup(address,string[],bytes,bytes4,bytes)',
                    transaction: t,
                    data: i,
                  }
                ),
                a.push(r);
            }
            const c = gh(o, 64);
            j.from(N(o, 100, 128)).isZero() ||
              rh.throwError(
                'CCIP Read callback selector included junk',
                f.errors.CALL_EXCEPTION,
                {
                  name: 'OffchainLookup',
                  signature:
                    'OffchainLookup(address,string[],bytes,bytes4,bytes)',
                  transaction: t,
                  data: i,
                }
              );
            const d = N(o, 96, 100),
              m = gh(o, 128),
              p = yield this.ccipReadFetch(t, c, a);
            null == p &&
              rh.throwError(
                'CCIP Read disabled or provided no URLs',
                f.errors.CALL_EXCEPTION,
                {
                  name: 'OffchainLookup',
                  signature:
                    'OffchainLookup(address,string[],bytes,bytes4,bytes)',
                  transaction: t,
                  data: i,
                }
              );
            const g = { to: n, data: P([d, bh([p, m])]) };
            return this._call(g, e, r + 1);
          } catch (t) {
            if (t.code === f.errors.SERVER_ERROR) throw t;
          }
        try {
          return k(i);
        } catch (r) {
          return rh.throwError(
            'bad result from backend',
            f.errors.SERVER_ERROR,
            {
              method: 'call',
              params: { transaction: t, blockTag: e },
              result: i,
              error: r,
            }
          );
        }
      });
    }
    call(t, e) {
      return eh(this, void 0, void 0, function* () {
        yield this.getNetwork();
        const r = yield ut({
          transaction: this._getTransactionRequest(t),
          blockTag: this._getBlockTag(e),
          ccipReadEnabled: Promise.resolve(t.ccipReadEnabled),
        });
        return this._call(
          r.transaction,
          r.blockTag,
          r.ccipReadEnabled ? 0 : -1
        );
      });
    }
    estimateGas(t) {
      return eh(this, void 0, void 0, function* () {
        yield this.getNetwork();
        const e = yield ut({ transaction: this._getTransactionRequest(t) }),
          r = yield this.perform('estimateGas', e);
        try {
          return j.from(r);
        } catch (t) {
          return rh.throwError(
            'bad result from backend',
            f.errors.SERVER_ERROR,
            { method: 'estimateGas', params: e, result: r, error: t }
          );
        }
      });
    }
    _getAddress(t) {
      return eh(this, void 0, void 0, function* () {
        'string' != typeof (t = yield t) &&
          rh.throwArgumentError('invalid address or ENS name', 'name', t);
        const e = yield this.resolveName(t);
        return (
          null == e &&
            rh.throwError(
              'ENS name not configured',
              f.errors.UNSUPPORTED_OPERATION,
              { operation: `resolveName(${JSON.stringify(t)})` }
            ),
          e
        );
      });
    }
    _getBlock(t, e) {
      return eh(this, void 0, void 0, function* () {
        yield this.getNetwork(), (t = yield t);
        let r = -128;
        const n = { includeTransactions: !!e };
        if (M(t, 32)) n.blockHash = t;
        else
          try {
            (n.blockTag = yield this._getBlockTag(t)),
              M(n.blockTag) && (r = parseInt(n.blockTag.substring(2), 16));
          } catch (e) {
            rh.throwArgumentError(
              'invalid block hash or block tag',
              'blockHashOrBlockTag',
              t
            );
          }
        return Du(
          () =>
            eh(this, void 0, void 0, function* () {
              const t = yield this.perform('getBlock', n);
              if (null == t)
                return (null != n.blockHash &&
                  null == this._emitted['b:' + n.blockHash]) ||
                  (null != n.blockTag && r > this._emitted.block)
                  ? null
                  : void 0;
              if (e) {
                let e = null;
                for (let r = 0; r < t.transactions.length; r++) {
                  const n = t.transactions[r];
                  if (null == n.blockNumber) n.confirmations = 0;
                  else if (null == n.confirmations) {
                    null == e &&
                      (e = yield this._getInternalBlockNumber(
                        100 + 2 * this.pollingInterval
                      ));
                    let t = e - n.blockNumber + 1;
                    t <= 0 && (t = 1), (n.confirmations = t);
                  }
                }
                const r = this.formatter.blockWithTransactions(t);
                return (
                  (r.transactions = r.transactions.map((t) =>
                    this._wrapTransaction(t)
                  )),
                  r
                );
              }
              return this.formatter.block(t);
            }),
          { oncePoll: this }
        );
      });
    }
    getBlock(t) {
      return this._getBlock(t, !1);
    }
    getBlockWithTransactions(t) {
      return this._getBlock(t, !0);
    }
    getTransaction(t) {
      return eh(this, void 0, void 0, function* () {
        yield this.getNetwork(), (t = yield t);
        const e = { transactionHash: this.formatter.hash(t, !0) };
        return Du(
          () =>
            eh(this, void 0, void 0, function* () {
              const r = yield this.perform('getTransaction', e);
              if (null == r)
                return null == this._emitted['t:' + t] ? null : void 0;
              const n = this.formatter.transactionResponse(r);
              if (null == n.blockNumber) n.confirmations = 0;
              else if (null == n.confirmations) {
                let t =
                  (yield this._getInternalBlockNumber(
                    100 + 2 * this.pollingInterval
                  )) -
                  n.blockNumber +
                  1;
                t <= 0 && (t = 1), (n.confirmations = t);
              }
              return this._wrapTransaction(n);
            }),
          { oncePoll: this }
        );
      });
    }
    getTransactionReceipt(t) {
      return eh(this, void 0, void 0, function* () {
        yield this.getNetwork(), (t = yield t);
        const e = { transactionHash: this.formatter.hash(t, !0) };
        return Du(
          () =>
            eh(this, void 0, void 0, function* () {
              const r = yield this.perform('getTransactionReceipt', e);
              if (null == r)
                return null == this._emitted['t:' + t] ? null : void 0;
              if (null == r.blockHash) return;
              const n = this.formatter.receipt(r);
              if (null == n.blockNumber) n.confirmations = 0;
              else if (null == n.confirmations) {
                let t =
                  (yield this._getInternalBlockNumber(
                    100 + 2 * this.pollingInterval
                  )) -
                  n.blockNumber +
                  1;
                t <= 0 && (t = 1), (n.confirmations = t);
              }
              return n;
            }),
          { oncePoll: this }
        );
      });
    }
    getLogs(t) {
      return eh(this, void 0, void 0, function* () {
        yield this.getNetwork();
        const e = yield ut({ filter: this._getFilter(t) }),
          r = yield this.perform('getLogs', e);
        return (
          r.forEach((t) => {
            null == t.removed && (t.removed = !1);
          }),
          Yu.arrayOf(this.formatter.filterLog.bind(this.formatter))(r)
        );
      });
    }
    getEtherPrice() {
      return eh(this, void 0, void 0, function* () {
        return yield this.getNetwork(), this.perform('getEtherPrice', {});
      });
    }
    _getBlockTag(t) {
      return eh(this, void 0, void 0, function* () {
        if ('number' == typeof (t = yield t) && t < 0) {
          t % 1 && rh.throwArgumentError('invalid BlockTag', 'blockTag', t);
          let e = yield this._getInternalBlockNumber(
            100 + 2 * this.pollingInterval
          );
          return (e += t), e < 0 && (e = 0), this.formatter.blockTag(e);
        }
        return this.formatter.blockTag(t);
      });
    }
    getResolver(t) {
      return eh(this, void 0, void 0, function* () {
        let e = t;
        for (;;) {
          if ('' === e || '.' === e) return null;
          if ('eth' !== t && 'eth' === e) return null;
          const r = yield this._getResolver(e, 'getResolver');
          if (null != r) {
            const n = new Ah(this, r, t);
            return e === t || (yield n.supportsWildcard()) ? n : null;
          }
          e = e.split('.').slice(1).join('.');
        }
      });
    }
    _getResolver(t, e) {
      return eh(this, void 0, void 0, function* () {
        null == e && (e = 'ENS');
        const r = yield this.getNetwork();
        r.ensAddress ||
          rh.throwError(
            'network does not support ENS',
            f.errors.UNSUPPORTED_OPERATION,
            { operation: e, network: r.name }
          );
        try {
          const e = yield this.call({
            to: r.ensAddress,
            data: '0x0178b8bf' + on(t).substring(2),
          });
          return this.formatter.callAddress(e);
        } catch (t) {}
        return null;
      });
    }
    resolveName(t) {
      return eh(this, void 0, void 0, function* () {
        t = yield t;
        try {
          return Promise.resolve(this.formatter.address(t));
        } catch (e) {
          if (M(t)) throw e;
        }
        'string' != typeof t &&
          rh.throwArgumentError('invalid ENS name', 'name', t);
        const e = yield this.getResolver(t);
        return e ? yield e.getAddress() : null;
      });
    }
    lookupAddress(t) {
      return eh(this, void 0, void 0, function* () {
        t = yield t;
        const e =
            (t = this.formatter.address(t)).substring(2).toLowerCase() +
            '.addr.reverse',
          r = yield this._getResolver(e, 'lookupAddress');
        if (null == r) return null;
        const n = ph(
          yield this.call({ to: r, data: '0x691f3431' + on(e).substring(2) }),
          0
        );
        return (yield this.resolveName(n)) != t ? null : n;
      });
    }
    getAvatar(t) {
      return eh(this, void 0, void 0, function* () {
        let e = null;
        if (M(t)) {
          const r =
              this.formatter.address(t).substring(2).toLowerCase() +
              '.addr.reverse',
            n = yield this._getResolver(r, 'getAvatar');
          if (!n) return null;
          e = new Ah(this, n, r);
          try {
            const t = yield e.getAvatar();
            if (t) return t.url;
          } catch (t) {
            if (t.code !== f.errors.CALL_EXCEPTION) throw t;
          }
          try {
            const t = ph(
              yield this.call({
                to: n,
                data: '0x691f3431' + on(r).substring(2),
              }),
              0
            );
            e = yield this.getResolver(t);
          } catch (t) {
            if (t.code !== f.errors.CALL_EXCEPTION) throw t;
            return null;
          }
        } else if (((e = yield this.getResolver(t)), !e)) return null;
        const r = yield e.getAvatar();
        return null == r ? null : r.url;
      });
    }
    perform(t, e) {
      return rh.throwError(t + ' not implemented', f.errors.NOT_IMPLEMENTED, {
        operation: t,
      });
    }
    _startEvent(t) {
      this.polling = this._events.filter((t) => t.pollable()).length > 0;
    }
    _stopEvent(t) {
      this.polling = this._events.filter((t) => t.pollable()).length > 0;
    }
    _addEventListener(t, e, r) {
      const n = new hh(oh(t), e, r);
      return this._events.push(n), this._startEvent(n), this;
    }
    on(t, e) {
      return this._addEventListener(t, e, !1);
    }
    once(t, e) {
      return this._addEventListener(t, e, !0);
    }
    emit(t, ...e) {
      let r = !1,
        n = [],
        i = oh(t);
      return (
        (this._events = this._events.filter(
          (t) =>
            t.tag !== i ||
            (setTimeout(() => {
              t.listener.apply(this, e);
            }, 0),
            (r = !0),
            !t.once || (n.push(t), !1))
        )),
        n.forEach((t) => {
          this._stopEvent(t);
        }),
        r
      );
    }
    listenerCount(t) {
      if (!t) return this._events.length;
      let e = oh(t);
      return this._events.filter((t) => t.tag === e).length;
    }
    listeners(t) {
      if (null == t) return this._events.map((t) => t.listener);
      let e = oh(t);
      return this._events.filter((t) => t.tag === e).map((t) => t.listener);
    }
    off(t, e) {
      if (null == e) return this.removeAllListeners(t);
      const r = [];
      let n = !1,
        i = oh(t);
      return (
        (this._events = this._events.filter(
          (t) =>
            t.tag !== i || t.listener != e || !!n || ((n = !0), r.push(t), !1)
        )),
        r.forEach((t) => {
          this._stopEvent(t);
        }),
        this
      );
    }
    removeAllListeners(t) {
      let e = [];
      if (null == t) (e = this._events), (this._events = []);
      else {
        const r = oh(t);
        this._events = this._events.filter(
          (t) => t.tag !== r || (e.push(t), !1)
        );
      }
      return (
        e.forEach((t) => {
          this._stopEvent(t);
        }),
        this
      );
    }
  }
  var kh = function (t, e, r, n) {
    return new (r || (r = Promise))(function (i, o) {
      function s(t) {
        try {
          u(n.next(t));
        } catch (t) {
          o(t);
        }
      }
      function a(t) {
        try {
          u(n.throw(t));
        } catch (t) {
          o(t);
        }
      }
      function u(t) {
        var e;
        t.done
          ? i(t.value)
          : ((e = t.value),
            e instanceof r
              ? e
              : new r(function (t) {
                  t(e);
                })).then(s, a);
      }
      u((n = n.apply(t, e || [])).next());
    });
  };
  const Sh = new f(Wu),
    Nh = ['call', 'estimateGas'];
  function Ph(t, e) {
    if (null == t) return null;
    if ('string' == typeof t.message && t.message.match('reverted')) {
      const r = M(t.data) ? t.data : null;
      if (!e || r) return { message: t.message, data: r };
    }
    if ('object' == typeof t) {
      for (const r in t) {
        const n = Ph(t[r], e);
        if (n) return n;
      }
      return null;
    }
    if ('string' == typeof t)
      try {
        return Ph(JSON.parse(t), e);
      } catch (t) {}
    return null;
  }
  function xh(t, e, r) {
    const n = r.transaction || r.signedTransaction;
    if ('call' === t) {
      const t = Ph(e, !0);
      if (t) return t.data;
      Sh.throwError(
        'missing revert data in call exception; Transaction reverted without a reason string',
        f.errors.CALL_EXCEPTION,
        { data: '0x', transaction: n, error: e }
      );
    }
    if ('estimateGas' === t) {
      let r = Ph(e.body, !1);
      null == r && (r = Ph(e, !1)),
        r &&
          Sh.throwError(
            'cannot estimate gas; transaction may fail or may require manual gas limit',
            f.errors.UNPREDICTABLE_GAS_LIMIT,
            { reason: r.message, method: t, transaction: n, error: e }
          );
    }
    let i = e.message;
    throw (
      (e.code === f.errors.SERVER_ERROR &&
      e.error &&
      'string' == typeof e.error.message
        ? (i = e.error.message)
        : 'string' == typeof e.body
        ? (i = e.body)
        : 'string' == typeof e.responseText && (i = e.responseText),
      (i = (i || '').toLowerCase()),
      i.match(
        /insufficient funds|base fee exceeds gas limit|InsufficientFunds/i
      ) &&
        Sh.throwError(
          'insufficient funds for intrinsic transaction cost',
          f.errors.INSUFFICIENT_FUNDS,
          { error: e, method: t, transaction: n }
        ),
      i.match(/nonce (is )?too low/i) &&
        Sh.throwError('nonce has already been used', f.errors.NONCE_EXPIRED, {
          error: e,
          method: t,
          transaction: n,
        }),
      i.match(
        /replacement transaction underpriced|transaction gas price.*too low/i
      ) &&
        Sh.throwError(
          'replacement fee too low',
          f.errors.REPLACEMENT_UNDERPRICED,
          { error: e, method: t, transaction: n }
        ),
      i.match(/only replay-protected/i) &&
        Sh.throwError(
          'legacy pre-eip-155 transactions not supported',
          f.errors.UNSUPPORTED_OPERATION,
          { error: e, method: t, transaction: n }
        ),
      Nh.indexOf(t) >= 0 &&
        i.match(
          /gas required exceeds allowance|always failing transaction|execution reverted|revert/
        ) &&
        Sh.throwError(
          'cannot estimate gas; transaction may fail or may require manual gas limit',
          f.errors.UNPREDICTABLE_GAS_LIMIT,
          { error: e, method: t, transaction: n }
        ),
      e)
    );
  }
  function Th(t) {
    return new Promise(function (e) {
      setTimeout(e, t);
    });
  }
  function Ch(t) {
    if (t.error) {
      const e = new Error(t.error.message);
      throw ((e.code = t.error.code), (e.data = t.error.data), e);
    }
    return t.result;
  }
  function Rh(t) {
    return t ? t.toLowerCase() : t;
  }
  const Ih = {};
  class Bh extends jn {
    constructor(t, e, r) {
      if ((super(), t !== Ih))
        throw new Error(
          'do not call the JsonRpcSigner constructor directly; use provider.getSigner'
        );
      st(this, 'provider', e),
        null == r && (r = 0),
        'string' == typeof r
          ? (st(this, '_address', this.provider.formatter.address(r)),
            st(this, '_index', null))
          : 'number' == typeof r
          ? (st(this, '_index', r), st(this, '_address', null))
          : Sh.throwArgumentError(
              'invalid address or index',
              'addressOrIndex',
              r
            );
    }
    connect(t) {
      return Sh.throwError(
        'cannot alter JSON-RPC Signer connection',
        f.errors.UNSUPPORTED_OPERATION,
        { operation: 'connect' }
      );
    }
    connectUnchecked() {
      return new Oh(Ih, this.provider, this._address || this._index);
    }
    getAddress() {
      return this._address
        ? Promise.resolve(this._address)
        : this.provider
            .send('eth_accounts', [])
            .then(
              (t) => (
                t.length <= this._index &&
                  Sh.throwError(
                    'unknown account #' + this._index,
                    f.errors.UNSUPPORTED_OPERATION,
                    { operation: 'getAddress' }
                  ),
                this.provider.formatter.address(t[this._index])
              )
            );
    }
    sendUncheckedTransaction(t) {
      t = lt(t);
      const e = this.getAddress().then((t) => (t && (t = t.toLowerCase()), t));
      if (null == t.gasLimit) {
        const r = lt(t);
        (r.from = e), (t.gasLimit = this.provider.estimateGas(r));
      }
      return (
        null != t.to &&
          (t.to = Promise.resolve(t.to).then((t) =>
            kh(this, void 0, void 0, function* () {
              if (null == t) return null;
              const e = yield this.provider.resolveName(t);
              return (
                null == e &&
                  Sh.throwArgumentError(
                    'provided ENS name resolves to null',
                    'tx.to',
                    t
                  ),
                e
              );
            })
          )),
        ut({ tx: ut(t), sender: e }).then(({ tx: e, sender: r }) => {
          null != e.from
            ? e.from.toLowerCase() !== r &&
              Sh.throwArgumentError('from address mismatch', 'transaction', t)
            : (e.from = r);
          const n = this.provider.constructor.hexlifyTransaction(e, {
            from: !0,
          });
          return this.provider.send('eth_sendTransaction', [n]).then(
            (t) => t,
            (t) => (
              'string' == typeof t.message &&
                t.message.match(/user denied/i) &&
                Sh.throwError(
                  'user rejected transaction',
                  f.errors.ACTION_REJECTED,
                  { action: 'sendTransaction', transaction: e }
                ),
              xh('sendTransaction', t, n)
            )
          );
        })
      );
    }
    signTransaction(t) {
      return Sh.throwError(
        'signing transactions is unsupported',
        f.errors.UNSUPPORTED_OPERATION,
        { operation: 'signTransaction' }
      );
    }
    sendTransaction(t) {
      return kh(this, void 0, void 0, function* () {
        const e = yield this.provider._getInternalBlockNumber(
            100 + 2 * this.provider.pollingInterval
          ),
          r = yield this.sendUncheckedTransaction(t);
        try {
          return yield Du(
            () =>
              kh(this, void 0, void 0, function* () {
                const t = yield this.provider.getTransaction(r);
                if (null !== t) return this.provider._wrapTransaction(t, r, e);
              }),
            { oncePoll: this.provider }
          );
        } catch (t) {
          throw ((t.transactionHash = r), t);
        }
      });
    }
    signMessage(t) {
      return kh(this, void 0, void 0, function* () {
        const e = 'string' == typeof t ? hr(t) : t,
          r = yield this.getAddress();
        try {
          return yield this.provider.send('personal_sign', [
            k(e),
            r.toLowerCase(),
          ]);
        } catch (e) {
          throw (
            ('string' == typeof e.message &&
              e.message.match(/user denied/i) &&
              Sh.throwError('user rejected signing', f.errors.ACTION_REJECTED, {
                action: 'signMessage',
                from: r,
                messageData: t,
              }),
            e)
          );
        }
      });
    }
    _legacySignMessage(t) {
      return kh(this, void 0, void 0, function* () {
        const e = 'string' == typeof t ? hr(t) : t,
          r = yield this.getAddress();
        try {
          return yield this.provider.send('eth_sign', [r.toLowerCase(), k(e)]);
        } catch (e) {
          throw (
            ('string' == typeof e.message &&
              e.message.match(/user denied/i) &&
              Sh.throwError('user rejected signing', f.errors.ACTION_REJECTED, {
                action: '_legacySignMessage',
                from: r,
                messageData: t,
              }),
            e)
          );
        }
      });
    }
    _signTypedData(t, e, r) {
      return kh(this, void 0, void 0, function* () {
        const n = yield Mn.resolveNames(t, e, r, (t) =>
            this.provider.resolveName(t)
          ),
          i = yield this.getAddress();
        try {
          return yield this.provider.send('eth_signTypedData_v4', [
            i.toLowerCase(),
            JSON.stringify(Mn.getPayload(n.domain, e, n.value)),
          ]);
        } catch (t) {
          throw (
            ('string' == typeof t.message &&
              t.message.match(/user denied/i) &&
              Sh.throwError('user rejected signing', f.errors.ACTION_REJECTED, {
                action: '_signTypedData',
                from: i,
                messageData: { domain: n.domain, types: e, value: n.value },
              }),
            t)
          );
        }
      });
    }
    unlock(t) {
      return kh(this, void 0, void 0, function* () {
        const e = this.provider,
          r = yield this.getAddress();
        return e.send('personal_unlockAccount', [r.toLowerCase(), t, null]);
      });
    }
  }
  class Oh extends Bh {
    sendTransaction(t) {
      return this.sendUncheckedTransaction(t).then((t) => ({
        hash: t,
        nonce: null,
        gasLimit: null,
        gasPrice: null,
        data: null,
        value: null,
        chainId: null,
        confirmations: 0,
        from: null,
        wait: (e) => this.provider.waitForTransaction(t, e),
      }));
    }
  }
  const Fh = {
    chainId: !0,
    data: !0,
    gasLimit: !0,
    gasPrice: !0,
    nonce: !0,
    to: !0,
    value: !0,
    type: !0,
    accessList: !0,
    maxFeePerGas: !0,
    maxPriorityFeePerGas: !0,
  };
  class Dh extends _h {
    constructor(t, e) {
      let r = e;
      null == r &&
        (r = new Promise((t, e) => {
          setTimeout(() => {
            this.detectNetwork().then(
              (e) => {
                t(e);
              },
              (t) => {
                e(t);
              }
            );
          }, 0);
        })),
        super(r),
        t || (t = at(this.constructor, 'defaultUrl')()),
        st(
          this,
          'connection',
          'string' == typeof t
            ? Object.freeze({ url: t })
            : Object.freeze(lt(t))
        ),
        (this._nextId = 42);
    }
    get _cache() {
      return (
        null == this._eventLoopCache && (this._eventLoopCache = {}),
        this._eventLoopCache
      );
    }
    static defaultUrl() {
      return 'http://localhost:8545';
    }
    detectNetwork() {
      return (
        this._cache.detectNetwork ||
          ((this._cache.detectNetwork = this._uncachedDetectNetwork()),
          setTimeout(() => {
            this._cache.detectNetwork = null;
          }, 0)),
        this._cache.detectNetwork
      );
    }
    _uncachedDetectNetwork() {
      return kh(this, void 0, void 0, function* () {
        yield Th(0);
        let t = null;
        try {
          t = yield this.send('eth_chainId', []);
        } catch (e) {
          try {
            t = yield this.send('net_version', []);
          } catch (t) {}
        }
        if (null != t) {
          const e = at(this.constructor, 'getNetwork');
          try {
            return e(j.from(t).toNumber());
          } catch (e) {
            return Sh.throwError(
              'could not detect network',
              f.errors.NETWORK_ERROR,
              { chainId: t, event: 'invalidNetwork', serverError: e }
            );
          }
        }
        return Sh.throwError(
          'could not detect network',
          f.errors.NETWORK_ERROR,
          { event: 'noNetwork' }
        );
      });
    }
    getSigner(t) {
      return new Bh(Ih, this, t);
    }
    getUncheckedSigner(t) {
      return this.getSigner(t).connectUnchecked();
    }
    listAccounts() {
      return this.send('eth_accounts', []).then((t) =>
        t.map((t) => this.formatter.address(t))
      );
    }
    send(t, e) {
      const r = { method: t, params: e, id: this._nextId++, jsonrpc: '2.0' };
      this.emit('debug', { action: 'request', request: mt(r), provider: this });
      const n = ['eth_chainId', 'eth_blockNumber'].indexOf(t) >= 0;
      if (n && this._cache[t]) return this._cache[t];
      const i = Fu(this.connection, JSON.stringify(r), Ch).then(
        (t) => (
          this.emit('debug', {
            action: 'response',
            request: r,
            response: t,
            provider: this,
          }),
          t
        ),
        (t) => {
          throw (
            (this.emit('debug', {
              action: 'response',
              error: t,
              request: r,
              provider: this,
            }),
            t)
          );
        }
      );
      return (
        n &&
          ((this._cache[t] = i),
          setTimeout(() => {
            this._cache[t] = null;
          }, 0)),
        i
      );
    }
    prepareRequest(t, e) {
      switch (t) {
        case 'getBlockNumber':
          return ['eth_blockNumber', []];
        case 'getGasPrice':
          return ['eth_gasPrice', []];
        case 'getBalance':
          return ['eth_getBalance', [Rh(e.address), e.blockTag]];
        case 'getTransactionCount':
          return ['eth_getTransactionCount', [Rh(e.address), e.blockTag]];
        case 'getCode':
          return ['eth_getCode', [Rh(e.address), e.blockTag]];
        case 'getStorageAt':
          return [
            'eth_getStorageAt',
            [Rh(e.address), C(e.position, 32), e.blockTag],
          ];
        case 'sendTransaction':
          return ['eth_sendRawTransaction', [e.signedTransaction]];
        case 'getBlock':
          return e.blockTag
            ? ['eth_getBlockByNumber', [e.blockTag, !!e.includeTransactions]]
            : e.blockHash
            ? ['eth_getBlockByHash', [e.blockHash, !!e.includeTransactions]]
            : null;
        case 'getTransaction':
          return ['eth_getTransactionByHash', [e.transactionHash]];
        case 'getTransactionReceipt':
          return ['eth_getTransactionReceipt', [e.transactionHash]];
        case 'call':
          return [
            'eth_call',
            [
              at(this.constructor, 'hexlifyTransaction')(e.transaction, {
                from: !0,
              }),
              e.blockTag,
            ],
          ];
        case 'estimateGas':
          return [
            'eth_estimateGas',
            [
              at(this.constructor, 'hexlifyTransaction')(e.transaction, {
                from: !0,
              }),
            ],
          ];
        case 'getLogs':
          return (
            e.filter &&
              null != e.filter.address &&
              (e.filter.address = Rh(e.filter.address)),
            ['eth_getLogs', [e.filter]]
          );
      }
      return null;
    }
    perform(t, e) {
      return kh(this, void 0, void 0, function* () {
        if ('call' === t || 'estimateGas' === t) {
          const t = e.transaction;
          if (
            t &&
            null != t.type &&
            j.from(t.type).isZero() &&
            null == t.maxFeePerGas &&
            null == t.maxPriorityFeePerGas
          ) {
            const r = yield this.getFeeData();
            null == r.maxFeePerGas &&
              null == r.maxPriorityFeePerGas &&
              (((e = lt(e)).transaction = lt(t)), delete e.transaction.type);
          }
        }
        const r = this.prepareRequest(t, e);
        null == r &&
          Sh.throwError(t + ' not implemented', f.errors.NOT_IMPLEMENTED, {
            operation: t,
          });
        try {
          return yield this.send(r[0], r[1]);
        } catch (r) {
          return xh(t, r, e);
        }
      });
    }
    _startEvent(t) {
      'pending' === t.tag && this._startPending(), super._startEvent(t);
    }
    _startPending() {
      if (null != this._pendingFilter) return;
      const t = this,
        e = this.send('eth_newPendingTransactionFilter', []);
      (this._pendingFilter = e),
        e
          .then(function (r) {
            return (
              (function n() {
                t.send('eth_getFilterChanges', [r])
                  .then(function (r) {
                    if (t._pendingFilter != e) return null;
                    let n = Promise.resolve();
                    return (
                      r.forEach(function (e) {
                        (t._emitted['t:' + e.toLowerCase()] = 'pending'),
                          (n = n.then(function () {
                            return t.getTransaction(e).then(function (e) {
                              return t.emit('pending', e), null;
                            });
                          }));
                      }),
                      n.then(function () {
                        return Th(1e3);
                      })
                    );
                  })
                  .then(function () {
                    if (t._pendingFilter == e)
                      return (
                        setTimeout(function () {
                          n();
                        }, 0),
                        null
                      );
                    t.send('eth_uninstallFilter', [r]);
                  })
                  .catch((t) => {});
              })(),
              r
            );
          })
          .catch((t) => {});
    }
    _stopEvent(t) {
      'pending' === t.tag &&
        0 === this.listenerCount('pending') &&
        (this._pendingFilter = null),
        super._stopEvent(t);
    }
    static hexlifyTransaction(t, e) {
      const r = lt(Fh);
      if (e) for (const t in e) e[t] && (r[t] = !0);
      ht(t, r);
      const n = {};
      return (
        [
          'chainId',
          'gasLimit',
          'gasPrice',
          'type',
          'maxFeePerGas',
          'maxPriorityFeePerGas',
          'nonce',
          'value',
        ].forEach(function (e) {
          if (null == t[e]) return;
          const r = x(j.from(t[e]));
          'gasLimit' === e && (e = 'gas'), (n[e] = r);
        }),
        ['from', 'to', 'data'].forEach(function (e) {
          null != t[e] && (n[e] = k(t[e]));
        }),
        t.accessList && (n.accessList = Ks(t.accessList)),
        n
      );
    }
  }
  let Lh = null;
  try {
    if (((Lh = WebSocket), null == Lh)) throw new Error('inject please');
  } catch (t) {
    const e = new f(Wu);
    Lh = function () {
      e.throwError(
        'WebSockets not supported in this environment',
        f.errors.UNSUPPORTED_OPERATION,
        { operation: 'new WebSocket()' }
      );
    };
  }
  var Uh = function (t, e, r, n) {
    return new (r || (r = Promise))(function (i, o) {
      function s(t) {
        try {
          u(n.next(t));
        } catch (t) {
          o(t);
        }
      }
      function a(t) {
        try {
          u(n.throw(t));
        } catch (t) {
          o(t);
        }
      }
      function u(t) {
        var e;
        t.done
          ? i(t.value)
          : ((e = t.value),
            e instanceof r
              ? e
              : new r(function (t) {
                  t(e);
                })).then(s, a);
      }
      u((n = n.apply(t, e || [])).next());
    });
  };
  const jh = new f(Wu);
  let zh = 1;
  class Gh extends Dh {
    constructor(t, e) {
      'any' === e &&
        jh.throwError(
          "WebSocketProvider does not support 'any' network yet",
          f.errors.UNSUPPORTED_OPERATION,
          { operation: 'network:any' }
        ),
        super('string' == typeof t ? t : '_websocket', e),
        (this._pollingInterval = -1),
        (this._wsReady = !1),
        st(
          this,
          '_websocket',
          'string' == typeof t ? new Lh(this.connection.url) : t
        ),
        st(this, '_requests', {}),
        st(this, '_subs', {}),
        st(this, '_subIds', {}),
        st(this, '_detectNetwork', super.detectNetwork()),
        (this.websocket.onopen = () => {
          (this._wsReady = !0),
            Object.keys(this._requests).forEach((t) => {
              this.websocket.send(this._requests[t].payload);
            });
        }),
        (this.websocket.onmessage = (t) => {
          const e = t.data,
            r = JSON.parse(e);
          if (null != r.id) {
            const t = String(r.id),
              n = this._requests[t];
            if ((delete this._requests[t], void 0 !== r.result))
              n.callback(null, r.result),
                this.emit('debug', {
                  action: 'response',
                  request: JSON.parse(n.payload),
                  response: r.result,
                  provider: this,
                });
            else {
              let t = null;
              r.error
                ? ((t = new Error(r.error.message || 'unknown error')),
                  st(t, 'code', r.error.code || null),
                  st(t, 'response', e))
                : (t = new Error('unknown error')),
                n.callback(t, void 0),
                this.emit('debug', {
                  action: 'response',
                  error: t,
                  request: JSON.parse(n.payload),
                  provider: this,
                });
            }
          } else if ('eth_subscription' === r.method) {
            const t = this._subs[r.params.subscription];
            t && t.processFunc(r.params.result);
          } else console.warn('this should not happen');
        });
      const r = setInterval(() => {
        this.emit('poll');
      }, 1e3);
      r.unref && r.unref();
    }
    get websocket() {
      return this._websocket;
    }
    detectNetwork() {
      return this._detectNetwork;
    }
    get pollingInterval() {
      return 0;
    }
    resetEventsBlock(t) {
      jh.throwError(
        'cannot reset events block on WebSocketProvider',
        f.errors.UNSUPPORTED_OPERATION,
        { operation: 'resetEventBlock' }
      );
    }
    set pollingInterval(t) {
      jh.throwError(
        'cannot set polling interval on WebSocketProvider',
        f.errors.UNSUPPORTED_OPERATION,
        { operation: 'setPollingInterval' }
      );
    }
    poll() {
      return Uh(this, void 0, void 0, function* () {
        return null;
      });
    }
    set polling(t) {
      t &&
        jh.throwError(
          'cannot set polling on WebSocketProvider',
          f.errors.UNSUPPORTED_OPERATION,
          { operation: 'setPolling' }
        );
    }
    send(t, e) {
      const r = zh++;
      return new Promise((n, i) => {
        const o = JSON.stringify({
          method: t,
          params: e,
          id: r,
          jsonrpc: '2.0',
        });
        this.emit('debug', {
          action: 'request',
          request: JSON.parse(o),
          provider: this,
        }),
          (this._requests[String(r)] = {
            callback: function (t, e) {
              return t ? i(t) : n(e);
            },
            payload: o,
          }),
          this._wsReady && this.websocket.send(o);
      });
    }
    static defaultUrl() {
      return 'ws://localhost:8546';
    }
    _subscribe(t, e, r) {
      return Uh(this, void 0, void 0, function* () {
        let n = this._subIds[t];
        null == n &&
          ((n = Promise.all(e).then((t) => this.send('eth_subscribe', t))),
          (this._subIds[t] = n));
        const i = yield n;
        this._subs[i] = { tag: t, processFunc: r };
      });
    }
    _startEvent(t) {
      switch (t.type) {
        case 'block':
          this._subscribe('block', ['newHeads'], (t) => {
            const e = j.from(t.number).toNumber();
            (this._emitted.block = e), this.emit('block', e);
          });
          break;
        case 'pending':
          this._subscribe('pending', ['newPendingTransactions'], (t) => {
            this.emit('pending', t);
          });
          break;
        case 'filter':
          this._subscribe(t.tag, ['logs', this._getFilter(t.filter)], (e) => {
            null == e.removed && (e.removed = !1),
              this.emit(t.filter, this.formatter.filterLog(e));
          });
          break;
        case 'tx': {
          const e = (t) => {
            const e = t.hash;
            this.getTransactionReceipt(e).then((t) => {
              t && this.emit(e, t);
            });
          };
          e(t),
            this._subscribe('tx', ['newHeads'], (t) => {
              this._events.filter((t) => 'tx' === t.type).forEach(e);
            });
          break;
        }
        case 'debug':
        case 'poll':
        case 'willPoll':
        case 'didPoll':
        case 'error':
          break;
        default:
          console.log('unhandled:', t);
      }
    }
    _stopEvent(t) {
      let e = t.tag;
      if ('tx' === t.type) {
        if (this._events.filter((t) => 'tx' === t.type).length) return;
        e = 'tx';
      } else if (this.listenerCount(t.event)) return;
      const r = this._subIds[e];
      r &&
        (delete this._subIds[e],
        r.then((t) => {
          this._subs[t] &&
            (delete this._subs[t], this.send('eth_unsubscribe', [t]));
        }));
    }
    destroy() {
      return Uh(this, void 0, void 0, function* () {
        this.websocket.readyState === Lh.CONNECTING &&
          (yield new Promise((t) => {
            (this.websocket.onopen = function () {
              t(!0);
            }),
              (this.websocket.onerror = function () {
                t(!1);
              });
          })),
          this.websocket.close(1e3);
      });
    }
  }
  var qh = function (t, e, r, n) {
    return new (r || (r = Promise))(function (i, o) {
      function s(t) {
        try {
          u(n.next(t));
        } catch (t) {
          o(t);
        }
      }
      function a(t) {
        try {
          u(n.throw(t));
        } catch (t) {
          o(t);
        }
      }
      function u(t) {
        var e;
        t.done
          ? i(t.value)
          : ((e = t.value),
            e instanceof r
              ? e
              : new r(function (t) {
                  t(e);
                })).then(s, a);
      }
      u((n = n.apply(t, e || [])).next());
    });
  };
  const Kh = new f(Wu);
  class Hh extends Dh {
    detectNetwork() {
      const t = Object.create(null, {
        detectNetwork: { get: () => super.detectNetwork },
      });
      return qh(this, void 0, void 0, function* () {
        let e = this.network;
        return (
          null == e &&
            ((e = yield t.detectNetwork.call(this)),
            e ||
              Kh.throwError('no network detected', f.errors.UNKNOWN_ERROR, {}),
            null == this._network &&
              (st(this, '_network', e), this.emit('network', e, null))),
          e
        );
      });
    }
  }
  class Vh extends Hh {
    constructor(t, e) {
      Kh.checkAbstract(new.target, Vh),
        (t = at(new.target, 'getNetwork')(t)),
        (e = at(new.target, 'getApiKey')(e));
      super(at(new.target, 'getUrl')(t, e), t),
        'string' == typeof e
          ? st(this, 'apiKey', e)
          : null != e &&
            Object.keys(e).forEach((t) => {
              st(this, t, e[t]);
            });
    }
    _startPending() {
      Kh.warn('WARNING: API provider does not support pending filters');
    }
    isCommunityResource() {
      return !1;
    }
    getSigner(t) {
      return Kh.throwError(
        'API provider does not support signing',
        f.errors.UNSUPPORTED_OPERATION,
        { operation: 'getSigner' }
      );
    }
    listAccounts() {
      return Promise.resolve([]);
    }
    static getApiKey(t) {
      return t;
    }
    static getUrl(t, e) {
      return Kh.throwError(
        'not implemented; sub-classes must override getUrl',
        f.errors.NOT_IMPLEMENTED,
        { operation: 'getUrl' }
      );
    }
  }
  const Jh = new f(Wu),
    Wh = '_gg7wSSi0KMBsdKnGVfHDueq6xMB9EkC';
  class Qh extends Gh {
    constructor(t, e) {
      const r = new Yh(t, e);
      super(
        r.connection.url
          .replace(/^http/i, 'ws')
          .replace('.alchemyapi.', '.ws.alchemyapi.'),
        r.network
      ),
        st(this, 'apiKey', r.apiKey);
    }
    isCommunityResource() {
      return this.apiKey === Wh;
    }
  }
  class Yh extends Vh {
    static getWebSocketProvider(t, e) {
      return new Qh(t, e);
    }
    static getApiKey(t) {
      return null == t
        ? Wh
        : (t &&
            'string' != typeof t &&
            Jh.throwArgumentError('invalid apiKey', 'apiKey', t),
          t);
    }
    static getUrl(t, e) {
      let r = null;
      switch (t.name) {
        case 'homestead':
          r = 'eth-mainnet.alchemyapi.io/v2/';
          break;
        case 'goerli':
          r = 'eth-goerli.g.alchemy.com/v2/';
          break;
        case 'matic':
          r = 'polygon-mainnet.g.alchemy.com/v2/';
          break;
        case 'maticmum':
          r = 'polygon-mumbai.g.alchemy.com/v2/';
          break;
        case 'arbitrum':
          r = 'arb-mainnet.g.alchemy.com/v2/';
          break;
        case 'arbitrum-goerli':
          r = 'arb-goerli.g.alchemy.com/v2/';
          break;
        case 'optimism':
          r = 'opt-mainnet.g.alchemy.com/v2/';
          break;
        case 'optimism-goerli':
          r = 'opt-goerli.g.alchemy.com/v2/';
          break;
        default:
          Jh.throwArgumentError('unsupported network', 'network', arguments[0]);
      }
      return {
        allowGzip: !0,
        url: 'https://' + r + e,
        throttleCallback: (t, r) => (e === Wh && th(), Promise.resolve(!0)),
      };
    }
    isCommunityResource() {
      return this.apiKey === Wh;
    }
  }
  const Zh = new f(Wu),
    Xh = '9f7d929b018cdffb338517efa06f58359e86ff1ffd350bc889738523659e7972';
  function $h(t) {
    switch (t) {
      case 'homestead':
        return 'rpc.ankr.com/eth/';
      case 'ropsten':
        return 'rpc.ankr.com/eth_ropsten/';
      case 'rinkeby':
        return 'rpc.ankr.com/eth_rinkeby/';
      case 'goerli':
        return 'rpc.ankr.com/eth_goerli/';
      case 'matic':
        return 'rpc.ankr.com/polygon/';
      case 'arbitrum':
        return 'rpc.ankr.com/arbitrum/';
    }
    return Zh.throwArgumentError('unsupported network', 'name', t);
  }
  class tl extends Vh {
    isCommunityResource() {
      return this.apiKey === Xh;
    }
    static getApiKey(t) {
      return null == t ? Xh : t;
    }
    static getUrl(t, e) {
      null == e && (e = Xh);
      const r = {
        allowGzip: !0,
        url: 'https://' + $h(t.name) + e,
        throttleCallback: (t, r) => (
          e.apiKey === Xh && th(), Promise.resolve(!0)
        ),
      };
      return (
        null != e.projectSecret &&
          ((r.user = ''), (r.password = e.projectSecret)),
        r
      );
    }
  }
  var el = function (t, e, r, n) {
    return new (r || (r = Promise))(function (i, o) {
      function s(t) {
        try {
          u(n.next(t));
        } catch (t) {
          o(t);
        }
      }
      function a(t) {
        try {
          u(n.throw(t));
        } catch (t) {
          o(t);
        }
      }
      function u(t) {
        var e;
        t.done
          ? i(t.value)
          : ((e = t.value),
            e instanceof r
              ? e
              : new r(function (t) {
                  t(e);
                })).then(s, a);
      }
      u((n = n.apply(t, e || [])).next());
    });
  };
  const rl = new f(Wu);
  class nl extends Vh {
    static getApiKey(t) {
      return (
        null != t &&
          rl.throwArgumentError(
            'apiKey not supported for cloudflare',
            'apiKey',
            t
          ),
        null
      );
    }
    static getUrl(t, e) {
      let r = null;
      if ('homestead' === t.name) r = 'https://cloudflare-eth.com/';
      else
        rl.throwArgumentError('unsupported network', 'network', arguments[0]);
      return r;
    }
    perform(t, e) {
      const r = Object.create(null, { perform: { get: () => super.perform } });
      return el(this, void 0, void 0, function* () {
        if ('getBlockNumber' === t) {
          return (yield r.perform.call(this, 'getBlock', {
            blockTag: 'latest',
          })).number;
        }
        return r.perform.call(this, t, e);
      });
    }
  }
  var il = function (t, e, r, n) {
    return new (r || (r = Promise))(function (i, o) {
      function s(t) {
        try {
          u(n.next(t));
        } catch (t) {
          o(t);
        }
      }
      function a(t) {
        try {
          u(n.throw(t));
        } catch (t) {
          o(t);
        }
      }
      function u(t) {
        var e;
        t.done
          ? i(t.value)
          : ((e = t.value),
            e instanceof r
              ? e
              : new r(function (t) {
                  t(e);
                })).then(s, a);
      }
      u((n = n.apply(t, e || [])).next());
    });
  };
  const ol = new f(Wu);
  function sl(t) {
    const e = {};
    for (let r in t) {
      if (null == t[r]) continue;
      let n = t[r];
      ('type' === r && 0 === n) ||
        ((n = {
          type: !0,
          gasLimit: !0,
          gasPrice: !0,
          maxFeePerGs: !0,
          maxPriorityFeePerGas: !0,
          nonce: !0,
          value: !0,
        }[r]
          ? x(k(n))
          : 'accessList' === r
          ? '[' +
            Ks(n)
              .map(
                (t) =>
                  `{address:"${t.address}",storageKeys:["${t.storageKeys.join(
                    '","'
                  )}"]}`
              )
              .join(',') +
            ']'
          : k(n)),
        (e[r] = n));
    }
    return e;
  }
  function al(t) {
    if (
      0 == t.status &&
      ('No records found' === t.message ||
        'No transactions found' === t.message)
    )
      return t.result;
    if (
      1 != t.status ||
      'string' != typeof t.message ||
      !t.message.match(/^OK/)
    ) {
      const e = new Error('invalid response');
      throw (
        ((e.result = JSON.stringify(t)),
        (t.result || '').toLowerCase().indexOf('rate limit') >= 0 &&
          (e.throttleRetry = !0),
        e)
      );
    }
    return t.result;
  }
  function ul(t) {
    if (
      t &&
      0 == t.status &&
      'NOTOK' == t.message &&
      (t.result || '').toLowerCase().indexOf('rate limit') >= 0
    ) {
      const e = new Error('throttled response');
      throw ((e.result = JSON.stringify(t)), (e.throttleRetry = !0), e);
    }
    if ('2.0' != t.jsonrpc) {
      const e = new Error('invalid response');
      throw ((e.result = JSON.stringify(t)), e);
    }
    if (t.error) {
      const e = new Error(t.error.message || 'unknown error');
      throw (
        (t.error.code && (e.code = t.error.code),
        t.error.data && (e.data = t.error.data),
        e)
      );
    }
    return t.result;
  }
  function hl(t) {
    if ('pending' === t) throw new Error('pending not supported');
    return 'latest' === t ? t : parseInt(t.substring(2), 16);
  }
  function ll(t, e, r) {
    if ('call' === t && e.code === f.errors.SERVER_ERROR) {
      const t = e.error;
      if (
        t &&
        (t.message.match(/reverted/i) || t.message.match(/VM execution error/i))
      ) {
        let r = t.data;
        if ((r && (r = '0x' + r.replace(/^.*0x/i, '')), M(r))) return r;
        ol.throwError(
          'missing revert data in call exception',
          f.errors.CALL_EXCEPTION,
          { error: e, data: '0x' }
        );
      }
    }
    let n = e.message;
    throw (
      (e.code === f.errors.SERVER_ERROR &&
        (e.error && 'string' == typeof e.error.message
          ? (n = e.error.message)
          : 'string' == typeof e.body
          ? (n = e.body)
          : 'string' == typeof e.responseText && (n = e.responseText)),
      (n = (n || '').toLowerCase()),
      n.match(/insufficient funds/) &&
        ol.throwError(
          'insufficient funds for intrinsic transaction cost',
          f.errors.INSUFFICIENT_FUNDS,
          { error: e, method: t, transaction: r }
        ),
      n.match(
        /same hash was already imported|transaction nonce is too low|nonce too low/
      ) &&
        ol.throwError('nonce has already been used', f.errors.NONCE_EXPIRED, {
          error: e,
          method: t,
          transaction: r,
        }),
      n.match(/another transaction with same nonce/) &&
        ol.throwError(
          'replacement fee too low',
          f.errors.REPLACEMENT_UNDERPRICED,
          { error: e, method: t, transaction: r }
        ),
      n.match(/execution failed due to an exception|execution reverted/) &&
        ol.throwError(
          'cannot estimate gas; transaction may fail or may require manual gas limit',
          f.errors.UNPREDICTABLE_GAS_LIMIT,
          { error: e, method: t, transaction: r }
        ),
      e)
    );
  }
  class cl extends _h {
    constructor(t, e) {
      super(t),
        st(this, 'baseUrl', this.getBaseUrl()),
        st(this, 'apiKey', e || null);
    }
    getBaseUrl() {
      switch (this.network ? this.network.name : 'invalid') {
        case 'homestead':
          return 'https://api.etherscan.io';
        case 'goerli':
          return 'https://api-goerli.etherscan.io';
        case 'sepolia':
          return 'https://api-sepolia.etherscan.io';
        case 'matic':
          return 'https://api.polygonscan.com';
        case 'maticmum':
          return 'https://api-testnet.polygonscan.com';
        case 'arbitrum':
          return 'https://api.arbiscan.io';
        case 'arbitrum-goerli':
          return 'https://api-goerli.arbiscan.io';
        case 'optimism':
          return 'https://api-optimistic.etherscan.io';
        case 'optimism-goerli':
          return 'https://api-goerli-optimistic.etherscan.io';
      }
      return ol.throwArgumentError(
        'unsupported network',
        'network',
        this.network.name
      );
    }
    getUrl(t, e) {
      const r = Object.keys(e).reduce((t, r) => {
          const n = e[r];
          return null != n && (t += `&${r}=${n}`), t;
        }, ''),
        n = this.apiKey ? `&apikey=${this.apiKey}` : '';
      return `${this.baseUrl}/api?module=${t}${r}${n}`;
    }
    getPostUrl() {
      return `${this.baseUrl}/api`;
    }
    getPostData(t, e) {
      return (e.module = t), (e.apikey = this.apiKey), e;
    }
    fetch(t, e, r) {
      return il(this, void 0, void 0, function* () {
        const n = r ? this.getPostUrl() : this.getUrl(t, e),
          i = r ? this.getPostData(t, e) : null,
          o = 'proxy' === t ? ul : al;
        this.emit('debug', { action: 'request', request: n, provider: this });
        const s = {
          url: n,
          throttleSlotInterval: 1e3,
          throttleCallback: (t, e) => (
            this.isCommunityResource() && th(), Promise.resolve(!0)
          ),
        };
        let a = null;
        i &&
          ((s.headers = {
            'content-type': 'application/x-www-form-urlencoded; charset=UTF-8',
          }),
          (a = Object.keys(i)
            .map((t) => `${t}=${i[t]}`)
            .join('&')));
        const u = yield Fu(s, a, o || ul);
        return (
          this.emit('debug', {
            action: 'response',
            request: n,
            response: mt(u),
            provider: this,
          }),
          u
        );
      });
    }
    detectNetwork() {
      return il(this, void 0, void 0, function* () {
        return this.network;
      });
    }
    perform(t, e) {
      const r = Object.create(null, { perform: { get: () => super.perform } });
      return il(this, void 0, void 0, function* () {
        switch (t) {
          case 'getBlockNumber':
            return this.fetch('proxy', { action: 'eth_blockNumber' });
          case 'getGasPrice':
            return this.fetch('proxy', { action: 'eth_gasPrice' });
          case 'getBalance':
            return this.fetch('account', {
              action: 'balance',
              address: e.address,
              tag: e.blockTag,
            });
          case 'getTransactionCount':
            return this.fetch('proxy', {
              action: 'eth_getTransactionCount',
              address: e.address,
              tag: e.blockTag,
            });
          case 'getCode':
            return this.fetch('proxy', {
              action: 'eth_getCode',
              address: e.address,
              tag: e.blockTag,
            });
          case 'getStorageAt':
            return this.fetch('proxy', {
              action: 'eth_getStorageAt',
              address: e.address,
              position: e.position,
              tag: e.blockTag,
            });
          case 'sendTransaction':
            return this.fetch(
              'proxy',
              { action: 'eth_sendRawTransaction', hex: e.signedTransaction },
              !0
            ).catch((t) => ll('sendTransaction', t, e.signedTransaction));
          case 'getBlock':
            if (e.blockTag)
              return this.fetch('proxy', {
                action: 'eth_getBlockByNumber',
                tag: e.blockTag,
                boolean: e.includeTransactions ? 'true' : 'false',
              });
            throw new Error('getBlock by blockHash not implemented');
          case 'getTransaction':
            return this.fetch('proxy', {
              action: 'eth_getTransactionByHash',
              txhash: e.transactionHash,
            });
          case 'getTransactionReceipt':
            return this.fetch('proxy', {
              action: 'eth_getTransactionReceipt',
              txhash: e.transactionHash,
            });
          case 'call': {
            if ('latest' !== e.blockTag)
              throw new Error(
                'EtherscanProvider does not support blockTag for call'
              );
            const t = sl(e.transaction);
            (t.module = 'proxy'), (t.action = 'eth_call');
            try {
              return yield this.fetch('proxy', t, !0);
            } catch (t) {
              return ll('call', t, e.transaction);
            }
          }
          case 'estimateGas': {
            const t = sl(e.transaction);
            (t.module = 'proxy'), (t.action = 'eth_estimateGas');
            try {
              return yield this.fetch('proxy', t, !0);
            } catch (t) {
              return ll('estimateGas', t, e.transaction);
            }
          }
          case 'getLogs': {
            const t = { action: 'getLogs' };
            if (
              (e.filter.fromBlock && (t.fromBlock = hl(e.filter.fromBlock)),
              e.filter.toBlock && (t.toBlock = hl(e.filter.toBlock)),
              e.filter.address && (t.address = e.filter.address),
              e.filter.topics &&
                e.filter.topics.length > 0 &&
                (e.filter.topics.length > 1 &&
                  ol.throwError(
                    'unsupported topic count',
                    f.errors.UNSUPPORTED_OPERATION,
                    { topics: e.filter.topics }
                  ),
                1 === e.filter.topics.length))
            ) {
              const r = e.filter.topics[0];
              ('string' == typeof r && 66 === r.length) ||
                ol.throwError(
                  'unsupported topic format',
                  f.errors.UNSUPPORTED_OPERATION,
                  { topic0: r }
                ),
                (t.topic0 = r);
            }
            const r = yield this.fetch('logs', t);
            let n = {};
            for (let t = 0; t < r.length; t++) {
              const e = r[t];
              if (null == e.blockHash) {
                if (null == n[e.blockNumber]) {
                  const t = yield this.getBlock(e.blockNumber);
                  t && (n[e.blockNumber] = t.hash);
                }
                e.blockHash = n[e.blockNumber];
              }
            }
            return r;
          }
          case 'getEtherPrice':
            return 'homestead' !== this.network.name
              ? 0
              : parseFloat(
                  (yield this.fetch('stats', { action: 'ethprice' })).ethusd
                );
        }
        return r.perform.call(this, t, e);
      });
    }
    getHistory(t, e, r) {
      return il(this, void 0, void 0, function* () {
        const n = {
          action: 'txlist',
          address: yield this.resolveName(t),
          startblock: null == e ? 0 : e,
          endblock: null == r ? 99999999 : r,
          sort: 'asc',
        };
        return (yield this.fetch('account', n)).map((t) => {
          ['contractAddress', 'to'].forEach(function (e) {
            '' == t[e] && delete t[e];
          }),
            null == t.creates &&
              null != t.contractAddress &&
              (t.creates = t.contractAddress);
          const e = this.formatter.transactionResponse(t);
          return t.timeStamp && (e.timestamp = parseInt(t.timeStamp)), e;
        });
      });
    }
    isCommunityResource() {
      return null == this.apiKey;
    }
  }
  var fl = function (t, e, r, n) {
    return new (r || (r = Promise))(function (i, o) {
      function s(t) {
        try {
          u(n.next(t));
        } catch (t) {
          o(t);
        }
      }
      function a(t) {
        try {
          u(n.throw(t));
        } catch (t) {
          o(t);
        }
      }
      function u(t) {
        var e;
        t.done
          ? i(t.value)
          : ((e = t.value),
            e instanceof r
              ? e
              : new r(function (t) {
                  t(e);
                })).then(s, a);
      }
      u((n = n.apply(t, e || [])).next());
    });
  };
  const dl = new f(Wu);
  function ml() {
    return new Date().getTime();
  }
  function pl(t) {
    let e = null;
    for (let r = 0; r < t.length; r++) {
      const n = t[r];
      if (null == n) return null;
      e
        ? (e.name === n.name &&
            e.chainId === n.chainId &&
            (e.ensAddress === n.ensAddress ||
              (null == e.ensAddress && null == n.ensAddress))) ||
          dl.throwArgumentError('provider mismatch', 'networks', t)
        : (e = n);
    }
    return e;
  }
  function gl(t, e) {
    t = t.slice().sort();
    const r = Math.floor(t.length / 2);
    if (t.length % 2) return t[r];
    const n = t[r - 1],
      i = t[r];
    return null != e && Math.abs(n - i) > e ? null : (n + i) / 2;
  }
  function yl(t) {
    if (null === t) return 'null';
    if ('number' == typeof t || 'boolean' == typeof t) return JSON.stringify(t);
    if ('string' == typeof t) return t;
    if (j.isBigNumber(t)) return t.toString();
    if (Array.isArray(t)) return JSON.stringify(t.map((t) => yl(t)));
    if ('object' == typeof t) {
      const e = Object.keys(t);
      return (
        e.sort(),
        '{' +
          e
            .map((e) => {
              let r = t[e];
              return (
                (r = 'function' == typeof r ? '[function]' : yl(r)),
                JSON.stringify(e) + ':' + r
              );
            })
            .join(',') +
          '}'
      );
    }
    throw new Error('unknown value type: ' + typeof t);
  }
  let vl = 1;
  function wl(t) {
    let e = null,
      r = null,
      n = new Promise((n) => {
        (e = function () {
          r && (clearTimeout(r), (r = null)), n();
        }),
          (r = setTimeout(e, t));
      });
    return {
      cancel: e,
      getPromise: function () {
        return n;
      },
      wait: (t) => ((n = n.then(t)), n),
    };
  }
  const bl = [
      f.errors.CALL_EXCEPTION,
      f.errors.INSUFFICIENT_FUNDS,
      f.errors.NONCE_EXPIRED,
      f.errors.REPLACEMENT_UNDERPRICED,
      f.errors.UNPREDICTABLE_GAS_LIMIT,
    ],
    Al = [
      'address',
      'args',
      'errorArgs',
      'errorSignature',
      'method',
      'transaction',
    ];
  function El(t, e) {
    const r = { weight: t.weight };
    return (
      Object.defineProperty(r, 'provider', { get: () => t.provider }),
      t.start && (r.start = t.start),
      e && (r.duration = e - t.start),
      t.done && (t.error ? (r.error = t.error) : (r.result = t.result || null)),
      r
    );
  }
  function Ml(t, e, r) {
    let n = yl;
    switch (e) {
      case 'getBlockNumber':
        return function (e) {
          const r = e.map((t) => t.result);
          let n = gl(
            e.map((t) => t.result),
            2
          );
          if (null != n)
            return (
              (n = Math.ceil(n)),
              r.indexOf(n + 1) >= 0 && n++,
              n >= t._highestBlockNumber && (t._highestBlockNumber = n),
              t._highestBlockNumber
            );
        };
      case 'getGasPrice':
        return function (t) {
          const e = t.map((t) => t.result);
          return e.sort(), e[Math.floor(e.length / 2)];
        };
      case 'getEtherPrice':
        return function (t) {
          return gl(t.map((t) => t.result));
        };
      case 'getBalance':
      case 'getTransactionCount':
      case 'getCode':
      case 'getStorageAt':
      case 'call':
      case 'estimateGas':
      case 'getLogs':
        break;
      case 'getTransaction':
      case 'getTransactionReceipt':
        n = function (t) {
          return null == t ? null : (((t = lt(t)).confirmations = -1), yl(t));
        };
        break;
      case 'getBlock':
        n = r.includeTransactions
          ? function (t) {
              return null == t
                ? null
                : (((t = lt(t)).transactions = t.transactions.map(
                    (t) => (((t = lt(t)).confirmations = -1), t)
                  )),
                  yl(t));
            }
          : function (t) {
              return null == t ? null : yl(t);
            };
        break;
      default:
        throw new Error('unknown method: ' + e);
    }
    return (function (t, e) {
      return function (r) {
        const n = {};
        r.forEach((e) => {
          const r = t(e.result);
          n[r] || (n[r] = { count: 0, result: e.result }), n[r].count++;
        });
        const i = Object.keys(n);
        for (let t = 0; t < i.length; t++) {
          const r = n[i[t]];
          if (r.count >= e) return r.result;
        }
      };
    })(n, t.quorum);
  }
  function _l(t, e) {
    return fl(this, void 0, void 0, function* () {
      const r = t.provider;
      return (null != r.blockNumber && r.blockNumber >= e) || -1 === e
        ? r
        : Du(
            () =>
              new Promise((n, i) => {
                setTimeout(function () {
                  return r.blockNumber >= e
                    ? n(r)
                    : t.cancelled
                    ? n(null)
                    : n(void 0);
                }, 0);
              }),
            { oncePoll: r }
          );
    });
  }
  function kl(t, e, r, n) {
    return fl(this, void 0, void 0, function* () {
      let i = t.provider;
      switch (r) {
        case 'getBlockNumber':
        case 'getGasPrice':
          return i[r]();
        case 'getEtherPrice':
          if (i.getEtherPrice) return i.getEtherPrice();
          break;
        case 'getBalance':
        case 'getTransactionCount':
        case 'getCode':
          return (
            n.blockTag && M(n.blockTag) && (i = yield _l(t, e)),
            i[r](n.address, n.blockTag || 'latest')
          );
        case 'getStorageAt':
          return (
            n.blockTag && M(n.blockTag) && (i = yield _l(t, e)),
            i.getStorageAt(n.address, n.position, n.blockTag || 'latest')
          );
        case 'getBlock':
          return (
            n.blockTag && M(n.blockTag) && (i = yield _l(t, e)),
            i[n.includeTransactions ? 'getBlockWithTransactions' : 'getBlock'](
              n.blockTag || n.blockHash
            )
          );
        case 'call':
        case 'estimateGas':
          return (
            n.blockTag && M(n.blockTag) && (i = yield _l(t, e)),
            'call' === r && n.blockTag
              ? i[r](n.transaction, n.blockTag)
              : i[r](n.transaction)
          );
        case 'getTransaction':
        case 'getTransactionReceipt':
          return i[r](n.transactionHash);
        case 'getLogs': {
          let r = n.filter;
          return (
            ((r.fromBlock && M(r.fromBlock)) || (r.toBlock && M(r.toBlock))) &&
              (i = yield _l(t, e)),
            i.getLogs(r)
          );
        }
      }
      return dl.throwError('unknown method error', f.errors.UNKNOWN_ERROR, {
        method: r,
        params: n,
      });
    });
  }
  class Sl extends _h {
    constructor(t, e) {
      0 === t.length &&
        dl.throwArgumentError('missing providers', 'providers', t);
      const r = t.map((t, e) => {
          if (On.isProvider(t)) {
            const e = Xu(t) ? 2e3 : 750,
              r = 1;
            return Object.freeze({
              provider: t,
              weight: 1,
              stallTimeout: e,
              priority: r,
            });
          }
          const r = lt(t);
          null == r.priority && (r.priority = 1),
            null == r.stallTimeout && (r.stallTimeout = Xu(t) ? 2e3 : 750),
            null == r.weight && (r.weight = 1);
          const n = r.weight;
          return (
            (n % 1 || n > 512 || n < 1) &&
              dl.throwArgumentError(
                'invalid weight; must be integer in [1, 512]',
                `providers[${e}].weight`,
                n
              ),
            Object.freeze(r)
          );
        }),
        n = r.reduce((t, e) => t + e.weight, 0);
      null == e
        ? (e = n / 2)
        : e > n &&
          dl.throwArgumentError(
            'quorum will always fail; larger than total weight',
            'quorum',
            e
          );
      let i = pl(r.map((t) => t.provider.network));
      null == i &&
        (i = new Promise((t, e) => {
          setTimeout(() => {
            this.detectNetwork().then(t, e);
          }, 0);
        })),
        super(i),
        st(this, 'providerConfigs', Object.freeze(r)),
        st(this, 'quorum', e),
        (this._highestBlockNumber = -1);
    }
    detectNetwork() {
      return fl(this, void 0, void 0, function* () {
        return pl(
          yield Promise.all(
            this.providerConfigs.map((t) => t.provider.getNetwork())
          )
        );
      });
    }
    perform(t, e) {
      return fl(this, void 0, void 0, function* () {
        if ('sendTransaction' === t) {
          const t = yield Promise.all(
            this.providerConfigs.map((t) =>
              t.provider.sendTransaction(e.signedTransaction).then(
                (t) => t.hash,
                (t) => t
              )
            )
          );
          for (let e = 0; e < t.length; e++) {
            const r = t[e];
            if ('string' == typeof r) return r;
          }
          throw t[0];
        }
        -1 === this._highestBlockNumber &&
          'getBlockNumber' !== t &&
          (yield this.getBlockNumber());
        const r = Ml(this, t, e),
          n = Va(this.providerConfigs.map(lt));
        n.sort((t, e) => t.priority - e.priority);
        const i = this._highestBlockNumber;
        let o = 0,
          s = !0;
        for (;;) {
          const a = ml();
          let u = n
            .filter((t) => t.runner && a - t.start < t.stallTimeout)
            .reduce((t, e) => t + e.weight, 0);
          for (; u < this.quorum && o < n.length; ) {
            const r = n[o++],
              s = vl++;
            (r.start = ml()),
              (r.staller = wl(r.stallTimeout)),
              r.staller.wait(() => {
                r.staller = null;
              }),
              (r.runner = kl(r, i, t, e).then(
                (n) => {
                  (r.done = !0),
                    (r.result = n),
                    this.listenerCount('debug') &&
                      this.emit('debug', {
                        action: 'request',
                        rid: s,
                        backend: El(r, ml()),
                        request: { method: t, params: mt(e) },
                        provider: this,
                      });
                },
                (n) => {
                  (r.done = !0),
                    (r.error = n),
                    this.listenerCount('debug') &&
                      this.emit('debug', {
                        action: 'request',
                        rid: s,
                        backend: El(r, ml()),
                        request: { method: t, params: mt(e) },
                        provider: this,
                      });
                }
              )),
              this.listenerCount('debug') &&
                this.emit('debug', {
                  action: 'request',
                  rid: s,
                  backend: El(r, null),
                  request: { method: t, params: mt(e) },
                  provider: this,
                }),
              (u += r.weight);
          }
          const h = [];
          n.forEach((t) => {
            !t.done &&
              t.runner &&
              (h.push(t.runner), t.staller && h.push(t.staller.getPromise()));
          }),
            h.length && (yield Promise.race(h));
          const l = n.filter((t) => t.done && null == t.error);
          if (l.length >= this.quorum) {
            const t = r(l);
            if (void 0 !== t)
              return (
                n.forEach((t) => {
                  t.staller && t.staller.cancel(), (t.cancelled = !0);
                }),
                t
              );
            s || (yield wl(100).getPromise()), (s = !1);
          }
          const c = n.reduce((t, e) => {
            if (!e.done || null == e.error) return t;
            const r = e.error.code;
            return (
              bl.indexOf(r) >= 0 &&
                (t[r] || (t[r] = { error: e.error, weight: 0 }),
                (t[r].weight += e.weight)),
              t
            );
          }, {});
          if (
            (Object.keys(c).forEach((t) => {
              const e = c[t];
              if (e.weight < this.quorum) return;
              n.forEach((t) => {
                t.staller && t.staller.cancel(), (t.cancelled = !0);
              });
              const r = e.error,
                i = {};
              Al.forEach((t) => {
                null != r[t] && (i[t] = r[t]);
              }),
                dl.throwError(r.reason || r.message, t, i);
            }),
            0 === n.filter((t) => !t.done).length)
          )
            break;
        }
        return (
          n.forEach((t) => {
            t.staller && t.staller.cancel(), (t.cancelled = !0);
          }),
          dl.throwError('failed to meet quorum', f.errors.SERVER_ERROR, {
            method: t,
            params: e,
            results: n.map((t) => El(t)),
            provider: this,
          })
        );
      });
    }
  }
  const Nl = new f(Wu),
    Pl = '84842078b09946638c03157f83405213';
  class xl extends Gh {
    constructor(t, e) {
      const r = new Tl(t, e),
        n = r.connection;
      n.password &&
        Nl.throwError(
          'INFURA WebSocket project secrets unsupported',
          f.errors.UNSUPPORTED_OPERATION,
          { operation: 'InfuraProvider.getWebSocketProvider()' }
        );
      super(n.url.replace(/^http/i, 'ws').replace('/v3/', '/ws/v3/'), t),
        st(this, 'apiKey', r.projectId),
        st(this, 'projectId', r.projectId),
        st(this, 'projectSecret', r.projectSecret);
    }
    isCommunityResource() {
      return this.projectId === Pl;
    }
  }
  class Tl extends Vh {
    static getWebSocketProvider(t, e) {
      return new xl(t, e);
    }
    static getApiKey(t) {
      const e = { apiKey: Pl, projectId: Pl, projectSecret: null };
      return (
        null == t ||
          ('string' == typeof t
            ? (e.projectId = t)
            : null != t.projectSecret
            ? (Nl.assertArgument(
                'string' == typeof t.projectId,
                'projectSecret requires a projectId',
                'projectId',
                t.projectId
              ),
              Nl.assertArgument(
                'string' == typeof t.projectSecret,
                'invalid projectSecret',
                'projectSecret',
                '[REDACTED]'
              ),
              (e.projectId = t.projectId),
              (e.projectSecret = t.projectSecret))
            : t.projectId && (e.projectId = t.projectId),
          (e.apiKey = e.projectId)),
        e
      );
    }
    static getUrl(t, e) {
      let r = null;
      switch (t ? t.name : 'unknown') {
        case 'homestead':
          r = 'mainnet.infura.io';
          break;
        case 'goerli':
          r = 'goerli.infura.io';
          break;
        case 'sepolia':
          r = 'sepolia.infura.io';
          break;
        case 'matic':
          r = 'polygon-mainnet.infura.io';
          break;
        case 'maticmum':
          r = 'polygon-mumbai.infura.io';
          break;
        case 'optimism':
          r = 'optimism-mainnet.infura.io';
          break;
        case 'optimism-goerli':
          r = 'optimism-goerli.infura.io';
          break;
        case 'arbitrum':
          r = 'arbitrum-mainnet.infura.io';
          break;
        case 'arbitrum-goerli':
          r = 'arbitrum-goerli.infura.io';
          break;
        default:
          Nl.throwError('unsupported network', f.errors.INVALID_ARGUMENT, {
            argument: 'network',
            value: t,
          });
      }
      const n = {
        allowGzip: !0,
        url: 'https://' + r + '/v3/' + e.projectId,
        throttleCallback: (t, r) => (
          e.projectId === Pl && th(), Promise.resolve(!0)
        ),
      };
      return (
        null != e.projectSecret &&
          ((n.user = ''), (n.password = e.projectSecret)),
        n
      );
    }
    isCommunityResource() {
      return this.projectId === Pl;
    }
  }
  const Cl = new f(Wu);
  class Rl extends Vh {
    static getApiKey(t) {
      return (
        t &&
          'string' != typeof t &&
          Cl.throwArgumentError('invalid apiKey', 'apiKey', t),
        t || 'ETHERS_JS_SHARED'
      );
    }
    static getUrl(t, e) {
      Cl.warn(
        'NodeSmith will be discontinued on 2019-12-20; please migrate to another platform.'
      );
      let r = null;
      switch (t.name) {
        case 'homestead':
          r = 'https://ethereum.api.nodesmith.io/v1/mainnet/jsonrpc';
          break;
        case 'ropsten':
          r = 'https://ethereum.api.nodesmith.io/v1/ropsten/jsonrpc';
          break;
        case 'rinkeby':
          r = 'https://ethereum.api.nodesmith.io/v1/rinkeby/jsonrpc';
          break;
        case 'goerli':
          r = 'https://ethereum.api.nodesmith.io/v1/goerli/jsonrpc';
          break;
        case 'kovan':
          r = 'https://ethereum.api.nodesmith.io/v1/kovan/jsonrpc';
          break;
        default:
          Cl.throwArgumentError('unsupported network', 'network', arguments[0]);
      }
      return r + '?apiKey=' + e;
    }
  }
  const Il = new f(Wu),
    Bl = '62e1ad51b37b8e00394bda3b';
  class Ol extends Vh {
    static getApiKey(t) {
      const e = {
        applicationId: null,
        loadBalancer: !0,
        applicationSecretKey: null,
      };
      return (
        null == t
          ? (e.applicationId = Bl)
          : 'string' == typeof t
          ? (e.applicationId = t)
          : null != t.applicationSecretKey
          ? ((e.applicationId = t.applicationId),
            (e.applicationSecretKey = t.applicationSecretKey))
          : t.applicationId
          ? (e.applicationId = t.applicationId)
          : Il.throwArgumentError(
              'unsupported PocketProvider apiKey',
              'apiKey',
              t
            ),
        e
      );
    }
    static getUrl(t, e) {
      let r = null;
      switch (t ? t.name : 'unknown') {
        case 'goerli':
          r = 'eth-goerli.gateway.pokt.network';
          break;
        case 'homestead':
          r = 'eth-mainnet.gateway.pokt.network';
          break;
        case 'kovan':
          r = 'poa-kovan.gateway.pokt.network';
          break;
        case 'matic':
          r = 'poly-mainnet.gateway.pokt.network';
          break;
        case 'maticmum':
          r = 'polygon-mumbai-rpc.gateway.pokt.network';
          break;
        case 'rinkeby':
          r = 'eth-rinkeby.gateway.pokt.network';
          break;
        case 'ropsten':
          r = 'eth-ropsten.gateway.pokt.network';
          break;
        default:
          Il.throwError('unsupported network', f.errors.INVALID_ARGUMENT, {
            argument: 'network',
            value: t,
          });
      }
      const n = { headers: {}, url: `https://${r}/v1/lb/${e.applicationId}` };
      return (
        null != e.applicationSecretKey &&
          ((n.user = ''), (n.password = e.applicationSecretKey)),
        n
      );
    }
    isCommunityResource() {
      return this.applicationId === Bl;
    }
  }
  const Fl = new f(Wu);
  let Dl = 1;
  function Ll(t, e) {
    const r = 'Web3LegacyFetcher';
    return function (t, n) {
      const i = { method: t, params: n, id: Dl++, jsonrpc: '2.0' };
      return new Promise((t, n) => {
        this.emit('debug', {
          action: 'request',
          fetcher: r,
          request: mt(i),
          provider: this,
        }),
          e(i, (e, o) => {
            if (e)
              return (
                this.emit('debug', {
                  action: 'response',
                  fetcher: r,
                  error: e,
                  request: i,
                  provider: this,
                }),
                n(e)
              );
            if (
              (this.emit('debug', {
                action: 'response',
                fetcher: r,
                request: i,
                response: o,
                provider: this,
              }),
              o.error)
            ) {
              const t = new Error(o.error.message);
              return (t.code = o.error.code), (t.data = o.error.data), n(t);
            }
            t(o.result);
          });
      });
    };
  }
  class Ul extends Dh {
    constructor(t, e) {
      null == t && Fl.throwArgumentError('missing provider', 'provider', t);
      let r = null,
        n = null,
        i = null;
      'function' == typeof t
        ? ((r = 'unknown:'), (n = t))
        : ((r = t.host || t.path || ''),
          !r && t.isMetaMask && (r = 'metamask'),
          (i = t),
          t.request
            ? ('' === r && (r = 'eip-1193:'),
              (n = (function (t) {
                return function (e, r) {
                  null == r && (r = []);
                  const n = { method: e, params: r };
                  return (
                    this.emit('debug', {
                      action: 'request',
                      fetcher: 'Eip1193Fetcher',
                      request: mt(n),
                      provider: this,
                    }),
                    t.request(n).then(
                      (t) => (
                        this.emit('debug', {
                          action: 'response',
                          fetcher: 'Eip1193Fetcher',
                          request: n,
                          response: t,
                          provider: this,
                        }),
                        t
                      ),
                      (t) => {
                        throw (
                          (this.emit('debug', {
                            action: 'response',
                            fetcher: 'Eip1193Fetcher',
                            request: n,
                            error: t,
                            provider: this,
                          }),
                          t)
                        );
                      }
                    )
                  );
                };
              })(t)))
            : t.sendAsync
            ? (n = Ll(0, t.sendAsync.bind(t)))
            : t.send
            ? (n = Ll(0, t.send.bind(t)))
            : Fl.throwArgumentError('unsupported provider', 'provider', t),
          r || (r = 'unknown:')),
        super(r, e),
        st(this, 'jsonRpcFetchFunc', n),
        st(this, 'provider', i);
    }
    send(t, e) {
      return this.jsonRpcFetchFunc(t, e);
    }
  }
  const jl = new f(Wu);
  function zl(t, e) {
    if ((null == t && (t = 'homestead'), 'string' == typeof t)) {
      const e = t.match(/^(ws|http)s?:/i);
      if (e)
        switch (e[1].toLowerCase()) {
          case 'http':
          case 'https':
            return new Dh(t);
          case 'ws':
          case 'wss':
            return new Gh(t);
          default:
            jl.throwArgumentError('unsupported URL scheme', 'network', t);
        }
    }
    const r = Pu(t);
    return (
      (r && r._defaultProvider) ||
        jl.throwError(
          'unsupported getDefaultProvider network',
          f.errors.NETWORK_ERROR,
          { operation: 'getDefaultProvider', network: t }
        ),
      r._defaultProvider(
        {
          FallbackProvider: Sl,
          AlchemyProvider: Yh,
          AnkrProvider: tl,
          CloudflareProvider: nl,
          EtherscanProvider: cl,
          InfuraProvider: Tl,
          JsonRpcProvider: Dh,
          NodesmithProvider: Rl,
          PocketProvider: Ol,
          Web3Provider: Ul,
          IpcProvider: null,
        },
        e
      )
    );
  }
  var Gl = Object.freeze({
    __proto__: null,
    Provider: On,
    BaseProvider: _h,
    Resolver: Ah,
    UrlJsonRpcProvider: Vh,
    FallbackProvider: Sl,
    AlchemyProvider: Yh,
    AlchemyWebSocketProvider: Qh,
    AnkrProvider: tl,
    CloudflareProvider: nl,
    EtherscanProvider: cl,
    InfuraProvider: Tl,
    InfuraWebSocketProvider: xl,
    JsonRpcProvider: Dh,
    JsonRpcBatchProvider: class extends Dh {
      send(t, e) {
        const r = { method: t, params: e, id: this._nextId++, jsonrpc: '2.0' };
        null == this._pendingBatch && (this._pendingBatch = []);
        const n = { request: r, resolve: null, reject: null },
          i = new Promise((t, e) => {
            (n.resolve = t), (n.reject = e);
          });
        return (
          this._pendingBatch.push(n),
          this._pendingBatchAggregator ||
            (this._pendingBatchAggregator = setTimeout(() => {
              const t = this._pendingBatch;
              (this._pendingBatch = null),
                (this._pendingBatchAggregator = null);
              const e = t.map((t) => t.request);
              return (
                this.emit('debug', {
                  action: 'requestBatch',
                  request: mt(e),
                  provider: this,
                }),
                Fu(this.connection, JSON.stringify(e)).then(
                  (r) => {
                    this.emit('debug', {
                      action: 'response',
                      request: e,
                      response: r,
                      provider: this,
                    }),
                      t.forEach((t, e) => {
                        const n = r[e];
                        if (n.error) {
                          const e = new Error(n.error.message);
                          (e.code = n.error.code),
                            (e.data = n.error.data),
                            t.reject(e);
                        } else t.resolve(n.result);
                      });
                  },
                  (r) => {
                    this.emit('debug', {
                      action: 'response',
                      error: r,
                      request: e,
                      provider: this,
                    }),
                      t.forEach((t) => {
                        t.reject(r);
                      });
                  }
                )
              );
            }, 10)),
          i
        );
      }
    },
    NodesmithProvider: Rl,
    PocketProvider: Ol,
    StaticJsonRpcProvider: Hh,
    Web3Provider: Ul,
    WebSocketProvider: Gh,
    IpcProvider: null,
    JsonRpcSigner: Bh,
    getDefaultProvider: zl,
    getNetwork: Pu,
    isCommunityResource: Xu,
    isCommunityResourcable: Zu,
    showThrottleMessage: th,
    Formatter: Yu,
  });
  const ql = 'abi/5.7.0',
    Kl = new f(ql),
    Hl = {};
  let Vl = { calldata: !0, memory: !0, storage: !0 },
    Jl = { calldata: !0, memory: !0 };
  function Wl(t, e) {
    if ('bytes' === t || 'string' === t) {
      if (Vl[e]) return !0;
    } else if ('address' === t) {
      if ('payable' === e) return !0;
    } else if ((t.indexOf('[') >= 0 || 'tuple' === t) && Jl[e]) return !0;
    return (
      (Vl[e] || 'payable' === e) &&
        Kl.throwArgumentError('invalid modifier', 'name', e),
      !1
    );
  }
  function Ql(t, e) {
    for (let r in e) st(t, r, e[r]);
  }
  const Yl = Object.freeze({
      sighash: 'sighash',
      minimal: 'minimal',
      full: 'full',
      json: 'json',
    }),
    Zl = new RegExp(/^(.*)\[([0-9]*)\]$/);
  class Xl {
    constructor(t, e) {
      t !== Hl &&
        Kl.throwError('use fromString', f.errors.UNSUPPORTED_OPERATION, {
          operation: 'new ParamType()',
        }),
        Ql(this, e);
      let r = this.type.match(Zl);
      Ql(
        this,
        r
          ? {
              arrayLength: parseInt(r[2] || '-1'),
              arrayChildren: Xl.fromObject({
                type: r[1],
                components: this.components,
              }),
              baseType: 'array',
            }
          : {
              arrayLength: null,
              arrayChildren: null,
              baseType: null != this.components ? 'tuple' : this.type,
            }
      ),
        (this._isParamType = !0),
        Object.freeze(this);
    }
    format(t) {
      if (
        (t || (t = Yl.sighash),
        Yl[t] || Kl.throwArgumentError('invalid format type', 'format', t),
        t === Yl.json)
      ) {
        let e = {
          type: 'tuple' === this.baseType ? 'tuple' : this.type,
          name: this.name || void 0,
        };
        return (
          'boolean' == typeof this.indexed && (e.indexed = this.indexed),
          this.components &&
            (e.components = this.components.map((e) =>
              JSON.parse(e.format(t))
            )),
          JSON.stringify(e)
        );
      }
      let e = '';
      return (
        'array' === this.baseType
          ? ((e += this.arrayChildren.format(t)),
            (e +=
              '[' +
              (this.arrayLength < 0 ? '' : String(this.arrayLength)) +
              ']'))
          : 'tuple' === this.baseType
          ? (t !== Yl.sighash && (e += this.type),
            (e +=
              '(' +
              this.components
                .map((e) => e.format(t))
                .join(t === Yl.full ? ', ' : ',') +
              ')'))
          : (e += this.type),
        t !== Yl.sighash &&
          (!0 === this.indexed && (e += ' indexed'),
          t === Yl.full && this.name && (e += ' ' + this.name)),
        e
      );
    }
    static from(t, e) {
      return 'string' == typeof t ? Xl.fromString(t, e) : Xl.fromObject(t);
    }
    static fromObject(t) {
      return Xl.isParamType(t)
        ? t
        : new Xl(Hl, {
            name: t.name || null,
            type: hc(t.type),
            indexed: null == t.indexed ? null : !!t.indexed,
            components: t.components ? t.components.map(Xl.fromObject) : null,
          });
    }
    static fromString(t, e) {
      return (
        (r = (function (t, e) {
          let r = t;
          function n(e) {
            Kl.throwArgumentError(
              `unexpected character at position ${e}`,
              'param',
              t
            );
          }
          function i(t) {
            let r = { type: '', name: '', parent: t, state: { allowType: !0 } };
            return e && (r.indexed = !1), r;
          }
          t = t.replace(/\s/g, ' ');
          let o = { type: '', name: '', state: { allowType: !0 } },
            s = o;
          for (let r = 0; r < t.length; r++) {
            let o = t[r];
            switch (o) {
              case '(':
                s.state.allowType && '' === s.type
                  ? (s.type = 'tuple')
                  : s.state.allowParams || n(r),
                  (s.state.allowType = !1),
                  (s.type = hc(s.type)),
                  (s.components = [i(s)]),
                  (s = s.components[0]);
                break;
              case ')':
                delete s.state,
                  'indexed' === s.name &&
                    (e || n(r), (s.indexed = !0), (s.name = '')),
                  Wl(s.type, s.name) && (s.name = ''),
                  (s.type = hc(s.type));
                let t = s;
                (s = s.parent),
                  s || n(r),
                  delete t.parent,
                  (s.state.allowParams = !1),
                  (s.state.allowName = !0),
                  (s.state.allowArray = !0);
                break;
              case ',':
                delete s.state,
                  'indexed' === s.name &&
                    (e || n(r), (s.indexed = !0), (s.name = '')),
                  Wl(s.type, s.name) && (s.name = ''),
                  (s.type = hc(s.type));
                let a = i(s.parent);
                s.parent.components.push(a), delete s.parent, (s = a);
                break;
              case ' ':
                s.state.allowType &&
                  '' !== s.type &&
                  ((s.type = hc(s.type)),
                  delete s.state.allowType,
                  (s.state.allowName = !0),
                  (s.state.allowParams = !0)),
                  s.state.allowName &&
                    '' !== s.name &&
                    ('indexed' === s.name
                      ? (e || n(r),
                        s.indexed && n(r),
                        (s.indexed = !0),
                        (s.name = ''))
                      : Wl(s.type, s.name)
                      ? (s.name = '')
                      : (s.state.allowName = !1));
                break;
              case '[':
                s.state.allowArray || n(r),
                  (s.type += o),
                  (s.state.allowArray = !1),
                  (s.state.allowName = !1),
                  (s.state.readArray = !0);
                break;
              case ']':
                s.state.readArray || n(r),
                  (s.type += o),
                  (s.state.readArray = !1),
                  (s.state.allowArray = !0),
                  (s.state.allowName = !0);
                break;
              default:
                s.state.allowType
                  ? ((s.type += o),
                    (s.state.allowParams = !0),
                    (s.state.allowArray = !0))
                  : s.state.allowName
                  ? ((s.name += o), delete s.state.allowArray)
                  : s.state.readArray
                  ? (s.type += o)
                  : n(r);
            }
          }
          return (
            s.parent && Kl.throwArgumentError('unexpected eof', 'param', t),
            delete o.state,
            'indexed' === s.name
              ? (e || n(r.length - 7),
                s.indexed && n(r.length - 7),
                (s.indexed = !0),
                (s.name = ''))
              : Wl(s.type, s.name) && (s.name = ''),
            (o.type = hc(o.type)),
            o
          );
        })(t, !!e)),
        Xl.fromObject({
          name: r.name,
          type: r.type,
          indexed: r.indexed,
          components: r.components,
        })
      );
      var r;
    }
    static isParamType(t) {
      return !(null == t || !t._isParamType);
    }
  }
  function $l(t, e) {
    return (function (t) {
      t = t.trim();
      let e = [],
        r = '',
        n = 0;
      for (let i = 0; i < t.length; i++) {
        let o = t[i];
        ',' === o && 0 === n
          ? (e.push(r), (r = ''))
          : ((r += o),
            '(' === o
              ? n++
              : ')' === o &&
                (n--,
                -1 === n &&
                  Kl.throwArgumentError('unbalanced parenthesis', 'value', t)));
      }
      r && e.push(r);
      return e;
    })(t).map((t) => Xl.fromString(t, e));
  }
  class tc {
    constructor(t, e) {
      t !== Hl &&
        Kl.throwError(
          'use a static from method',
          f.errors.UNSUPPORTED_OPERATION,
          { operation: 'new Fragment()' }
        ),
        Ql(this, e),
        (this._isFragment = !0),
        Object.freeze(this);
    }
    static from(t) {
      return tc.isFragment(t)
        ? t
        : 'string' == typeof t
        ? tc.fromString(t)
        : tc.fromObject(t);
    }
    static fromObject(t) {
      if (tc.isFragment(t)) return t;
      switch (t.type) {
        case 'function':
          return sc.fromObject(t);
        case 'event':
          return ec.fromObject(t);
        case 'constructor':
          return oc.fromObject(t);
        case 'error':
          return uc.fromObject(t);
        case 'fallback':
        case 'receive':
          return null;
      }
      return Kl.throwArgumentError('invalid fragment object', 'value', t);
    }
    static fromString(t) {
      return 'event' ===
        (t = (t = (t = t.replace(/\s/g, ' '))
          .replace(/\(/g, ' (')
          .replace(/\)/g, ') ')
          .replace(/\s+/g, ' ')).trim()).split(' ')[0]
        ? ec.fromString(t.substring(5).trim())
        : 'function' === t.split(' ')[0]
        ? sc.fromString(t.substring(8).trim())
        : 'constructor' === t.split('(')[0].trim()
        ? oc.fromString(t.trim())
        : 'error' === t.split(' ')[0]
        ? uc.fromString(t.substring(5).trim())
        : Kl.throwArgumentError('unsupported fragment', 'value', t);
    }
    static isFragment(t) {
      return !(!t || !t._isFragment);
    }
  }
  class ec extends tc {
    format(t) {
      if (
        (t || (t = Yl.sighash),
        Yl[t] || Kl.throwArgumentError('invalid format type', 'format', t),
        t === Yl.json)
      )
        return JSON.stringify({
          type: 'event',
          anonymous: this.anonymous,
          name: this.name,
          inputs: this.inputs.map((e) => JSON.parse(e.format(t))),
        });
      let e = '';
      return (
        t !== Yl.sighash && (e += 'event '),
        (e +=
          this.name +
          '(' +
          this.inputs.map((e) => e.format(t)).join(t === Yl.full ? ', ' : ',') +
          ') '),
        t !== Yl.sighash && this.anonymous && (e += 'anonymous '),
        e.trim()
      );
    }
    static from(t) {
      return 'string' == typeof t ? ec.fromString(t) : ec.fromObject(t);
    }
    static fromObject(t) {
      if (ec.isEventFragment(t)) return t;
      'event' !== t.type &&
        Kl.throwArgumentError('invalid event object', 'value', t);
      const e = {
        name: cc(t.name),
        anonymous: t.anonymous,
        inputs: t.inputs ? t.inputs.map(Xl.fromObject) : [],
        type: 'event',
      };
      return new ec(Hl, e);
    }
    static fromString(t) {
      let e = t.match(fc);
      e || Kl.throwArgumentError('invalid event string', 'value', t);
      let r = !1;
      return (
        e[3].split(' ').forEach((t) => {
          switch (t.trim()) {
            case 'anonymous':
              r = !0;
              break;
            case '':
              break;
            default:
              Kl.warn('unknown modifier: ' + t);
          }
        }),
        ec.fromObject({
          name: e[1].trim(),
          anonymous: r,
          inputs: $l(e[2], !0),
          type: 'event',
        })
      );
    }
    static isEventFragment(t) {
      return t && t._isFragment && 'event' === t.type;
    }
  }
  function rc(t, e) {
    e.gas = null;
    let r = t.split('@');
    return 1 !== r.length
      ? (r.length > 2 &&
          Kl.throwArgumentError(
            'invalid human-readable ABI signature',
            'value',
            t
          ),
        r[1].match(/^[0-9]+$/) ||
          Kl.throwArgumentError(
            'invalid human-readable ABI signature gas',
            'value',
            t
          ),
        (e.gas = j.from(r[1])),
        r[0])
      : t;
  }
  function nc(t, e) {
    (e.constant = !1),
      (e.payable = !1),
      (e.stateMutability = 'nonpayable'),
      t.split(' ').forEach((t) => {
        switch (t.trim()) {
          case 'constant':
            e.constant = !0;
            break;
          case 'payable':
            (e.payable = !0), (e.stateMutability = 'payable');
            break;
          case 'nonpayable':
            (e.payable = !1), (e.stateMutability = 'nonpayable');
            break;
          case 'pure':
            (e.constant = !0), (e.stateMutability = 'pure');
            break;
          case 'view':
            (e.constant = !0), (e.stateMutability = 'view');
            break;
          case 'external':
          case 'public':
          case '':
            break;
          default:
            console.log('unknown modifier: ' + t);
        }
      });
  }
  function ic(t) {
    let e = { constant: !1, payable: !0, stateMutability: 'payable' };
    return (
      null != t.stateMutability
        ? ((e.stateMutability = t.stateMutability),
          (e.constant =
            'view' === e.stateMutability || 'pure' === e.stateMutability),
          null != t.constant &&
            !!t.constant !== e.constant &&
            Kl.throwArgumentError(
              'cannot have constant function with mutability ' +
                e.stateMutability,
              'value',
              t
            ),
          (e.payable = 'payable' === e.stateMutability),
          null != t.payable &&
            !!t.payable !== e.payable &&
            Kl.throwArgumentError(
              'cannot have payable function with mutability ' +
                e.stateMutability,
              'value',
              t
            ))
        : null != t.payable
        ? ((e.payable = !!t.payable),
          null != t.constant ||
            e.payable ||
            'constructor' === t.type ||
            Kl.throwArgumentError(
              'unable to determine stateMutability',
              'value',
              t
            ),
          (e.constant = !!t.constant),
          e.constant
            ? (e.stateMutability = 'view')
            : (e.stateMutability = e.payable ? 'payable' : 'nonpayable'),
          e.payable &&
            e.constant &&
            Kl.throwArgumentError(
              'cannot have constant payable function',
              'value',
              t
            ))
        : null != t.constant
        ? ((e.constant = !!t.constant),
          (e.payable = !e.constant),
          (e.stateMutability = e.constant ? 'view' : 'payable'))
        : 'constructor' !== t.type &&
          Kl.throwArgumentError(
            'unable to determine stateMutability',
            'value',
            t
          ),
      e
    );
  }
  class oc extends tc {
    format(t) {
      if (
        (t || (t = Yl.sighash),
        Yl[t] || Kl.throwArgumentError('invalid format type', 'format', t),
        t === Yl.json)
      )
        return JSON.stringify({
          type: 'constructor',
          stateMutability:
            'nonpayable' !== this.stateMutability
              ? this.stateMutability
              : void 0,
          payable: this.payable,
          gas: this.gas ? this.gas.toNumber() : void 0,
          inputs: this.inputs.map((e) => JSON.parse(e.format(t))),
        });
      t === Yl.sighash &&
        Kl.throwError(
          'cannot format a constructor for sighash',
          f.errors.UNSUPPORTED_OPERATION,
          { operation: 'format(sighash)' }
        );
      let e =
        'constructor(' +
        this.inputs.map((e) => e.format(t)).join(t === Yl.full ? ', ' : ',') +
        ') ';
      return (
        this.stateMutability &&
          'nonpayable' !== this.stateMutability &&
          (e += this.stateMutability + ' '),
        e.trim()
      );
    }
    static from(t) {
      return 'string' == typeof t ? oc.fromString(t) : oc.fromObject(t);
    }
    static fromObject(t) {
      if (oc.isConstructorFragment(t)) return t;
      'constructor' !== t.type &&
        Kl.throwArgumentError('invalid constructor object', 'value', t);
      let e = ic(t);
      e.constant &&
        Kl.throwArgumentError('constructor cannot be constant', 'value', t);
      const r = {
        name: null,
        type: t.type,
        inputs: t.inputs ? t.inputs.map(Xl.fromObject) : [],
        payable: e.payable,
        stateMutability: e.stateMutability,
        gas: t.gas ? j.from(t.gas) : null,
      };
      return new oc(Hl, r);
    }
    static fromString(t) {
      let e = { type: 'constructor' },
        r = (t = rc(t, e)).match(fc);
      return (
        (r && 'constructor' === r[1].trim()) ||
          Kl.throwArgumentError('invalid constructor string', 'value', t),
        (e.inputs = $l(r[2].trim(), !1)),
        nc(r[3].trim(), e),
        oc.fromObject(e)
      );
    }
    static isConstructorFragment(t) {
      return t && t._isFragment && 'constructor' === t.type;
    }
  }
  class sc extends oc {
    format(t) {
      if (
        (t || (t = Yl.sighash),
        Yl[t] || Kl.throwArgumentError('invalid format type', 'format', t),
        t === Yl.json)
      )
        return JSON.stringify({
          type: 'function',
          name: this.name,
          constant: this.constant,
          stateMutability:
            'nonpayable' !== this.stateMutability
              ? this.stateMutability
              : void 0,
          payable: this.payable,
          gas: this.gas ? this.gas.toNumber() : void 0,
          inputs: this.inputs.map((e) => JSON.parse(e.format(t))),
          outputs: this.outputs.map((e) => JSON.parse(e.format(t))),
        });
      let e = '';
      return (
        t !== Yl.sighash && (e += 'function '),
        (e +=
          this.name +
          '(' +
          this.inputs.map((e) => e.format(t)).join(t === Yl.full ? ', ' : ',') +
          ') '),
        t !== Yl.sighash &&
          (this.stateMutability
            ? 'nonpayable' !== this.stateMutability &&
              (e += this.stateMutability + ' ')
            : this.constant && (e += 'view '),
          this.outputs &&
            this.outputs.length &&
            (e +=
              'returns (' +
              this.outputs.map((e) => e.format(t)).join(', ') +
              ') '),
          null != this.gas && (e += '@' + this.gas.toString() + ' ')),
        e.trim()
      );
    }
    static from(t) {
      return 'string' == typeof t ? sc.fromString(t) : sc.fromObject(t);
    }
    static fromObject(t) {
      if (sc.isFunctionFragment(t)) return t;
      'function' !== t.type &&
        Kl.throwArgumentError('invalid function object', 'value', t);
      let e = ic(t);
      const r = {
        type: t.type,
        name: cc(t.name),
        constant: e.constant,
        inputs: t.inputs ? t.inputs.map(Xl.fromObject) : [],
        outputs: t.outputs ? t.outputs.map(Xl.fromObject) : [],
        payable: e.payable,
        stateMutability: e.stateMutability,
        gas: t.gas ? j.from(t.gas) : null,
      };
      return new sc(Hl, r);
    }
    static fromString(t) {
      let e = { type: 'function' },
        r = (t = rc(t, e)).split(' returns ');
      r.length > 2 &&
        Kl.throwArgumentError('invalid function string', 'value', t);
      let n = r[0].match(fc);
      if (
        (n || Kl.throwArgumentError('invalid function signature', 'value', t),
        (e.name = n[1].trim()),
        e.name && cc(e.name),
        (e.inputs = $l(n[2], !1)),
        nc(n[3].trim(), e),
        r.length > 1)
      ) {
        let n = r[1].match(fc);
        ('' == n[1].trim() && '' == n[3].trim()) ||
          Kl.throwArgumentError('unexpected tokens', 'value', t),
          (e.outputs = $l(n[2], !1));
      } else e.outputs = [];
      return sc.fromObject(e);
    }
    static isFunctionFragment(t) {
      return t && t._isFragment && 'function' === t.type;
    }
  }
  function ac(t) {
    const e = t.format();
    return (
      ('Error(string)' !== e && 'Panic(uint256)' !== e) ||
        Kl.throwArgumentError(
          `cannot specify user defined ${e} error`,
          'fragment',
          t
        ),
      t
    );
  }
  class uc extends tc {
    format(t) {
      if (
        (t || (t = Yl.sighash),
        Yl[t] || Kl.throwArgumentError('invalid format type', 'format', t),
        t === Yl.json)
      )
        return JSON.stringify({
          type: 'error',
          name: this.name,
          inputs: this.inputs.map((e) => JSON.parse(e.format(t))),
        });
      let e = '';
      return (
        t !== Yl.sighash && (e += 'error '),
        (e +=
          this.name +
          '(' +
          this.inputs.map((e) => e.format(t)).join(t === Yl.full ? ', ' : ',') +
          ') '),
        e.trim()
      );
    }
    static from(t) {
      return 'string' == typeof t ? uc.fromString(t) : uc.fromObject(t);
    }
    static fromObject(t) {
      if (uc.isErrorFragment(t)) return t;
      'error' !== t.type &&
        Kl.throwArgumentError('invalid error object', 'value', t);
      const e = {
        type: t.type,
        name: cc(t.name),
        inputs: t.inputs ? t.inputs.map(Xl.fromObject) : [],
      };
      return ac(new uc(Hl, e));
    }
    static fromString(t) {
      let e = { type: 'error' },
        r = t.match(fc);
      return (
        r || Kl.throwArgumentError('invalid error signature', 'value', t),
        (e.name = r[1].trim()),
        e.name && cc(e.name),
        (e.inputs = $l(r[2], !1)),
        ac(uc.fromObject(e))
      );
    }
    static isErrorFragment(t) {
      return t && t._isFragment && 'error' === t.type;
    }
  }
  function hc(t) {
    return (
      t.match(/^uint($|[^1-9])/)
        ? (t = 'uint256' + t.substring(4))
        : t.match(/^int($|[^1-9])/) && (t = 'int256' + t.substring(3)),
      t
    );
  }
  const lc = new RegExp('^[a-zA-Z$_][a-zA-Z0-9$_]*$');
  function cc(t) {
    return (
      (t && t.match(lc)) ||
        Kl.throwArgumentError(`invalid identifier "${t}"`, 'value', t),
      t
    );
  }
  const fc = new RegExp('^([^)(]*)\\((.*)\\)([^)(]*)$');
  const dc = new f(ql);
  class mc {
    constructor(t, e, r, n) {
      (this.name = t),
        (this.type = e),
        (this.localName = r),
        (this.dynamic = n);
    }
    _throwError(t, e) {
      dc.throwArgumentError(t, this.localName, e);
    }
  }
  class pc {
    constructor(t) {
      st(this, 'wordSize', t || 32),
        (this._data = []),
        (this._dataLength = 0),
        (this._padding = new Uint8Array(t));
    }
    get data() {
      return P(this._data);
    }
    get length() {
      return this._dataLength;
    }
    _writeData(t) {
      return this._data.push(t), (this._dataLength += t.length), t.length;
    }
    appendWriter(t) {
      return this._writeData(b(t._data));
    }
    writeBytes(t) {
      let e = w(t);
      const r = e.length % this.wordSize;
      return r && (e = b([e, this._padding.slice(r)])), this._writeData(e);
    }
    _getValue(t) {
      let e = w(j.from(t));
      return (
        e.length > this.wordSize &&
          dc.throwError('value out-of-bounds', f.errors.BUFFER_OVERRUN, {
            length: this.wordSize,
            offset: e.length,
          }),
        e.length % this.wordSize &&
          (e = b([this._padding.slice(e.length % this.wordSize), e])),
        e
      );
    }
    writeValue(t) {
      return this._writeData(this._getValue(t));
    }
    writeUpdatableValue() {
      const t = this._data.length;
      return (
        this._data.push(this._padding),
        (this._dataLength += this.wordSize),
        (e) => {
          this._data[t] = this._getValue(e);
        }
      );
    }
  }
  class gc {
    constructor(t, e, r, n) {
      st(this, '_data', w(t)),
        st(this, 'wordSize', e || 32),
        st(this, '_coerceFunc', r),
        st(this, 'allowLoose', n),
        (this._offset = 0);
    }
    get data() {
      return k(this._data);
    }
    get consumed() {
      return this._offset;
    }
    static coerce(t, e) {
      let r = t.match('^u?int([0-9]+)$');
      return r && parseInt(r[1]) <= 48 && (e = e.toNumber()), e;
    }
    coerce(t, e) {
      return this._coerceFunc ? this._coerceFunc(t, e) : gc.coerce(t, e);
    }
    _peekBytes(t, e, r) {
      let n = Math.ceil(e / this.wordSize) * this.wordSize;
      return (
        this._offset + n > this._data.length &&
          (this.allowLoose && r && this._offset + e <= this._data.length
            ? (n = e)
            : dc.throwError('data out-of-bounds', f.errors.BUFFER_OVERRUN, {
                length: this._data.length,
                offset: this._offset + n,
              })),
        this._data.slice(this._offset, this._offset + n)
      );
    }
    subReader(t) {
      return new gc(
        this._data.slice(this._offset + t),
        this.wordSize,
        this._coerceFunc,
        this.allowLoose
      );
    }
    readBytes(t, e) {
      let r = this._peekBytes(0, t, !!e);
      return (this._offset += r.length), r.slice(0, t);
    }
    readValue() {
      return j.from(this.readBytes(this.wordSize));
    }
  }
  class yc extends mc {
    constructor(t) {
      super('address', 'address', t, !1);
    }
    defaultValue() {
      return '0x0000000000000000000000000000000000000000';
    }
    encode(t, e) {
      try {
        e = Re(e);
      } catch (t) {
        this._throwError(t.message, e);
      }
      return t.writeValue(e);
    }
    decode(t) {
      return Re(C(t.readValue().toHexString(), 20));
    }
  }
  class vc extends mc {
    constructor(t) {
      super(t.name, t.type, void 0, t.dynamic), (this.coder = t);
    }
    defaultValue() {
      return this.coder.defaultValue();
    }
    encode(t, e) {
      return this.coder.encode(t, e);
    }
    decode(t) {
      return this.coder.decode(t);
    }
  }
  const wc = new f(ql);
  function bc(t, e, r) {
    let n = null;
    if (Array.isArray(r)) n = r;
    else if (r && 'object' == typeof r) {
      let t = {};
      n = e.map((e) => {
        const n = e.localName;
        return (
          n ||
            wc.throwError(
              'cannot encode object for signature with missing names',
              f.errors.INVALID_ARGUMENT,
              { argument: 'values', coder: e, value: r }
            ),
          t[n] &&
            wc.throwError(
              'cannot encode object for signature with duplicate names',
              f.errors.INVALID_ARGUMENT,
              { argument: 'values', coder: e, value: r }
            ),
          (t[n] = !0),
          r[n]
        );
      });
    } else wc.throwArgumentError('invalid tuple value', 'tuple', r);
    e.length !== n.length &&
      wc.throwArgumentError('types/value length mismatch', 'tuple', r);
    let i = new pc(t.wordSize),
      o = new pc(t.wordSize),
      s = [];
    e.forEach((t, e) => {
      let r = n[e];
      if (t.dynamic) {
        let e = o.length;
        t.encode(o, r);
        let n = i.writeUpdatableValue();
        s.push((t) => {
          n(t + e);
        });
      } else t.encode(i, r);
    }),
      s.forEach((t) => {
        t(i.length);
      });
    let a = t.appendWriter(i);
    return (a += t.appendWriter(o)), a;
  }
  function Ac(t, e) {
    let r = [],
      n = t.subReader(0);
    e.forEach((e) => {
      let i = null;
      if (e.dynamic) {
        let r = t.readValue(),
          o = n.subReader(r.toNumber());
        try {
          i = e.decode(o);
        } catch (t) {
          if (t.code === f.errors.BUFFER_OVERRUN) throw t;
          (i = t),
            (i.baseType = e.name),
            (i.name = e.localName),
            (i.type = e.type);
        }
      } else
        try {
          i = e.decode(t);
        } catch (t) {
          if (t.code === f.errors.BUFFER_OVERRUN) throw t;
          (i = t),
            (i.baseType = e.name),
            (i.name = e.localName),
            (i.type = e.type);
        }
      null != i && r.push(i);
    });
    const i = e.reduce((t, e) => {
      const r = e.localName;
      return r && (t[r] || (t[r] = 0), t[r]++), t;
    }, {});
    e.forEach((t, e) => {
      let n = t.localName;
      if (!n || 1 !== i[n]) return;
      if (('length' === n && (n = '_length'), null != r[n])) return;
      const o = r[e];
      o instanceof Error
        ? Object.defineProperty(r, n, {
            enumerable: !0,
            get: () => {
              throw o;
            },
          })
        : (r[n] = o);
    });
    for (let t = 0; t < r.length; t++) {
      const e = r[t];
      e instanceof Error &&
        Object.defineProperty(r, t, {
          enumerable: !0,
          get: () => {
            throw e;
          },
        });
    }
    return Object.freeze(r);
  }
  class Ec extends mc {
    constructor(t, e, r) {
      super(
        'array',
        t.type + '[' + (e >= 0 ? e : '') + ']',
        r,
        -1 === e || t.dynamic
      ),
        (this.coder = t),
        (this.length = e);
    }
    defaultValue() {
      const t = this.coder.defaultValue(),
        e = [];
      for (let r = 0; r < this.length; r++) e.push(t);
      return e;
    }
    encode(t, e) {
      Array.isArray(e) || this._throwError('expected array value', e);
      let r = this.length;
      -1 === r && ((r = e.length), t.writeValue(e.length)),
        wc.checkArgumentCount(
          e.length,
          r,
          'coder array' + (this.localName ? ' ' + this.localName : '')
        );
      let n = [];
      for (let t = 0; t < e.length; t++) n.push(this.coder);
      return bc(t, n, e);
    }
    decode(t) {
      let e = this.length;
      -1 === e &&
        ((e = t.readValue().toNumber()),
        32 * e > t._data.length &&
          wc.throwError('insufficient data length', f.errors.BUFFER_OVERRUN, {
            length: t._data.length,
            count: e,
          }));
      let r = [];
      for (let t = 0; t < e; t++) r.push(new vc(this.coder));
      return t.coerce(this.name, Ac(t, r));
    }
  }
  class Mc extends mc {
    constructor(t) {
      super('bool', 'bool', t, !1);
    }
    defaultValue() {
      return !1;
    }
    encode(t, e) {
      return t.writeValue(e ? 1 : 0);
    }
    decode(t) {
      return t.coerce(this.type, !t.readValue().isZero());
    }
  }
  class _c extends mc {
    constructor(t, e) {
      super(t, t, e, !0);
    }
    defaultValue() {
      return '0x';
    }
    encode(t, e) {
      e = w(e);
      let r = t.writeValue(e.length);
      return (r += t.writeBytes(e)), r;
    }
    decode(t) {
      return t.readBytes(t.readValue().toNumber(), !0);
    }
  }
  class kc extends _c {
    constructor(t) {
      super('bytes', t);
    }
    decode(t) {
      return t.coerce(this.name, k(super.decode(t)));
    }
  }
  class Sc extends mc {
    constructor(t, e) {
      let r = 'bytes' + String(t);
      super(r, r, e, !1), (this.size = t);
    }
    defaultValue() {
      return '0x0000000000000000000000000000000000000000000000000000000000000000'.substring(
        0,
        2 + 2 * this.size
      );
    }
    encode(t, e) {
      let r = w(e);
      return (
        r.length !== this.size && this._throwError('incorrect data length', e),
        t.writeBytes(r)
      );
    }
    decode(t) {
      return t.coerce(this.name, k(t.readBytes(this.size)));
    }
  }
  class Nc extends mc {
    constructor(t) {
      super('null', '', t, !1);
    }
    defaultValue() {
      return null;
    }
    encode(t, e) {
      return null != e && this._throwError('not null', e), t.writeBytes([]);
    }
    decode(t) {
      return t.readBytes(0), t.coerce(this.name, null);
    }
  }
  class Pc extends mc {
    constructor(t, e, r) {
      const n = (e ? 'int' : 'uint') + 8 * t;
      super(n, n, r, !1), (this.size = t), (this.signed = e);
    }
    defaultValue() {
      return 0;
    }
    encode(t, e) {
      let r = j.from(e),
        n = Ze.mask(8 * t.wordSize);
      if (this.signed) {
        let t = n.mask(8 * this.size - 1);
        (r.gt(t) || r.lt(t.add(We).mul(Ve))) &&
          this._throwError('value out-of-bounds', e);
      } else (r.lt(Je) || r.gt(n.mask(8 * this.size))) && this._throwError('value out-of-bounds', e);
      return (
        (r = r.toTwos(8 * this.size).mask(8 * this.size)),
        this.signed && (r = r.fromTwos(8 * this.size).toTwos(8 * t.wordSize)),
        t.writeValue(r)
      );
    }
    decode(t) {
      let e = t.readValue().mask(8 * this.size);
      return (
        this.signed && (e = e.fromTwos(8 * this.size)), t.coerce(this.name, e)
      );
    }
  }
  class xc extends _c {
    constructor(t) {
      super('string', t);
    }
    defaultValue() {
      return '';
    }
    encode(t, e) {
      return super.encode(t, hr(e));
    }
    decode(t) {
      return fr(super.decode(t));
    }
  }
  class Tc extends mc {
    constructor(t, e) {
      let r = !1;
      const n = [];
      t.forEach((t) => {
        t.dynamic && (r = !0), n.push(t.type);
      });
      super('tuple', 'tuple(' + n.join(',') + ')', e, r), (this.coders = t);
    }
    defaultValue() {
      const t = [];
      this.coders.forEach((e) => {
        t.push(e.defaultValue());
      });
      const e = this.coders.reduce((t, e) => {
        const r = e.localName;
        return r && (t[r] || (t[r] = 0), t[r]++), t;
      }, {});
      return (
        this.coders.forEach((r, n) => {
          let i = r.localName;
          i &&
            1 === e[i] &&
            ('length' === i && (i = '_length'), null == t[i] && (t[i] = t[n]));
        }),
        Object.freeze(t)
      );
    }
    encode(t, e) {
      return bc(t, this.coders, e);
    }
    decode(t) {
      return t.coerce(this.name, Ac(t, this.coders));
    }
  }
  const Cc = new f(ql),
    Rc = new RegExp(/^bytes([0-9]*)$/),
    Ic = new RegExp(/^(u?int)([0-9]*)$/);
  class Bc {
    constructor(t) {
      st(this, 'coerceFunc', t || null);
    }
    _getCoder(t) {
      switch (t.baseType) {
        case 'address':
          return new yc(t.name);
        case 'bool':
          return new Mc(t.name);
        case 'string':
          return new xc(t.name);
        case 'bytes':
          return new kc(t.name);
        case 'array':
          return new Ec(this._getCoder(t.arrayChildren), t.arrayLength, t.name);
        case 'tuple':
          return new Tc(
            (t.components || []).map((t) => this._getCoder(t)),
            t.name
          );
        case '':
          return new Nc(t.name);
      }
      let e = t.type.match(Ic);
      if (e) {
        let r = parseInt(e[2] || '256');
        return (
          (0 === r || r > 256 || r % 8 != 0) &&
            Cc.throwArgumentError(
              'invalid ' + e[1] + ' bit length',
              'param',
              t
            ),
          new Pc(r / 8, 'int' === e[1], t.name)
        );
      }
      if (((e = t.type.match(Rc)), e)) {
        let r = parseInt(e[1]);
        return (
          (0 === r || r > 32) &&
            Cc.throwArgumentError('invalid bytes length', 'param', t),
          new Sc(r, t.name)
        );
      }
      return Cc.throwArgumentError('invalid type', 'type', t.type);
    }
    _getWordSize() {
      return 32;
    }
    _getReader(t, e) {
      return new gc(t, this._getWordSize(), this.coerceFunc, e);
    }
    _getWriter() {
      return new pc(this._getWordSize());
    }
    getDefaultValue(t) {
      const e = t.map((t) => this._getCoder(Xl.from(t)));
      return new Tc(e, '_').defaultValue();
    }
    encode(t, e) {
      t.length !== e.length &&
        Cc.throwError(
          'types/values length mismatch',
          f.errors.INVALID_ARGUMENT,
          {
            count: { types: t.length, values: e.length },
            value: { types: t, values: e },
          }
        );
      const r = t.map((t) => this._getCoder(Xl.from(t))),
        n = new Tc(r, '_'),
        i = this._getWriter();
      return n.encode(i, e), i.data;
    }
    decode(t, e, r) {
      const n = t.map((t) => this._getCoder(Xl.from(t)));
      return new Tc(n, '_').decode(this._getReader(w(e), r));
    }
  }
  const Oc = new Bc(),
    Fc = new f(ql);
  class Dc extends pt {}
  class Lc extends pt {}
  class Uc extends pt {}
  class jc extends pt {
    static isIndexed(t) {
      return !(!t || !t._isIndexed);
    }
  }
  const zc = {
    '0x08c379a0': {
      signature: 'Error(string)',
      name: 'Error',
      inputs: ['string'],
      reason: !0,
    },
    '0x4e487b71': {
      signature: 'Panic(uint256)',
      name: 'Panic',
      inputs: ['uint256'],
    },
  };
  function Gc(t, e) {
    const r = new Error(
      `deferred error during ABI decoding triggered accessing ${t}`
    );
    return (r.error = e), r;
  }
  const qc = new RegExp('^bytes([0-9]+)$'),
    Kc = new RegExp('^(u?int)([0-9]*)$'),
    Hc = new RegExp('^(.*)\\[([0-9]*)\\]$'),
    Vc = new f('solidity/5.7.0');
  function Jc(t, e, r) {
    switch (t) {
      case 'address':
        return r ? E(e, 32) : w(e);
      case 'string':
        return hr(e);
      case 'bytes':
        return w(e);
      case 'bool':
        return (e = e ? '0x01' : '0x00'), r ? E(e, 32) : w(e);
    }
    let n = t.match(Kc);
    if (n) {
      let i = parseInt(n[2] || '256');
      return (
        ((n[2] && String(i) !== n[2]) || i % 8 != 0 || 0 === i || i > 256) &&
          Vc.throwArgumentError('invalid number type', 'type', t),
        r && (i = 256),
        E((e = j.from(e).toTwos(i)), i / 8)
      );
    }
    if (((n = t.match(qc)), n)) {
      const i = parseInt(n[1]);
      return (
        (String(i) !== n[1] || 0 === i || i > 32) &&
          Vc.throwArgumentError('invalid bytes type', 'type', t),
        w(e).byteLength !== i &&
          Vc.throwArgumentError(`invalid value for ${t}`, 'value', e),
        r
          ? w(
              (
                e +
                '0000000000000000000000000000000000000000000000000000000000000000'
              ).substring(0, 66)
            )
          : e
      );
    }
    if (((n = t.match(Hc)), n && Array.isArray(e))) {
      const r = n[1];
      parseInt(n[2] || String(e.length)) != e.length &&
        Vc.throwArgumentError(`invalid array length for ${t}`, 'value', e);
      const i = [];
      return (
        e.forEach(function (t) {
          i.push(Jc(r, t, !0));
        }),
        b(i)
      );
    }
    return Vc.throwArgumentError('invalid type', 'type', t);
  }
  function Wc(t, e) {
    t.length != e.length &&
      Vc.throwArgumentError(
        'wrong number of values; expected ${ types.length }',
        'values',
        e
      );
    const r = [];
    return (
      t.forEach(function (t, n) {
        r.push(Jc(t, e[n]));
      }),
      k(b(r))
    );
  }
  const Qc = new f('units/5.7.0'),
    Yc = ['wei', 'kwei', 'mwei', 'gwei', 'szabo', 'finney', 'ether'];
  function Zc(t, e) {
    if ('string' == typeof e) {
      const t = Yc.indexOf(e);
      -1 !== t && (e = 3 * t);
    }
    return X(t, null != e ? e : 18);
  }
  function Xc(t, e) {
    if (
      ('string' != typeof t &&
        Qc.throwArgumentError('value must be a string', 'value', t),
      'string' == typeof e)
    ) {
      const t = Yc.indexOf(e);
      -1 !== t && (e = 3 * t);
    }
    return $(t, null != e ? e : 18);
  }
  var $c = Object.freeze({
    __proto__: null,
    AbiCoder: Bc,
    defaultAbiCoder: Oc,
    Fragment: tc,
    ConstructorFragment: oc,
    ErrorFragment: uc,
    EventFragment: ec,
    FunctionFragment: sc,
    ParamType: Xl,
    FormatTypes: Yl,
    checkResultErrors: function (t) {
      const e = [],
        r = function (t, n) {
          if (Array.isArray(n))
            for (let i in n) {
              const o = t.slice();
              o.push(i);
              try {
                r(o, n[i]);
              } catch (t) {
                e.push({ path: o, error: t });
              }
            }
        };
      return r([], t), e;
    },
    Logger: f,
    RLP: Se,
    _fetchData: Ou,
    fetchJson: Fu,
    poll: Du,
    checkProperties: ht,
    deepCopy: mt,
    defineReadOnly: st,
    getStatic: at,
    resolveProperties: ut,
    shallowCopy: lt,
    arrayify: w,
    concat: b,
    stripZeros: A,
    zeroPad: E,
    isBytes: v,
    isBytesLike: g,
    defaultPath: Da,
    HDNode: La,
    SigningKey: Cs,
    Interface: class {
      constructor(t) {
        let e = [];
        (e = 'string' == typeof t ? JSON.parse(t) : t),
          st(
            this,
            'fragments',
            e.map((t) => tc.from(t)).filter((t) => null != t)
          ),
          st(this, '_abiCoder', at(new.target, 'getAbiCoder')()),
          st(this, 'functions', {}),
          st(this, 'errors', {}),
          st(this, 'events', {}),
          st(this, 'structs', {}),
          this.fragments.forEach((t) => {
            let e = null;
            switch (t.type) {
              case 'constructor':
                return this.deploy
                  ? void Fc.warn('duplicate definition - constructor')
                  : void st(this, 'deploy', t);
              case 'function':
                e = this.functions;
                break;
              case 'event':
                e = this.events;
                break;
              case 'error':
                e = this.errors;
                break;
              default:
                return;
            }
            let r = t.format();
            e[r] ? Fc.warn('duplicate definition - ' + r) : (e[r] = t);
          }),
          this.deploy ||
            st(this, 'deploy', oc.from({ payable: !1, type: 'constructor' })),
          st(this, '_isInterface', !0);
      }
      format(t) {
        t || (t = Yl.full),
          t === Yl.sighash &&
            Fc.throwArgumentError(
              'interface does not support formatting sighash',
              'format',
              t
            );
        const e = this.fragments.map((e) => e.format(t));
        return t === Yl.json ? JSON.stringify(e.map((t) => JSON.parse(t))) : e;
      }
      static getAbiCoder() {
        return Oc;
      }
      static getAddress(t) {
        return Re(t);
      }
      static getSighash(t) {
        return N(Tr(t.format()), 0, 4);
      }
      static getEventTopic(t) {
        return Tr(t.format());
      }
      getFunction(t) {
        if (M(t)) {
          for (const e in this.functions)
            if (t === this.getSighash(e)) return this.functions[e];
          Fc.throwArgumentError('no matching function', 'sighash', t);
        }
        if (-1 === t.indexOf('(')) {
          const e = t.trim(),
            r = Object.keys(this.functions).filter(
              (t) => t.split('(')[0] === e
            );
          return (
            0 === r.length
              ? Fc.throwArgumentError('no matching function', 'name', e)
              : r.length > 1 &&
                Fc.throwArgumentError('multiple matching functions', 'name', e),
            this.functions[r[0]]
          );
        }
        const e = this.functions[sc.fromString(t).format()];
        return (
          e || Fc.throwArgumentError('no matching function', 'signature', t), e
        );
      }
      getEvent(t) {
        if (M(t)) {
          const e = t.toLowerCase();
          for (const t in this.events)
            if (e === this.getEventTopic(t)) return this.events[t];
          Fc.throwArgumentError('no matching event', 'topichash', e);
        }
        if (-1 === t.indexOf('(')) {
          const e = t.trim(),
            r = Object.keys(this.events).filter((t) => t.split('(')[0] === e);
          return (
            0 === r.length
              ? Fc.throwArgumentError('no matching event', 'name', e)
              : r.length > 1 &&
                Fc.throwArgumentError('multiple matching events', 'name', e),
            this.events[r[0]]
          );
        }
        const e = this.events[ec.fromString(t).format()];
        return (
          e || Fc.throwArgumentError('no matching event', 'signature', t), e
        );
      }
      getError(t) {
        if (M(t)) {
          const e = at(this.constructor, 'getSighash');
          for (const r in this.errors) {
            if (t === e(this.errors[r])) return this.errors[r];
          }
          Fc.throwArgumentError('no matching error', 'sighash', t);
        }
        if (-1 === t.indexOf('(')) {
          const e = t.trim(),
            r = Object.keys(this.errors).filter((t) => t.split('(')[0] === e);
          return (
            0 === r.length
              ? Fc.throwArgumentError('no matching error', 'name', e)
              : r.length > 1 &&
                Fc.throwArgumentError('multiple matching errors', 'name', e),
            this.errors[r[0]]
          );
        }
        const e = this.errors[sc.fromString(t).format()];
        return (
          e || Fc.throwArgumentError('no matching error', 'signature', t), e
        );
      }
      getSighash(t) {
        if ('string' == typeof t)
          try {
            t = this.getFunction(t);
          } catch (e) {
            try {
              t = this.getError(t);
            } catch (t) {
              throw e;
            }
          }
        return at(this.constructor, 'getSighash')(t);
      }
      getEventTopic(t) {
        return (
          'string' == typeof t && (t = this.getEvent(t)),
          at(this.constructor, 'getEventTopic')(t)
        );
      }
      _decodeParams(t, e) {
        return this._abiCoder.decode(t, e);
      }
      _encodeParams(t, e) {
        return this._abiCoder.encode(t, e);
      }
      encodeDeploy(t) {
        return this._encodeParams(this.deploy.inputs, t || []);
      }
      decodeErrorResult(t, e) {
        'string' == typeof t && (t = this.getError(t));
        const r = w(e);
        return (
          k(r.slice(0, 4)) !== this.getSighash(t) &&
            Fc.throwArgumentError(
              `data signature does not match error ${t.name}.`,
              'data',
              k(r)
            ),
          this._decodeParams(t.inputs, r.slice(4))
        );
      }
      encodeErrorResult(t, e) {
        return (
          'string' == typeof t && (t = this.getError(t)),
          k(b([this.getSighash(t), this._encodeParams(t.inputs, e || [])]))
        );
      }
      decodeFunctionData(t, e) {
        'string' == typeof t && (t = this.getFunction(t));
        const r = w(e);
        return (
          k(r.slice(0, 4)) !== this.getSighash(t) &&
            Fc.throwArgumentError(
              `data signature does not match function ${t.name}.`,
              'data',
              k(r)
            ),
          this._decodeParams(t.inputs, r.slice(4))
        );
      }
      encodeFunctionData(t, e) {
        return (
          'string' == typeof t && (t = this.getFunction(t)),
          k(b([this.getSighash(t), this._encodeParams(t.inputs, e || [])]))
        );
      }
      decodeFunctionResult(t, e) {
        'string' == typeof t && (t = this.getFunction(t));
        let r = w(e),
          n = null,
          i = '',
          o = null,
          s = null,
          a = null;
        switch (r.length % this._abiCoder._getWordSize()) {
          case 0:
            try {
              return this._abiCoder.decode(t.outputs, r);
            } catch (t) {}
            break;
          case 4: {
            const t = k(r.slice(0, 4)),
              e = zc[t];
            if (e)
              (o = this._abiCoder.decode(e.inputs, r.slice(4))),
                (s = e.name),
                (a = e.signature),
                e.reason && (n = o[0]),
                'Error' === s
                  ? (i = `; VM Exception while processing transaction: reverted with reason string ${JSON.stringify(
                      o[0]
                    )}`)
                  : 'Panic' === s &&
                    (i = `; VM Exception while processing transaction: reverted with panic code ${o[0]}`);
            else
              try {
                const e = this.getError(t);
                (o = this._abiCoder.decode(e.inputs, r.slice(4))),
                  (s = e.name),
                  (a = e.format());
              } catch (t) {}
            break;
          }
        }
        return Fc.throwError(
          'call revert exception' + i,
          f.errors.CALL_EXCEPTION,
          {
            method: t.format(),
            data: k(e),
            errorArgs: o,
            errorName: s,
            errorSignature: a,
            reason: n,
          }
        );
      }
      encodeFunctionResult(t, e) {
        return (
          'string' == typeof t && (t = this.getFunction(t)),
          k(this._abiCoder.encode(t.outputs, e || []))
        );
      }
      encodeFilterTopics(t, e) {
        'string' == typeof t && (t = this.getEvent(t)),
          e.length > t.inputs.length &&
            Fc.throwError(
              'too many arguments for ' + t.format(),
              f.errors.UNEXPECTED_ARGUMENT,
              { argument: 'values', value: e }
            );
        let r = [];
        t.anonymous || r.push(this.getEventTopic(t));
        const n = (t, e) =>
          'string' === t.type
            ? Tr(e)
            : 'bytes' === t.type
            ? ye(k(e))
            : ('bool' === t.type &&
                'boolean' == typeof e &&
                (e = e ? '0x01' : '0x00'),
              t.type.match(/^u?int/) && (e = j.from(e).toHexString()),
              'address' === t.type && this._abiCoder.encode(['address'], [e]),
              C(k(e), 32));
        for (
          e.forEach((e, i) => {
            let o = t.inputs[i];
            o.indexed
              ? null == e
                ? r.push(null)
                : 'array' === o.baseType || 'tuple' === o.baseType
                ? Fc.throwArgumentError(
                    'filtering with tuples or arrays not supported',
                    'contract.' + o.name,
                    e
                  )
                : Array.isArray(e)
                ? r.push(e.map((t) => n(o, t)))
                : r.push(n(o, e))
              : null != e &&
                Fc.throwArgumentError(
                  'cannot filter non-indexed parameters; must be null',
                  'contract.' + o.name,
                  e
                );
          });
          r.length && null === r[r.length - 1];

        )
          r.pop();
        return r;
      }
      encodeEventLog(t, e) {
        'string' == typeof t && (t = this.getEvent(t));
        const r = [],
          n = [],
          i = [];
        return (
          t.anonymous || r.push(this.getEventTopic(t)),
          e.length !== t.inputs.length &&
            Fc.throwArgumentError(
              'event arguments/values mismatch',
              'values',
              e
            ),
          t.inputs.forEach((t, o) => {
            const s = e[o];
            if (t.indexed)
              if ('string' === t.type) r.push(Tr(s));
              else if ('bytes' === t.type) r.push(ye(s));
              else {
                if ('tuple' === t.baseType || 'array' === t.baseType)
                  throw new Error('not implemented');
                r.push(this._abiCoder.encode([t.type], [s]));
              }
            else n.push(t), i.push(s);
          }),
          { data: this._abiCoder.encode(n, i), topics: r }
        );
      }
      decodeEventLog(t, e, r) {
        if (
          ('string' == typeof t && (t = this.getEvent(t)),
          null != r && !t.anonymous)
        ) {
          let e = this.getEventTopic(t);
          (M(r[0], 32) && r[0].toLowerCase() === e) ||
            Fc.throwError(
              'fragment/topic mismatch',
              f.errors.INVALID_ARGUMENT,
              { argument: 'topics[0]', expected: e, value: r[0] }
            ),
            (r = r.slice(1));
        }
        let n = [],
          i = [],
          o = [];
        t.inputs.forEach((t, e) => {
          t.indexed
            ? 'string' === t.type ||
              'bytes' === t.type ||
              'tuple' === t.baseType ||
              'array' === t.baseType
              ? (n.push(Xl.fromObject({ type: 'bytes32', name: t.name })),
                o.push(!0))
              : (n.push(t), o.push(!1))
            : (i.push(t), o.push(!1));
        });
        let s = null != r ? this._abiCoder.decode(n, b(r)) : null,
          a = this._abiCoder.decode(i, e, !0),
          u = [],
          h = 0,
          l = 0;
        t.inputs.forEach((t, e) => {
          if (t.indexed)
            if (null == s) u[e] = new jc({ _isIndexed: !0, hash: null });
            else if (o[e]) u[e] = new jc({ _isIndexed: !0, hash: s[l++] });
            else
              try {
                u[e] = s[l++];
              } catch (t) {
                u[e] = t;
              }
          else
            try {
              u[e] = a[h++];
            } catch (t) {
              u[e] = t;
            }
          if (t.name && null == u[t.name]) {
            const r = u[e];
            r instanceof Error
              ? Object.defineProperty(u, t.name, {
                  enumerable: !0,
                  get: () => {
                    throw Gc(`property ${JSON.stringify(t.name)}`, r);
                  },
                })
              : (u[t.name] = r);
          }
        });
        for (let t = 0; t < u.length; t++) {
          const e = u[t];
          e instanceof Error &&
            Object.defineProperty(u, t, {
              enumerable: !0,
              get: () => {
                throw Gc(`index ${t}`, e);
              },
            });
        }
        return Object.freeze(u);
      }
      parseTransaction(t) {
        let e = this.getFunction(t.data.substring(0, 10).toLowerCase());
        return e
          ? new Lc({
              args: this._abiCoder.decode(
                e.inputs,
                '0x' + t.data.substring(10)
              ),
              functionFragment: e,
              name: e.name,
              signature: e.format(),
              sighash: this.getSighash(e),
              value: j.from(t.value || '0'),
            })
          : null;
      }
      parseLog(t) {
        let e = this.getEvent(t.topics[0]);
        return !e || e.anonymous
          ? null
          : new Dc({
              eventFragment: e,
              name: e.name,
              signature: e.format(),
              topic: this.getEventTopic(e),
              args: this.decodeEventLog(e, t.data, t.topics),
            });
      }
      parseError(t) {
        const e = k(t);
        let r = this.getError(e.substring(0, 10).toLowerCase());
        return r
          ? new Uc({
              args: this._abiCoder.decode(r.inputs, '0x' + e.substring(10)),
              errorFragment: r,
              name: r.name,
              signature: r.format(),
              sighash: this.getSighash(r),
            })
          : null;
      }
      static isInterface(t) {
        return !(!t || !t._isInterface);
      }
    },
    LogDescription: Dc,
    TransactionDescription: Lc,
    base58: pa,
    base64: Br,
    hexlify: k,
    isHexString: M,
    hexConcat: P,
    hexStripZeros: T,
    hexValue: x,
    hexZeroPad: C,
    hexDataLength: S,
    hexDataSlice: N,
    nameprep: function (t) {
      if (t.match(/^[a-z0-9-]*$/i) && t.length <= 59) return t.toLowerCase();
      let e = dr(t);
      var r;
      (r = e.map((t) => {
        if (vr.indexOf(t) >= 0) return [];
        if (t >= 65024 && t <= 65039) return [];
        let e = (function (t) {
          let e = gr(t, wr);
          if (e) return [t + e.s];
          let r = br[t];
          if (r) return r;
          let n = Ar[t];
          return n ? [t + n[0]] : Er[t] || null;
        })(t);
        return e || [t];
      })),
        (e = r.reduce(
          (t, e) => (
            e.forEach((e) => {
              t.push(e);
            }),
            t
          ),
          []
        )),
        (e = dr(cr(e), ir.NFKC)),
        e.forEach((t) => {
          if (gr(t, Mr)) throw new Error('STRINGPREP_CONTAINS_PROHIBITED');
        }),
        e.forEach((t) => {
          if (gr(t, yr)) throw new Error('STRINGPREP_CONTAINS_UNASSIGNED');
        });
      let n = cr(e);
      if (
        '-' === n.substring(0, 1) ||
        '--' === n.substring(2, 4) ||
        '-' === n.substring(n.length - 1)
      )
        throw new Error('invalid hyphen');
      return n;
    },
    _toEscapedUtf8String: function (t, e) {
      return (
        '"' +
        ur(t, e)
          .map((t) => {
            if (t < 256) {
              switch (t) {
                case 8:
                  return '\\b';
                case 9:
                  return '\\t';
                case 10:
                  return '\\n';
                case 13:
                  return '\\r';
                case 34:
                  return '\\"';
                case 92:
                  return '\\\\';
              }
              if (t >= 32 && t < 127) return String.fromCharCode(t);
            }
            return t <= 65535
              ? lr(t)
              : lr(55296 + (((t -= 65536) >> 10) & 1023)) +
                  lr(56320 + (1023 & t));
          })
          .join('') +
        '"'
      );
    },
    toUtf8Bytes: hr,
    toUtf8CodePoints: dr,
    toUtf8String: fr,
    Utf8ErrorFuncs: ar,
    formatBytes32String: function (t) {
      const e = hr(t);
      if (e.length > 31)
        throw new Error('bytes32 string must be less than 32 bytes');
      return k(b([e, tr]).slice(0, 32));
    },
    parseBytes32String: function (t) {
      const e = w(t);
      if (32 !== e.length)
        throw new Error('invalid bytes32 - not 32 bytes long');
      if (0 !== e[31])
        throw new Error('invalid bytes32 string - no null terminator');
      let r = 31;
      for (; 0 === e[r - 1]; ) r--;
      return fr(e.slice(0, r));
    },
    dnsEncode: sn,
    hashMessage: an,
    namehash: on,
    isValidName: function (t) {
      try {
        return 0 !== nn(t).length;
      } catch (t) {}
      return !1;
    },
    id: Tr,
    _TypedDataEncoder: Mn,
    getAddress: Re,
    getIcapAddress: function (t) {
      let e = ((r = Re(t).substring(2)),
      new O(r, 16).toString(36)).toUpperCase();
      for (var r; e.length < 30; ) e = '0' + e;
      return 'XE' + Ce('XE00' + e) + e;
    },
    getContractAddress: Ie,
    getCreate2Address: function (t, e, r) {
      return (
        32 !== S(e) &&
          Ne.throwArgumentError('salt must be 32 bytes', 'salt', e),
        32 !== S(r) &&
          Ne.throwArgumentError(
            'initCodeHash must be 32 bytes',
            'initCodeHash',
            r
          ),
        Re(N(ye(b(['0xff', Re(t), e, r])), 12))
      );
    },
    isAddress: function (t) {
      try {
        return Re(t), !0;
      } catch (t) {}
      return !1;
    },
    formatEther: function (t) {
      return Zc(t, 18);
    },
    parseEther: function (t) {
      return Xc(t, 18);
    },
    formatUnits: Zc,
    parseUnits: Xc,
    commify: function (t) {
      const e = String(t).split('.');
      (e.length > 2 ||
        !e[0].match(/^-?[0-9]*$/) ||
        (e[1] && !e[1].match(/^[0-9]*$/)) ||
        '.' === t ||
        '-.' === t) &&
        Qc.throwArgumentError('invalid value', 'value', t);
      let r = e[0],
        n = '';
      for (
        '-' === r.substring(0, 1) && ((n = '-'), (r = r.substring(1)));
        '0' === r.substring(0, 1);

      )
        r = r.substring(1);
      '' === r && (r = '0');
      let i = '';
      for (
        2 === e.length && (i = '.' + (e[1] || '0'));
        i.length > 2 && '0' === i[i.length - 1];

      )
        i = i.substring(0, i.length - 1);
      const o = [];
      for (; r.length; ) {
        if (r.length <= 3) {
          o.unshift(r);
          break;
        }
        {
          const t = r.length - 3;
          o.unshift(r.substring(t)), (r = r.substring(0, t));
        }
      }
      return n + o.join(',') + i;
    },
    computeHmac: ba,
    keccak256: ye,
    ripemd160: va,
    sha256: wa,
    sha512: function (t) {
      return '0x' + Kn.sha512().update(w(t)).digest('hex');
    },
    randomBytes: Ha,
    shuffled: Va,
    solidityPack: Wc,
    solidityKeccak256: function (t, e) {
      return ye(Wc(t, e));
    },
    soliditySha256: function (t, e) {
      return wa(Wc(t, e));
    },
    splitSignature: R,
    joinSignature: I,
    accessListify: Ks,
    parseTransaction: Ys,
    serializeTransaction: Ws,
    get TransactionTypes() {
      return Os;
    },
    getJsonWalletAddress: function (t) {
      if (iu(t))
        try {
          return Re(JSON.parse(t).ethaddr);
        } catch (t) {
          return null;
        }
      if (ou(t))
        try {
          return Re(JSON.parse(t).address);
        } catch (t) {
          return null;
        }
      return null;
    },
    computeAddress: js,
    recoverAddress: zs,
    computePublicKey: Is,
    recoverPublicKey: Rs,
    verifyMessage: function (t, e) {
      return zs(an(t), e);
    },
    verifyTypedData: function (t, e, r, n) {
      return zs(Mn.hash(t, e, r), n);
    },
    getAccountPath: function (t) {
      return (
        ('number' != typeof t || t < 0 || t >= Ca || t % 1) &&
          Pa.throwArgumentError('invalid account index', 'index', t),
        `m/44'/60'/${t}'/0/0`
      );
    },
    mnemonicToEntropy: ja,
    entropyToMnemonic: za,
    isValidMnemonic: function (t, e) {
      try {
        return ja(t, e), !0;
      } catch (t) {}
      return !1;
    },
    mnemonicToSeed: Ua,
    get SupportedAlgorithm() {
      return ga;
    },
    get UnicodeNormalizationForm() {
      return ir;
    },
    get Utf8ErrorReason() {
      return or;
    },
    Indexed: jc,
  });
  const tf = 'ethers/5.7.2',
    ef = new f(tf);
  var rf = Object.freeze({
    __proto__: null,
    Signer: jn,
    Wallet: bu,
    VoidSigner: zn,
    getDefaultProvider: zl,
    providers: Gl,
    BaseContract: ca,
    Contract: fa,
    ContractFactory: da,
    BigNumber: j,
    FixedNumber: et,
    constants: er,
    get errors() {
      return l;
    },
    logger: ef,
    utils: $c,
    wordlists: Na,
    version: tf,
    Wordlist: Ma,
  });
  try {
    const t = window;
    null == t._ethers && (t._ethers = rf);
  } catch (t) {}
  var nf = Object.freeze({
    __proto__: null,
    ethers: rf,
    Signer: jn,
    Wallet: bu,
    VoidSigner: zn,
    getDefaultProvider: zl,
    providers: Gl,
    BaseContract: ca,
    Contract: fa,
    ContractFactory: da,
    BigNumber: j,
    FixedNumber: et,
    constants: er,
    get errors() {
      return l;
    },
    logger: ef,
    utils: $c,
    wordlists: Na,
    version: tf,
    Wordlist: Ma,
  });
  'undefined' == typeof globalThis && (globalThis = global),
    (globalThis.ethers = nf);
});
