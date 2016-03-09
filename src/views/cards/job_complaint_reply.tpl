       <div class="ui card p10 full">
            <h4 class="ui teal header">投诉回复</h4>
            <div class="ui grid p10">
                <div class="row">
                    <div class="nine wide column">
                        <label>回复人：</label>
                         {{.User.Name}}
                    </div> 
                </div>

                <div class="row">
                    <div class="nine wide column">
                        <label>回复内容：</label>
                        {{.Message}}
                    </div> 
                </div> 

                <div class="row">
                    <div class="sixteen wide column">
                            <label>作业附件：</label>
                            {{range .Files}}
                            <a href="{{.Url}}" class="pr20" download="{{.Name}}" target="_blank">{{.Name}}</a>
                            {{end}}
                    </div>
                </div>

            </div>
        <div>
             <label class="ui teal right ribbon label">回复投诉时间：{{.Created}}</label>
        </div>
        </div> 
 