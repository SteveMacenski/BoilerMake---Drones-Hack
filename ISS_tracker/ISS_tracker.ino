
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
   Serial.flush();
   Serial.print(altitude);Serial.print(' ');Serial.println(asmuth);
}

//to ensure altitude is positive since servo is 0-180
altitude = (int)altitude;
myServo.write(altitude+90);
//myServo.write(90+20); //to make it move in demo for Apple TODO

// defining variables
double step_deg = 1.8; //degrees
int step_count = 0; //number of steps taken


//calculating need for change in stepper motor 
double asmuth_change = asmuth - asmuth_old;
int stepper_move = asmuth_change / step_deg;

// in case moving N-s or S-n
if (stepper_move > 1) {
  myMotor->step(stepper_move,FORWARD, SINGLE);
//  if (stepper_move != 0){
//    Serial.println(stepper_move);
//}
}
if (stepper_move < 0) {
  myMotor->step(-1*stepper_move,BACKWARD, SINGLE);
//  if (stepper_move != 0){
//    Serial.println(stepper_move);
//  }
}

// assigning asmuth to old one to store for delta calculation
asmuth_old = asmuth;

}
