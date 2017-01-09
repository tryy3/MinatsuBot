package helpplugin

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"

	"strings"

	"regexp"
	"text/tabwriter"

	"github.com/bwmarrin/discordgo"
	"github.com/tryy3/minatsubot"
)

var (
	Log minatsubot.Logger
	Reg *regexp.Regexp
)

type Help struct {
}

func (s Help) Init() {
	Log = minatsubot.Logger{
		Writers: []io.Writer{os.Stdout},
		Prefix:  "HelpPlugin",
		Level:   minatsubot.GetLoggingLevel(minatsubot.PluginAPI.GetSettings().Logging),
	}
	Reg = regexp.MustCompile("<command>")
}

func (s Help) Enable() {
	minatsubot.PluginAPI.RegisterCommand(
		minatsubot.CommandDescription{
			Name:        "help",
			Description: "List all commands",
			Usage:       "<command> [command|page]",
		}, &HelpCMD{})

}

func (s Help) Disable() {

}

type HelpCMD struct {
}

func (s HelpCMD) Run(command string, label string, args []string, message *discordgo.MessageCreate) {
	page := 1
	if len(args) > 0 {
		num, err := strconv.Atoi(args[0])
		if err != nil {
			command, ok := minatsubot.PluginAPI.GetCommand(args[0])
			if !ok {
				minatsubot.PluginAPI.Session.ChannelMessageSend(message.ChannelID, "The plugin "+args[0]+" is not a valid command")
			}

			w := &tabwriter.Writer{}
			buf := &bytes.Buffer{}
			w.Init(buf, 0, 4, 0, ' ', 0)

			fmt.Fprintf(w, "```Help Information\n")
			fmt.Fprintf(w, "Name: \t%s\n", command.Name)
			fmt.Fprintf(w, "Description: \t%s\n", command.Description)
			fmt.Fprintf(w, "Usage: \t%s%s\n", minatsubot.PluginAPI.GetSettings().Prefix, Reg.ReplaceAllString(command.Usage, command.Name))
			fmt.Fprintf(w, "Aliases: \t%s\n", joinString(command.Aliases))
			fmt.Fprintf(w, "```")

			w.Flush()
			out := buf.String()

			minatsubot.PluginAPI.Session.ChannelMessageSend(message.ChannelID, out)
			return
		}
		page = num
	}

	pages := NewPaginator(10)
	pages.AddList(minatsubot.PluginAPI.GetAllCommands())

	w := &tabwriter.Writer{}
	buf := &bytes.Buffer{}
	w.Init(buf, 0, 4, 0, ' ', 0)

	fmt.Fprintf(w, "```Help List\n")
	for _, desc := range pages.GetPage(page).GetText() {
		fmt.Fprintf(w, "Name: \t%s\n", desc.Name)
		fmt.Fprintf(w, "Description: \t%s\n", desc.Description)
		fmt.Fprintf(w, "\n")
	}
	fmt.Fprintf(w, "```")

	w.Flush()
	out := buf.String()

	minatsubot.PluginAPI.Session.ChannelMessageSend(message.ChannelID, out)
}

func joinString(str []string) string {
	s := strings.Join(str, ", ")
	if s != "" {
		return s
	}
	return "None"
}

func NewPaginator(limit int) *paginator {
	return &paginator{
		pages: []*page{},
		limit: limit,
	}
}

type page struct {
	text []minatsubot.CommandDescription
}

func (p *page) addText(s minatsubot.CommandDescription) {
	p.text = append(p.text, s)
}

func (p *page) GetText() []minatsubot.CommandDescription {
	return p.text
}

type paginator struct {
	pages []*page
	limit int
}

func (p *paginator) createPage() *page {
	pa := &page{text: []minatsubot.CommandDescription{}}
	p.pages = append(p.pages, pa)
	return pa
}

func (p *paginator) AddText(s minatsubot.CommandDescription) {
	var pa *page
	if len(p.pages) <= 0 {
		pa = p.createPage()
	}
	pa = p.pages[len(p.pages)-1]

	if len(pa.text)-1 >= p.limit {
		pa = p.createPage()
	}
	pa.addText(s)
}

func (p *paginator) AddList(s []minatsubot.CommandDescription) {
	for _, str := range s {
		p.AddText(str)
	}
}

func (p *paginator) GetPage(i int) *page {
	if i >= len(p.pages) {
		i = len(p.pages) - 1
	}

	return p.pages[i]
}

func (p *paginator) GetAmountPages() int {
	return len(p.pages)
}
