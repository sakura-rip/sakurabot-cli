# sakurabot cli

This Cli tool is used for the management of SakuraBot.

SakuraBot is the brand name for LINE BOT.

It's operated by [Sakura](https://github.com/sakura-rip).  
You can't use this Cli tool because it depends on many secret projects.

The purpose of this tool is to automate the management of buyers, servers, etc., which was previously done manually.


related organisations:  
[bot-sakura](https://github.com/bot-sakura),
[line-org](https://github.com/line-org),
[sakura-biz](https://github.com/sakura-biz),


Big thanks to: [frugal](https://github.com/Workiva/frugal)
played a big part in speeding up my LINE BOT.




### Cobra children command file template for goland livetemplate
```html
<template name="cobrachile" value="&#10;import (&#10;&#9;&quot;github.com/spf13/cobra&quot;&#10;&#9;&quot;github.com/spf13/pflag&quot;&#10;)&#10;&#10;var $CMD_NAME_LOWER$Param = new($CMD_NAME_LOWER$Params)&#10;&#10;// $CMD_NAME$Command base command for &quot;$BASE_CMD$ $CMD_NAME_LOWER$&quot;&#10;func $CMD_NAME$Command() *cobra.Command {&#10;&#9;cmd := &amp;cobra.Command{&#10;&#9;&#9;Use:   &quot;$CMD_NAME_LOWER$&quot;,&#10;&#9;&#9;Short: &quot;$CMD_NAME_LOWER$ $BASE_CMD$&quot;,&#10;&#9;&#9;Run:   run$CMD_NAME$Command,&#10;&#9;}&#10;&#9;cmd.Flags().AddFlagSet($CMD_NAME_LOWER$Param.getFlagSet())&#10;&#9;return cmd&#10;}&#10;&#10;// $CMD_NAME_LOWER$Params add commands parameter&#10;type $CMD_NAME_LOWER$Params struct {&#10;&#9;$END$&#9;&#10;}&#10;&#10;// getFlagSet returns the flagSet for $CMD_NAME_LOWER$Params&#10;func (p *$CMD_NAME_LOWER$Params) getFlagSet() *pflag.FlagSet {&#10;&#9;fs := new(pflag.FlagSet)&#10;&#10;&#9;return fs&#10;}&#10;&#10;// validate validate parameters&#10;func (p *$CMD_NAME_LOWER$Params) validate() error {&#10;&#9;return validator.New().Struct(p)&#10;}&#10;&#10;// processParams process parameters variable&#10;func (p *$CMD_NAME_LOWER$Params) processParams(args []string) {&#10;&#9;if err := p.validate(); err != nil {&#10;&#9;&#9;utils.Logger.Fatal().Err(err).Msg(&quot;&quot;)&#10;&#9;}&#10;}&#10;&#10;// processInteract process interact parameter initializer&#10;func (p *$CMD_NAME_LOWER$Params) processInteract(args []string) {&#10;&#10;}&#10;&#10;// run$CMD_NAME$Command execute &quot;$BASE_CMD$ $CMD_NAME_LOWER$&quot; command&#10;func run$CMD_NAME$Command(cmd *cobra.Command, args []string) {&#10;&#9;if cmd.Flags().NFlag() == 0 {&#10;&#9;&#9;$CMD_NAME_LOWER$Param.processInteract(args)&#10;&#9;}&#10;&#9;$CMD_NAME_LOWER$Param.processParams(args)&#10;&#10;}&#10;" description="" toReformat="false" toShortenFQNames="true">
  <variable name="CMD_NAME" expression="" defaultValue="" alwaysStopAt="true" />
  <variable name="CMD_NAME_LOWER" expression="camelCase(CMD_NAME)" defaultValue="" alwaysStopAt="true" />
  <variable name="BASE_CMD" expression="" defaultValue="" alwaysStopAt="true" />
  <context>
    <option name="GO" value="true" />
  </context>
</template>
```