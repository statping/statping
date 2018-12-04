/*
 * Statping
 * Copyright (C) 2018.  Hunter Long and the project contributors
 * Written by Hunter Long <info@socialeck.com> and the project contributors
 *
 * https://github.com/hunterlong/statping
 *
 * The licenses for most software and other practical works are designed
 * to take away your freedom to share and change the works.  By contrast,
 * the GNU General Public License is intended to guarantee your freedom to
 * share and change all versions of a program--to make sure it remains free
 * software for all its users.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

context('Asset Tests', () => {

    beforeEach(function () {
        cy.visit('http://localhost:8080/dashboard')
        cy.get('input[name="username"]').type('admin')
        cy.get('input[name="password"]').type('admin')
        cy.get('form').submit()
    })

    it('should create local assets', () => {
        cy.visit('http://localhost:8080/settings/build')
        cy.get('#v-pills-style-tab').click()
        cy.wait(500)
        cy.get(':nth-child(2) > .CodeMirror-line').should('contain', '$background-color')
    })

    it('should save assets form', () => {
        cy.request({method: 'POST', url: 'http://localhost:8080/settings/css', form: true, body: {
            variables: '$tester: #bababa',
            theme: '@import \'variables\';    .test-var { color: $tester; }'
        }})
    })

    it('should confirm sass variable in css', () => {
        cy.request('http://localhost:8080/css/base.css').its('body').should('contain', '.test-var')
    })

    it('should delete assets', () => {
        cy.visit('http://localhost:8080/settings')
        cy.get('#v-pills-style-tab').click()
        cy.wait(500)
        cy.get('.btn-danger').click()
    })

    it('should check css file after delete', () => {
        cy.request('http://localhost:8080/css/base.css').its('body').should('contain', 'BODY')
    })

});
