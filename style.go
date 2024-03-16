package badge

var flatTemplate = stripXmlWhitespace(`
<svg xmlns="http://www.w3.org/2000/svg"                                            <text aria-hidden="true" x="{{.Bounds.SubjectX}}" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="{{.Bounds.SubjectLength}}">{{.Subject | html}}</text>
                                                <text x="{{.Bounds.SubjectX}}" y="140" transform="scale(.1)" fill="#fff" textLength="{{.Bounds.SubjectLength}}">{{.Subject | html}}</text>
                                                    <text aria-hidden="true" x="{{.Bounds.StatusX}}" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="{{.Bounds.StatusLength}}">{{.Status | html}}</text>
                                                        <text x="{{.Bounds.StatusX}}" y="140" transform="scale(.1)" fill="#fff" textLength="{{.Bounds.StatusLength}}">{{.Status | html}}</text>
                                                          </g>
                                                          </svg>
`)
