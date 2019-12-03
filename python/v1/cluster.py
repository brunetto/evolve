#!/usr/bin/python

from __future__ import division

import operator
import pickle
import numpy as Npy
from numpy import *
from numpy.random import *
from random import seed, random
from math import sqrt, pow, sin, cos, atan, fabs
from sys import exit, stdout
from copy import deepcopy

PI = 4*atan(1)

class Cluster() :
    def __init__(self, N) :
        self.m = zeros((N,1))
        self.r = zeros((N,3))
        self.v = zeros((N,3))
    def __iter__(self):
        return self
    def __repr__(self) :
        tt = 'mass:\n%s\n' % (self.m)
        tt += 'position:\n%s\n' % (self.r)
        tt += 'velocity:\n%s' % (self.v)
        return tt
    def sort(self) :
        cl_w_key = [ (dot(self.r[i], self.r[i]),
                     copy(self.m[i]), copy(self.r[i]), copy(self.v[i]))
                         for i in range(len(self.m))]
        cl_w_key.sort(key=operator.itemgetter(0))
        for i in range(len(self.m)) :
            self.m[i] = cl_w_key[i][1]
            self.r[i] = cl_w_key[i][2]
            self.v[i] = cl_w_key[i][3]

    def get_acc(self) :
        acc = zeros(shape(self.r))
        for i in range(len(self.r)) :
            Ri = self.r[i]
            mi = self.m[i][0]
            for j in range(i+1, len(self.r)) :
                Rj = self.r[j]
                mj = self.m[j][0]
                Rij = Ri-Rj
                apre = pow(dot(Rij,Rij), -3.0/2)
                acc[i] -= mj*apre*Rij
                acc[j] += mi*apre*Rij
        return acc

    def energy(self) :
        KE = 0.5*sum(self.v*self.v*self.m)
        PE = 0.0
        for i in range(len(self.r)) :
            Ri = self.r[i]
            mi = self.m[i][0]
            for j in range(i+1, len(self.r)) :
                Rj = self.r[j]
                mj = self.m[j][0]
                Rij = Ri-Rj
                PE -= mi*mj/sqrt(dot(Rij,Rij))
        return PE+KE,KE,PE

    def LagrangianRadii(self) :
        LR_perc = array([0.01, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 0.99]
                  )
        clu_cop = deepcopy(self)
        clu_cop.sort()
        LR_vals = LagRad(clu_cop, LR_perc)
        lr = zeros(len(LR_perc))
        for i in range(len(LR_perc)) :
            lr[i] = LR_vals[i]
        ll = column_stack((LR_perc[:,newaxis],lr[:,newaxis]))
        return ll

    def scale(self) :
        # this function scales positions, such that PE = -1/2
        # and scales velocities, such that KE = 1/4
        En,KE,PE = self.energy()
        self.r *= -PE*2.0
        self.v *= 1.0/sqrt(4.0*KE)

    def set_CM_to_zero(self) :
        Mtot = self.m.sum()
        CMr = (self.r*self.m).sum(axis=0) / Mtot
        CMv = (self.v*self.m).sum(axis=0) / Mtot
        self.r = self.r - CMr
        self.v = self.v - CMv
        CMr = (self.r*self.m).sum(axis=0) / Mtot
        CMv = (self.v*self.m).sum(axis=0) / Mtot

    def print_star(t) :
        EN,KE,PE = energy(clu)
        print t,
        print EN+0.25, KE-0.25, PE+0.5,
        for i in range(len(self.m)) :
            for k in range(3) :
                 print self.r[i][k],
                 for k in range(3) :
                     print self.v[i][k],
         print
         stdout.flush()
 
 def lin_interp(x_y, x) :
     # given a list of values x_y (both monotonically
     # increasing), find the corresponding y value for x
     if x<x_y[0][0] or x>=x_y[-1][0] :
         print "out of bounds"
         exit(-1)
     b = 0
     e = len(x_y)-1
     while 1: # loop invariant x_y[b][1]<= y < x_y[e][1]
         m = int((e+b)/2)
         if x_y[m][0] > x : e=m
         else : b=m
         if e-b==1 : break
     # XXX diagnostics only
     if e-b!=1 or x>=x_y[e][0] or x<x_y[b][0] :
         print 'problem with binary search!'
         exit(-2)
     # XXX end of diagnostics
     return (x_y[b][1] +
         (x_y[e][1]-x_y[b][1])*(x-x_y[b][0])/(x_y[e][0]-x_y[b][0]))
 
 def LagRad(clu, list=[0.01, 0.1, 0.2, 0.3, 0.4,
               0.5, 0.6, 0.7, 0.8, 0.9, 0.99]) :
     # calculates the lagrange radii of the cluster,
     # for the values given in the list
     # this routine assumes that the cluster is sorted
     # it does not assume that the total mass is unity
     m_r = [[0.0, 0.0]]
     mtot = 0.0
     for i in range(len(clu.m)) :
         r = sqrt(dot(clu.r[i], clu.r[i]))
         mtot += clu.m[i][0]
         m_r.append([mtot,r])
     m_r = [[x[0]/mtot, x[1]] for x in m_r]
     # XXX diagnostics only
 #    for x in m_r:
 #        print x
     # end of diagnostics
     radii = []
     for LR in list :
         radii.append(lin_interp(m_r, LR))
     return radii
 
 if __name__=="__main__":
 
     fd = open('Plummer.in', 'r')
     cl = pickle.load(fd)
     ET,KE,PE = cl.energy()
     print ET, KE, PE
     print cl.LagrangianRadii()