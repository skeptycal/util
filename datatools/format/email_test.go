package format

import (
	"reflect"
	"testing"
)

func TestGetDomainNames(t *testing.T) {

    const (
        sample = `news@e.lenscrafters.com OR digital@masterdynamic.com OR marketingcampaigns@sproutsocial.com OR cb2@mail.cb2.com OR gnc@rewards.gnc.com OR bathandbodyworks@e2.bathandbodyworks.com OR email@promotions.overstock.com OR email@e.academy.com OR info@bloglovin.com OR info@yourstory.com OR offers@wish.com OR newsletters@communications.eplans.com OR cnbcprimetime@response.cnbc.com OR avis@e.avis.com OR targetnews@em.target.com OR offers@your.offers.dominos.com OR autozone@e.autozone.com OR notifications@texastribune.org OR floorplans@communications.homeplans.com OR nytimes@e.newyorktimes.com OR hello@newsela.com OR GameStop@em.gamestop.com OR email@usa.uniqlo.com`

        sampleSolved = `e.lenscrafters.com OR masterdynamic.com OR sproutsocial.com OR mail.cb2.com OR rewards.gnc.com OR e2.bathandbodyworks.com OR promotions.overstock.com OR e.academy.com OR bloglovin.com OR yourstory.com OR wish.com OR communications.eplans.com OR response.cnbc.com OR e.avis.com OR em.target.com OR your.offers.dominos.com OR e.autozone.com OR texastribune.org OR communications.homeplans.com OR e.newyorktimes.com OR newsela.com OR em.gamestop.com OR usa.uniqlo.com`
    )

	type args struct {
		list string
	}
	tests := []struct {
		name       string
		args       args
		wantRetval string
	}{
        // TODO: Add test cases.
        {"sample",args{sample}, sampleSolved},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetval := GetDomainNames(tt.args.list); !reflect.DeepEqual(gotRetval, tt.wantRetval) {
				t.Errorf("GetDomainNames() = %v, want %v", gotRetval, tt.wantRetval)
			}
		})
	}
}
