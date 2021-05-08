#!/usr/bin/swift

// # <xbar.title>My Azure Devops Pull Requests</xbar.title>
// # <xbar.version>v1.0</xbar.version>
// # <xbar.author>Evan Anger</xbar.author>
// # <xbar.author.github>devandanger</xbar.author.github>
// # <xbar.desc>Access Pull Requests of Interests</xbar.desc>
// # <xbar.dependencies>swift</xbar.dependencies>

import Foundation

var displayCount = 15 // Min: 10, Max: 25

func printOutput() {
    print("My Pull Requests")
    print("---")
    print("Assigned Pull Requests")
    print("---")
}

