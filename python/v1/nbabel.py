#!/usr/bin/python

from __future__ import division

import operator
from numpy import *
from math import pow
from sys import exit, stdout
import pickle
import Cluster as sc

def drift_cluster(clu,dt) :
    clu.r += clu.v*dt

def kick_cluster(cl,dt) :
    clu.v += cl.get_acc(clu)*dt

def evolve_cluster_DKD_leapfrog(clu, dt) :
    # Drift-Kick-Drift leapfrog
    drift_cluster(clu,dt/2.0)
    kick_cluster(clu,dt)
    drift_cluster(clu,dt/2.0)

if __name__=="__main__":
    fd = open('Plummer.in', 'r')
    cl = pickle.load(fd)
    ET0,KE0,PE0 = cl.energy()

    t = 0.0
    tend = 1.0
    dt = 1e-3
    k = 0
    # Predictor-Corrector Leapfrog
    acc_0 = cl.get_acc()
    while t<tend :
        cl.r += dt*cl.v + 0.5*dt*dt*acc_0
        acc_1 = cl.get_acc()
        cl.v += 0.5*dt*(acc_0+acc_1)
        acc_0 = acc_1
        t += dt
        k += 1
        if k%10==0 :
            ET,KE,PE = cl.energy()
            print "t= ", t, " E= ", ET, KE, PE, \
                  " dE= ", (ET-ET0)/ET0, (KE-KE0)/KE0, (PE-PE0)/PE0
            print cl.LagrangianRadii()
