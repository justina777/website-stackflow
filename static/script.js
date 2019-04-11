
$( document ).ready(function() {
    $(".gradient").height($(document).height());
});

$("a.show-detail").on("click", function(){
    $("#divMask").height($(document).height());
    $("#divMask").show();
    $("#divItem").css({ left: ($(window).width() - $("#divItem").width())/2+'px'});
    $("#divItem").show();
    
    $.get( "/item?url="+$(this).data('link'), function(data) {
        $("#divItem #qloading").hide();
        $("#divItem #qid").html(data);
      })
    .fail(function() {
      alert( "Can't fetch page from stack overflow. The url is "+ $(this).data('link'));
    })
    .always(function() {
    //   alert( "finished" );
    });
});

$("#divMask").on("click", function(){
    $("#divItem").hide();
    $("#divMask").hide();
    $("#divItem #qloading").show();
    $("#divItem #qid").html('');
});

$("a.item-close").on("click", function(){
    $("#divItem").hide();
    $("#divMask").hide();
    $("#divItem #qloading").show();
    $("#divItem #qid").html('');
});