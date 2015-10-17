#include <Adafruit_MotorShield.h>
#include <Time.h>

#define MESSAGE_LENGTH 30

struct State {
  double issLat;
  double issLon;
};

Adafruit_MotorShield AFMS = AdaFruit_MotorShield();
Adafruit_StepperMotor *Executor;

void setup() {
  Serial.begin(9600);
  Executor = AFMS.getStepper(48, 1);
}

void loop() {
  // TODO decide what to do
  while (Serial.available() >= MESSAGE_LENGTH) {
  }
}
