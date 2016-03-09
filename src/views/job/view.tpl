{{if eq .error ""}}
{{str2html .cards}}
{{else}}
<div class="ui p10 full red message">
    错误： {{.error}}
</div>
{{end}}

<script>
    window.onload=function(){
        $(function(){

        });
    }
</script>