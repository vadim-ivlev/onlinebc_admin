
// подравниваем высоты textarea
$("textarea").each(function (textarea) {
    $(this).height($(this)[0].scrollHeight);
});


var b64array =[]


function inputFileChangeHandler(event) {
    clear()
    var files = event.target.files
    for (var file of files) {
        var reader = new FileReader()
        reader.fileName = file.name
        reader.addEventListener("load", function (e) {
            var res = e.target.result
            b64 = res.substr(res.search("base64,") + "base64,".length)
            b64array.push({ fileName: e.target.fileName, base64: b64 })
            addImage(res)
        }, false)
        reader.readAsDataURL(file)
    }
}


function createQueries(b64array) {
    var query = ""
    b64array.forEach(function (f, i) {
        query += createOneQuery(i, f.fileName, f.base64)
    });
    document.querySelector('#form0 textarea[name="query"]').value = "mutation {\n"+query+"\n}\n"
}


function addImage(src) {
    $('#img_container').append(`<img style="max-width: 150px;" src="${src}">`)
}


function createOneQuery(i, fileName, b64string) {
    return `
    new${i}: createMedium( post_id: ${ document.querySelector('#formA input[name="post_id"]').getAttribute('value') }, 
        source: "${ document.querySelector('#formA input[name="source"]').getAttribute('value') }", 
        filename: "${fileName}",
        base64: "${b64string}"
    ) 
    {   
        id 
        post_id  
        source 
        thumb  
        uri  
    }
    `
}

function clear(){
    $('#img_container').html('')
    b64array=[]
}

