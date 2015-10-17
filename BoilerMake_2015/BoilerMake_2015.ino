// code for BoilerMake 2015
// Steven Macenski
// making a 2D write plotter using arduino
// standard 2 link robotics conventions apply and use forward and inverse kinematic solutions to the classical problem
#include <math.h>
#include <Servo.h>

double l1 = 25; //cm
double l2 = 14; //cm

double theta1 = 0;
double theta2 = 0;

double x;
double y;

Servo link1;
Servo link2;

void setup() {
  Serial.begin(9600);
  link1.attach(9);
  link2.attach(10);
}

void loop() {
  // inverse kinematics
theta2 = atan2(sqrt(1-((x*x + y*y - l1*l1 - l2*l2)/(2*l1*l2))), (x*x + y*y - l1*l1 - l2*l2)/(2*l1*l2));
theta1 = atan2(y,x) - atan2(l1 + l2*cos(theta2), l2*sin(theta2));

//converting to degrees for writing
theta2 = theta2*(57296/1000);
theta1 = theta1*(57296/1000);

//writing to links
link1.write(theta1);
link2.write(theta2);

}
