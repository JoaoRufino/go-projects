var $mapster;

(function(window, $) {
  
  $mapster = $('#map-canvas').mapster(Mapster.MAP_OPTIONS);
  ws = new WebSocket('ws://localhost:4000/')
  ws.onopen = onOpen;
  ws.onmessage =onMessage;
  ws.onclose = onOpen;
  ws.onmessage =onMessage;

}(window,jQuery));

  function onOpen (evt) {
      ws.send(JSON.stringify({ 'type' : 1 , 'id' : 1,'code' : 2, 'subcode': 0, 'lat': 40.53,'lng': -8.43,'createdAt': "16-05-2017" }));
   }
  
  function onMessage (evt)   {
    if(isJson(evt.data)) {
    ops = JSON.parse(evt.data);
      switch (ops.type) {
        case 1:
        view=viewDENM(ops.code,ops.subcode,ops.createdAt)
        $mapster.mapster('addMarker',{
            lat: ops.lat,
            lng: ops.lng,
            content: view.content,
            icon: view.icon,
            events: ['click'],
            id: ops.id
        });
        break;
        case 2:

        break;

        case 0:
        break;

        default:
        break;
      }
  }
  console.log("ERROR MESSAGE NOT IN JSON")
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




  function viewDENM (code,subcode,timestamp) {
    switch (code) {
        case 2:
          switch(subcode) {
            //Common Accident
            case 0:
            content = "Normal car accident"
            icon = "/components/images/accident.svg"
            break;

            //Multiple Vehicle Accident
            case 1:
            content = "Car accident involving multiple vehicles"
            icon = "/components/images/accident_multiple_vehicles.svg"
            break;

            //Heavy Accident
            case 2:
            content = "Heavy car accident"
            icon = "/components/images/accident_heavy.svg"
            break;

            //Lorry Accident
            case 3:
            content = "Car accident involving lorry"
            icon = "/components/images/accident_lorry.svg"
            break;

            //Bus Accident
            case 4:
            content = "Bus accident"
            icon = "/components/images/accident_bus.svg"
            break;

            //Hazardous Accident
            case 5:
            content = "Car accident involving hazardous materials"
            icon = "/components/images/accident_hazardous_materials.svg"
            break;

            //Unsecured Accident
            case 7:
            content = "Unsecured accident"
            icon = "/components/images/accident_unsecure.svg"
            break;

            default:
            content = "Car accident"
            icon = "/components/images/default.svg"
            break;

        //RoadWorks
        case 3:
          switch(subcode) {
            case 0:
            content = "Road works"
            icon = "/components/images/road_works.svg"
            break;
            default:
            content = "Road works"
            icon = "/components/images/default.svg"
            break;
          }
          break;

        //Objects on the road
        case 10:
          switch(subcode) {
            case 0:
            content = "Objects on the road"
            icon = "/components/images/obstacle.svg"
            break;
            default:
            content = "Objects on the road"
            icon = "/components/images/default.svg"
            break;
          }
          break;

          
        //Animals on the road
        case 11:
          switch(subcode) {
            case 0:
            content = "Animal on the road"
            icon = "/components/images/animals_on_the_road.svg"
            break;
            case 1:
            content = "Wild animal on the road"
            icon = "/components/images/animals_on_the_road_wild.svg"
            break;
            case 2:
            content = "Herd of animals on the road"
            icon = "/components/images/animals_on_the_road_herd.svg"
            break;
            case 3:
            content = "Small animal on the road"
            icon = "/components/images/animals_on_the_road_small.svg"
            break;
            case 4:
            content = "Large animal on the road"
            icon = "/components/images/animals_on_the_road_large.svg"
            break;
            default:
            content = "Animals on the road"
            icon = "/components/images/default.svg"
            break;
          }
          }
          break;

        //People on the road 
        case 12:
          switch(subcode) {
            case 0:
            content = "Human presence on the road"
            icon = "/components/images/human_on_the_road.svg"
            break;
            default:
            content = "Human presence on the road"
            icon = "/components/images/default.svg"
            break;
          }
          break;

        //Reduced visibility
        case 18:
          switch(subcode) {
            case 0: 
            content = "Visibility Low"
            icon = "/components/images/visibility_low.svg"
            break;
            case 1:
            content = "Fog - Visibiliy Reduced"
            icon = "/components/images/visibility_fog.svg"
            break;
            default:
            content = "Reduced Visibility"
            icon = "/components/images/default.svg"
            break;
          }
          break;

        //Dangerous situation
        case 99:
          switch(subcode) {
            case 0:
            content = "Emergency eletronic break lights"
            icon = "/components/images/eletronic_breaks.svg"
            break;
            default:
            content = "Dangerous Situation"
            icon = "/components/images/default.svg"
            break;
          }
          break;  
        default:
          content = "DENM code unkown"
            icon = "/components/images/default.svg"
            break;
      }
      return { 'content' : '<div id="google-popup"><center><p><h1>' + content + '</h1> ' +timestamp + '</p></center>', 
        'icon' : icon }
  }

  function viewCAM (code) {
      switch (code) {
        default:
      }
    }