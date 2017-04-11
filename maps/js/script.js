(function(window,mapster) {
  
  // map options
  var options= mapster.MAP_OPTIONS,
  element = document.getElementById('map-canvas'),
  // map
  map = Mapster.create(element, options);
//  map.zoom(18);
//  alert(map.zoom());()

/*    event:{
      name: 'dragend',
      callback: function(){
        alert('Hello!')
      }
    },*/
  var marker= map.addMarker({
    lat:  40.5333333,
    lng: -8.435883,
    visible: true,
    draggable: true,
    id: 1,
    content:'Acidente normal',
/*    event:{
      name: 'click',
      callback: function(){
        map._removeMarker(marker);
      }
    },*/
    icon: 'images/accident.svg'
    });

    var marker2= map.addMarker({
    lat:  40.6333333,
    lng: -8.435883,
    visible: true,
    draggable: true,
    id: 2,
    content:'Acidente autocarro',
    icon: 'images/accident_bus.svg'
    });

    var found = map.findBy(function(marker){
      return marker.draggable === true;
    });

    console.log(found);

}(window,window.Mapster));