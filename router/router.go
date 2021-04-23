package router

import (
	"github.com/gogf/gf/frame/g"
)

func init() {
	public := g.Server().Group("")
	{
		NewClusterRouter(public).Init()
		NewHostRouter(public).Init()
		NewApptemplateRouter(public).Init()
		NewClusterAppRouter(public).Init()
		NewClusterNodesRouter(public).Init()
		NewClusterLogRouter(public).Init()
		NewClusterSecretRouter(public).Init()
		NewClusterSpecRouter(public).Init()
		NewClusterStatusRouter(public).Init()
		NewCredentialRouter(public).Init()
		NewRegionRouter(public).Init()
		NewSystemSettingRouter(public).Init()
		NewTenantMemberRouter(public).Init()
		NewTenantResourceRouter(public).Init()
		NewTenantRouter(public).Init()
		NewUserClusterResourceLimitRouter(public).Init()
		NewUserRouter(public).Init()
		NewVolumeRouter(public).Init()
		NewZoneRouter(public).Init()
	}

}
