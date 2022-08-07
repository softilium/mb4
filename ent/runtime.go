// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/rs/xid"
	"github.com/softilium/mb4/ent/divpayout"
	"github.com/softilium/mb4/ent/emission"
	"github.com/softilium/mb4/ent/emitent"
	"github.com/softilium/mb4/ent/industry"
	"github.com/softilium/mb4/ent/investaccount"
	"github.com/softilium/mb4/ent/investaccountcashflow"
	"github.com/softilium/mb4/ent/investaccountvaluation"
	"github.com/softilium/mb4/ent/quote"
	"github.com/softilium/mb4/ent/report"
	"github.com/softilium/mb4/ent/schema"
	"github.com/softilium/mb4/ent/strategy"
	"github.com/softilium/mb4/ent/strategyfactor"
	"github.com/softilium/mb4/ent/strategyfilter"
	"github.com/softilium/mb4/ent/strategyfixedticker"
	"github.com/softilium/mb4/ent/ticker"
	"github.com/softilium/mb4/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	divpayoutFields := schema.DivPayout{}.Fields()
	_ = divpayoutFields
	// divpayoutDescForYear is the schema descriptor for ForYear field.
	divpayoutDescForYear := divpayoutFields[0].Descriptor()
	// divpayout.ForYearValidator is a validator for the "ForYear" field. It is called by the builders before save.
	divpayout.ForYearValidator = divpayoutDescForYear.Validators[0].(func(int) error)
	// divpayoutDescForQuarter is the schema descriptor for ForQuarter field.
	divpayoutDescForQuarter := divpayoutFields[1].Descriptor()
	// divpayout.ForQuarterValidator is a validator for the "ForQuarter" field. It is called by the builders before save.
	divpayout.ForQuarterValidator = divpayoutDescForQuarter.Validators[0].(func(int) error)
	// divpayoutDescStatus is the schema descriptor for Status field.
	divpayoutDescStatus := divpayoutFields[3].Descriptor()
	// divpayout.StatusValidator is a validator for the "Status" field. It is called by the builders before save.
	divpayout.StatusValidator = divpayoutDescStatus.Validators[0].(func(int) error)
	emissionFields := schema.Emission{}.Fields()
	_ = emissionFields
	// emissionDescID is the schema descriptor for id field.
	emissionDescID := emissionFields[0].Descriptor()
	// emission.DefaultID holds the default value on creation for the id field.
	emission.DefaultID = emissionDescID.Default.(func() xid.ID)
	// emission.IDValidator is a validator for the "id" field. It is called by the builders before save.
	emission.IDValidator = func() func(string) error {
		validators := emissionDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	emitentFields := schema.Emitent{}.Fields()
	_ = emitentFields
	// emitentDescDescr is the schema descriptor for Descr field.
	emitentDescDescr := emitentFields[1].Descriptor()
	// emitent.DescrValidator is a validator for the "Descr" field. It is called by the builders before save.
	emitent.DescrValidator = func() func(string) error {
		validators := emitentDescDescr.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(_Descr string) error {
			for _, fn := range fns {
				if err := fn(_Descr); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// emitentDescID is the schema descriptor for id field.
	emitentDescID := emitentFields[0].Descriptor()
	// emitent.DefaultID holds the default value on creation for the id field.
	emitent.DefaultID = emitentDescID.Default.(func() xid.ID)
	// emitent.IDValidator is a validator for the "id" field. It is called by the builders before save.
	emitent.IDValidator = func() func(string) error {
		validators := emitentDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	industryFields := schema.Industry{}.Fields()
	_ = industryFields
	// industryDescDescr is the schema descriptor for Descr field.
	industryDescDescr := industryFields[1].Descriptor()
	// industry.DescrValidator is a validator for the "Descr" field. It is called by the builders before save.
	industry.DescrValidator = func() func(string) error {
		validators := industryDescDescr.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(_Descr string) error {
			for _, fn := range fns {
				if err := fn(_Descr); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// industryDescID is the schema descriptor for id field.
	industryDescID := industryFields[0].Descriptor()
	// industry.IDValidator is a validator for the "id" field. It is called by the builders before save.
	industry.IDValidator = func() func(string) error {
		validators := industryDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	investaccountFields := schema.InvestAccount{}.Fields()
	_ = investaccountFields
	// investaccountDescDescr is the schema descriptor for Descr field.
	investaccountDescDescr := investaccountFields[1].Descriptor()
	// investaccount.DescrValidator is a validator for the "Descr" field. It is called by the builders before save.
	investaccount.DescrValidator = func() func(string) error {
		validators := investaccountDescDescr.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(_Descr string) error {
			for _, fn := range fns {
				if err := fn(_Descr); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// investaccountDescID is the schema descriptor for id field.
	investaccountDescID := investaccountFields[0].Descriptor()
	// investaccount.DefaultID holds the default value on creation for the id field.
	investaccount.DefaultID = investaccountDescID.Default.(func() xid.ID)
	// investaccount.IDValidator is a validator for the "id" field. It is called by the builders before save.
	investaccount.IDValidator = func() func(string) error {
		validators := investaccountDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	investaccountcashflowFields := schema.InvestAccountCashflow{}.Fields()
	_ = investaccountcashflowFields
	// investaccountcashflowDescID is the schema descriptor for id field.
	investaccountcashflowDescID := investaccountcashflowFields[0].Descriptor()
	// investaccountcashflow.DefaultID holds the default value on creation for the id field.
	investaccountcashflow.DefaultID = investaccountcashflowDescID.Default.(func() xid.ID)
	// investaccountcashflow.IDValidator is a validator for the "id" field. It is called by the builders before save.
	investaccountcashflow.IDValidator = func() func(string) error {
		validators := investaccountcashflowDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	investaccountvaluationFields := schema.InvestAccountValuation{}.Fields()
	_ = investaccountvaluationFields
	// investaccountvaluationDescID is the schema descriptor for id field.
	investaccountvaluationDescID := investaccountvaluationFields[0].Descriptor()
	// investaccountvaluation.DefaultID holds the default value on creation for the id field.
	investaccountvaluation.DefaultID = investaccountvaluationDescID.Default.(func() xid.ID)
	// investaccountvaluation.IDValidator is a validator for the "id" field. It is called by the builders before save.
	investaccountvaluation.IDValidator = func() func(string) error {
		validators := investaccountvaluationDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	quoteFields := schema.Quote{}.Fields()
	_ = quoteFields
	// quoteDescID is the schema descriptor for id field.
	quoteDescID := quoteFields[0].Descriptor()
	// quote.DefaultID holds the default value on creation for the id field.
	quote.DefaultID = quoteDescID.Default.(func() xid.ID)
	// quote.IDValidator is a validator for the "id" field. It is called by the builders before save.
	quote.IDValidator = func() func(string) error {
		validators := quoteDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	reportFields := schema.Report{}.Fields()
	_ = reportFields
	// reportDescReportYear is the schema descriptor for ReportYear field.
	reportDescReportYear := reportFields[1].Descriptor()
	// report.ReportYearValidator is a validator for the "ReportYear" field. It is called by the builders before save.
	report.ReportYearValidator = reportDescReportYear.Validators[0].(func(int) error)
	// reportDescReportQuarter is the schema descriptor for ReportQuarter field.
	reportDescReportQuarter := reportFields[2].Descriptor()
	// report.ReportQuarterValidator is a validator for the "ReportQuarter" field. It is called by the builders before save.
	report.ReportQuarterValidator = reportDescReportQuarter.Validators[0].(func(int) error)
	// reportDescURL is the schema descriptor for Url field.
	reportDescURL := reportFields[17].Descriptor()
	// report.URLValidator is a validator for the "Url" field. It is called by the builders before save.
	report.URLValidator = reportDescURL.Validators[0].(func(string) error)
	// reportDescID is the schema descriptor for id field.
	reportDescID := reportFields[0].Descriptor()
	// report.DefaultID holds the default value on creation for the id field.
	report.DefaultID = reportDescID.Default.(func() xid.ID)
	// report.IDValidator is a validator for the "id" field. It is called by the builders before save.
	report.IDValidator = func() func(string) error {
		validators := reportDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	strategyFields := schema.Strategy{}.Fields()
	_ = strategyFields
	// strategyDescDescr is the schema descriptor for Descr field.
	strategyDescDescr := strategyFields[1].Descriptor()
	// strategy.DescrValidator is a validator for the "Descr" field. It is called by the builders before save.
	strategy.DescrValidator = func() func(string) error {
		validators := strategyDescDescr.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_Descr string) error {
			for _, fn := range fns {
				if err := fn(_Descr); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// strategyDescMaxTickers is the schema descriptor for MaxTickers field.
	strategyDescMaxTickers := strategyFields[2].Descriptor()
	// strategy.DefaultMaxTickers holds the default value on creation for the MaxTickers field.
	strategy.DefaultMaxTickers = strategyDescMaxTickers.Default.(int)
	// strategy.MaxTickersValidator is a validator for the "MaxTickers" field. It is called by the builders before save.
	strategy.MaxTickersValidator = strategyDescMaxTickers.Validators[0].(func(int) error)
	// strategyDescMaxTickersPerIndustry is the schema descriptor for MaxTickersPerIndustry field.
	strategyDescMaxTickersPerIndustry := strategyFields[3].Descriptor()
	// strategy.DefaultMaxTickersPerIndustry holds the default value on creation for the MaxTickersPerIndustry field.
	strategy.DefaultMaxTickersPerIndustry = strategyDescMaxTickersPerIndustry.Default.(int)
	// strategy.MaxTickersPerIndustryValidator is a validator for the "MaxTickersPerIndustry" field. It is called by the builders before save.
	strategy.MaxTickersPerIndustryValidator = strategyDescMaxTickersPerIndustry.Validators[0].(func(int) error)
	// strategyDescBaseIndex is the schema descriptor for BaseIndex field.
	strategyDescBaseIndex := strategyFields[4].Descriptor()
	// strategy.BaseIndexValidator is a validator for the "BaseIndex" field. It is called by the builders before save.
	strategy.BaseIndexValidator = strategyDescBaseIndex.Validators[0].(func(string) error)
	// strategyDescLastYearInventResult is the schema descriptor for LastYearInventResult field.
	strategyDescLastYearInventResult := strategyFields[5].Descriptor()
	// strategy.DefaultLastYearInventResult holds the default value on creation for the LastYearInventResult field.
	strategy.DefaultLastYearInventResult = strategyDescLastYearInventResult.Default.(float64)
	// strategyDescLastYearYield is the schema descriptor for LastYearYield field.
	strategyDescLastYearYield := strategyFields[6].Descriptor()
	// strategy.DefaultLastYearYield holds the default value on creation for the LastYearYield field.
	strategy.DefaultLastYearYield = strategyDescLastYearYield.Default.(float64)
	// strategyDescLast3YearsInvertResult is the schema descriptor for Last3YearsInvertResult field.
	strategyDescLast3YearsInvertResult := strategyFields[7].Descriptor()
	// strategy.DefaultLast3YearsInvertResult holds the default value on creation for the Last3YearsInvertResult field.
	strategy.DefaultLast3YearsInvertResult = strategyDescLast3YearsInvertResult.Default.(float64)
	// strategyDescLast3YearsYield is the schema descriptor for Last3YearsYield field.
	strategyDescLast3YearsYield := strategyFields[8].Descriptor()
	// strategy.DefaultLast3YearsYield holds the default value on creation for the Last3YearsYield field.
	strategy.DefaultLast3YearsYield = strategyDescLast3YearsYield.Default.(float64)
	// strategyDescWeekRefillAmount is the schema descriptor for WeekRefillAmount field.
	strategyDescWeekRefillAmount := strategyFields[9].Descriptor()
	// strategy.WeekRefillAmountValidator is a validator for the "WeekRefillAmount" field. It is called by the builders before save.
	strategy.WeekRefillAmountValidator = strategyDescWeekRefillAmount.Validators[0].(func(float64) error)
	// strategyDescStartAmount is the schema descriptor for StartAmount field.
	strategyDescStartAmount := strategyFields[10].Descriptor()
	// strategy.StartAmountValidator is a validator for the "StartAmount" field. It is called by the builders before save.
	strategy.StartAmountValidator = strategyDescStartAmount.Validators[0].(func(float64) error)
	// strategyDescStartSimulation is the schema descriptor for StartSimulation field.
	strategyDescStartSimulation := strategyFields[11].Descriptor()
	// strategy.DefaultStartSimulation holds the default value on creation for the StartSimulation field.
	strategy.DefaultStartSimulation = strategyDescStartSimulation.Default.(time.Time)
	// strategyDescBuyOnlyLowPrice is the schema descriptor for BuyOnlyLowPrice field.
	strategyDescBuyOnlyLowPrice := strategyFields[12].Descriptor()
	// strategy.DefaultBuyOnlyLowPrice holds the default value on creation for the BuyOnlyLowPrice field.
	strategy.DefaultBuyOnlyLowPrice = strategyDescBuyOnlyLowPrice.Default.(bool)
	// strategyDescAllowLossWhenSell is the schema descriptor for AllowLossWhenSell field.
	strategyDescAllowLossWhenSell := strategyFields[13].Descriptor()
	// strategy.DefaultAllowLossWhenSell holds the default value on creation for the AllowLossWhenSell field.
	strategy.DefaultAllowLossWhenSell = strategyDescAllowLossWhenSell.Default.(bool)
	// strategyDescSameEmitent is the schema descriptor for SameEmitent field.
	strategyDescSameEmitent := strategyFields[14].Descriptor()
	// strategy.DefaultSameEmitent holds the default value on creation for the SameEmitent field.
	strategy.DefaultSameEmitent = strategyDescSameEmitent.Default.(int)
	// strategy.SameEmitentValidator is a validator for the "SameEmitent" field. It is called by the builders before save.
	strategy.SameEmitentValidator = strategyDescSameEmitent.Validators[0].(func(int) error)
	// strategyDescID is the schema descriptor for id field.
	strategyDescID := strategyFields[0].Descriptor()
	// strategy.DefaultID holds the default value on creation for the id field.
	strategy.DefaultID = strategyDescID.Default.(func() xid.ID)
	// strategy.IDValidator is a validator for the "id" field. It is called by the builders before save.
	strategy.IDValidator = func() func(string) error {
		validators := strategyDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	strategyfactorFields := schema.StrategyFactor{}.Fields()
	_ = strategyfactorFields
	// strategyfactorDescLineNum is the schema descriptor for LineNum field.
	strategyfactorDescLineNum := strategyfactorFields[1].Descriptor()
	// strategyfactor.DefaultLineNum holds the default value on creation for the LineNum field.
	strategyfactor.DefaultLineNum = strategyfactorDescLineNum.Default.(int)
	// strategyfactor.LineNumValidator is a validator for the "LineNum" field. It is called by the builders before save.
	strategyfactor.LineNumValidator = strategyfactorDescLineNum.Validators[0].(func(int) error)
	// strategyfactorDescIsUsed is the schema descriptor for IsUsed field.
	strategyfactorDescIsUsed := strategyfactorFields[2].Descriptor()
	// strategyfactor.DefaultIsUsed holds the default value on creation for the IsUsed field.
	strategyfactor.DefaultIsUsed = strategyfactorDescIsUsed.Default.(bool)
	// strategyfactorDescInverse is the schema descriptor for Inverse field.
	strategyfactorDescInverse := strategyfactorFields[7].Descriptor()
	// strategyfactor.DefaultInverse holds the default value on creation for the Inverse field.
	strategyfactor.DefaultInverse = strategyfactorDescInverse.Default.(bool)
	// strategyfactorDescK is the schema descriptor for K field.
	strategyfactorDescK := strategyfactorFields[8].Descriptor()
	// strategyfactor.DefaultK holds the default value on creation for the K field.
	strategyfactor.DefaultK = strategyfactorDescK.Default.(float64)
	// strategyfactorDescGist is the schema descriptor for Gist field.
	strategyfactorDescGist := strategyfactorFields[9].Descriptor()
	// strategyfactor.DefaultGist holds the default value on creation for the Gist field.
	strategyfactor.DefaultGist = strategyfactorDescGist.Default.(float64)
	// strategyfactorDescID is the schema descriptor for id field.
	strategyfactorDescID := strategyfactorFields[0].Descriptor()
	// strategyfactor.DefaultID holds the default value on creation for the id field.
	strategyfactor.DefaultID = strategyfactorDescID.Default.(func() xid.ID)
	// strategyfactor.IDValidator is a validator for the "id" field. It is called by the builders before save.
	strategyfactor.IDValidator = func() func(string) error {
		validators := strategyfactorDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	strategyfilterFields := schema.StrategyFilter{}.Fields()
	_ = strategyfilterFields
	// strategyfilterDescLineNum is the schema descriptor for LineNum field.
	strategyfilterDescLineNum := strategyfilterFields[1].Descriptor()
	// strategyfilter.DefaultLineNum holds the default value on creation for the LineNum field.
	strategyfilter.DefaultLineNum = strategyfilterDescLineNum.Default.(int)
	// strategyfilter.LineNumValidator is a validator for the "LineNum" field. It is called by the builders before save.
	strategyfilter.LineNumValidator = strategyfilterDescLineNum.Validators[0].(func(int) error)
	// strategyfilterDescIsUsed is the schema descriptor for IsUsed field.
	strategyfilterDescIsUsed := strategyfilterFields[2].Descriptor()
	// strategyfilter.DefaultIsUsed holds the default value on creation for the IsUsed field.
	strategyfilter.DefaultIsUsed = strategyfilterDescIsUsed.Default.(bool)
	// strategyfilterDescLeftValueKind is the schema descriptor for LeftValueKind field.
	strategyfilterDescLeftValueKind := strategyfilterFields[3].Descriptor()
	// strategyfilter.DefaultLeftValueKind holds the default value on creation for the LeftValueKind field.
	strategyfilter.DefaultLeftValueKind = strategyfilterDescLeftValueKind.Default.(int)
	// strategyfilter.LeftValueKindValidator is a validator for the "LeftValueKind" field. It is called by the builders before save.
	strategyfilter.LeftValueKindValidator = strategyfilterDescLeftValueKind.Validators[0].(func(int) error)
	// strategyfilterDescLeftValue is the schema descriptor for LeftValue field.
	strategyfilterDescLeftValue := strategyfilterFields[4].Descriptor()
	// strategyfilter.LeftValueValidator is a validator for the "LeftValue" field. It is called by the builders before save.
	strategyfilter.LeftValueValidator = func() func(string) error {
		validators := strategyfilterDescLeftValue.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_LeftValue string) error {
			for _, fn := range fns {
				if err := fn(_LeftValue); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// strategyfilterDescOperation is the schema descriptor for Operation field.
	strategyfilterDescOperation := strategyfilterFields[6].Descriptor()
	// strategyfilter.DefaultOperation holds the default value on creation for the Operation field.
	strategyfilter.DefaultOperation = strategyfilterDescOperation.Default.(int)
	// strategyfilter.OperationValidator is a validator for the "Operation" field. It is called by the builders before save.
	strategyfilter.OperationValidator = strategyfilterDescOperation.Validators[0].(func(int) error)
	// strategyfilterDescRightValue is the schema descriptor for RightValue field.
	strategyfilterDescRightValue := strategyfilterFields[7].Descriptor()
	// strategyfilter.RightValueValidator is a validator for the "RightValue" field. It is called by the builders before save.
	strategyfilter.RightValueValidator = func() func(string) error {
		validators := strategyfilterDescRightValue.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_RightValue string) error {
			for _, fn := range fns {
				if err := fn(_RightValue); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// strategyfilterDescID is the schema descriptor for id field.
	strategyfilterDescID := strategyfilterFields[0].Descriptor()
	// strategyfilter.DefaultID holds the default value on creation for the id field.
	strategyfilter.DefaultID = strategyfilterDescID.Default.(func() xid.ID)
	// strategyfilter.IDValidator is a validator for the "id" field. It is called by the builders before save.
	strategyfilter.IDValidator = func() func(string) error {
		validators := strategyfilterDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	strategyfixedtickerFields := schema.StrategyFixedTicker{}.Fields()
	_ = strategyfixedtickerFields
	// strategyfixedtickerDescLineNum is the schema descriptor for LineNum field.
	strategyfixedtickerDescLineNum := strategyfixedtickerFields[1].Descriptor()
	// strategyfixedticker.DefaultLineNum holds the default value on creation for the LineNum field.
	strategyfixedticker.DefaultLineNum = strategyfixedtickerDescLineNum.Default.(int)
	// strategyfixedticker.LineNumValidator is a validator for the "LineNum" field. It is called by the builders before save.
	strategyfixedticker.LineNumValidator = strategyfixedtickerDescLineNum.Validators[0].(func(int) error)
	// strategyfixedtickerDescIsUsed is the schema descriptor for IsUsed field.
	strategyfixedtickerDescIsUsed := strategyfixedtickerFields[2].Descriptor()
	// strategyfixedticker.DefaultIsUsed holds the default value on creation for the IsUsed field.
	strategyfixedticker.DefaultIsUsed = strategyfixedtickerDescIsUsed.Default.(bool)
	// strategyfixedtickerDescID is the schema descriptor for id field.
	strategyfixedtickerDescID := strategyfixedtickerFields[0].Descriptor()
	// strategyfixedticker.DefaultID holds the default value on creation for the id field.
	strategyfixedticker.DefaultID = strategyfixedtickerDescID.Default.(func() xid.ID)
	// strategyfixedticker.IDValidator is a validator for the "id" field. It is called by the builders before save.
	strategyfixedticker.IDValidator = func() func(string) error {
		validators := strategyfixedtickerDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	tickerFields := schema.Ticker{}.Fields()
	_ = tickerFields
	// tickerDescDescr is the schema descriptor for Descr field.
	tickerDescDescr := tickerFields[1].Descriptor()
	// ticker.DescrValidator is a validator for the "Descr" field. It is called by the builders before save.
	ticker.DescrValidator = func() func(string) error {
		validators := tickerDescDescr.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(_Descr string) error {
			for _, fn := range fns {
				if err := fn(_Descr); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// tickerDescKind is the schema descriptor for Kind field.
	tickerDescKind := tickerFields[2].Descriptor()
	// ticker.DefaultKind holds the default value on creation for the Kind field.
	ticker.DefaultKind = tickerDescKind.Default.(int32)
	// tickerDescID is the schema descriptor for id field.
	tickerDescID := tickerFields[0].Descriptor()
	// ticker.IDValidator is a validator for the "id" field. It is called by the builders before save.
	ticker.IDValidator = func() func(string) error {
		validators := tickerDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUserName is the schema descriptor for UserName field.
	userDescUserName := userFields[1].Descriptor()
	// user.UserNameValidator is a validator for the "UserName" field. It is called by the builders before save.
	user.UserNameValidator = func() func(string) error {
		validators := userDescUserName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(_UserName string) error {
			for _, fn := range fns {
				if err := fn(_UserName); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescAdmin is the schema descriptor for Admin field.
	userDescAdmin := userFields[3].Descriptor()
	// user.DefaultAdmin holds the default value on creation for the Admin field.
	user.DefaultAdmin = userDescAdmin.Default.(bool)
	// userDescHowManyTickersOnHomepage is the schema descriptor for HowManyTickersOnHomepage field.
	userDescHowManyTickersOnHomepage := userFields[5].Descriptor()
	// user.DefaultHowManyTickersOnHomepage holds the default value on creation for the HowManyTickersOnHomepage field.
	user.DefaultHowManyTickersOnHomepage = userDescHowManyTickersOnHomepage.Default.(int)
	// user.HowManyTickersOnHomepageValidator is a validator for the "HowManyTickersOnHomepage" field. It is called by the builders before save.
	user.HowManyTickersOnHomepageValidator = userDescHowManyTickersOnHomepage.Validators[0].(func(int) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() xid.ID)
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = func() func(string) error {
		validators := userDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
