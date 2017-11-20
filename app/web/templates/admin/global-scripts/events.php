<script type='text/javascript'>
    $("body").on('click', 'a', function(event)
    {
        var link = $(this).attr("href");

        if( template_prevent_link( link ) )
        {
            event.preventDefault();
            template_get( link );
        }
        
    });

    $("body").on('submit', 'form', function(event)
    {
        event.preventDefault();
        if( $(this).attr('method') == 'post' )
            template_post( $(this).attr('action'), $(this).serialize() );
    });

    TEMPLATE_BACK_BUTTON.click(function()
    {
        template_back_button_was_clicked();
    });

    TEMPLATE_REFRESH_BUTTON.click(function()
    {
        f5_was_clicked();
    });

    document.onkeydown = checkKey;

    function checkKey(e) 
    {
        e = e || window.event;

        if (e.keyCode == '38') 
        {
            // up arrow
        }
        else if (e.keyCode == '40') 
        {
            // down arrow
        }
        else if (e.keyCode == '37') 
        {
            //left_arrow_was_clicked();
        }
        else if (e.keyCode == '39') 
        {
        // right arrow
        }
        else if (e.keyCode == '13') 
        {
            //enter_was_clicked();
        }
        else if (e.keyCode == 116) 
        {

            e.preventDefault();
            f5_was_clicked();
        }

    }
</script>