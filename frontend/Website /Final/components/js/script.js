var $mapster;

(function(window, $) {
  
  $mapster = $('#map-canvas').mapster(Mapster.MAP_OPTIONS);
  ws = new WebSocket('ws://echo.websocket.org/')
  ws.onopen = onOpen;
  ws.onmessage =onMessage;

}(window,jQuery));

  function onOpen (evt) {
      ws.send(JSON.stringify({ 'type' : 1 , 'id' : 1,'code' : 2, 'subcode': 0, 'lat': 40.53,'lng': -8.43,'createdAt': "16-05-2017" }));
   }
  
  function onMessage (evt)   {
    if(isJson(evt.data)) {
    ops = JSON.parse(evt.data);
    $mapster.mapster('addMarker',{
    lat: 40.52,
    lng: -8.45,
    content: "<p></b> Big Accident </b></p>"+ ops.createdAt,
    icon: "/components/images/accident_bus.svg",
    events: ['click'],
    id: ops.id
  });
      switch (ops.type) {
        case 1:
        $mapster.mapster('addMarker',createDENM(ops));
        break;
        case 2:

        break;

        case 0:
        break;
      }
  }
}

  function isJson(str) 
{
    try 
    {
        JSON.parse(str);
    } 
    catch (e) 
    {
        return false;
    }
    return true;
}

function DENM(ops) 
  switch (ops.code) {
    case 2:
      switch(ops.subcode) {

      }
      break;
    case 3:
      switch(ops.subcode) {

      }
      break;
    case 10:
      switch(ops.subcode) {

      }
      break;
    case 10:
      switch(ops.subcode) {

      }
      break; 
      
    //Animals on the road
    case 11:
      switch(ops.subcode) {

      }
      break;

    //People on the road 
    case 12:
      switch(ops.subcode) {

      }
      break;

    //Reduced visibility
    case 18:
      switch(ops.subcode) {

      }
      break;

    //Dangerous situation
    case 99:
      switch(ops.subcode) {

      }
      break;  




      accident="<p><b>Normal Accident</b></p>";
      ico="/components/images/accident.svg"
    break;
    default:
  }

 return  { lat: ops.lat,
    lng: ops.lng,
    content: accident + ops.createdAt,
    icon: ico,
    events: ['click'],
    id: ops.id
  }
}

function



