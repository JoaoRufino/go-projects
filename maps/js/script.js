(function(window, $) {
  
  var $mapster = $('#map-canvas').mapster(Mapster.MAP_OPTIONS);

  $mapster.mapster('addMarker',{
    lat: 40.5,
    lng: -8.3,
    content: 'Acidente Normal',
  })

}(window,jQuery));