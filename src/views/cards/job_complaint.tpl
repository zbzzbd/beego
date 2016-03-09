 
        <div class="ui card p10 full">
            {{if eq (printf "%d" .EditStatus) "1" }}
            <h4 class="ui teal header">修改投诉</h4>
            {{else if eq (printf "%d" .EditStatus) "0"}}
             <h4 class="ui teal header">客户投诉</h4>
            {{end}}

            <div class="ui grid p10">
                <div class="row">
                    <div class="nine wide column">
                        <label>业务担当：</label>
                         {{.CreateUser.Name}}
                    </div>
                    <div class="five wide column">
                        <label>客诉事项：</label>
                        {{if eq (printf  "%s" .Type)  "1" }}
                        <label>延时</label>
                        {{else if eq (printf "%s" .Type) "2"}}
                        <label>逻辑错误</label>
                        {{else if eq (printf "%s" .Type) "3"}}
                        <label>其他重大错误</label>
                        {{end}}
                    </div>
                </div> 
                <div class="row">
                    <div class="nine wide column">
                        <label>客诉描述：</label>
                        {{.Complain}}
                    </div> 
                </div> 
            </div>
            <div>
             <label class="ui teal right ribbon label">投诉时间：{{.Created}}</label>
            </div>
        </div> 
 

