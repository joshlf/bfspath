bfspath/example
===============

This folder contains an example implementation of a bfspath.Node type, and constructs graphs from graph definition files (example files in graphs folder). A graph definition file consists of one line listing the start and end nodes followed by lines defining directional edges and their lengths:

&lt;start&gt; &lt;end&gt;<br>
&lt;head node&gt; &lt;tail node&gt; &lt;edge length&gt;<br>
&lt;head node&gt; &lt;tail node&gt; &lt;edge length&gt;<br>
&lt;head node&gt; &lt;tail node&gt; &lt;edge length&gt;<br>

...etc

###Use

To use, remove the "_test.go" extension before building (the suffix prevents the "go get" command from installing this example directory).