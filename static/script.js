
$("a.show-detail").on("click", function(){
    console.log($(this).data('link'));
    
    $("#divItem #qid").html($(this).data('id'));
    $("#divMask").height($(window).height());
    $("#divMask").show();
    $("#divItem").css({ left: ($(window).width() - $("#divItem").width())/2+'px'});
    $("#divItem").show();
    // $("#divItem #iframe-page").attr('src',"/item?url="+$(this).data('link'));
    $.get( "/item?url="+$(this).data('link'), function(data) {
        $("#divItem #qid").html(data);
      })
    .fail(function() {
    //   alert( "error" );
    })
    .always(function() {
    //   alert( "finished" );
    });

    
});


$("#divMask").on("click", function(){
    $("#divItem").hide();
    $("#divMask").hide();
    $("#divItem #qid").html('');
});

$("a.item-close").on("click", function(){
    $("#divItem").hide();
    $("#divMask").hide();
    $("#divItem #qid").html('');
});