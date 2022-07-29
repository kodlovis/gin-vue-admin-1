<template>
    <div>
        <div class="gva-search-box">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item label="部门">
                    <el-select v-model="setInfo.department_value" class="m-2" placeholder="Select" @change="getCycle">
                        <el-option v-for="item in department_options" :key="item.value" :label="item.label"
                            :value="item.value" />
                    </el-select>
                </el-form-item>
                <el-form-item label="角色">
                    <el-select v-model="setInfo.role_value" class="m-2" placeholder="Select" @change="getCycle">
                        <el-option v-for="item in role_options" :key="item.value" :label="item.label"
                            :value="item.value" />
                    </el-select>
                </el-form-item>
                <el-form-item label="填报日期">
                    <el-date-picker v-model="setInfo.up_date" type="date" placeholder="Pick a day" :size="'default'"
                        @change="getCycle" />
                </el-form-item>
                <el-form-item>
                    <!-- <el-tag  :type="active == true ? 'success' : 'danger'">{{active == true ? '已确认' : '未确认'}}</el-tag> -->
                </el-form-item>



                <el-form-item>
                    <el-popover placement="bottom" width="240" v-model="copy_visible">
                        <!-- <el-form :inline="true" :model="searchInfo" class="demo-form-inline"> -->
                        <el-form-item>
                            <el-date-picker v-model="down_date" type="date" placeholder="Pick a day"
                                :size="'default'" />
                        </el-form-item>
                        <!-- </el-form> -->

                        <div style="text-align: right; margin: 0">
                            <el-button size="mini" type="text" @click="copy_visible = false">取消</el-button>
                            <el-button type="primary" size="mini" @click="copyCycle">确定</el-button>
                        </div>
                        <el-button slot="reference" icon="copyDocument" :type="is_copy_allow == 0 ? 'primary' : 'info'"
                            :disabled="is_copy_allow == 0 ? false : true">复制</el-button>

                    </el-popover>


                </el-form-item>
                <el-form-item>

                    <el-button icon="Top" :type="active == true ? 'success' : 'danger'" :disabled="active"
                        @click="openConfirm">{{ active == true ? '已确认' : '未确认' }}</el-button>
                </el-form-item>


            </el-form>
        </div>
        <div class="gva-table-box">
            <el-calendar>
                <template #dateCell="{ data }">
                    <div class="confirm" :class="addClass(data.day)">
                        <ul>
                            <li style="padding-top:6px;">
                                <p :class="data.isSelected ? 'is-selected' : ''" @click="openDraw(data)">
                                    {{ data.day.split("-").slice(1).join("-") }}
                                    {{ data.isSelected ? "✔️" : "" }}
                                </p>
                            </li>
                            <li v-if="(new Date(data.day).getTime() > new Date().getTime()) && in_out[data.day] != undefined"
                                style="padding-top:6px;" :class="addClassText(data.day, '收入')">
                                收入：{{ in_out[data.day] == undefined ? '0' : in_out[data.day][0] }}万</li>
                            <li v-if="(new Date(data.day).getTime() > new Date().getTime()) && in_out[data.day] != undefined"
                                style="padding-top:6px ;" :class="addClassText(data.day, '支出')">
                                支出：{{ in_out[data.day] == undefined ? '0' : in_out[data.day][1] }}万</li>
                        </ul>
                    </div>
                </template>
            </el-calendar>
        </div>
        <!-- 添加数据 -->
        <el-drawer :visible.sync="drawer2" :direction="'rtl'" :before-close="cancelClick" :show-close="false"
            :size="'80%'" custom-class="demo-drawer" ref="drawer">
            <template #title>
                <h4>
                    预算数据录入&nbsp;&nbsp;{{ choseDate }}&nbsp;&nbsp;&nbsp;{{
                            setInfo.department_value
                    }}&nbsp;&nbsp;&nbsp;{{ setInfo.role_value }}
                </h4>

                <el-button @click="cancelClick">关闭</el-button>
                <el-button style="margin-right:3%" :type="is_allow == 0 ? 'primary' : 'info'"
                    :disabled="is_allow == 0 ? false : true" @click="confirmClick">保存</el-button>


            </template>
            <div class="demo-drawer__content">
                <div>
                    <el-form :label-position="'right'" label-width="150px" :model="formBudget">
                        <el-row justify="center">
                            <el-col :span="12" v-show="setInfo.role_value == '商务员'">
                                <el-card class="box-card">
                                    <template #header>
                                        <div class="card-header">
                                            <span>支出-现货采购</span>
                                        </div>
                                    </template>
                                    <el-form-item label="采购货款" style="width: 90%">
                                        <el-input v-model="formBudget.cghk" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.cghkbz" type="textarea"
                                            placeholder="所有日常业务采购付款（流量业务除外），包括保证金、预付款、货款等" />
                                    </el-form-item>

                                    <el-form-item label="买货解质押" style="width: 90%">
                                        <el-input v-model="formBudget.mhjzy" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.mhjzybz" type="textarea"
                                            placeholder="签合同销售（含期转现业务、交易所结算业务）仓单，提前从期货账户解质押做的入金" />
                                    </el-form-item>

                                    <el-form-item label="流量业务" style="width: 90%">
                                        <el-input v-model="formBudget.outllyw" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.outllywbz" type="textarea"
                                            placeholder="所有流量业务采购付款（日常业务除外），包括保证金、预付款、货款等" />
                                    </el-form-item>

                                    <el-form-item label="协助集团融资借款" style="width: 90%">
                                        <el-input v-model="formBudget.xzjtrzjk" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.xzjtrzjkbz" type="textarea"
                                            placeholder="一般贸易融资项下协助集团融资给集团的借款支出" />
                                    </el-form-item>

                                    <el-form-item label="到期付汇" style="width: 90%">
                                        <el-input v-model="formBudget.dqfh" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.dqfhbz" type="textarea"
                                            placeholder="信用证项下到期付汇支出" />
                                    </el-form-item>

                                    <el-form-item label="开证保证金" style="width: 90%">
                                        <el-input v-model="formBudget.kzbzj" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.kzbzjbz" type="textarea"
                                            placeholder="开证业务保证金支出" />
                                    </el-form-item>

                                    <el-form-item label="税金" style="width: 90%">
                                        <el-input v-model="formBudget.sj" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.sjbz" type="textarea"
                                            placeholder="一般贸易进口关税、增值税支付" />
                                    </el-form-item>
                                </el-card>
                            </el-col>
                            <el-col :span="12" v-show="setInfo.role_value == '商务员'">
                                <el-card class="box-card">
                                    <template #header>
                                        <div class="card-header">
                                            <span>收入-现货销售</span>
                                        </div>
                                    </template>
                                    <el-form-item label="正常销售" style="width: 90%">
                                        <el-input v-model="formBudget.zcxs" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.zcxsbz" type="textarea"
                                            placeholder="所有日常业务销售回款（流量业务除外），包括保证金、预收款、货款等" />
                                    </el-form-item>
                                    <el-form-item label="买货质押" style="width: 90%">
                                        <el-input v-model="formBudget.mhzy" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.mhzybz" type="textarea"
                                            placeholder="签合同采购（含期转现业务、交易所结算业务）得到的仓单，在期货账户质押出金" />
                                    </el-form-item>
                                    <el-form-item label="流量业务" style="width: 90%">
                                        <el-input v-model="formBudget.inllyw" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.inllywbz" type="textarea"
                                            placeholder="所有流量业务销售回款（日常业务除外），包括保证金、预收款、货款等" />
                                    </el-form-item>

                                    <el-form-item label="协助集团借款回款" style="width: 90%">
                                        <el-input v-model="formBudget.xzjtjkhk" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.xzjtjkhkbz" type="textarea"
                                            placeholder="一般贸易项下协助集团融资借款还回，包含本金和利息" />
                                    </el-form-item>

                                    <el-form-item label="开证保证金退回" style="width: 90%">
                                        <el-input v-model="formBudget.kzbzjth" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.kzbzjthbz" type="textarea"
                                            placeholder="开证保证金退回" />
                                    </el-form-item>
                                </el-card>
                            </el-col>
                            <el-col :span="12" v-show="setInfo.role_value == '交易员'">
                                <el-card class="box-card">
                                    <template #header>
                                        <div class="card-header">
                                            <span>支出-期货入金</span>
                                        </div>
                                    </template>
                                    <el-form-item label="开仓" style="width: 90%">
                                        <el-input v-model="formBudget.kc" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.kcbz" type="textarea"
                                            placeholder="新建立头寸、计划建头寸而做的入金" />
                                    </el-form-item>

                                    <el-form-item label="银行解质押" style="width: 90%">
                                        <el-input v-model="formBudget.yhjzy" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.yhjzybz" type="textarea"
                                            placeholder="通过上期仓单交易平台向银行赎回仓单的付款本息" />
                                    </el-form-item>

                                    <el-form-item label="卖交割解质押" style="width: 90%">
                                        <el-input v-model="formBudget.mjgjzy" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.mjgjzybz" type="textarea"
                                            placeholder="为配合卖交割而解质押仓单所做的入金" />
                                    </el-form-item>

                                    <el-form-item label="买交割" style="width: 90%">
                                        <el-input v-model="formBudget.outmjg" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.outmjgbz" type="textarea" placeholder="支付买交割货款" />
                                    </el-form-item>

                                    <el-form-item label="提保入金" style="width: 90%">
                                        <el-input v-model="formBudget.tbrj" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.tbrjbz" type="textarea"
                                            placeholder="因节假日、涨跌停板导致的临时提高保证金" />
                                    </el-form-item>
                                </el-card>
                            </el-col>
                            <el-col :span="12" v-show="setInfo.role_value == '交易员'">
                                <el-card class="box-card">
                                    <template #header>
                                        <div class="card-header">
                                            <span>收入-期货出金</span>
                                        </div>
                                    </template>
                                    <el-form-item label="平仓" style="width: 90%">
                                        <el-input v-model="formBudget.pc" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.pcbz" type="textarea" placeholder="头寸了结后出金" />
                                    </el-form-item>

                                    <el-form-item label="银行质押" style="width: 90%">
                                        <el-input v-model="formBudget.yhzy" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.yhzybz" type="textarea"
                                            placeholder="仓单质押上期仓单交易平台，所收到的款项" />
                                    </el-form-item>

                                    <el-form-item label="买交割质押" style="width: 90%">
                                        <el-input v-model="formBudget.mjgzy" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.mjgzybz" type="textarea"
                                            placeholder="通过交割获得的仓单，质押期货账户后出金" />
                                    </el-form-item>

                                    <el-form-item label="卖交割" style="width: 90%">
                                        <el-input v-model="formBudget.inmjg" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.inmjgbz" type="textarea" placeholder="卖交割收到的货款" />
                                    </el-form-item>

                                    <el-form-item label="释放保证金" style="width: 90%">
                                        <el-input v-model="formBudget.sfbzj" />
                                    </el-form-item>
                                    <el-form-item label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.sfbzjbz" type="textarea"
                                            placeholder="因节假日、涨跌停板导致的临时提高保证金，恢复正常后出金" />
                                    </el-form-item>
                                </el-card>


                            </el-col>
                            <el-col :span="12" v-show="setInfo.role_value == '资金助理'">
                                <el-card class="box-card">
                                    <template #header>
                                        <div class="card-header">
                                            <span>支出</span>
                                        </div>
                                    </template>

                                    <el-form-item v-show="this.is_all_show('ywck')" label="业务存款" style="width: 90%">
                                        <el-input v-model="formBudget.ywck" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('ywck')" label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.ywckbz" type="textarea"
                                            placeholder="融资、回报、套利业务存款" />
                                    </el-form-item>

                                    <el-form-item v-show="this.is_all_show('ywhk')" label="业务还款" style="width: 90%">
                                        <el-input v-model="formBudget.ywhk" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('ywhk')" label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.ywhkbz" type="textarea"
                                            placeholder="套利业务到期周is_all_show转还款" />
                                    </el-form-item>

                                    <el-form-item v-show="this.is_all_show('cwsxf')" label="财务手续费" style="width: 90%">
                                        <el-input v-model="formBudget.cwsxf" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('cwsxf')" label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.cwsxfbz" type="textarea"
                                            placeholder="开票费、开证费、贴现费等财务费用" />
                                    </el-form-item>

                                    <el-form-item v-show="this.is_all_show('rzhk')" label="融资还款" style="width: 90%">
                                        <el-input v-model="formBudget.rzhk" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('rzhk')" label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.rzhkbz" type="textarea" placeholder="融资业务到期还款" />
                                    </el-form-item>

                                    <el-form-item v-show="this.is_all_show('jkcc')" label="借款拆出" style="width: 90%">
                                        <el-input v-model="formBudget.jkcc" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('jkcc')" label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.jkccbz" type="textarea" placeholder="借款给其他公司" />
                                    </el-form-item>

                                    <el-form-item v-show="this.is_all_show('outcjhk')" label="拆借还款" style="width: 90%">
                                        <el-input v-model="formBudget.outcjhk" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('outcjhk')" label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.outcjhkbz" type="textarea"
                                            placeholder="向其他公司借款后还款" />
                                    </el-form-item>

                                    <el-form-item v-show="this.is_all_show('rzck')" label="融资存款" style="width: 90%">
                                        <el-input v-model="formBudget.rzck" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('rzck')" label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.rzckbz" type="textarea" placeholder="融资业务保证金存入" />
                                    </el-form-item>

                                    <el-form-item v-show="this.is_all_show('hbck')" label="回报存款" style="width: 90%">
                                        <el-input v-model="formBudget.hbck" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('hbck')" label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.hbckbz" type="textarea" placeholder="回报业务保证金存入" />
                                    </el-form-item>

                                    <el-form-item v-show="this.is_all_show('jjfz')" label="工资奖金和房租" style="width: 90%">
                                        <el-input v-model="formBudget.jjfz" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('jjfz')" label="备注" style="width:c 90%">
                                        <el-input v-model="formBudget.jjfzbz" type="textarea"
                                            placeholder="月工资奖金、季度房租，人事提供数据" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('outqt')" label="其他" style="width: 90%">
                                        <el-input v-model="formBudget.outqt" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('outqt')" label="备注" style="width:c 90%">
                                        <el-input v-model="formBudget.outqtbz" type="textarea" placeholder="临时其他类别" />
                                    </el-form-item>
                                </el-card>
                            </el-col>
                            <el-col :span="12" v-show="setInfo.role_value == '资金助理'">
                                <el-card class="box-card">
                                    <template #header>
                                        <div class="card-header">
                                            <span>收入</span>
                                        </div>
                                    </template>

                                    <el-form-item v-show="this.is_all_show('lxsr')" label="利息收入" style="width: 90%">
                                        <el-input v-model="formBudget.lxsr" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('lxsr')" label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.lxsrbz" type="textarea"
                                            placeholder="融资、回报、套利业务到期后收入" />
                                    </el-form-item>

                                    <el-form-item v-show="this.is_all_show('zxf')" label="咨询费" style="width: 90%">
                                        <el-input v-model="formBudget.zxf" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('zxf')" label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.zxfbz" type="textarea" placeholder="咨询费收入" />
                                    </el-form-item>

                                    <el-form-item v-show="this.is_all_show('bzjjsyth')" label="保证金及收益退回"
                                        style="width: 90%">
                                        <el-input v-model="formBudget.bzjjsyth" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('bzjjsyth')" label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.bzjjsythbz" type="textarea"
                                            placeholder="融资、回报业务到期本金及收益" />
                                    </el-form-item>

                                    <el-form-item v-show="this.is_all_show('jkcr')" label="借款拆入" style="width: 90%">
                                        <el-input v-model="formBudget.jkcr" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('jkcr')" label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.jkcrbz" type="textarea" placeholder="向其他公司借款" />
                                    </el-form-item>

                                    <el-form-item v-show="this.is_all_show('incjhk')" label="拆借还款" style="width: 90%">
                                        <el-input v-model="formBudget.incjhk" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('incjhk')" label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.incjhkbz" type="textarea"
                                            placeholder="其他公司借款后还款" />
                                    </el-form-item>

                                    <el-form-item v-show="this.is_all_show('rztx')" label="融资贴现" style="width: 90%">
                                        <el-input v-model="formBudget.rztx" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('rztx')" label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.rztxbz" type="textarea"
                                            placeholder="融资业务银票、国内证贴现资金" />
                                    </el-form-item>

                                    <el-form-item v-show="this.is_all_show('hbtx')" label="回报贴现" style="width: 90%">
                                        <el-input v-model="formBudget.hbtx" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('hbtx')" label="备注" style="width: 90%">
                                        <el-input v-model="formBudget.hbtxbz" type="textarea"
                                            placeholder="回报业务银票、国内证贴现资金" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('inqt')" label="其他" style="width: 90%">
                                        <el-input v-model="formBudget.inqt" />
                                    </el-form-item>
                                    <el-form-item v-show="this.is_all_show('inqt')" label="备注" style="width:c 90%">
                                        <el-input v-model="formBudget.inqtbz" type="textarea" placeholder="临时其他类别" />
                                    </el-form-item>

                                </el-card>
                            </el-col>
                        </el-row>
                    </el-form>
                </div>

            </div>

        </el-drawer>
        <!-- 数据确认 -->
        <el-dialog :visible.sync="confirm_dialog">
            <template #title>
                <h4>
                    数据确认&nbsp;&nbsp;{{ setInfo.up_date | formatDate }}&nbsp;&nbsp;&nbsp;{{
                            setInfo.department_value
                    }}&nbsp;&nbsp;&nbsp;{{ setInfo.role_value }}
                </h4>
            </template>
            <div class="gva-table-box">
                <el-table :data="tableData[0].detail" border style="width: 100%;" class="table-info bgf" width="1080"
                    :header-cell-style="rowClass" cell-class-name="lm-cell" empty-text="暂无数据" tooltip-effect="dark">
                    <template v-for="(item, index) in tableData">
                        <el-table-column :key="index" :label="item.title" v-if="item.title == '类型'" fixed="left"
                            width="255">
                            <template slot-scope="scope">
                                <div>{{ item.detail[scope.$index] }}</div>
                            </template>
                        </el-table-column>
                    </template>

                    <template v-for="(item, index) in tableData">
                        <el-table-column :key="index" :label="item.title" v-if="item.title != '类型'">
                            <template slot-scope="scope">
                                <div>{{ item.detail[scope.$index] }}</div>
                            </template>
                        </el-table-column>
                    </template>
                </el-table>
            </div>
            <template #footer>
                <el-button @click="confirm_dialog = false">取消</el-button>
                <el-button style="margin-right:3%" :type="is_allow == 0 ? 'primary' : 'info'"
                    :disabled="is_allow == 0 ? false : true" @click="upConfirm">确认</el-button>
            </template>
        </el-dialog>
    </div>
</template>

<script>

import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
import {
    clickBudgetData,
    createBudgetDatas,
    getOneConfirmData,
    upConfirmData,
    getCycleData,
} from "@/api/internalSystem/fund/fund_budget_data"; //  此处请自行替换地址
import {
    getFundRole
} from "@/api/internalSystem/fund/remainingSum";
import {
    getConfirmSetting,
    getCycleConfirmDate,
} from "@/api/internalSystem/fund/fund_confirm_date";  //  此处请自行替换地址

export default {
    name: "BudgetData",
    mixins: [infoList],

    data() {
        return {
            down_date: new Date().getTime() - 86400000,
            copy_visible: false,
            choseDate: new Date(),
            date_list: [],
            department_options: [

            ],
            is_show_department: ['杉贸本部', '总经理一室', '智维本部'],
            is_show: {
                "杉贸本部": ['rzhk', 'jkcc', 'outcjhk', 'rzck', 'hbck', 'cwsxf', 'bzjjsyth', 'jkcr', 'incjhk', 'rztx', 'hbtx', 'outqt', 'inqt'],
                "总经理一室": ['ywck', 'ywhk', 'cwsxf', 'hbtx', 'lxsr', 'zxf', 'outqt', 'inqt'],
                "智维本部": ['jkcc', 'outcjhk', 'jkcr', 'incjhk', 'outqt', 'inqt'],
            },
            role_options: [

            ],
            formBudget: {
                cghk: "",
                cghkbz: "",
                mhjzy: "",
                mhjzybz: "",
                outllyw: "",
                outllywbz: "",
                xzjtrzjk: "",
                xzjtrzjkbz: "",
                dqfh: "",
                dqfhbz: "",
                kzbzj: "",
                kzbzjbz: "",
                sj: "",
                sjbz: "",
                zcxs: "",
                zcxsbz: "",
                mhzy: "",
                mhzybz: "",
                inllyw: "",
                inllywbz: "",
                xzjtjkhk: "",
                xzjtjkhkbz: "",
                kzbzjth: "",
                kzbzjthbz: "",
                kc: "",
                kcbz: "",
                yhjzy: "",
                yhjzybz: "",
                mjgjzy: "",
                mjgjzybz: "",
                outmjg: "",
                outmjgbz: "",
                tbrj: "",
                tbrjbz: "",
                pc: "",
                pcbz: "",
                yhzy: "",
                yhzybz: "",
                mjgzy: "",
                mjgzybz: "",
                inmjg: "",
                inmjgbz: "",
                sfbzj: "",
                sfbzjbz: "",
                jjfz: "",
                jjfzbz: "",
                ywck: "",
                ywckbz: "",
                ywhk: "",
                ywhkbz: "",
                cwsxf: "",
                cwsxfbz: "",
                rzhk: "",
                rzhkbz: "",
                jkcc: "",
                jkccbz: "",
                outcjhk: "",
                outcjhkbz: "",
                rzck: "",
                rzckbz: "",
                hbck: "",
                hbckbz: "",
                lxsr: "",
                lxsrbz: "",
                zxf: "",
                zxfbz: "",
                bzjjsyth: "",
                bzjjsythbz: "",
                jkcr: "",
                jkcrbz: "",
                incjhk: "",
                incjhkbz: "",
                rztx: "",
                rztxbz: "",
                hbtx: "",
                hbtxbz: "",
                outqt: "",
                outqtbz: "",
                inqt: "",
                inqtbz: "",
            },
            // 获取local中的身份
            setInfo: {
                department_value: "",
                role_value: "",
                up_date: new Date().getTime(),
            },
            drawer2: false,
            confirm_dialog: false,
            all_type: {
                cghk: "支出-现货采购-采购货款",
                cghkbz: "",
                mhjzy: "支出-现货采购-卖货解质押",
                mhjzybz: "",
                outllyw: "支出-现货采购-流量业务",
                outllywbz: "",
                xzjtrzjk: "支出-现货采购-协助集团融资借款",
                xzjtrzjkbz: "",
                dqfh: "支出-现货采购-到期付汇",
                dqfhbz: "",
                kzbzj: "支出-现货采购-开证保证金",
                kzbzjbz: "",
                sj: "支出-现货采购-税金",
                sjbz: "",
                zcxs: "收入-现货销售-正常销售",
                zcxsbz: "",
                mhzy: "收入-现货销售-买货质押",
                mhzybz: "",
                inllyw: "收入-现货销售-流量业务",
                inllywbz: "",
                xzjtjkhk: "收入-现货销售-协助集团融资借款回款",
                xzjtjkhkbz: "",
                kzbzjth: "收入-现货销售-开保证金及收益退回",
                kzbzjthbz: "",
                kc: "支出-期货入金-开仓",
                kcbz: "",
                yhjzy: "支出-期货入金-银行解质押",
                yhjzybz: "",
                mjgjzy: "支出-期货入金-卖交割解质押",
                mjgjzybz: "",
                outmjg: "支出-期货入金-买交割",
                outmjgbz: "",
                tbrj: "支出-期货入金-提保入金",
                tbrjbz: "",
                pc: "收入-期货出金-平仓",
                pcbz: "",
                yhzy: "收入-期货出金-银行质押",
                yhzybz: "",
                mjgzy: "收入-期货出金-买交割质押",
                mjgzybz: "",
                inmjg: "收入-期货出金-卖交割",
                inmjgbz: "",
                sfbzj: "收入-期货出金-释放保证金",
                sfbzjbz: "",
                jjfz: "支出-工资奖金和房租",
                jjfzbz: "",
                ywck: "支出-业务存款",
                ywckbz: "",
                ywhk: "支出-业务还款",
                ywhkbz: "",
                cwsxf: "支出-财务手续费",
                cwsxfbz: "",
                rzhk: "支出-融资还款",
                rzhkbz: "",
                jkcc: "支出-借款拆出",
                jkccbz: "",
                outcjhk: "支出-拆借还款",
                outcjhkbz: "",
                rzck: "支出-融资存款",
                rzckbz: "",
                hbck: "支出-回报存款",
                hbckbz: "",
                lxsr: "收入-利息收入",
                lxsrbz: "",
                zxf: "收入-咨询费",
                zxfbz: "",
                bzjjsyth: "收入-保证金及收益退回",
                bzjjsythbz: "",
                jkcr: "收入-借款拆入",
                jkcrbz: "",
                incjhk: "收入-拆借还款",
                incjhkbz: "",
                rztx: "收入-融资贴现",
                rztxbz: "",
                hbtx: "收入-回报贴现",
                hbtxbz: "",
                outqt: "支出-其他",
                outqtbz: "",
                inqt: "收入-其他",
                inqtbz: "",
            },
            tableData: [
                {
                    title: "类型",
                    isshow: true,
                    detail: [

                    ]
                },
            ],
            out_type: ['cghk', 'mhjzy', 'outllyw', 'xzjtrzjk', 'dqfh', 'kzbzj', 'sj', 'kc', 'yhjzy', 'mjgjzy', 'outmjg', 'tbrj', 'jjfz', 'ywck', 'ywhk', 'cwsxf', 'rzhk', 'jkcc', 'outcjhk', 'rzck', 'hbck', 'outqt'],
            in_type: [],
            in_out: {},
            active: false,
            is_allow: 0,
            is_copy_allow: 0,

        };
    },
    filters: {
        formatDate: function (time) {
            if (time != null && time != "") {
                var date = new Date(time);
                return formatTimeToStr(date, "yyyy-MM-dd");
            } else {
                return "";
            }
        },
        formatBoolean: function (bool) {
            if (bool != null) {
                return bool ? "是" : "否";
            } else {
                return "";
            }
        },
    },
    methods: {
        //打开数据确认窗口
        async openConfirm() {
            this.confirm_dialog = true
            this.getReport()
        },
        //加载数据确认表格
        async getReport() {
            this.tableData = this.$options.data().tableData
            switch (this.setInfo.role_value) {
                case '商务员':
                    this.tableData[0].detail = [
                        "支出-现货采购-采购货款",
                        "支出-现货采购-卖货解质押",
                        "支出-现货采购-流量业务",
                        "支出-现货采购-协助集团融资借款",
                        "支出-现货采购-到期付汇",
                        "支出-现货采购-开证保证金",
                        "支出-现货采购-税金",
                        "支出小计",
                        "收入-现货销售-正常销售",
                        "收入-现货销售-买货质押",
                        "收入-现货销售-流量业务",
                        "收入-现货销售-协助集团融资借款回款",
                        "收入-现货销售-开保证金及收益退回",
                        "收入小计",
                    ]
                    break
                case '交易员':
                    this.tableData[0].detail = [
                        "支出-期货入金-开仓",
                        "支出-期货入金-银行解质押",
                        "支出-期货入金-卖交割解质押",
                        "支出-期货入金-买交割",
                        "支出-期货入金-提保入金",
                        "支出小计",
                        "收入-期货出金-平仓",
                        "收入-期货出金-银行质押",
                        "收入-期货出金-买交割质押",
                        "收入-期货出金-卖交割",
                        "收入-期货出金-释放保证金",
                        "收入小计",
                    ]
                    break
                case '资金助理':
                    switch (this.setInfo.department_value) {
                        case '杉贸本部':
                            this.tableData[0].detail = [
                                "支出-融资还款",
                                "支出-借款拆出",
                                "支出-拆借还款",
                                "支出-融资存款",
                                "支出-回报存款",
                                "支出-财务手续费",
                                "支出-其他",
                                "支出小计",

                                "收入-保证金及收益退回",
                                "收入-借款拆入",
                                "收入-拆借还款",
                                "收入-融资贴现",
                                "收入-回报贴现",
                                "收入-其他",
                                "收入小计"
                            ]
                            break
                        case '总经理一室':
                            this.tableData[0].detail = [
                                "支出-业务存款",
                                "支出-业务还款",
                                "支出-财务手续费",
                                "支出-其他",
                                "支出小计",

                                "收入-利息收入",
                                "收入-咨询费",
                                "收入-回报贴现",
                                "收入-其他",
                                "收入小计"
                            ]
                            break
                        case '智维本部':
                            this.tableData[0].detail = [
                                "支出-借款拆出",
                                "支出-拆借还款",
                                "支出-其他",
                                "支出小计",

                                "收入-借款拆入",
                                "收入-拆借还款",
                                "收入-其他",
                                "收入小计"
                            ]
                            break
                        default:
                            this.tableData[0].detail = []

                    }
                    break
                default:
                    this.tableData[0].detail = []
            }
            var start_dict = {
                'start_date': formatTimeToStr(this.setInfo.up_date, "yyyy-MM-dd")
            }
            let res = await getCycleConfirmDate(start_dict)
            var i = 1
            const datelist = []
            var end_dateStr
            var start_dateStr = formatTimeToStr(new Date(start_dict.start_date).getTime() + 86400000, "yyyy-MM-dd")
            if (res.data.list) {
                end_dateStr = formatTimeToStr(new Date(res.data.list[0].end_date).getTime() + 86400000, "yyyy-MM-dd")
            } else {
                end_dateStr = formatTimeToStr(new Date(start_dict.start_date).getTime() + 86400000 * 22, "yyyy-MM-dd")

            }
            while (start_dateStr != end_dateStr) {
                i = i + 1
                var month_data = formatTimeToStr(new Date(start_dateStr), "MM-dd")
                datelist.push(month_data)
                start_dateStr = formatTimeToStr(new Date(start_dict.start_date).getTime() + 86400000 * i, "yyyy-MM-dd")
            }
            for (var k in datelist) {
                var table_dict = {
                    title: datelist[k],
                    isshow: true,
                    detail: [],
                }
                this.tableData.push(table_dict)
            }
            let confirm_values = {
                department_code: this.setInfo.department_value,
                create_role: this.setInfo.role_value,
                create_date: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd"),
            }
            let res_data = await getCycleData(confirm_values)
            console.log(res_data)
            for (var j in res_data.data.list) {
                var list = res_data.data.list[j]
                var list_key = datelist.indexOf(formatTimeToStr(new Date(list.in_data), "MM-dd")) + 1
                if (list_key > 0) {
                    this.$set(this.tableData[list_key].detail, this.fill_confirm_data(list), list.in_value)
                }
                console.log(list_key)
            }
            switch (this.setInfo.role_value) {
                case '商务员':
                    for (var o in this.tableData) {
                        if (o > 0) {
                            this.$set(this.tableData[o].detail, 7, this.sum_details(this.tableData[o].detail)['out_sum'])
                            this.$set(this.tableData[o].detail, 13, this.sum_details(this.tableData[o].detail)['in_sum'])
                        }
                    }
                    break
                case '交易员':
                    for (var p in this.tableData) {
                        if (p > 0) {
                            this.$set(this.tableData[p].detail, 5, this.sum_details(this.tableData[p].detail)['out_sum'])
                            this.$set(this.tableData[p].detail, 11, this.sum_details(this.tableData[p].detail)['in_sum'])
                        }
                    }
                    break
                case '资金助理':
                    if (this.setInfo.department_value == '杉贸本部') {
                        for (var q in this.tableData) {
                            if (q > 0) {
                                this.$set(this.tableData[q].detail, 7, this.sum_details(this.tableData[q].detail)['out_sum'])
                                this.$set(this.tableData[q].detail, 14, this.sum_details(this.tableData[q].detail)['in_sum'])
                            }
                        }
                    }
                    if (this.setInfo.department_value == '总经理一室') {
                        for (var w in this.tableData) {
                            if (w > 0) {
                                this.$set(this.tableData[w].detail, 4, this.sum_details(this.tableData[w].detail)['out_sum'])
                                this.$set(this.tableData[w].detail, 9, this.sum_details(this.tableData[w].detail)['in_sum'])
                            }
                        }
                    }
                    if (this.setInfo.department_value == '智维本部') {
                        for (var e in this.tableData) {
                            if (e > 0) {
                                this.$set(this.tableData[e].detail, 3, this.sum_details(this.tableData[e].detail)['out_sum'])
                                this.$set(this.tableData[e].detail, 7, this.sum_details(this.tableData[e].detail)['in_sum'])
                            }
                        }
                    }

                    break
                default:
                    break
            }
            console.log(this.tableData)

        },
        sum_details(details_data) {
            console.log(details_data)

            switch (this.setInfo.role_value) {
                case '商务员':
                    var out_sum = 0
                    var in_sum = 0
                    for (var i = 0; i < 8; i++) {
                        if (details_data[i]) {
                            out_sum += details_data[i]
                        }
                    }
                    for (var j = 8; j < 13; j++) {
                        if (details_data[j]) {
                            in_sum += details_data[j]
                        }
                    }
                    var res_dict = {
                        in_sum: in_sum,
                        out_sum: out_sum
                    }
                    return res_dict
                case '交易员':
                    var out_sum1 = 0
                    var in_sum1 = 0
                    for (var i1 = 0; i1 < 5; i1++) {
                        if (details_data[i1]) {
                            out_sum1 += details_data[i1]
                        }
                    }
                    for (var j1 = 6; j1 < 11; j1++) {
                        if (details_data[j1]) {
                            in_sum1 += details_data[j1]
                        }
                    }
                    var res_dict1 = {
                        in_sum: in_sum1,
                        out_sum: out_sum1
                    }
                    return res_dict1
                case '资金助理':
                    switch (this.setInfo.department_value) {
                        case '杉贸本部':
                            var out_sum2 = 0
                            var in_sum2 = 0
                            for (var i2 = 0; i2 < 7; i2++) {
                                if (details_data[i2]) {
                                    out_sum2 += details_data[i2]
                                }
                            }
                            for (var j2 = 8; j2 < 14; j2++) {
                                if (details_data[j2]) {
                                    in_sum2 += details_data[j2]
                                }
                            }
                            var res_dict2 = {
                                in_sum: in_sum2,
                                out_sum: out_sum2
                            }
                            return res_dict2
                        case '总经理一室':
                            var out_sum3 = 0
                            var in_sum3 = 0
                            for (var i3 = 0; i3 < 4; i3++) {
                                if (details_data[i3]) {
                                    out_sum3 += details_data[i3]
                                }
                            }
                            for (var j3 = 5; j3 < 9; j3++) {
                                if (details_data[j3]) {
                                    in_sum3 += details_data[j3]
                                }
                            }
                            var res_dict3 = {
                                in_sum: in_sum3,
                                out_sum: out_sum3
                            }
                            return res_dict3
                        case '智维本部':
                            var out_sum4 = 0
                            var in_sum4 = 0
                            for (var i4 = 0; i4 < 3; i4++) {
                                if (details_data[i4]) {
                                    out_sum4 += details_data[i4]
                                }
                            }
                            for (var j4 = 4; j4 < 7; j4++) {
                                if (details_data[j4]) {
                                    in_sum4 += details_data[j4]
                                }
                            }
                            var res_dict4 = {
                                in_sum: in_sum4,
                                out_sum: out_sum4
                            }
                            return res_dict4
                        default:
                            return
                    }

                default:
                    return
            }

        },
        //填写数据确认table
        fill_confirm_data(data) {
            if (this.setInfo.role_value == '商务员') {
                switch (data.type_code) {
                    case 'cghk':
                        return 0
                    case 'mhjzy':
                        return 1
                    case 'outllyw':
                        return 2
                    case 'xzjtrzjk':
                        return 3
                    case 'dqfh':
                        return 4
                    case 'kzbzj':
                        return 5
                    case 'sj':
                        return 6
                    case 'zcxs':
                        return 8
                    case 'mhzy':
                        return 9
                    case 'inllyw':
                        return 10
                    case 'xzjtjkhk':
                        return 11
                    case 'kzbzjth':
                        return 12
                    default:
                        return
                }
            }
            if (this.setInfo.role_value == '交易员') {
                switch (data.type_code) {
                    case 'kc':
                        return 0
                    case 'yhjzy':
                        return 1
                    case 'mjgjzy':
                        return 2
                    case 'outmjg':
                        return 3
                    case 'tbrj':
                        return 4

                    case 'pc':
                        return 6
                    case 'yhzy':
                        return 7
                    case 'mjgzy':
                        return 8
                    case 'inmjg':
                        return 9
                    case 'sfbzj':
                        return 10
                    default:
                        return
                }
            }
            if (this.setInfo.role_value == '资金助理') {
                if (this.setInfo.department_value == '杉贸本部')
                    switch (data.type_code) {
                        case 'rzhk':
                            return 0
                        case 'jkcc':
                            return 1
                        case 'outcjhk':
                            return 2
                        case 'rzck':
                            return 3
                        case 'hbck':
                            return 4
                        case 'cwsxf':
                            return 5
                        case 'outqt':
                            return 6

                        case 'bzjjsyth':
                            return 8
                        case 'jkcr':
                            return 9
                        case 'incjhk':
                            return 10
                        case 'rztx':
                            return 11
                        case 'hbtx':
                            return 12
                        case 'inqt':
                            return 13
                        default:
                            return
                    }
                if (this.setInfo.department_value == '总经理一室')
                    switch (data.type_code) {
                        case 'ywck':
                            return 0
                        case 'ywhk':
                            return 1
                        case 'cwsxf':
                            return 2
                        case 'outqt':
                            return 3

                        case 'lxsr':
                            return 5
                        case 'zxf':
                            return 6
                        case 'hbtx':
                            return 7
                        case 'inqt':
                            return 8
                        default:
                            return
                    }
                if (this.setInfo.department_value == '智维本部')
                    switch (data.type_code) {
                        case 'jkcc':
                            return 0
                        case 'outcjhk':
                            return 1
                        case 'outqt':
                            return 2

                        case 'jkcr':
                            return 4
                        case 'incjhk':
                            return 5
                        case 'inqt':
                            return 6

                        default:
                            return
                    }
            }

        },

        //填报日期是否确认颜色
        async is_confirm_color() {
            let confirm_value = {
                chose_date: formatTimeToStr(this.setInfo.up_date, "yyyy-MM-dd"),
            }
            let res = await getOneConfirmData(confirm_value)
            if (res.data.list) {
                this.active = false
                for (var i in res.data.list) {
                    if (res.data.list[i]['department_code'] == this.setInfo.department_value &&
                        res.data.list[i]['create_role'] == this.setInfo.role_value &&
                        formatTimeToStr(res.data.list[i]['confirm_date'], "yyyy-MM-dd") == confirm_value.chose_date) {
                        this.active = true
                    }
                }
            } else {
                this.active = false
            }
        },
        //资金助理填写内容
        is_all_show(type) {
            if (this.is_show_department.includes(this.setInfo.department_value)) {
                return this.is_show[this.setInfo.department_value].includes(type)
            } else {
                return true
            }
        },
        //日历颜色样式
        addClass(days) {
            if (this.date_list.includes(days)) {
                return "will-date"
            } else {
                return ""
            }
        },
        addClassText(days, type) {
            if (this.date_list.includes(days)) {
                if (type == '收入') { return "text-red" } else {
                    return "text-green"
                }
            } else {
                return ""
            }
        },
        //获取日历周期
        async getCalenderDate() {
            var start_dict = {
                'start_date': formatTimeToStr(this.setInfo.up_date, "yyyy-MM-dd")
            }
            let res = await getCycleConfirmDate(start_dict)
            var i = 1
            const datelist = []
            var end_dateStr
            var start_dateStr = formatTimeToStr(new Date(start_dict.start_date).getTime() + 86400000, "yyyy-MM-dd")
            console.log(res)
            if (res.data.list) {
                end_dateStr = formatTimeToStr(new Date(res.data.list[0].end_date).getTime() + 86400000, "yyyy-MM-dd")
            } else {
                end_dateStr = formatTimeToStr(new Date(start_dict.start_date).getTime() + 86400000 * 22, "yyyy-MM-dd")

            }
            while (start_dateStr != end_dateStr) {
                i = i + 1
                datelist.push(start_dateStr)
                start_dateStr = formatTimeToStr(new Date(start_dict.start_date).getTime() + 86400000 * i, "yyyy-MM-dd")
            }
            this.date_list = datelist
        },
        //数据确认提交
        async upConfirm() {
            let confirm_value = {
                department_code: this.setInfo.department_value,
                create_role: this.setInfo.role_value,
                create_user: this.setInfo.role_value,
                confirm_date: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd"),
                is_confirm: "✔️"
            }
            let res = await upConfirmData(confirm_value)
            // console.log(res)
            if (res.code == 0) {
                this.$message({
                    type: "success",
                    message: "数据确认成功"
                })
                this.active = true
                this.confirm_dialog = false
            } else {
                this.$message({
                    type: "error",
                    message: "数据确认失败"
                })
            }
        },
        //获取周期数据
        async getCycle() {
            let confirm_value = {
                department_code: this.setInfo.department_value,
                create_role: this.setInfo.role_value,
                create_date: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd"),
            }
            this.in_out = {}
            let res = await getCycleData(confirm_value)
            //获取有录入数据的日期
            var list_date = []
            for (var i in res.data.list) {
                list_date.push(formatTimeToStr(new Date(res.data.list[i]['in_data']), "yyyy-MM-dd"))
            }
            var one_list = Array.from(new Set(list_date))
            for (var k in one_list) {
                //申明变量  
                var key = one_list[k]
                var out_data = 0
                var in_data = 0
                var day_lists = []
                for (var j in res.data.list) {
                    if (formatTimeToStr(new Date(res.data.list[j]['in_data']), "yyyy-MM-dd") == key) {
                        day_lists.push(res.data.list[j])
                    }
                }
                for (var l in day_lists) {
                    if (this.out_type.includes(day_lists[l]['type_code'])) {
                        out_data = out_data + day_lists[l]['in_value']
                    } else {
                        in_data = in_data + day_lists[l]['in_value']
                    }
                }
                this.$set(this.in_out, key, [out_data, in_data])
            }
            this.is_confirm_color()
            this.getSetting()
        },
        //赋值选中日期数据
        async copyCycle() {
            let confirm_value = {
                department_code: this.setInfo.department_value,
                create_role: this.setInfo.role_value,
                create_date: formatTimeToStr(new Date(this.down_date), "yyyy-MM-dd"),
            }
            let res = await getCycleData(confirm_value)
            var up_cycle_data = []
            for (var i in res.data.list) {
                if (new Date(res.data.list[i]['in_data']) > new Date().getTime()) {
                    var up_dict = {
                        in_data: formatTimeToStr(new Date(res.data.list[i]['in_data']), "yyyy-MM-dd"),
                        type_code: res.data.list[i]['type_code'],
                        in_value: Number(res.data.list[i]['in_value']),
                        comment: res.data.list[i]['comment'],
                        currency: "RMB",
                        department_code: res.data.list[i]['department_code'],
                        create_user: res.data.list[i]['create_user'],
                        create_role: res.data.list[i]['create_role'],
                        up_date: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd")
                    }
                    up_cycle_data.push(up_dict)
                }
            }
            let up_res = await createBudgetDatas(up_cycle_data);
            if (up_res.code == 0) {
                this.$message({
                    type: "success",
                    message: "复制成功"
                })
                this.getCycle()
                this.copy_visible = false
            }

        },

        //打开添加数据的抽屉
        async openDraw(data) {
            this.choseDate = data.day;
            let search_value = {
                department_code: this.setInfo.department_value,
                create_role: this.setInfo.role_value,
                in_data: this.choseDate,
                up_date: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd"),
            }

            let res = await clickBudgetData(search_value)

            if (res.data.list) {
                this.arrangeReturnData(res.data.list)
                this.drawer2 = true;
            } else {
                this.drawer2 = true;
            }
        },
        //关闭抽屉
        cancelClick() {
            this.drawer2 = false;
            this.formBudget = this.$options.data().formBudget
        },
        async confirmClick() {
            let up_value = this.arrangeData(this.formBudget);
            if (up_value.length != 0) {
                let res = await createBudgetDatas(up_value);
                if (res.code == 0) {
                    this.$message({
                        type: "success",
                        message: "填报成功"
                    })
                    this.getCycle()
                    this.cancelClick()
                }
            } else {
                this.$message({
                    type: "error",
                    message: "未填写数据"
                })

            }
        },

        confirmEnding(string, target) {
            return string.substr(-target.length) === target ? true : false;
        },

        //将表单数据变成一条一条得数据类型
        arrangeData(data) {
            let lists = [];
            for (var k in data) {
                if (!this.confirmEnding(k, "bz")) {
                    if (Number(data[k]) != 0) {
                        const one = {
                            in_data: this.choseDate,
                            type_code: k,
                            in_value: Number(data[k]),
                            comment: data[k + "bz"],
                            currency: "RMB",
                            department_code: this.setInfo.department_value,
                            create_user: this.setInfo.role_value,
                            create_role: this.setInfo.role_value,
                            up_date: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd")
                        };
                        lists.push(one);
                    }
                }
            }
            return lists;
        },
        arrangeReturnData(data) {

            data.forEach((item) => {
                this.formBudget[item.type_code] = item.in_value
                this.formBudget[item.type_code + "bz"] = item.comment
            })
        },
        //获取填写状态
        async getSetting() {
            let setting_value = {
                confirm_date: formatTimeToStr(new Date(this.setInfo.up_date), "yyyy-MM-dd"),
            }
            let res_setting = await getConfirmSetting(setting_value)
            this.is_copy_allow = this.is_allow = res_setting.data.list[0].is_ok
        },
        getrole() {
            let aaa = eval(JSON.parse(localStorage.getItem("fund_role"))["roles"])
            console.log(aaa)
        },
        cellStyle() {
            return "text-align:center";
        },
        rowClass() {
            return "text-align:center";
        },

        async getfunduser() {
            var dict = {
                uid: JSON.parse(localStorage.getItem("vuex"))['user']['userInfo']['ID']
            }
            let res = await getFundRole(dict)
            console.log(res)
            var role_List = eval(res.data.list[0].roles)
            role_List.forEach((item) => {
                var dict = {
                    value: item,
                    label: item,
                }
                this.role_options.push(dict)

            })
            var department_List = eval(res.data.list[0].departments)
            department_List.forEach((item) => {
                var dict = {
                    value: item,
                    label: item,
                }
                this.department_options.push(dict)
            })
            this.setInfo.department_value = department_List[0]
            this.setInfo.role_value = role_List[0]
            this.getCalenderDate()
            this.getCalenderDate()
            this.getCycle()
            this.is_confirm_color()
            this.getSetting()


        },

    },

    created() {
        this.getfunduser()

    },

}

</script>

<style>
.my-label {
    background: #E1F3D8;
}

.my-content {
    background: #FDE2E2;
}

.is-selected {
    color: #1989fa;
}

.el-row {
    margin-bottom: 20px;
}

.el-row:last-child {
    margin-bottom: 0;
}

.el-col {
    border-radius: 4px;
}

.grid-content {
    border-radius: 4px;
    min-height: 36px;
}

.confirm {
    height: 100%;
    width: 100%;
    text-align: center;
}

.will-date {
    background-color: #eee;
}

.el-calendar-table .el-calendar-day {
    padding: 2px;
}

.text-green {
    color: green;
}

.text-red {
    color: red;
}

.backgroud-green {
    background-color: green;
}

.myTable {
    border-collapse: collapse;
    margin: 0 auto;
    text-align: center;
}


.myTable td,
.myTable th {
    border: 1px solid #cad9ea;
    color: #666;
    height: 60px;
}
</style>
