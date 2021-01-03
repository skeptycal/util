// package diophantine is a fun diversion into the world of theoretical mathematics.
/*

Reference: https://arxiv.org/pdf/1903.04284.pdf


1. Introduction

Let k be a positive integer with k 6≡ ±4 (mod 9). Then Heath-Brown [HB92] has conjectured that there are infinitely many triples (x, y, z) ∈ Z^3 such that

    (1)                 k = x^3 + y^3 + z^3

Various numerical investigations of (1) have been carried out, beginning as early as 1954
[MW55]; see [BPTYJ07] for a thorough account of the history of these investigations
up to 2000. The computations performed since that time have been dominated by an
algorithm due to Elkies [Elk00]. The latest that we are aware of is the paper of Huisman
[Hui16] (based on the implementation by Elsenhans and Jahnel [EJ09]), which determined
all solutions to (1) with k < 1000 and max{|x|, |y|, |z|} ≤ 1015. In particular, Huisman
reports that solutions are known for all but 13 values of k < 1000:
    (2) 33, 42, 114, 165, 390, 579, 627, 633, 732, 795, 906, 921, 975.

Elkies’ algorithm works by finding rational points near the Fermat curve X^3 + Y^3 = 1
using lattice basis reduction; it is well suited to finding solutions for many values of k
simultaneously. In this paper we describe a different approach that is more efficient when
k is fixed. It has the advantage of provably finding all solutions with a bound on the
smallest coordinate, rather than the largest as in Elkies’ algorithm. This always yields a
nontrivial expansion of the search range since, apart from finitely many exceptions that
can be accounted for separately, one has

  max{|x|, |y|, |z|} > 3√2 min{|x|, |y|, |z|}

Moreover, empirically it is often the case that one of the variables is much smaller than
the other two, so we expect the gain to be even greater in practice.

Our strategy is similar to some earlier approaches (see especially [HBLtR93], [Bre95],
[KTS97] and [BPTYJ07]), and is based on the observation that in any solution,

    k − z^3 = x^3 + y^3

has

    x + y

as a factor.  Our main contribution over the earlier investigations is to
note that with some time-space tradeoffs, the running time is very nearly linear in the
height bound, and it is quite practical when implemented on modern 64-bit computers.

In more detail, suppose that (x, y, z) is a solution to (1), and assume without loss of
generality that |x| ≥ |y| ≥ |z|. Then we have

    k − z^3 = x^3 + y^3 = (x + y)(x^2 − xy + y^2)



*/
package diophantine

func check() bool {
	return false
}
