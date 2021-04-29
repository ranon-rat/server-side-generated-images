const receiveVideo =new WebSocket(`ws://${window.location.host}/ws/receiveVideo/`)
receiveVideo.onmessage=function(event){
 console.log(event.data.slice(0,30))
    document.getElementById("transmission").src= event.data;
    
}