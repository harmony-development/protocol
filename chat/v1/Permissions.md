---
title: Permissions
---

The permissions system of Harmony resembles that of the permissions system common
to open-source Minecraft server software, featuring rules and querying a node string
against those rules.

## Permission

In a community, there will be features, management functionality, bot commands, and other things which exist.
Most of these actions have a permission associated with them, allowing you to control which members have access to each feature.

A permission is just a string, such as `roles.user.manage`, separated into parts using periods.

This string is also referred to as a permission node, or just node for short.

## Rule

A rule is a string that matches permissions, using UNIX-style globbing and parameter expansion.
This matching string is paired with a mode, which will either allow or deny permissions that match the string.

The rule `roles.*` will match `roles.user.manage` and `roles.user.view`.
This is a star-expression, meaning that anything beginning with the text before the star and ending with the text after the star will match.
Only one star-expression is permitted within a rule.

The rule `roles.user.{manage,view}` will match `roles.user.manage` and `roles.user.view`, but not `roles.user.share`.
This is an or-expression, meaning that `a.{b,c}` will match both `a.b` and `a.c`.
Any amount of items can be put in the or-expression's braces: `a.{b,c,d}` will match both `a.b`, `a.c`, and `a.d`.
Multiple or-expressions can compound, in which they effectively "multiply" the amount of items they expand to.

`a.{b,c}.{d,e}` will expand to `a.b.d`, `a.b.e`, `a.c.d`, and `a.c.e`.

## Role

A role is a collection of rules, which can be assigned to users.
These are named and coloured for aesthetic reasons.

Each guild MUST have a default role. This role MUST have the ID `0`.
The default role will not be included in `GetGuildRoles` or `GetUserRoles`.
Clients are free to show it however they want.

## Overrides

A role can also have channel and category-specific overrides, which are a collection of rules that take precedence over the guild-level set of rules.
Actions may query permissions in the context of a specific channel.
For example, you may want to allow `@everyone` to `messages.send` in most channels, but deny them that permission to send messages in an announcements channel.

Some actions may not query permissions in a channel or category, such as changing the name of a community.

If an action queries in the context of a channel, it'll first check the overrides of the channel, then the rules of the category of the channel if applicable, then the rules associated with a guild.

## Evaluation

```go
// Mode determines whether a permission will glob or not
type Mode int

// RoleID is the ID of a role
type RoleID uint64

// ChannelID is the ID of a channel
type ChannelID uint64

const (
    // Allow permission
    Allow Mode = iota
    // Deny permission
    Deny
)

type PermissionGlob struct {
}

func (p *PermissionGlob) Match(str string) bool {
    // return true if the string matches this glob,
    // otherwise false
}

type PermissionNode struct {
    Glob PermissionGlob
    Mode
}

type GuildState struct {
    // map of role IDs to list of permission nodes
    Roles      map[RoleID][]PermissionNode

    // map of channel IDs to parent channel IDs (categories)
    Categories map[ChannelID]ChannelID

    // map of channel IDs to map of role IDs to list of permission nodes
    Channels   map[ChannelID]map[RoleID][]PermissionNode
}

func (g GuildState) Check(permission string, userRoles []uint64, in ChannelID) bool {
    userRoles = append(userRoles, uint64(Everyone))

    if in != 0 {
        if channelData, ok := g.Channels[in]; ok {
            for _, role := range userRoles {
                nodes, ok := channelData[RoleID(role)]
                _ = ok
                for _, node := range nodes {
                    if node.Glob.Match(permission) {
                        return node.Mode == Allow
                    }
                }
            }
        }

        if category, ok := g.Categories[in]; ok {
            if channelData, ok := g.Channels[category]; ok {
                for _, role := range userRoles {
                    nodes, ok := channelData[RoleID(role)]
                    _ = ok
                    for _, node := range nodes {
                        if node.Glob.Match(permission) {
                            return node.Mode == Allow
                        }
                    }
                }
            }
        }
    }

    for _, role := range userRoles {
        nodes, ok := g.Roles[RoleID(role)]
        _ = ok
        for _, node := range nodes {
            if node.Glob.Match(permission) {
                return node.Mode == Allow
            }
        }
    }

    return false
}
```