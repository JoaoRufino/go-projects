(function(window, $) {
  
  var $mapster = $('#map-canvas').mapster(Mapster.MAP_OPTIONS);

  $mapster.mapster('addMarker',{
    lat: 40.5,
    lng: -8.3,
    content: 'Acidente Normal',
    id:1
  });

  $mapster.mapster('addMarker',{
    lat: 40.5,
    lng: -8.4,
    content: 'Acidente Normal',
    icon: 'images/accident.svg',
    events: ['click'],
    id:2
  });

}(window,jQuery));