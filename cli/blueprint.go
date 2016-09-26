package cli

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/hdc-cli/client/blueprints"
	"github.com/hortonworks/hdc-cli/models"
	"github.com/urfave/cli"
	"strconv"
	"strings"
	"sync"
	"time"
)

var BlueprintHeader []string = []string{"Cluster Type", "HDP Version"}

var BlueprintMap map[string]string

func init() {
	BlueprintMap = make(map[string]string)
	BlueprintMap["hdp-etl-edw"] = "EDW-ETL: Apache Hive 1.2.1, Apache Spark 1.6"
	BlueprintMap["hdp-edw-analytics"] = "EDW-Analytics: Apache Hive 2 LLAP, Apache Zeppelin"
	BlueprintMap["hdp-data-science"] = "Data Science: Apache Spark 1.6, Zeppelin"
	BlueprintMap["hdp-etl-edw-tp"] = "EDW-ETL: Apache Spark 2.0-preview, Apache Hive 2.0"
	BlueprintMap["hdp-etl-edw-spark2"] = "EDW-ETL: Apache Spark 2.0-preview, Apache Hive 2.0"
	BlueprintMap["shared-services"] = "Enterprise Services: Apache Atlas, Apache Ranger"
	BlueprintMap["hdp25-etl-edw-shared"] = "Enterprise ETL-EDW: Apache Hive 1.2.1"

	BlueprintMap["EDW-ETL: Apache Spark 2.0-preview, Apache Hive 2.0"] = "EDW-ETL: Apache Spark 2.0-preview"
}

type Blueprint struct {
	ClusterType string `json:"ClusterType" yaml:"ClusterType"`
	HDPVersion  string `json:"HDPVersion" yaml:"HDPVersion"`
}

func (b *Blueprint) DataAsStringArray() []string {
	return []string{b.ClusterType, b.HDPVersion}
}

func ListBlueprints(c *cli.Context) error {
	oAuth2Client := NewOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))

	// make the request to get all items
	respBlueprints, err := oAuth2Client.Cloudbreak.Blueprints.GetPublics(&blueprints.GetPublicsParams{})
	if err != nil {
		logErrorAndExit(ListBlueprints, err.Error())
	}

	var tableRows []Row
	for _, blueprint := range respBlueprints.Payload {
		// this is a workaround, needs to be hidden, by not storing them as public
		if !strings.HasPrefix(blueprint.Name, "b") {
			row := &Blueprint{ClusterType: getFancyBlueprintName(blueprint), HDPVersion: blueprint.AmbariBlueprint.Blueprint.StackVersion}
			tableRows = append(tableRows, row)
		}
	}
	output := Output{Format: c.String(FlOutput.Name)}
	output.WriteList(BlueprintHeader, tableRows)

	return nil
}

func (c *Cloudbreak) GetBlueprintByName(name string) *models.BlueprintResponse {
	defer timeTrack(time.Now(), "get blueprint by name")
	log.Infof("[GetBlueprintByName] get blueprint by name: %s", name)

	resp, err := c.Cloudbreak.Blueprints.GetPrivate(&blueprints.GetPrivateParams{Name: getRealBlueprintName(name)})
	if err != nil {
		logErrorAndExit(c.GetBlueprintByName, err.Error())
	}

	id, _ := strconv.Atoi(*resp.Payload.ID)
	id64 := int64(id)
	log.Infof("[GetBlueprintByName] '%s' blueprint id: %d", name, id64)
	return resp.Payload
}

func (c *Cloudbreak) CreateBlueprint(skeleton ClusterSkeleton, blueprint *models.BlueprintResponse, channel chan int64, wg *sync.WaitGroup) {
	defer wg.Done()

	if len(skeleton.Configurations) == 0 {
		log.Info("[CreateBlueprint] there are no custom configurations, use the default blueprint")
		channel <- getBlueprintId(blueprint)
		return
	}

	defer timeTrack(time.Now(), "create blueprint")

	blueprintName := "b" + strconv.FormatInt(time.Now().UnixNano(), 10)

	bpRequest := models.BlueprintRequest{
		Name:            blueprintName,
		AmbariBlueprint: blueprint.AmbariBlueprint,
		Properties:      skeleton.Configurations,
	}

	resp, err := c.Cloudbreak.Blueprints.PostPublic(&blueprints.PostPublicParams{Body: &bpRequest})

	if err != nil {
		logErrorAndExit(c.CreateBlueprint, err.Error())
	}

	log.Infof("[CreateBlueprint] blueprint created, id: %d", resp.Payload.ID)
	channel <- resp.Payload.ID
}

func getBlueprintId(blueprint *models.BlueprintResponse) int64 {
	id, _ := strconv.Atoi(*blueprint.ID)
	return int64(id)
}

func getFancyBlueprintName(blueprint *models.BlueprintResponse) string {
	var name string
	ambariBpName := *blueprint.AmbariBlueprint.Blueprint.Name
	if len(ambariBpName) > 0 {
		fancyName := BlueprintMap[ambariBpName]
		if len(fancyName) > 0 {
			name = fancyName
		} else {
			name = blueprint.Name
		}
	} else {
		name = blueprint.Name
	}
	return name
}

func getRealBlueprintName(name string) string {
	realName := BlueprintMap[name]
	if len(realName) > 0 {
		return realName
	}
	return name
}
