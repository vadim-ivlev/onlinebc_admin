
            // подравниваем высоты textarea
            function adjustTextarea() {
                $("textarea").each(function (textarea) {
                    $(this).height($(this)[0].scrollHeight);
                });
            }
    
            // Accordion
            function accordionize(){
                $( "ol" ).accordion({
                    header: "li > div.comment",
                    active: false,
                    heightStyle: "content",
                    collapsible: true,
                    animate: 50,
                    classes: {
                        "ui-accordion-header": "header",
                        "ui-accordion-header-collapsed": "header-collapsed",
                        "ui-accordion-content": "content"
                    }
                })
    
                
            }
    
            adjustTextarea()
            accordionize()
            // setTimeout('accordionize()',3000)



