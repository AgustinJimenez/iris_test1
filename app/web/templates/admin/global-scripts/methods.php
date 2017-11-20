<script type='text/javascript'>
    function template_prevent_link( link )
    {
        prevent = true;

        if( link.slice(-1) == "#" )
            prevent =  false;

        else if( link.indexOf('admin/logout') >= 0 )
            prevent =  false;

        return prevent;
    }

    function template_back_button_was_clicked()
    {
        left_arrow_was_clicked();
    }

    function show_loading()
    {
        TEMPLATE_LOADING_MODAL.modal("show");
    }

    function hide_loading()
    {
        TEMPLATE_LOADING_MODAL.modal("hide");
    }

    function template_get( url )
    {
        if( template_use_loading )
            show_loading();

        $.get( url, function( data ) 
        {
            template_change_content_container( data, url, true);
        });
    }

    

    function template_post( url, params )
    {
        //console.log(url, params);

        $.ajax
        ({
            type: "POST",
            url: url,
            data: params,
            dataType: "json"
        })
        .always(function(data) 
        {
            //console.log( "finished", data );
            template_change_content_container( data.responseText, null );
        });
        
    }

    function template_get_last_url_visited()
    {
        return TEMPLATE_VISITED_URLS.slice(-1)[0];
    }

    function template_drop_last_url()
    {
        TEMPLATE_VISITED_URLS.pop();
    }

    function left_arrow_was_clicked()
    {
        if( TEMPLATE_VISITED_URLS.length > 1 )
        {
            template_drop_last_url();
            template_get( template_get_last_url_visited() );
        }
            
        //console.log(TEMPLATE_VISITED_URLS, 'last= '+last_url);

        /* 
        
        if( TEMPLATE_VISITED_URLS.length > 1 )
        TEMPLATE_VISITED_URLS.forEach(function(item, index)
        {
        
            if( TEMPLATE_VISITED_URLS[index + 1] )
                if( item == TEMPLATE_VISITED_URLS[index + 1])
                    TEMPLATE_VISITED_URLS.splice(index + 1, 1);
        });

        */

        
    }

    function template_add_visited_url(url)
    {
        if(url != template_get_last_url_visited() )
            TEMPLATE_VISITED_URLS.push( url );
    }


    function f5_was_clicked()
    {

        var last_url = template_get_last_url_visited();
        if(last_url)
            template_get( last_url );

        refresh_icon_animate();
    }

    function refresh_icon_animate()
    {
        //console.log( TEMPLATE_REFRESH_BUTTON.children("i").attr("class") );
        //TEMPLATE_REFRESH_BUTTON.children("i").addClass('fa-spin').delay(4000).removeClass('fa-spin');
        TEMPLATE_REFRESH_BUTTON.children("i").addClass('fa-spin');
        setTimeout(function()
        {
            TEMPLATE_REFRESH_BUTTON.children("i").removeClass('fa-spin');
        },1000);
        
    }


    function template_change_content_container( html, url, from_link = false)
    {
        TEMPLATE_CONTENT_CONTAINER.stop().hide().html( html ).fadeIn('slow');
        if(url)
            template_add_visited_url(url);

        set_input_autocomplete_off();
        if( template_use_loading )
            hide_loading();
    }

    function set_input_autocomplete_off()
    {
        $("input").attr("autocomplete", 'off');
    }

</script>