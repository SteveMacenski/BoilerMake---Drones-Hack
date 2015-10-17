
//includes 
#include <Wire.h>
#include <Adafruit_MotorShield.h>
#include <math.h>
#include <Servo.h>
#include "utility/Adafruit_PWMServoDriver.h"

//setting up motor shield
Adafruit_MotorShield AFMS = Adafruit_MotorShield(); 
Adafruit_StepperMotor *myMotor = AFMS.getStepper(200, 1);

//servos
Servo myServo;
double asmuth_old = 0; //variable to store i-1 value
double asmuth = 0; //start pointing north
double altitude = 90;//start pointing up

//attach pins and speeds
void setup() {
Serial.begin(9600);
Serial.println("ISS Tracker Initializing");
AFMS.begin();
myMotor->setSpeed(10);
myServo.attach(9);

}

//main loop
void loop() {
while (Serial.available())  {
   altitude = Serial.parseFloat();
   asmuth = Serial.parseFloat();
}

//to ensure altitude is positive since servo is 0-180
altitude = (int)altitude;
if (altitude >= 0 && altitude <=180){
  myServo.write(altitude);
}
if (altitude > 180 && altitude <360){
  myServo.write(altitude - 180);
}

// defining variables
double step_deg = 1.8; //degrees
int step_count = 0; //number of steps taken

//trial cases when rotating back around the 0-> 360 and 360->0 
if (abs(asmuth_old-asmuth) > 300 && asmuth_old > 300) {
  asmuth_old = asmuth_old - 360;
}
if (abs(asmuth_old-asmuth) > 300 && asmuth > 300) {
  asmuth_old = asmuth_old + 360;
}

//calculating need for change in stepper motor 
double asmuth_change = asmuth - asmuth_old;
int stepper_move = asmuth_change / step_deg;


// in case moving N-s or S-n
if (stepper_move > 1) {
  myMotor->step(stepper_move,FORWARD, SINGLE);
  if (stepper_move != 0){
    Serial.println(stepper_move);
}
}
if (stepper_move < -1) {
  myMotor->step(-1*stepper_move,BACKWARD, SINGLE);
  if (stepper_move != 0){
    Serial.println(stepper_move);
  }
}

// assigning asmuth to old one to store for delta calculation
asmuth_old = asmuth;

}
