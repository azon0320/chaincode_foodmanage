package models

const ModelVersion int = 1

// -----  BEFORE AUTHENTICATE  -----
const UnAuthRegisterSeller = "reg_seller"
const UnAuthRegisterBuyer = "reg_buyer"
const UnAuthRegisterTransporter = "reg_transporter"

// -----  BEGIN SELLER  -----
const OPERATE_ADDPRODUCT = "add_prod"
const OPERATE_UPDATE_PRODUCT = "update_prod"
const OPERATE_TAKEONSELL = "sell_on"
const OPERATE_TAKEOFFSELL = "sell_off"
const OPERATE_TRANSMIT = "transmit"
const OPERATE_CANCELORDER = "cancel_order"

// -----  BEGIN BUYER  -----
const OPERATE_PURCHASE = "buy_prod"
const OPERATE_CONFIRM = "confirm"
//const OPERATE_CANCELORDER = "cancel_order"

// -----  BEGIN TRANSPORTER  -----
const OPERATE_CANCELTRANSPORT = "cancel_transport"
const OPERATE_UPDATE_TRANSPORT = "update_transport"
const OPERATE_COMPLETE_TRANSPORT = "complete_transport"

const DefaultInitialBalance = 1000