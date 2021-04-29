"use strict";

const video = document.getElementById("video");
const canvas = document.getElementById("canvas");
const errorMsgElement = document.querySelector("span#errorMsg");

const sendVideo = new WebSocket(`ws://${window.location.host}/ws/sendVideo/`);
const init = async () => {
  while (true) {
    try {
      const stream = await navigator.mediaDevices.getUserMedia({
        audio: false,
        video: {
          facingMode: "user",
        },
      });
      success(stream);
    } catch (e) {}
  }
};

const success = async (stream) => {
  window.stream = stream;
  video.srcObject = stream;

  let context = canvas.getContext("2d");

  setInterval(async () => {
    // decode the images
    context.drawImage(video, 0, 0, 640, 480);

    let canvasData = canvas
      .toDataURL("image/png")
      .replace("image/png", "image/octet-stream");
    // here needs to send the image
  
    sendVideo.send(canvasData);
  }, 1000);}
  sendVideo.onopen=init();




  

init()