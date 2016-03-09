<div class="ui right floated pagination menu">
    <a class="icon item" href="{{.paginator.PageLinkPrev}}">
        <i class="left chevron icon"></i>
    </a>
    {{range $index, $page := .paginator.Pages}}
    {{if $.paginator.IsActive .}}
    <a class="item active" href="{{$.paginator.PageLink $page}}">{{$page}}</a>
    {{else}}
    <a class="item" href="{{$.paginator.PageLink $page}}">{{$page}}</a>
    {{end}}
    {{end}}
    <a class="icon item" href="{{.paginator.PageLinkNext}}">
        <i class="right chevron icon"></i>
    </a>
</div>
