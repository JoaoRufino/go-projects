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

    maxZoom:20,
    minZoom:1,
    zoomControlOptions: {
      position: google.maps.ControlPosition.TOP_LEFT,
      style: google.maps.ZoomControlStyle.DEFAULT
    },
    cluster: {
      options:{
        averageCenter: true,
        styles:[{
          url: 'components/images/m2.svg',
          height: 56,
          width: 55,
          textSize:19
          },
          {
          url: 'components/images/m1.svg',
          height: 56,
          width: 55,
          textColor: '#F00',
          textSize:18

          
      }]
    }
  }


  };

}(window,google,window.Mapster || (window.Mapster={})))