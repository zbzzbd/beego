<div class="main-menu">
    <div >
        <div class="caption">
            <img src="/img/chinarun.png">
            <div style="font-size: 18px;">项目管理系统</div>
        </div>

        <div class="ui secondary vertical menu">
            <a class="item" href="/">
                <i class="home icon"></i>
                首页
            </a>

            <!--/////////////// 业务单元 ///////////////-->
            <div class='{{ShowMenu "/job" .user_info}}'>
                <a class="item" href-tag="/job">
                    <i class="file text outline icon"></i>
                    业务单元
                </a>

                <a class="sub item" href="/project/job/list">
                    <i class="circle thin icon"></i>
                    作业总单
                </a>

                <a class="sub item" href="/job/create">
                    <i class="circle thin icon"></i>
                    作业登记
                </a>
                <a class="sub item" href="/job/progress">
                    <i class="circle thin icon"></i>
                    作业进程
                </a>
                <a class="sub item" href="/job/complaint/new">
                    <i class="circle thin icon"></i>
                    投诉登记
                </a>
                <a class="sub item" href="/job/complaint/view">
                    <i class="circle thin icon"></i>
                    投诉进程
                </a>
            </div>
<!--/////////////// 客服单元 ///////////////-->
            <div class='{{ShowMenu "/customer" .user_info}}'>
                <a class="item" href-tag="/customer">
                    <i class="file text outline icon"></i>
                    客服单元
                </a>

                <a class="sub item" href="/project/job/list">
                    <i class="circle thin icon"></i>
                    作业总单
                </a>
                 <a class="sub item" href="/produce/job/claim">
                    <i class="circle thin icon"></i>
                    作业认领
                </a>
                <a class="sub item" href="/produce/job/submit">
                    <i class="circle thin icon"></i>
                    提交作业
                </a>
                <a class="sub item" href="/produce/complaint/view">
                    <i class="circle thin icon"></i>
                    客户投诉
                </a>
                
            </div>


            <!--/////////////// 订单管理 ///////////////-->
            <div class='{{ShowMenu "/project" .user_info}}'>
                <a class="item" href-tag="/project">
                    <i class="book icon"></i>
                    项目管理
                </a>
                <a class='sub item' href="/project/create">
                    <i class="circle thin icon"></i>
                    项目登记
                </a>
                <a class='sub item' href="/project/list">
                    <i class="circle thin icon"></i>
                    项目进程
                </a>
                <a class='sub item' href="/project/job/valid">
                    <i class="circle thin icon"></i>
                    作业审核
                </a>
                <a class="sub item" href="/project/job/list">
                    <i class="circle thin icon"></i>
                    作业总单
                </a>

                <a class="sub item" href="/user/list">
                    <i class="circle thin icon"></i>
                    用户列表
                </a>
                <a class="sub item" href="/require/list">
                    <i class="circle thin icon"></i>
                    作业要求列表
                </a>
                <a class="sub item" href="/project/job/dellist">
                <i class="circle thin icon"></i>
                    作业删除列表                    
                </a>

                <a class="sub item" href="/job/complaint/view">
                    <i class="circle thin icon"></i>
                    投诉进程
                </a>
            </div>

            <!--/////////////// 制作单元 ///////////////-->
            <div class='{{ShowMenu "/produce" .user_info}}'>
                <a class="item" href-tag="/produce">
                    <i class="write icon"></i>
                    制作单元
                </a>

                <a class="sub item" href="/project/job/list">
                    <i class="circle thin icon"></i>
                    作业总单
                </a>

                <a class="sub item" href="/produce/job/claim">
                    <i class="circle thin icon"></i>
                    作业认领
                </a>
                <a class="sub item" href="/produce/job/submit">
                    <i class="circle thin icon"></i>
                    提交作业
                </a>
                <a class="sub item" href="/produce/complaint/view">
                    <i class="circle thin icon"></i>
                    客户投诉
                </a>
            </div>
        </div>
    </div>
</div>
