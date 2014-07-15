package cvss

import "testing"

func TestFromVector(t *testing.T) {
	var _, err = Parse("AV:N/AC:H/I:N/A:N")
	if err != nil {
		t.Errorf("New from vector failed: %v", err)
	}
}

func TestFromVectorFails(t *testing.T) {
	var _, err = Parse("This is not valid")
	if err == nil {
		t.Errorf("New from vector should have failed but did not")
	}
}

func TestCve_2002_0392_HighEnv(t *testing.T) {
	var c, err = Parse("AV:N/AC:L/Au:N/C:N/I:N/A:C")
	if err != nil {
		t.Errorf("New from vector failed: %v", err)
	}
	// Temp
	c = append(c, Exploitability_Functional)
	c = append(c, RemediationLevel_OfficialFix)
	c = append(c, ReportConfidence_Confirmed)
	// Env
	c = append(c, CollateralDamagePotential_High)
	c = append(c, TargetDistribution_High)
	c = append(c, ConfidentialityRequirement_High)
	c = append(c, IntegrityRequirement_High)
	c = append(c, AvailabilityRequirement_High)

	s := c.Score()
	expected := Score{7.8, 6.4, 9.2}
	if s != expected {
		t.Errorf("Score differ, expected %v got %v", expected, s)
	}
}

func TestCve_2002_0392_LowEnv(t *testing.T) {
	var c, err = Parse("AV:N/AC:L/Au:N/C:N/I:N/A:C")
	if err != nil {
		t.Errorf("New from vector failed: %v", err)
	}
	// Temp
	c = append(c, Exploitability_Functional)
	c = append(c, RemediationLevel_OfficialFix)
	c = append(c, ReportConfidence_Confirmed)
	// Env
	c = append(c, CollateralDamagePotential_None)
	c = append(c, TargetDistribution_None)
	c = append(c, ConfidentialityRequirement_Medium)
	c = append(c, IntegrityRequirement_Medium)
	c = append(c, AvailabilityRequirement_High)

	s := c.Score()
	expected := Score{7.8, 6.4, 0.0}
	if s != expected {
		t.Errorf("Score differ, expected %v got %v\n%s", expected, s, c.collateralDamagePotential())
	}
}

func TestCve_2003_0818_HighEnv(t *testing.T) {
	var c, err = Parse("AV:N/AC:L/Au:N/C:C/I:C/A:C")
	if err != nil {
		t.Errorf("New from vector failed: %v", err)
	}
	// Temp
	c = append(c, Exploitability_Functional)
	c = append(c, RemediationLevel_OfficialFix)
	c = append(c, ReportConfidence_Confirmed)
	// Env
	c = append(c, CollateralDamagePotential_High)
	c = append(c, TargetDistribution_High)
	c = append(c, ConfidentialityRequirement_Medium)
	c = append(c, IntegrityRequirement_Medium)
	c = append(c, AvailabilityRequirement_Low)

	s := c.Score()
	expected := Score{10.0, 8.3, 9.0}
	if s != expected {
		t.Errorf("Score differ, expected %v got %v", expected, s)
	}
}

func TestCve_2003_0818_LowEnv(t *testing.T) {
	var c, err = Parse("AV:N/AC:L/Au:N/C:C/I:C/A:C")
	if err != nil {
		t.Errorf("New from vector failed: %v", err)
	}
	// Temp
	c = append(c, Exploitability_Functional)
	c = append(c, RemediationLevel_OfficialFix)
	c = append(c, ReportConfidence_Confirmed)
	// Env
	c = append(c, CollateralDamagePotential_None)
	c = append(c, TargetDistribution_None)
	c = append(c, ConfidentialityRequirement_Medium)
	c = append(c, IntegrityRequirement_Medium)
	c = append(c, AvailabilityRequirement_Low)

	s := c.Score()
	expected := Score{10.0, 8.3, 0.0}
	if s != expected {
		t.Errorf("Score differ, expected %v got %v", expected, s)
	}
}

func TestCve_2003_0062_HighEnv(t *testing.T) {
	var c, err = Parse("AV:L/AC:H/Au:N/C:C/I:C/A:C")
	if err != nil {
		t.Errorf("New from vector failed: %v", err)
	}
	// Temp
	c = append(c, Exploitability_ProofOfConcept)
	c = append(c, RemediationLevel_OfficialFix)
	c = append(c, ReportConfidence_Confirmed)
	// Env
	c = append(c, CollateralDamagePotential_High)
	c = append(c, TargetDistribution_High)
	c = append(c, ConfidentialityRequirement_High)
	c = append(c, IntegrityRequirement_High)
	c = append(c, AvailabilityRequirement_High)

	s := c.Score()
	expected := Score{6.2, 4.9, 7.5}
	if s != expected {
		t.Errorf("Score differ, expected %v got %v", expected, s)
	}
}

func TestCve_2003_0062_LowEnv(t *testing.T) {
	var c, err = Parse("AV:L/AC:H/Au:N/C:C/I:C/A:C")
	if err != nil {
		t.Errorf("New from vector failed: %v", err)
	}
	// Temp
	c = append(c, Exploitability_ProofOfConcept)
	c = append(c, RemediationLevel_OfficialFix)
	c = append(c, ReportConfidence_Confirmed)
	// Env
	c = append(c, CollateralDamagePotential_None)
	c = append(c, TargetDistribution_None)
	c = append(c, ConfidentialityRequirement_Medium)
	c = append(c, IntegrityRequirement_Medium)
	c = append(c, AvailabilityRequirement_Medium)

	s := c.Score()
	expected := Score{6.2, 4.9, 0.0}
	if s != expected {
		t.Errorf("Score differ, expected %v got %v", expected, s)
	}
}

func TestToVector(t *testing.T) {
	originalVec := "AV:L/AC:H/Au:N/C:C/I:C/A:C"
	// We always sort the output to be predictable.
	expected := "A:C/AC:H/AV:L/Au:N/C:C/I:C"
	var c, err = Parse(originalVec)
	if err != nil {
		t.Errorf("New from vector failed: %v", err)
	}
	v := c.ToStringVector()
	if v != expected {
		t.Errorf("ToVector result differs, expected %v got %v", expected, v)
	}
}

func TestToString(t *testing.T) {
	s := new(CVSS).String()
	if len(s) == 0 {
		t.Errorf("Should have produced a string of lenght >0")
	}
}
