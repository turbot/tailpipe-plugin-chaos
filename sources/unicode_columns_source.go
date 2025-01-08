package sources

import (
	"context"
	"fmt"
	"time"

	"github.com/turbot/tailpipe-plugin-chaos/config"
	"github.com/turbot/tailpipe-plugin-chaos/rows"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
	"github.com/turbot/tailpipe-plugin-sdk/schema"
	"github.com/turbot/tailpipe-plugin-sdk/types"
)

const UnicodeColumnsSourceIdentifier = "chaos_unicode_columns"

func init() {
	row_source.RegisterRowSource[*UnicodeColumnsSource]()
}

type UnicodeColumnsSource struct {
	row_source.RowSourceImpl[*UnicodeColumnsSourceConfig, *config.ChaosConnection]
}

func (s *UnicodeColumnsSource) Collect(ctx context.Context) error {
	sn := UnicodeColumnsSourceIdentifier
	ts := time.Now().Add(-time.Second * 23)
	se := &schema.SourceEnrichment{
		CommonFields: schema.CommonFields{
			TpSourceName: &sn,
			TpSourceType: sn,
		},
	}
	for i := 0; i < s.Config.RowCount; i++ {
		rowData := s.populateRowData(i, ts)
		row := &types.RowData{Data: rowData, SourceEnrichment: se}
		if err := s.OnRow(ctx, row, nil); err != nil {
			return fmt.Errorf("error processing row: %w", err)
		}
	}
	return nil
}

func (s *UnicodeColumnsSource) Identifier() string {
	return UnicodeColumnsSourceIdentifier
}

func (s *UnicodeColumnsSource) populateRowData(i int, ts time.Time) *rows.UnicodeColumns {
	longString := "\u003cp\u003e\u003cstrong\u003ePost Incident Review (PIR) – Failure to move or delete resources\u003c/strong\u003e\u003cspan\u003e\u0026nbsp;\u003c/span\u003e\u003c/p\u003e\n\u003cp\u003e\u003cstrong\u003eWhat happened?\u003c/strong\u003e\u003cspan\u003e\u0026nbsp;\u003c/span\u003e\u003c/p\u003e\n\u003cp\u003eBetween 19:49 UTC on 25 November 2024 and 03:45 UTC on 29 November 2024, customers may have experienced failures when attempting to move resources between resource groups and/or between subscriptions. Additionally, some resource groups may have temporarily become stuck in a locked state after attempting move operations. This issue was limited to move operation attempts, previously communicated delete operations were not directly affected by this event. The vast majority of impact was alleviated by 22:56 UTC on 27 November when the fix was rolled out to all regions, minus some low traffic areas. All remaining regions received the fix by 03:45 UTC on 29 November.\u0026nbsp;\u003c/p\u003e\n\u003cp\u003e\u0026nbsp;\u003c/p\u003e\n\u003cp\u003e\u003cstrong\u003eWhat went wrong and why?\u003c/strong\u003e\u003cspan\u003e\u0026nbsp;\u003c/span\u003e\u003c/p\u003e\n\u003cp\u003e\u003cspan\u003eAzure Resource Manager (ARM) is a management layer that handles requests to Azure services, while Azure Advisor is a cloud consultant that provides recommendations to improve posture on security, reliability, operational excellence, performance and cost. Azure Advisor had subscribed to notifications from ARM when resources are moved. When move operations are attempted the resource ID changes, which the Advisor RP would need to track the new ID. The Advisor RP subscribed to the ARM notification must have been acknowledged for the operation to be completed. This enabled Advisor to move any resources near real time (NRT) to improve relevancy of recommendations.\u0026nbsp;\u003c/span\u003e\u003c/p\u003e\n\u003cp\u003e\u003cspan\u003eDuring the event, the notifications between the Advisor Resource Provider (RP) layer and ARM experienced issues. The operations were successfully submitted by ARM and notifications sent to Advisor; however, Advisor did not acknowledge these notifications as expected. When the Advisor RP did not acknowledge the request, resource groups where the operations were performed would become locked for some time, until the entire move operation failed.\u0026nbsp;\u003c/span\u003e\u003c/p\u003e\n\u003cp\u003e\u003cspan\u003ePart of the Advisor stack listening to ARM was erroring out, which uncovered a configuration with ARM and the Advisor RP that made certain notifications a required action to complete operation requests. So, when those attempts were not acknowledged by Advisor, the required actions were not met according to ARM, resulting in timeouts or errors.\u0026nbsp;\u003c/span\u003e\u003c/p\u003e\n\u003cp\u003e\u003cspan\u003eNo changes were made to either service that resulted in this event. Additionally, the ARM notification requirements were not effectively adhered to initially, but the configuration requiring certain notifications has since been removed, preventing any additional changes to avoid future occurrences related to this class of issue.\u0026nbsp;\u003c/span\u003e\u003c/p\u003e\n\u003cp\u003e\u003cspan\u003e\u0026nbsp;\u003c/span\u003e\u003c/p\u003e\n\u003cp\u003e\u003cstrong\u003eHow did we respond?\u003c/strong\u003e\u003cspan\u003e\u0026nbsp;\u003c/span\u003e\u003c/p\u003e\n\u003cul\u003e\u003cli\u003e\u003cspan\u003e19:49 UTC on 25 Nov 2024 - Customer impact began.\u0026nbsp;\u003c/span\u003e\u003c/li\u003e\u003cli\u003e\u003cspan\u003e19:50 UTC on 25 Nov 2024 – Advisor monitoring detected increased timeouts between ARM and the Advisor RP layer. The alert triggered did not indicate customer impact at the time. We began investigating potential errors linked to timeouts.\u0026nbsp;\u003c/span\u003e\u003c/li\u003e\u003cli\u003e\u003cspan\u003e21:53 UTC on 25 Nov 2024 – Identified an unhealthy component application potentially contributing to the timeouts and engaged the component team to help investigate. It was verified the upstream component was not contributing to the timeouts.\u0026nbsp;\u003c/span\u003e\u003c/li\u003e\u003cli\u003e\u003cspan\u003e00:16 UTC on 26 Nov 2024 – Began restarting Compute resources underlying the unhealthy component to restore the service. Telemetry indicated timeouts had stopped, the service recovered and steadily remained healthy. We continued investigating for contributing factors.\u0026nbsp;\u003c/span\u003e\u003c/li\u003e\u003cli\u003e\u003cspan\u003e12:59 UTC on 26 Nov 2024 - Received an external customer report alerting us of impact. Shortly after, we began several troubleshooting workstreams and recovery tests including additional resource restarts, scaling out the service, pausing, restarting, and redeploying components.\u0026nbsp;\u003c/span\u003e\u003c/li\u003e\u003cli\u003e\u003cspan\u003e09:05 UTC on 27 Nov 2024 – Identified the configuration contributing to the Advisor stack component errors. A new configuration was tested, removing the notification requirement from ARM. The test was successful, and ARM logs validated no issues moving resources with the latest configuration. A formal change was made to the ARM configuration to prepare for deployment across production.\u0026nbsp;\u003c/span\u003e\u003c/li\u003e\u003cli\u003e\u003cspan\u003e16:33 UTC on 27 Nov 2024 – Began rolling out the configuration update. Some tooling limitations resulted in a slower, more granular rollout; however, standard safe deployment practices were followed to limit risk.\u0026nbsp;\u003c/span\u003e\u003c/li\u003e\u003cli\u003e\u003cspan\u003e22:56 UTC on 27 Nov 2024 - Completed rolling out the fix to high and medium traffic regions.\u0026nbsp;\u003c/span\u003e\u003c/li\u003e\u003cli\u003e\u003cspan\u003e03:45 UTC on 29 November – Completed rolling out the fix to all regions, confirmed failures returned to normal and, impact mitigated.\u0026nbsp;\u003c/span\u003e\u003c/li\u003e\u003c/ul\u003e\n\u003cp\u003e\u0026nbsp;\u003c/p\u003e\n\u003cp\u003e\u003cstrong\u003eHow are we making incidents like this less likely or less impactful?\u003c/strong\u003e\u003cspan\u003e\u0026nbsp;\u003c/span\u003e\u003c/p\u003e\n\u003cp\u003e\u003c/p\u003e\u003cul\u003e\u003cli\u003e\u003cspan\u003eRemoved ARM requirement of certain notifications for operation execution to prevent this type of issue in the future (Completed).\u0026nbsp;\u003c/span\u003e\u003c/li\u003e\u003cli\u003e\u003cspan\u003eEnhance monitoring of key resource performance metrics to help detect component issues sooner, prior to customer impact (Estimated completion: Jan 2025).\u0026nbsp;\u003c/span\u003e\u003c/li\u003e\u003cli\u003e\u003cspan\u003eImprove system diagnostics and logging to help identify and resolve potential issues sooner (Estimated completion: Feb 2025).\u0026nbsp;\u003c/span\u003e\u003c/li\u003e\u003cli\u003e\u003cspan\u003eExpand troubleshooting guides with more detailed dependency investigation guidance for faster diagnosis and identification of customer impact (Estimated completion: Jan 2025).\u0026nbsp;\u003c/span\u003e\u003c/li\u003e\u003cli\u003e\u003cspan\u003eVerify Azure Advisor monitoring and logging precision for other dependencies to help ensure a comprehensive signal across the service (Estimated completion: Feb 2025)\u0026nbsp;\u003c/span\u003e\u003c/li\u003e\u003c/ul\u003e\u003cp\u003e\u0026nbsp;\u003c/p\u003e\u003cp\u003e\u003cstrong\u003eHow can customers make incidents like this less impactful?\u003c/strong\u003e\u003cspan\u003e\u0026nbsp;\u003c/span\u003e\u003c/p\u003e\u003cul\u003e\u003cli\u003eGenerally, consider evaluating the reliability of your applications using guidance from the Azure Well-Architected Framework and its interactive Well-Architected Review:\u0026nbsp;\u003ca href=\"https://aka.ms/AzPIR/WAF\" target=\"_blank\"\u003ehttps://aka.ms/AzPIR/WAF\u003c/a\u003e\u0026nbsp;\u003c/li\u003e\u003cli\u003eThe impact times above represent the full incident duration, so are not specific to any individual customer. Actual impact to service availability varied between customers and resources – for guidance on implementing monitoring to understand granular impact:\u0026nbsp;\u003ca href=\"https://aka.ms/AzPIR/Monitoring\" target=\"_blank\"\u003ehttps://aka.ms/AzPIR/Monitoring\u003c/a\u003e\u0026nbsp;\u003c/li\u003e\u003cli\u003eFinally, consider ensuring that the right people in your organization will be notified about any future service issues – by configuring Azure Service Health alerts. These can trigger emails, SMS, push notifications, webhooks, and more:\u0026nbsp;\u003ca href=\"https://aka.ms/AzPIR/Alerts\" target=\"_blank\"\u003ehttps://aka.ms/AzPIR/Alerts\u003c/a\u003e\u0026nbsp;\u003c/li\u003e\u003c/ul\u003e\u003cp\u003e\u003cbr\u003e\u003c/p\u003e\u003cp\u003e\u003cstrong\u003eHow can we make our incident communications more useful?\u003c/strong\u003e\u003cspan\u003e\u0026nbsp;\u003c/span\u003e\u0026nbsp;\u003c/p\u003e\u003cp\u003eYou can rate this PIR and provide any feedback using our quick 3-question survey: \u003ca href=\"https://aka.ms/AzPIR/SMX1-FR8\" target=\"_blank\"\u003ehttps://aka.ms/AzPIR/SMX1-FR8\u003c/a\u003e\u003c/p\u003e\u003cp\u003e\u003c/p\u003e"
	jsonObject := make(map[string]interface{})
	jsonObject["someKey"] = "https://someurl.com/somepath\u0026somequery"
	jsonObject["someOtherKey"] = "https://someurl.com/somepath\\u0026somequery"

	subJsonObject := make(map[string]interface{})
	subJsonObject["subKey"] = "https://someurl.com/somepath\u0026somequery"
	subJsonObject["longValue"] = longString
	subJsonObject["properties"] = longString

	jsonObject["subObject"] = subJsonObject

	return &rows.UnicodeColumns{
		Title:          fmt.Sprintf("Title %d", i),
		TopLevelString: longString,
		JsonObject:     &jsonObject,
		Timestamp:      ts,
	}
}
