import {
  ElButton, ElForm, ElFormItem, ElInput, ElMessage, ElContainer, ElAside, ElMain, ElHeader, ElMenu, ElSubmenu,
  ElMenuItem, ElIcon, ElBreadcrumb, ElBreadcrumbItem, ElCard, ElRow, ElCol, ElTable, ElTableColumn, ElSwitch,
  ElTooltip, ElPagination, ElDialog, ElMessageBox
} from 'element-plus'

export default (app) => {
  app.use(ElButton)
  app.use(ElForm)
  app.use(ElFormItem)
  app.use(ElInput)
  app.use(ElRow)
  app.use(ElContainer)
  app.use(ElAside)
  app.use(ElMain)
  app.use(ElHeader)
  app.use(ElMenu)
  app.use(ElSubmenu)
  app.use(ElMenuItem)
  app.use(ElIcon)
  app.use(ElBreadcrumb)
  app.use(ElBreadcrumbItem)
  app.use(ElCard)
  app.use(ElCol)
  app.use(ElTable)
  app.use(ElTableColumn)
  app.use(ElSwitch)
  app.use(ElTooltip)
  app.use(ElPagination)
  app.use(ElDialog)

  app.config.globalProperties.$message = ElMessage
  app.config.globalProperties.$confirm = ElMessageBox.confirm
}
