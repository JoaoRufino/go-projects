(function(window,google,mapster) {

mapster.MAP_OPTIONS = {
    center: {
      lat:  40.6333333,
      lng: -8.435883
    },
    zoom: 12,
    disableDefaultUI: false,
    scrollwheel:true,
    draggable: true,

    maxZoom:13,
    minZoom:11,
    zoomControlOptions: {
      position: google.maps.ControlPosition.TOP_LEFT,
      style: google.maps.ZoomControlStyle.DEFAULT
    }

  };

}(window,google,window.Mapster || (window.Mapster={})))