import serial
import math
import time
from datetime import datetime
import ephem

ser = serial.Serial("COM13",9600) #initializes connection with the arduino
time.sleep(5.0)

deg2rad = 180.0 / math.pi

home = ephem.Observer()
home.lon = '86.9290' #+E
home.lat = '40.4230' #+N
home.elevation = 185 #meters

#Always get the latest ISS TLE dat from http://spaceflight.nasa.gov/realdata/sighting
iss = ephem.readtle('ISS',
                        '1 25544U 98067A   15289.54097117  .00016717  00000-0  10270-3 0  9006',
                        '2 25544  51.6434 198.3749 0006726  41.6054 318.5610 15.54354269  6922')

while True:
    home.date = datetime.utcnow()
    iss.compute(home)
    altitude = iss.alt*deg2rad
    otherthing = iss.az*deg2rad
    string_altitude = str(altitude) + " " + str(otherthing)# + "\n" 
    ser.write(string_altitude)
    ser.flush()
    print(string_altitude)
    time.sleep(60.0)

#notes, the horizon is mapped wrong, should be tall to floor not horizon to horizon. Also make sure in set up that the stepper is in the right cardinal orientation of the ISS
